package deagon

import "strings"
import "golang.org/x/text/cases"
import "golang.org/x/text/language"

type Formatter interface {
	Format(firstname, surname string) string
}

type EmptyFormatter struct{}

func NewEmptyFormatter() *EmptyFormatter {
	return &EmptyFormatter{}
}

func (*EmptyFormatter) Format(_, _ string) string {
	return ""
}

type UppercaseSpaceFormatter struct{}

func NewUppercaseSpaceFormatter() *UppercaseSpaceFormatter {
	return &UppercaseSpaceFormatter{}
}

func (*UppercaseSpaceFormatter) Format(firstname, surname string) string {
	return firstname + " " + surname
}

type CapitalizedSpaceFormatter struct{}

func NewCapitalizedSpaceFormatter() *CapitalizedSpaceFormatter {
	return &CapitalizedSpaceFormatter{}
}

func (*CapitalizedSpaceFormatter) Format(firstname, surname string) string {
	caser := cases.Title(language.Dutch)
	return caser.String(strings.ToLower(firstname)) + " " + caser.String(strings.ToLower(surname))
}

type LowercaseDashFormatter struct{}

func NewLowercaseDashFormatter() *LowercaseDashFormatter {
	return &LowercaseDashFormatter{}
}

func (*LowercaseDashFormatter) Format(firstname, surname string) string {
	return strings.ToLower(firstname) + "-" + strings.ToLower(surname)
}
