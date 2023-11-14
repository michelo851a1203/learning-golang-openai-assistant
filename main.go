package main

import (
	"os"
	"sync"
	"testf/openAiAssistant"
	"testf/openAiFile"
	"testf/openAiMessages"
	"testf/openAiThreadRun"
	"testf/openAiType"
	"testf/openAiType/openAiModel"
	"testf/openaiThreads"

	_ "github.com/joho/godotenv/autoload"
)

func CreateAssistant(assistantImpl *openAiAssistant.AssistantImpl) {
	assistantName := "測試用 assistant"
	description := "測試用的 assistant，隨後會刪掉"
	instruction := ""

	request := openAiType.CreateAssistantRequest{
		Model:        openAiModel.Gpt4TurboPreview,
		Name:         &assistantName,
		Description:  &description,
		Instructions: &instruction,
	}
	assistantImpl.CreateAssistant(&request)
}

func main() {
	apiKey := os.Getenv("API_KEY")
	wg := sync.WaitGroup{}
	assistantImpl := openAiAssistant.AssistantImpl{ApiKey: apiKey}
	assistantFileImpl := openAiAssistant.AssistantFileImpl{ApiKey: apiKey}
	openAiFileImpl := openAiFile.OpenAiFileImpl{ApiKey: apiKey}
	threadsImpl := openaiThreads.ThreadsImpl{ApiKey: apiKey}
	messagesImpl := openAiMessages.MessagesImpl{ApiKey: apiKey}
	threadRunImpl := openAiThreadRun.ThreadRunImpl{ApiKey: apiKey}

	go func() {

	}()

	wg.Wait()
}
