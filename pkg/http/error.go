package http

type HttpClientError struct {
	Status int
	Err error
	Body string
}

func (e *HttpClientError) Error() string {
	return e.Err.Error()
}