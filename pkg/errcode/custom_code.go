package errcode

import "net/http"

var (
	Success                   = NewError(0, "Success", http.StatusOK)
	ServerError               = NewError(100000, "Server error", http.StatusInternalServerError)
	InvalidParams             = NewError(100001, "Invalid params", http.StatusBadRequest)
	NotFound                  = NewError(100002, "Not found", http.StatusNotFound)
	UnauthorizedAuthNotExist  = NewError(100003, "Auth failed, not exist", http.StatusUnauthorized)
	UnauthorizedTokenError    = NewError(100004, "Auth failed, token error", http.StatusUnauthorized)
	UnauthorizedTokenTimeout  = NewError(100005, "Auth failed, token timeout", http.StatusUnauthorized)
	UnauthorizedTokenGenerate = NewError(100006, "Auth failed, token generate", http.StatusUnauthorized)
	TooManyRequests           = NewError(100007, "Too many requests", http.StatusTooManyRequests)
	NotImplemented            = NewError(100010, "Not Implemented Yet.", http.StatusNotImplemented)
	InvalidRequestBody        = NewError(100011, "Invalid request body", http.StatusBadRequest)
)

var (
	ErrorGetTaskListFail = NewError(20010001, "Get Task list fail", http.StatusInternalServerError)
	ErrorCreateTaskFail  = NewError(20010002, "Create Task fail", http.StatusInternalServerError)
	ErrorUpdateTaskFail  = NewError(20010003, "Update Task fail", http.StatusInternalServerError)
	ErrorDeleteTaskFail  = NewError(20010004, "Delete Task fail", http.StatusInternalServerError)
	ErrorCountTaskFail   = NewError(20010005, "Count Task fail", http.StatusInternalServerError)
	ErrorGetTaskFail     = NewError(20010006, "Get Task list fail", http.StatusNotFound)
)
