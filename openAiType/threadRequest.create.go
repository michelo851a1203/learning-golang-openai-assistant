package openAiType

type ThreadCreateRequest struct {
	Messages []*OpenAiMessagesObject `json:"messages,omitempty"`
	MetaData *OpenAiMetaData         `json:"metadata,omitempty"`
}
