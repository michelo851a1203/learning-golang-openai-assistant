package openAiTool

type Tool string

const (
	CodeInterpreter Tool = "code_interpreter"
	Retrieval       Tool = "retrieval"
	FunctionCalling Tool = "function"
)
