package openAiLastError

type LastErrorCode string

const (
	ServerError       LastErrorCode = "server_error"
	RateLimitExceeded LastErrorCode = "rate_limit_exceeded"
)
