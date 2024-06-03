//go:build go1.22

// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.1-0.20240601101045-cb61b77eea50 DO NOT EDIT.
package ports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error Basic error
type Error struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

// User Retrieved / Created user
type User struct {
	// Email User email
	Email string `json:"email"`

	// Name User name
	Name string `json:"name"`

	// Password user password
	Password *string `json:"password,omitempty"`
}

// UserDetail Retrieved / created user detail
type UserDetail struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// ErrInternal Basic error
type ErrInternal = Error

// ErrUnauthorized Basic error
type ErrUnauthorized = Error

// ErrUserNotFound Basic error
type ErrUserNotFound = Error

// PutUserDetailUpdateFirstNameParams defines parameters for PutUserDetailUpdateFirstName.
type PutUserDetailUpdateFirstNameParams struct {
	// FirstName user first name used to update
	FirstName string `form:"firstName" json:"firstName"`
}

// PutUserDetailUpdateLastNameParams defines parameters for PutUserDetailUpdateLastName.
type PutUserDetailUpdateLastNameParams struct {
	// LastName user last name used to update
	LastName string `form:"lastName" json:"lastName"`
}

// PutUserUpdateEmailParams defines parameters for PutUserUpdateEmail.
type PutUserUpdateEmailParams struct {
	// Email email used to update
	Email string `form:"email" json:"email"`
}

// PutUserUpdateNameParams defines parameters for PutUserUpdateName.
type PutUserUpdateNameParams struct {
	// Name name used for update
	Name string `form:"name" json:"name"`
}

// PutUserUpdatePasswordParams defines parameters for PutUserUpdatePassword.
type PutUserUpdatePasswordParams struct {
	// Password password used to update
	Password string `form:"password" json:"password"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = User

// PostUserDetailJSONRequestBody defines body for PostUserDetail for application/json ContentType.
type PostUserDetailJSONRequestBody = UserDetail

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /user)
	DeleteUser(w http.ResponseWriter, r *http.Request)

	// (GET /user)
	GetUser(w http.ResponseWriter, r *http.Request)

	// (POST /user)
	PostUser(w http.ResponseWriter, r *http.Request)

	// (DELETE /user/detail)
	DeleteUserDetail(w http.ResponseWriter, r *http.Request)

	// (GET /user/detail)
	GetUserDetail(w http.ResponseWriter, r *http.Request)

	// (POST /user/detail)
	PostUserDetail(w http.ResponseWriter, r *http.Request)

	// (PUT /user/detail/update-first-name)
	PutUserDetailUpdateFirstName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateFirstNameParams)

	// (PUT /user/detail/update-last-name)
	PutUserDetailUpdateLastName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateLastNameParams)

	// (PUT /user/update-email)
	PutUserUpdateEmail(w http.ResponseWriter, r *http.Request, params PutUserUpdateEmailParams)

	// (PUT /user/update-name)
	PutUserUpdateName(w http.ResponseWriter, r *http.Request, params PutUserUpdateNameParams)

	// (PUT /user/update-password)
	PutUserUpdatePassword(w http.ResponseWriter, r *http.Request, params PutUserUpdatePasswordParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUser operation middleware
func (siw *ServerInterfaceWrapper) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUserDetail operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUserDetail(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserDetail operation middleware
func (siw *ServerInterfaceWrapper) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserDetail(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUserDetail operation middleware
func (siw *ServerInterfaceWrapper) PostUserDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserDetail(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUserDetailUpdateFirstName operation middleware
func (siw *ServerInterfaceWrapper) PutUserDetailUpdateFirstName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PutUserDetailUpdateFirstNameParams

	// ------------- Required query parameter "firstName" -------------

	if paramValue := r.URL.Query().Get("firstName"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "firstName"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "firstName", r.URL.Query(), &params.FirstName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "firstName", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUserDetailUpdateFirstName(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUserDetailUpdateLastName operation middleware
func (siw *ServerInterfaceWrapper) PutUserDetailUpdateLastName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PutUserDetailUpdateLastNameParams

	// ------------- Required query parameter "lastName" -------------

	if paramValue := r.URL.Query().Get("lastName"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "lastName"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "lastName", r.URL.Query(), &params.LastName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "lastName", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUserDetailUpdateLastName(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUserUpdateEmail operation middleware
func (siw *ServerInterfaceWrapper) PutUserUpdateEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PutUserUpdateEmailParams

	// ------------- Required query parameter "email" -------------

	if paramValue := r.URL.Query().Get("email"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "email"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "email", r.URL.Query(), &params.Email)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "email", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUserUpdateEmail(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUserUpdateName operation middleware
func (siw *ServerInterfaceWrapper) PutUserUpdateName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PutUserUpdateNameParams

	// ------------- Required query parameter "name" -------------

	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "name"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "name", r.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUserUpdateName(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUserUpdatePassword operation middleware
func (siw *ServerInterfaceWrapper) PutUserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params PutUserUpdatePasswordParams

	// ------------- Required query parameter "password" -------------

	if paramValue := r.URL.Query().Get("password"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "password"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "password", r.URL.Query(), &params.Password)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "password", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUserUpdatePassword(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("DELETE "+options.BaseURL+"/user", wrapper.DeleteUser)
	m.HandleFunc("GET "+options.BaseURL+"/user", wrapper.GetUser)
	m.HandleFunc("POST "+options.BaseURL+"/user", wrapper.PostUser)
	m.HandleFunc("DELETE "+options.BaseURL+"/user/detail", wrapper.DeleteUserDetail)
	m.HandleFunc("GET "+options.BaseURL+"/user/detail", wrapper.GetUserDetail)
	m.HandleFunc("POST "+options.BaseURL+"/user/detail", wrapper.PostUserDetail)
	m.HandleFunc("PUT "+options.BaseURL+"/user/detail/update-first-name", wrapper.PutUserDetailUpdateFirstName)
	m.HandleFunc("PUT "+options.BaseURL+"/user/detail/update-last-name", wrapper.PutUserDetailUpdateLastName)
	m.HandleFunc("PUT "+options.BaseURL+"/user/update-email", wrapper.PutUserUpdateEmail)
	m.HandleFunc("PUT "+options.BaseURL+"/user/update-name", wrapper.PutUserUpdateName)
	m.HandleFunc("PUT "+options.BaseURL+"/user/update-password", wrapper.PutUserUpdatePassword)

	return m
}
