package openAiType

type OpenAiThreadObject struct {
	ID        string          `json:"id"`
	Object    string          `json:"object"`
	CreatedAt int64           `json:"created_at"`
	Metadata  *OpenAiMetaData `json:"metadata"`
}
