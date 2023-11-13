package openAiMessages

import "testf/openAiType"

type MessageFile interface {
	GetMessageFileList(threadID string, messageID string, listRequest *openAiType.QueryListRequest) (*openAiType.ListResponse[openAiType.OpenAiMessagesFileObject], error)
	GetMessageFile(threadID string, messageID string, fileID string) (*openAiType.OpenAiMessagesFileObject, error)
}
