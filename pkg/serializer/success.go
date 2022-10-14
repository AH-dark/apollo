package serializer

func NewSuccessResponse(data interface{}) *Response {
	return NewResponse(0, "success", data)
}
