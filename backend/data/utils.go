package data

import "fmt"

type DbError struct {
	EmptyEmail *emptyEmail
	MissingEnv *missingEnv
}

type emptyEmail struct{}
type missingEnv struct {
	Keys []string
}

var Err = &DbError{
	EmptyEmail: &emptyEmail{},
	MissingEnv: &missingEnv{},
}

func (e *missingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e *emptyEmail) Error() string {
	return "User email is empty"
}

func (e *missingEnv) Error() string {
	var missCount int
	var missingEnvs string
	for _, env := range e.Keys {
		missCount++
		missingEnvs += fmt.Sprintf("%s, ", env)
	}
	return fmt.Sprintf("%d count of envs are missing: %s", missCount, missingEnvs)
}
