package openAiType

type CreateAssistantRequest struct {
	Model        string              `json:"model"`
	Name         string              `json:"name,omitempty"`
	Description  string              `json:"description,omitempty"`
	Instructions string              `json:"instructions,omitempty"`
	Tools        []*OpenAiTool       `json:"tools,omitempty"`
	FileIds      []*OpenAiFileObject `json:"file_ids,omitempty"`
	Metadata     OpenAiMetaData      `json:"metadata,omitempty"`
}
