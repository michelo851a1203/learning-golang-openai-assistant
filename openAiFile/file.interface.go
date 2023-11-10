package openAiFile

import "testf/openAiType"

type OpenAiFile interface {
	ListAll() (*openAiType.ListFileResponse, error)
	UploadFile() (*openAiType.OpenAiFileObject, error)
	DeleteFile() (*openAiType.DeleteFileResponse, error)
	GetFile() (*openAiType.OpenAiFileObject, error)
	GetFileContent() (string, error)
}
