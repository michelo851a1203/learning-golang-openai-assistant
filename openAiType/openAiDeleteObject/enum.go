package openaiDeleteObject

type ObjectStatus string

const (
	Assistant     ObjectStatus = "assistant.deleted"
	AssistantFile ObjectStatus = "assistant.file.deleted"
	File          ObjectStatus = "file"
	Threads       ObjectStatus = "thread.deleted"
)
