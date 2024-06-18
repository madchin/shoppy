package server

import (
	custom_error "backend/internal/common/errors"
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Middleware func(http.Handler) http.Handler

const BearerAuthScopes = "bearerAuth.Scopes"

type Auth interface {
	Verify(token string) (UserInfo, error)
	Sign(UserInfo) (string, error)
}

type key int

var userKey key

type UserInfo struct {
	Uuid string
}

func UserInfoFromContext(ctx context.Context) (UserInfo, bool) {
	userInfo, ok := ctx.Value(userKey).(UserInfo)
	return userInfo, ok
}

func NewContext(ctx context.Context, userInfo UserInfo) context.Context {
	return context.WithValue(ctx, userKey, userInfo)
}

func NewUserInfo(uuid string) UserInfo {
	return UserInfo{uuid}
}

func AuthMiddleware(auth Auth) Middleware {
	return func(next http.Handler) http.Handler {
		return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if middleware is not required, omit this middleware
			if r.Context().Value(BearerAuthScopes) == nil {
				next.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				ErrorHandler(w, r, custom_error.NewAuthorizationError("Unauthorized", "missing 'Authorization' header"))
				return
			}
			authHeaderContent := strings.Split(authHeader, " ")
			if len(authHeaderContent) != 2 || authHeaderContent[0] != "Bearer" {
				ErrorHandler(w, r, custom_error.NewAuthorizationError("Unauthorized", "wrong 'Authorization' header format"))
				return
			}
			token := authHeaderContent[1]
			userInfo, err := auth.Verify(token)
			if err != nil {
				ErrorHandler(w, r, custom_error.NewAuthorizationError("Unauthorized", fmt.Sprintf("authorization token is wrong or expired %s", err.Error())))
				return
			}
			ctx := NewContext(r.Context(), userInfo)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}))
	}
}
