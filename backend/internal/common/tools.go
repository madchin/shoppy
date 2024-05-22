//go:build tools
// +build tools

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@master --config=../users/port/http/api.yaml ../users/port/http/oapi_spec.yaml
package tools

import (
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
)
