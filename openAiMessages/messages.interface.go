package openAiMessages

import "testf/openAiType"

type Messages interface {
	CreateMessages(threadID string, createRequest *openAiType.CreateMessagesRequest) (*openAiType.OpenAiMessagesObject, error)
	ModifyMessages(threadID string, messagesID string) (*openAiType.OpenAiMessagesObject, error)
	GetMessagesList(threadID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.OpenAiMessagesObject], error)
	GetMessages(threadID string, MessagesID string) (*openAiType.OpenAiMessagesObject, error)
}
