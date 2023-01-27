package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ConvertToTitleCase(s string) string {
	csr := cases.Title(language.Ukrainian)
	return csr.String(s)
}
