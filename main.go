package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testf/openAiFile"
	"testf/openAiType"
	"testf/openAiType/openAiFilePurpose"

	_ "github.com/joho/godotenv/autoload"
)

var filePool *sync.Pool

func main() {

	apiKey := os.Getenv("API_KEY")
	// assistantImpl := openAiAssistant.AssistantImpl{ApiKey: apiKey}
	// assistantFileImpl := openAiAssistant.AssistantFileImpl{ApiKey: apiKey}
	openAiFileImpl := openAiFile.OpenAiFileImpl{ApiKey: apiKey}
	// threadsImpl := openaiThreads.ThreadsImpl{ApiKey: apiKey}
	// messagesImpl := openAiMessages.MessagesImpl{ApiKey: apiKey}
	// threadRunImpl := openAiThreadRun.ThreadRunImpl{ApiKey: apiKey}

	testFilePath := "./filesForTest"
	fileInfo, err := os.ReadDir(testFilePath)
	if err != nil {
		panic(err)
	}

	openAiFileChan := make(chan *openAiType.OpenAiFileObject)
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

	// go func() {
	// 	request := openAiType.CreateAssistantRequest{
	// 		Model:        openAiModel.Gpt4TurboPreview,
	// 		Name:         &trainingParameters.AssistantName,
	// 		Description:  &trainingParameters.Description,
	// 		Instructions: &trainingParameters.Instruction,
	// 	}
	//
	// }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGQUIT)

	go func() {
		<-quit
		fmt.Println("結束...")

		deleteFileWaitGroup := sync.WaitGroup{}
		deleteFileMap := *filePool.Get().(*map[string]*openAiType.OpenAiFileObject)
		deleteFileWaitGroup.Add(len(deleteFileMap))

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
		deleteFileWaitGroup.Wait()
		os.Exit(0)
	}()

	for {
		fmt.Println("請輸入文字")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("==================")
		result := filePool.Get().(*map[string]*openAiType.OpenAiFileObject)
		fmt.Println(result)
		filePool.Put(result)
		fmt.Println("==================")
		fmt.Printf("result : %s", text)
	}
}
