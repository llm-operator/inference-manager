package runtime

import (
	"context"
	"testing"

	"github.com/llm-operator/inference-manager/engine/internal/config"
	mv1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestDeployRuntimeParams(t *testing.T) {

	commonClient := &commonClient{
		RuntimeConfig: config.RuntimeConfig{
			ModelResources: map[string]config.Resources{
				"TinyLlama-TinyLlama-1.1B-Chat-v1.0": {
					Limits: map[string]string{
						"cpu":            "1",
						"nvidia.com/gpu": "2",
					},
				},
			},
		},
	}

	tcs := []struct {
		name     string
		modelID  string
		resp     *mv1.GetBaseModelPathResponse
		wantArgs []string
	}{
		{
			name:    "both formats",
			modelID: "TinyLlama-TinyLlama-1.1B-Chat-v1.0",
			resp: &mv1.GetBaseModelPathResponse{
				Path: "path",
				Formats: []mv1.ModelFormat{
					mv1.ModelFormat_MODEL_FORMAT_GGUF,
					mv1.ModelFormat_MODEL_FORMAT_HUGGING_FACE,
				},
			},
			wantArgs: []string{
				"--port", "80",
				"--model", "/models",
				"--served-model-name", "TinyLlama-TinyLlama-1.1B-Chat-v1.0",
				"--chat-template", "\n<|begin_of_text|>\n{% for message in messages %}\n{{'<|start_header_id|>' + message['role'] + '<|end_header_id|>\\n' + message['content'] + '\\n<|eot_id|>\\n'}}\n{% endfor %}\n",
				"--tensor-parallel-size", "2",
			},
		},
		{
			name:    "gguf only",
			modelID: "TinyLlama-TinyLlama-1.1B-Chat-v1.0",
			resp: &mv1.GetBaseModelPathResponse{
				Path: "path",
				Formats: []mv1.ModelFormat{
					mv1.ModelFormat_MODEL_FORMAT_GGUF,
				},
			},
			wantArgs: []string{
				"--port", "80",
				"--model", "/models/TinyLlama-TinyLlama-1.1B-Chat-v1.0.gguf",
				"--served-model-name", "TinyLlama-TinyLlama-1.1B-Chat-v1.0",
				"--chat-template", "\n<|begin_of_text|>\n{% for message in messages %}\n{{'<|start_header_id|>' + message['role'] + '<|end_header_id|>\\n' + message['content'] + '\\n<|eot_id|>\\n'}}\n{% endfor %}\n",
				"--tensor-parallel-size", "2",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			v := &vllmClient{
				commonClient: commonClient,
				modelClient: &fakeModelClient{
					resp: tc.resp,
				},
			}

			got, err := v.deployRuntimeParams(context.Background(), tc.modelID)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantArgs, got.args)
		})
	}
}

func TestNumGPUs(t *testing.T) {
	v := &vllmClient{
		commonClient: &commonClient{
			RuntimeConfig: config.RuntimeConfig{
				ModelResources: map[string]config.Resources{
					"model0": {
						Limits: map[string]string{
							nvidiaGPUResource: "2",
							"cpu":             "4",
						},
					},
					"model1": {
						Limits: map[string]string{
							"cpu": "8",
						},
					},
				},
			},
		},
	}

	tcs := []struct {
		name    string
		modelID string
		want    int
	}{
		{
			name:    "model0",
			modelID: "model0",
			want:    2,
		},
		{
			name:    "model1",
			modelID: "model1",
			want:    0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := v.numGPUs(tc.modelID)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

type fakeModelClient struct {
	resp *mv1.GetBaseModelPathResponse
}

func (c *fakeModelClient) GetBaseModelPath(ctx context.Context, in *mv1.GetBaseModelPathRequest, opts ...grpc.CallOption) (*mv1.GetBaseModelPathResponse, error) {
	return c.resp, nil
}
