package openAiType

import (
	"testf/openAiType/openAiAnnotation"
	"testf/openAiType/openAiContentType"
	"testf/openAiType/openAiRole"
)

type OpenAiFileCitationsObject struct {
	FileID string `json:"file_id"`
	Quote  string `json:"quote"`
}

type OpenAiFilePathObject struct {
	FileID string `json:"file_id"`
}

type OpenAiAnnotationsObject struct {
	Type          openAiAnnotation.AnnotationType `json:"type"`
	Text          string                          `json:"text"`
	FileCitations *OpenAiFileCitationsObject      `json:"file_citations,omitempty"`
	FilePath      *OpenAiFilePathObject           `json:"file_path,omitempty"`
	StartIndex    int                             `json:"start_index"`
	EndIndex      int                             `json:"end_index"`
}

func (obj *OpenAiAnnotationsObject) CheckValidObject() bool {
	switch obj.Type {
	case openAiAnnotation.FileCitation:
		return obj.FileCitations != nil && obj.FilePath == nil
	case openAiAnnotation.FilePath:
		return obj.FileCitations == nil && obj.FilePath != nil
	}
	return false
}

type OpenAiImageFileObject struct {
	FileID string `json:"file_id"`
}

type OpenAiTextObject struct {
	Value       string                     `json:"value"`
	Annotations []*OpenAiAnnotationsObject `json:"annotations"`
}

type OpenAiContent struct {
	Type      openAiContentType.ContentStatusType `json:"type"`
	ImageFile *OpenAiImageFileObject              `json:"image_file,omitempty"`
	Text      *OpenAiTextObject                   `json:"text,omitempty"`
}

type OpenAiMessagesObject struct {
	ID          string           `json:"id"`
	Object      string           `json:"object"`
	CreatedAt   int64            `json:"created_at"`
	ThreadID    string           `json:"thread_id"`
	Role        openAiRole.Role  `json:"role"`
	Content     []*OpenAiContent `json:"content"`
	FileIDs     []string         `json:"file_ids"`
	AssistantID *string          `json:"assistant_id"`
	RunID       *string          `json:"run_id"`
	Metadata    *OpenAiMetaData  `json:"metadata"`
}
