package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpErr struct {
	code int
	msg  string
}

var HttpErrInternal = httpErr{code: http.StatusInternalServerError, msg: "Oops! Something went wrong"}
var HttpErrWrongBody = httpErr{code: http.StatusBadRequest, msg: "Request body"}

func ErrorResponse(w http.ResponseWriter, httpErr httpErr) {
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(httpErr.toJsonStruct())
	if err != nil {
		errorJsonResponse(w)
		return
	}
	w.WriteHeader(httpErr.code)
	w.Write(msg)
}

func errorJsonResponse(w http.ResponseWriter) {
	w.WriteHeader(HttpErrInternal.code)
	w.Write([]byte(fmt.Sprintf("{code: %d, msg: %s}", HttpErrInternal.code, HttpErrInternal.msg)))
}

func (h httpErr) toJsonStruct() struct {
	Code int
	Msg  string
} {
	return struct {
		Code int
		Msg  string
	}{h.code, h.msg}
}
