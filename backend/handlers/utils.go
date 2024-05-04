package handler

import "net/http"

type UserHandlerFunc func(userService userService, uuid string, w http.ResponseWriter, r *http.Request)
