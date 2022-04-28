package deagon

import "strings"

type Formatter interface {
	Format(firstname, surname string) string
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
	return strings.Title(strings.ToLower(firstname)) + " " + strings.Title(strings.ToLower(surname))
}

type LowercaseDashFormatter struct{}

func NewLowercaseDashFormatter() *LowercaseDashFormatter {
	return &LowercaseDashFormatter{}
}

func (*LowercaseDashFormatter) Format(firstname, surname string) string {
	return strings.ToLower(firstname) + "-" + strings.ToLower(surname)
}
