package openAiType

type OpenAiAssistantFileListResponse struct {
	Object  string                        `json:"object"`
	Data    []OpenAiAssistantFileResponse `json:"data"`
	FirstID string                        `json:"first_id"`
	LastID  string                        `json:"last_id"`
	HasMore bool                          `json:"has_more"`
}
