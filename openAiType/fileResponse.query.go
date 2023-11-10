package openAiType

type ListFileResponse struct {
	Data   []OpenAiFileObject `json:"data"`
	Object string             `json:"object"`
}
