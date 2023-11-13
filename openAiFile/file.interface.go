package openAiFile

import (
	"os"
	"testf/openAiType"
	"testf/openAiType/openAiFilePurpose"
)

type OpenAiFile interface {
	GetFileList(purpose string) (*openAiType.ListFileResponse, error)
	UploadFile(purpose openAiFilePurpose.PurposeStatus, preparedFile *os.File) (*openAiType.OpenAiFileObject, error)
	DeleteFile(fileID string) (*openAiType.DeleteResponse, error)
	GetFile(fileID string) (*openAiType.OpenAiFileObject, error)
	GetFileContent(fileID string) (string, error)
}
