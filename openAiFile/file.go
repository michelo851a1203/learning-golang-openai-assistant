package openAiFile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testf/openAiType"
)

type OpenAiFileImpl struct {
	ApiKey       string
	FileID       string
	PreparedFile *os.File
	Purpose      string
}

func (openAiFileImpl *OpenAiFileImpl) GetFileList() (
	*openAiType.ListFileResponse,
	error,
) {

	request, err := http.NewRequest(
		http.MethodGet,
		"https://api.openai.com/v1/files",
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	openAiFileClient := &http.Client{}

	response, err := openAiFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.ListFileResponse{}
	json.Unmarshal(body, result)

	return result, nil

}

func (openAiFileImpl *OpenAiFileImpl) UploadFile() (
	*openAiType.OpenAiFileObject,
	error,
) {

	currentFile := openAiFileImpl.PreparedFile

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	defer multipartWriter.Close()

	// file
	fileWriter, err := multipartWriter.CreateFormFile("file", currentFile.Name())
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, currentFile)
	if err != nil {
		return nil, err
	}
	// purpose
	err = multipartWriter.WriteField("purpose", openAiFileImpl.Purpose)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/files",
		&requestBody,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	openAiFileClient := &http.Client{}

	response, err := openAiFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiFileObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) DeleteFile() (
	*openAiType.DeleteFileResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/files/%s", openAiFileImpl.FileID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	deleteFileClient := &http.Client{}

	response, err := deleteFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.DeleteFileResponse{}
	json.Unmarshal(body, result)

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) GetFile() (
	*openAiType.OpenAiFileObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/files/%s", openAiFileImpl.FileID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	getFileClient := &http.Client{}

	response, err := getFileClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &openAiType.OpenAiFileObject{}
	json.Unmarshal(body, result)

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) GetFileContent() (string, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/files/%s/content", openAiFileImpl.FileID),
		nil,
	)

	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	getFileClient := &http.Client{}

	response, err := getFileClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
