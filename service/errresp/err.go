package errresp

type Error struct {
	Error string `json:"error"`
}

// ErrorResp All use wrapping packing will use this
// #wrapping:"err"#
func ErrorResp(rawErr error) (e *Error /*#code:"400"#*/, err error) {
	return &Error{rawErr.Error()}, nil
}
