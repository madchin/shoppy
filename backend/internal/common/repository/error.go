package repository

import (
	"errors"
	"fmt"
)

var ErrInternal = errors.New("internal error")

type ErrMissingEnv struct {
	Keys []string
}

func (e *ErrMissingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e *ErrMissingEnv) Error() string {
	var missCount int
	var missingEnvs string
	for _, env := range e.Keys {
		missCount++
		missingEnvs += fmt.Sprintf("%s, ", env)
	}
	return fmt.Sprintf("%d count of envs are missing: %s", missCount, missingEnvs)
}
