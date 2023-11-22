package openAiType

import "testf/openAiType/openAiListObject"

type ListFileResponse struct {
	Data   []*OpenAiFileObject           `json:"data"`
	Object openAiListObject.ObjectStatus `json:"object"`
}
