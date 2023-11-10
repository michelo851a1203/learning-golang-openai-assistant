package openAiType

type UpdateAssistantRequest struct {
	Model        string              `json:"model,omitempty"`
	Name         string              `json:"name,omitempty"`
	Description  string              `json:"description,omitempty"`
	Instructions string              `json:"instructions,omitempty"`
	Tools        []*OpenAiTool       `json:"tools,omitempty"`
	FileIds      []*OpenAiAttachFile `json:"file_ids,omitempty"`
	Metadata     OpenAiMetaData      `json:"metadata,omitempty"`
}
