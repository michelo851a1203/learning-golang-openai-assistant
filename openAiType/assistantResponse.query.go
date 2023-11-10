package openAiType

type OpenAiListAssistantResponse struct {
	Object  string             `json:"object"`
	Data    []*AssistantObject `json:"data"`
	FirstID string             `json:"first_id"`
	LastID  string             `json:"last_id"`
	HasMore bool               `json:"has_more"`
}
