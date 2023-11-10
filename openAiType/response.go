package openAiType

type AssistantObject struct {
	ID           string              `json:"id"`
	Object       string              `json:"object"`
	CreatedAt    int64               `json:"created_at"`
	Name         string              `json:"name,omitempty"`
	Description  string              `json:"description,omitempty"`
	Model        string              `json:"model"`
	Instructions string              `json:"instructions,omitempty"`
	Tools        []*OpenAiTool       `json:"tools,omitempty"`
	FileIDs      []*OpenAiAttachFile `json:"file_ids,omitempty"`
	Metadata     OpenAiMetaData      `json:"metadata"`
}
