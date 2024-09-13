package modelsyncer

import (
	"context"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/llm-operator/inference-manager/engine/internal/ollama"
	mv1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestPullModel(t *testing.T) {
	tcs := []struct {
		name                 string
		modelID              string
		model                *mv1.Model
		wantCreated          []string
		wantRegisteredModels []string
	}{
		{
			name:    "system model",
			modelID: "google-gemma-2b",
			model: &mv1.Model{
				Id:      "google-gemma-2b",
				OwnedBy: "system",
			},
			wantCreated: []string{
				"google-gemma-2b",
			},
			wantRegisteredModels: []string{
				"google-gemma-2b",
			},
		},
		{
			name:    "non-system model",
			modelID: "ft:google-gemma-2b:fine-tuning-wpsd9kb5nl",
			model: &mv1.Model{

				Id:      "ft:google-gemma-2b:fine-tuning-wpsd9kb5nl",
				OwnedBy: "fake-tenant-id",
			},
			wantCreated: []string{
				"google-gemma-2b",
				"google-gemma-2b:fine-tuning-wpsd9kb5nl",
			},
			wantRegisteredModels: []string{
				"google-gemma-2b",
				"ft:google-gemma-2b:fine-tuning-wpsd9kb5nl",
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fom := &fakeOllamaManager{}
			om, err := New(
				fom,
				&noopS3Client{},
				&fakeModelInternalClient{model: tc.model},
			)
			assert.NoError(t, err)
			err = om.PullModel(context.Background(), tc.modelID)
			assert.NoError(t, err)

			assert.ElementsMatch(t, tc.wantCreated, fom.created)

			var registered []string
			for k := range om.registeredModels {
				registered = append(registered, k)
			}
			assert.ElementsMatch(t, tc.wantRegisteredModels, registered)
			assert.Empty(t, om.inProgressModels)
		})
	}
}

func TestPullModelInProgress(t *testing.T) {
	const (
		modelID = "google-gemma-2b"
	)
	waitCh := make(chan struct{})

	fom := &fakeOllamaManager{}
	om, err := New(
		fom,
		&blockingS3Client{
			waitcCh: waitCh,
		},
		&fakeModelInternalClient{
			model: &mv1.Model{
				Id:      modelID,
				OwnedBy: "system",
			},
		},
	)
	assert.NoError(t, err)

	// Start two goroutines to pull the same model.
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err := om.PullModel(context.Background(), modelID)
		assert.NoError(t, err)
		wg.Done()
	}()
	go func() {
		err := om.PullModel(context.Background(), modelID)
		assert.NoError(t, err)
		wg.Done()
	}()

	// Wait until the pull starts.
	assert.Eventually(
		t,
		func() bool {
			om.mu.Lock()
			b := len(om.inProgressModels) == 1 || len(om.registeredModels) == 1
			om.mu.Unlock()
			return b
		},
		time.Second,
		10*time.Millisecond,
	)

	close(waitCh)
	wg.Wait()

	assert.ElementsMatch(t, []string{modelID}, fom.created)

	var registered []string
	for k := range om.registeredModels {
		registered = append(registered, k)
	}
	assert.ElementsMatch(t, []string{modelID}, registered)
	assert.Empty(t, om.inProgressModels)
}

type fakeOllamaManager struct {
	created []string
}

func (n *fakeOllamaManager) CreateNewModelOfGGUF(modelID string, spec *ollama.ModelSpec) error {
	n.created = append(n.created, modelID)
	return nil
}

func (n *fakeOllamaManager) DownloadAndCreateNewModel(ctx context.Context, modelID string, resp *mv1.GetBaseModelPathResponse) error {
	n.created = append(n.created, modelID)
	return nil
}

func (n *fakeOllamaManager) UpdateModelTemplateToLatest(modelID string) error {
	return nil
}

type noopS3Client struct {
}

func (n *noopS3Client) Download(ctx context.Context, f io.WriterAt, path string) error {
	return nil
}

type fakeModelInternalClient struct {
	model *mv1.Model
}

func (n *fakeModelInternalClient) GetModel(ctx context.Context, in *mv1.GetModelRequest, opts ...grpc.CallOption) (*mv1.Model, error) {

	return n.model, nil

}

func (n *fakeModelInternalClient) GetBaseModelPath(ctx context.Context, in *mv1.GetBaseModelPathRequest, opts ...grpc.CallOption) (*mv1.GetBaseModelPathResponse, error) {
	return &mv1.GetBaseModelPathResponse{
		Path:          "fake-path",
		Formats:       []mv1.ModelFormat{mv1.ModelFormat_MODEL_FORMAT_GGUF},
		GgufModelPath: "fake-gguf-path",
	}, nil
}

func (n *fakeModelInternalClient) GetModelPath(ctx context.Context, in *mv1.GetModelPathRequest, opts ...grpc.CallOption) (*mv1.GetModelPathResponse, error) {
	return &mv1.GetModelPathResponse{
		Path: "fake-path",
	}, nil
}

type blockingS3Client struct {
	waitcCh chan struct{}
}

func (b *blockingS3Client) Download(ctx context.Context, f io.WriterAt, path string) error {
	<-b.waitcCh
	return nil
}
