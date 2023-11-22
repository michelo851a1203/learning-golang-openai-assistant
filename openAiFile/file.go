package openAiFile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testf/openAiError"
	"testf/openAiError/openAiErrorCode"
	"testf/openAiType"
	"testf/openAiType/openAiFilePurpose"
)

type OpenAiFileImpl struct {
	ApiKey string
}

func (openAiFileImpl *OpenAiFileImpl) GetFileList(
	purpose openAiFilePurpose.PurposeStatus,
) (
	*openAiType.ListFileResponse,
	error,
) {
	queryString := ""
	if purpose != "" {
		queryString = "?purpose=" + string(purpose)
	}

	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/files%s", queryString),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileListNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetFileList",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	openAiFileClient := &http.Client{}

	response, err := openAiFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileListSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetFileList",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileListReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetFileList",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	result := &openAiType.ListFileResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileListResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetFileList",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	if result.Object == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.FileError]{
				OpenStatusCode: openAiErrorCode.GetFileListErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetFileList",
				RawError:       err.Error(),
				Details:        &openAiError.FileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileListOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetFileList",
			RawError:       errorResult.String(),
			Details:        &openAiError.FileError{},
		}
	}

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) UploadFile(
	purpose openAiFilePurpose.PurposeStatus,
	preparedFile *os.File,
) (
	*openAiType.OpenAiFileObject,
	error,
) {
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// file
	fileWriter, err := multipartWriter.CreateFormFile("file", preparedFile.Name())
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileCreateFormFileError,
			Message:        "CreateFormFile: file Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	_, err = io.Copy(fileWriter, preparedFile)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileCopyFileError,
			Message:        "File Copy Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}
	// purpose
	err = multipartWriter.WriteField("purpose", string(purpose))

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFilePurposeFieldError,
			Message:        "WriteField: purpose Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	err = multipartWriter.Close()

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileMultipartWriterError,
			Message:        "multipartWriter Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/files",
		&requestBody,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileNewRequestError,
			Message:        "NewRequest Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	request.Header.Add("Content-Type", multipartWriter.FormDataContentType())
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	openAiFileClient := &http.Client{}

	response, err := openAiFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	result := &openAiType.OpenAiFileObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "UploadFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.FileError]{
				OpenStatusCode: openAiErrorCode.UploadFileErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "UploadFile",
				RawError:       err.Error(),
				Details:        &openAiError.FileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.UploadFileOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "UploadFile",
			RawError:       errorResult.String(),
			Details:        &openAiError.FileError{},
		}
	}

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) DeleteFile(fileID string) (
	*openAiType.DeleteResponse,
	error,
) {
	request, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.openai.com/v1/files/%s", fileID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.DeleteFileNewRequestError,
			Message:        "NewRequest Error",
			Method:         "DeleteFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	deleteFileClient := &http.Client{}

	response, err := deleteFileClient.Do(request)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.DeleteFileSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "DeleteFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.DeleteFileReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "DeleteFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	result := &openAiType.DeleteResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.DeleteFileResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "DeleteFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.FileError]{
				OpenStatusCode: openAiErrorCode.DeleteFileErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "DeleteFile",
				RawError:       err.Error(),
				Details:        &openAiError.FileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.DeleteFileOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "DeleteFile",
			RawError:       errorResult.String(),
			Details:        &openAiError.FileError{},
		}
	}

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) GetFile(fileID string) (
	*openAiType.OpenAiFileObject,
	error,
) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/files/%s", fileID),
		nil,
	)

	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
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
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	result := &openAiType.OpenAiFileObject{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileResponseJSONError,
			Message:        "Response JSON Error",
			Method:         "GetFile",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	if result.ID == "" {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return nil, &openAiError.OpenAiError[openAiError.FileError]{
				OpenStatusCode: openAiErrorCode.GetFileErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetFile",
				RawError:       err.Error(),
				Details:        &openAiError.FileError{},
			}
		}

		return nil, &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetFile",
			RawError:       errorResult.String(),
			Details:        &openAiError.FileError{},
		}
	}

	return result, nil
}

func (openAiFileImpl *OpenAiFileImpl) GetFileContent(fileID string) (string, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api.openai.com/v1/files/%s/content", fileID),
		nil,
	)

	if err != nil {
		return "", &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileContentNewRequestError,
			Message:        "NewRequest Error",
			Method:         "GetFileContent",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", openAiFileImpl.ApiKey))

	getFileClient := &http.Client{}

	response, err := getFileClient.Do(request)
	if err != nil {
		return "", &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileContentSendHTTPRequestError,
			Message:        "Send Http Request Error",
			Method:         "GetFileContent",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileContentReadResponseBodyError,
			Message:        "Read Response Body Error",
			Method:         "GetFileContent",
			RawError:       err.Error(),
			Details:        &openAiError.FileError{},
		}
	}

	if len(body) == 0 {
		errorResult := &openAiError.OpenAiNativeApiError{}
		err = json.Unmarshal(body, errorResult)
		if err != nil {
			return "", &openAiError.OpenAiError[openAiError.FileError]{
				OpenStatusCode: openAiErrorCode.GetFileContentErrorResponseJSONError,
				Message:        "Error Response JSON Error",
				Method:         "GetFileContent",
				RawError:       err.Error(),
				Details:        &openAiError.FileError{},
			}
		}

		return "", &openAiError.OpenAiError[openAiError.FileError]{
			OpenStatusCode: openAiErrorCode.GetFileContentOpenAIError,
			Message:        "OpenAI Response Error",
			Method:         "GetFileContent",
			RawError:       errorResult.String(),
			Details:        &openAiError.FileError{},
		}
	}

	return string(body), nil
}
