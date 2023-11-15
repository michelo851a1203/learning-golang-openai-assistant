package openAiMessages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testf/openAiType"
)

type MessageFileImpl struct {
	ApiKey string
}

func (MessageFileImpl *MessageFileImpl) GetMessageFileList(
	threadID string,
	messageID string,
	listRequest *openAiType.QueryListRequest,
) (
	*openAiType.ListResponse[openAiType.OpenAiMessagesFileObject],
	error,
) {

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/messages/%s/files%s",
			threadID,
			messageID,
			listRequest.ToQueryString(),
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessageFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	MessageFileClient := &http.Client{}

	response, err := MessageFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListResponse[openAiType.OpenAiMessagesFileObject]{}
	json.Unmarshal(body, result)

	return result, nil
}

func (MessageFileImpl *MessageFileImpl) GetMessageFile(
	threadID string,
	messageID string,
	fileID string,
) (
	*openAiType.OpenAiMessagesFileObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"https://api.openai.com/v1/threads/%s/messages/%s/files/%s",
			threadID,
			messageID,
			fileID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", MessageFileImpl.ApiKey))
	request.Header.Add("OpenAI-Beta", "assistants=v1")

	detailMessageFileClient := &http.Client{}

	response, err := detailMessageFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiMessagesFileObject{}
	json.Unmarshal(body, result)

	return result, nil
}
