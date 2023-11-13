package openAiType

type OpenAiMessagesFileObject struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	MessageID string `json:"message_id"`
	FileID    string `json:"file_id"`
}
