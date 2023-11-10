package openAiType

type DeleteAssistantResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Delete bool   `json:"deleted"`
}
