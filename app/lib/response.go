package lib

type SuccessResponse struct {
	Code    int
	Message string
	Info    any
}

type ErrorResponse struct {
	Code    int
	Message string
}
