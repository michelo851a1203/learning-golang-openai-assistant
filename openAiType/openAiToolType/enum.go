package openAiToolType

type OpenAiToolType string

const (
	CodeInterpreter OpenAiToolType = "code_interpreter"
	Retrieval       OpenAiToolType = "retrieval"
	FunctionCalling OpenAiToolType = "function"
)
