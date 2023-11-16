package openAiType

import "fmt"

type OpenAiThreadObject struct {
	ID        string          `json:"id"`
	Object    string          `json:"object"`
	CreatedAt int64           `json:"created_at"`
	Metadata  *OpenAiMetaData `json:"metadata"`
}

func (openAiThreadObject *OpenAiThreadObject) String() string {
	return fmt.Sprintf(
		"\n  ID: %s, \n  Object: %s, \n  CreatedAt: %d, \n  Metadata: %v",
		openAiThreadObject.ID,
		openAiThreadObject.Object,
		openAiThreadObject.CreatedAt,
		openAiThreadObject.Metadata,
	)
}
