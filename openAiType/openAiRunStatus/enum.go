package openAiRunStatus

type RunStatus string

const (
	Queued         RunStatus = "queued"
	InProgress     RunStatus = "in_progress"
	RequiresAction RunStatus = "requires_action"
	Cancelling     RunStatus = "cancelling"
	Cancelled      RunStatus = "cancelled"
	Failed         RunStatus = "failed"
	Completed      RunStatus = "completed"
	Expired        RunStatus = "expired"
)
