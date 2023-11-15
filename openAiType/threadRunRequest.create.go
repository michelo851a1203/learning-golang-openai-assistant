package openAiType

import "testf/openAiType/openAiModel"

type CreateThreadRunRequest struct {
	AssistantID  string                  `json:"assistant_id"`
	Model        openAiModel.OpenAiModel `json:"model,omitempty"`
	Instructions string                  `json:"instructions,omitempty"`
	Tools        []*OpenAiTool           `json:"tools,omitempty"`
	Metadata     *OpenAiMetaData         `json:"metadata,omitempty"`
}

type RunRequestThreadMessage struct {
	Role     string          `json:"role"`
	Content  string          `json:"content"`
	FileIDs  string          `json:"file_ids"`
	Metadata *OpenAiMetaData `json:"metadata,omitempty"`
}

type RunRequestThread struct {
	Messages []*RunRequestThreadMessage `json:"messages"`
	Metadata *OpenAiMetaData            `json:"metadata,omitempty"`
}

type ThreadAndRunRequest struct {
	AssistantID  string                  `json:"assistant_id"`
	Thread       string                  `json:"thread_id"`
	Model        openAiModel.OpenAiModel `json:"model"`
	Instructions *string                 `json:"instructions,omitempty"`
	Tools        []*OpenAiTool           `json:"tools,omitempty"`
	Metadata     *OpenAiMetaData         `json:"metadata,omitempty"`
}
