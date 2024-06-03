package common_adapter

import "strings"

func IsDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "duplicate")
}
