package middlewares

type ExceptionResult struct {
	HttpStatus int
	Data       map[string]interface{}
}
