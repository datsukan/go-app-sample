package usecase

type ResultStatus struct {
	Error      error
	StatusCode int
}

func NewResultStatus(statusCode int, err error) *ResultStatus {
	return &ResultStatus{
		Error:      err,
		StatusCode: statusCode,
	}
}
