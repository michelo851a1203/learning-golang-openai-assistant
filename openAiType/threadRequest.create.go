package openAiType

type ThreadCreateRequest struct {
	Messages []*OpenAiMessages `json:"messages,omitempty"`
	MetaData *OpenAiMetaData   `json:"metadata,omitempty"`
}
