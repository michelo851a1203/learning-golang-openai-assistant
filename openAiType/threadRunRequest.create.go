package openAiType

import "testf/openAiType/openAiModel"

type CreateThreadRunRequest struct {
	AssistantID  string                  `json:"assistant_id"`
	Model        openAiModel.OpenAiModel `json:"model"`
	Instructions string                  `json:"instructions,omitempty"`
	Tools        []*OpenAiTool           `json:"tools,omitempty"`
	Metadata     *OpenAiMetaData         `json:"metadata,omitempty"`
}
