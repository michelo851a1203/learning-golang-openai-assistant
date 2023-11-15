package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testf/openAiAssistant"
	"testf/openAiFile"
	"testf/openAiMessages"
	"testf/openAiThreadRun"
	"testf/openAiType"
	"testf/openAiType/openAiFilePurpose"
	"testf/openAiType/openAiModel"
	"testf/openAiType/openAiRole"
	"testf/openAiType/openAiRunStatus"
	"testf/openAiType/openAiTool"
	"testf/openaiThreads"
	"testf/trainingParameters"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var filePool *sync.Pool
var assistantPool *sync.Pool
var threadPool *sync.Pool

func main() {
	apiKey := os.Getenv("API_KEY")
	assistantImpl := openAiAssistant.AssistantImpl{ApiKey: apiKey}
	openAiFileImpl := openAiFile.OpenAiFileImpl{ApiKey: apiKey}
	threadsImpl := openaiThreads.ThreadsImpl{ApiKey: apiKey}
	messagesImpl := openAiMessages.MessagesImpl{ApiKey: apiKey}
	threadRunImpl := openAiThreadRun.ThreadRunImpl{ApiKey: apiKey}

	testFilePath := "./filesForTest"
	fileInfo, err := os.ReadDir(testFilePath)
	if err != nil {
		panic(err)
	}

	openAiFileChan := make(chan *openAiType.OpenAiFileObject, 5)
	fileWaitGroup := sync.WaitGroup{}
	fileWaitGroup.Add(len(fileInfo))
	for _, info := range fileInfo {
		go func(
			fileName string,
			fileApi *openAiFile.OpenAiFileImpl,
			wg *sync.WaitGroup,
		) {
			defer wg.Done()

			preTrainFile, err := os.Open(fmt.Sprintf("./%s/%s", testFilePath, fileName))

			if err != nil {
				fmt.Printf("打開預訓練檔案錯誤 -> [%s]\n", err.Error())
				fmt.Println("=======================")
				return
			}
			defer preTrainFile.Close()
			openAiFile, err := fileApi.UploadFile(openAiFilePurpose.Assistant, preTrainFile)
			if err != nil {
				fmt.Printf("上傳預訓練檔案錯誤 -> [%s]\n", err.Error())
				fmt.Println("=======================")
				return
			}
			fmt.Printf("上傳預訓練檔案成功 fileID -> [%s], 檔案名稱:[%s]\n", openAiFile.ID, fileName)
			fmt.Println("=======================")
			openAiFileChan <- openAiFile
		}(info.Name(), &openAiFileImpl, &fileWaitGroup)
	}

	go func() {
		fileWaitGroup.Wait()
		close(openAiFileChan)
	}()

	fileMap := map[string]*openAiType.OpenAiFileObject{}
	for openAiFile := range openAiFileChan {
		fileMap[openAiFile.ID] = openAiFile
	}
	filePool = &sync.Pool{
		New: func() interface{} {
			return &fileMap
		},
	}

	go func(api *openAiAssistant.AssistantImpl) {
		poolFileMap := filePool.Get().(*map[string]*openAiType.OpenAiFileObject)
		defer filePool.Put(poolFileMap)
		fileIDList := []string{}

		for fileID := range *poolFileMap {
			fileIDList = append(fileIDList, fileID)
		}

		request := openAiType.CreateAssistantRequest{
			Model:        openAiModel.Gpt4TurboPreview,
			Name:         &trainingParameters.AssistantName,
			Instructions: &trainingParameters.Instruction,
			Tools: []*openAiType.OpenAiTool{
				{
					Type: openAiTool.Retrieval,
				},
			},
			FileIds: fileIDList,
		}
		assistant, err := api.CreateAssistant(&request)
		if err != nil {
			fmt.Printf("創建助理發生錯誤 : [%s]\n", err.Error())
			fmt.Println("================")
			return
		}
		fmt.Println("創建助理成功：================")
		fmt.Println(assistant)
		fmt.Println("================")
		assistantPool = &sync.Pool{
			New: func() interface{} {
				return assistant
			},
		}
	}(&assistantImpl)

	go func(api *openaiThreads.ThreadsImpl) {
		request := openAiType.ThreadCreateRequest{}
		thread, err := api.CreateThread(&request)
		if err != nil {
			fmt.Printf("創建線程發生錯誤 : [%s]\n", err.Error())
			fmt.Println("================")
			return
		}
		fmt.Println("創建線程成功：================")
		fmt.Println(thread)
		fmt.Println("================")
		threadPool = &sync.Pool{
			New: func() interface{} {
				return thread
			},
		}

	}(&threadsImpl)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGQUIT)

	go func() {
		<-quit
		fmt.Println("結束...")

		deleteFileWaitGroup := sync.WaitGroup{}
		deleteFileMap := *filePool.Get().(*map[string]*openAiType.OpenAiFileObject)
		deleteFileWaitGroup.Add(len(deleteFileMap) + 2)

		for openAiFileID := range deleteFileMap {
			go func(fileID string) {
				defer deleteFileWaitGroup.Done()
				deleteResponse, err := openAiFileImpl.DeleteFile(fileID)
				fmt.Println("================")
				if err != nil {
					fmt.Printf("刪除檔案發生錯誤 : 檔案 ID : %s -> 錯誤為:[%s]\n", fileID, err.Error())
					fmt.Println("================")
					return
				}
				if deleteResponse.Deleted {
					fmt.Printf("刪除檔案成功 : 檔案 ID : %s \n", fileID)
					fmt.Println("================")
					return
				}

				fmt.Printf("刪除檔案失敗 : 檔案 ID : %s \n", fileID)
				fmt.Println("================")
			}(openAiFileID)
		}

		go func() {
			defer deleteFileWaitGroup.Done()
			currentAssistant, ok := assistantPool.Get().(*openAiType.AssistantObject)
			fmt.Println(currentAssistant)
			fmt.Println(ok)

			if currentAssistant != nil && ok {
				deleteAssistantResult, err := assistantImpl.DeleteAssistant(currentAssistant.ID)
				if err != nil {
					fmt.Printf("刪除助理發生錯誤 : [%s]\n", err.Error())
					fmt.Println("================")
				}
				if deleteAssistantResult.Deleted {
					fmt.Println("刪除助理成功")
					fmt.Println("================")
				} else {
					fmt.Println("刪除助理失敗")
					fmt.Println("================")
				}
			}
		}()

		go func() {
			defer deleteFileWaitGroup.Done()
			currentThread, ok := threadPool.Get().(*openAiType.OpenAiThreadObject)
			if currentThread != nil && ok {
				threadDeletedResult, err := threadsImpl.DeleteThread(currentThread.ID)
				if err != nil {
					fmt.Printf("刪除線程發生錯誤 : [%s]\n", err.Error())
					fmt.Println("================")
				}
				if threadDeletedResult.Deleted {
					fmt.Println("刪除線程成功")
					fmt.Println("================")
				} else {
					fmt.Println("刪除線程失敗")
					fmt.Println("================")
				}
			}
		}()

		deleteFileWaitGroup.Wait()

		os.Exit(0)
	}()

	for {
		fmt.Println("請輸入文字")
		reader := bufio.NewReader(os.Stdin)
		useInputText, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		currentAssistant := assistantPool.Get().(*openAiType.AssistantObject)
		currentThread := threadPool.Get().(*openAiType.OpenAiThreadObject)
		assistantID := currentAssistant.ID
		threadID := currentThread.ID

		createMessageRequest := openAiType.CreateMessagesRequest{
			Role:    openAiRole.User,
			Content: useInputText,
		}

		messageResponse, err := messagesImpl.CreateMessages(threadID, &createMessageRequest)
		if err != nil {
			fmt.Printf("創建訊息發生錯誤 : [%s]\n", err.Error())
			fmt.Println("================")
			continue
		}

		runRequest := openAiType.CreateThreadRunRequest{
			AssistantID: assistantID,
		}

		replyObject, err := threadRunImpl.CreateRun(threadID, &runRequest)
		if err != nil {
			fmt.Printf("執行訊息發生錯誤 : [%s]\n", err.Error())
			fmt.Println("================")
			continue
		}
		runID := replyObject.ID
		for range time.Tick(time.Millisecond * 500) {
			stepListResponse, err := threadRunImpl.GetRunStepList(threadID, runID, nil)
			if err != nil {
				fmt.Printf("取得執行步驟列表發生錯誤 : [%s]\n", err.Error())
				fmt.Println("================")
				break
			}
			if len(stepListResponse.Data) == 0 {
				continue
			}
			currentRunObject, err := threadRunImpl.GetRun(threadID, runID)
			if err != nil {
				fmt.Printf("取得執行狀態發生錯誤 : [%s]\n", err.Error())
				fmt.Println("================")
				break
			}
			if currentRunObject.Status == openAiRunStatus.InProgress {
				continue
			}
			fmt.Printf("對話任務執行完成... 狀態: [%s]\n", currentRunObject.Status)
			fmt.Println("================")
			break
		}

		queryRequest := openAiType.QueryListRequest{
			Limit:  10,
			Before: messageResponse.ID,
		}
		messageListResponse, err := messagesImpl.GetMessagesList(threadID, &queryRequest)
		if err != nil {
			fmt.Printf("取得對話列表發生錯誤 : [%s]\n", err.Error())
			fmt.Println("================")
			continue
		}
		messageDataList := messageListResponse.Data
		if len(messageDataList) == 0 {
			fmt.Println("沒有取得對話錯誤")
			fmt.Println("================")
			return
		}
		contentList := messageDataList[0].Content
		if len(contentList) == 0 {
			fmt.Println("沒有取得對話錯誤2")
			fmt.Println("================")
			return
		}

		fmt.Println("對話結果================")
		outputText := contentList[0].Text

		fmt.Println(outputText.Value)
		fmt.Println("================")

		threadPool.Put(currentThread)
		threadPool.Put(currentAssistant)

		fmt.Println("==================")
	}
}
