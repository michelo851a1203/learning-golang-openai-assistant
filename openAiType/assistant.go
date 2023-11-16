package openAiType

import "fmt"

type AssistantObject struct {
	ID           string              `json:"id"`
	Object       string              `json:"object"`
	CreatedAt    int64               `json:"created_at"`
	Name         *string             `json:"name,omitempty"`
	Description  *string             `json:"description,omitempty"`
	Model        string              `json:"model"`
	Instructions *string             `json:"instructions,omitempty"`
	Tools        []*OpenAiTool       `json:"tools,omitempty"`
	FileIDs      []*OpenAiFileObject `json:"file_ids,omitempty"`
	Metadata     *OpenAiMetaData     `json:"metadata"`
}

func (assistantObject *AssistantObject) String() string {
	return fmt.Sprintf(
		"\n  ID: %s, \n  Object: %s, \n  CreatedAt: %d, \n  Name: %v, \n  Description: %v, \n  Model: %s, \n  Instructions: %v, \n  Tools: %v, \n  FileIDs: %v, \n  Metadata: %v",
		assistantObject.ID,
		assistantObject.Object,
		assistantObject.CreatedAt,
		assistantObject.Name,
		assistantObject.Description,
		assistantObject.Model,
		assistantObject.Instructions,
		assistantObject.Tools,
		assistantObject.FileIDs,
		assistantObject.Metadata,
	)
}
