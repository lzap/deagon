package deagon

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Formatter defines a common interface for formatting names.
type Formatter interface {
	Format(firstname, surname string) string
}

// EmptyFormatter always returns an empty string.
type EmptyFormatter struct{}

// NewEmptyFormatter creates a new EmptyFormatter.
func NewEmptyFormatter() *EmptyFormatter {
	return &EmptyFormatter{}
}

// Format ignores inputs and returns an empty string.
func (*EmptyFormatter) Format(_, _ string) string {
	return ""
}

// UppercaseSpaceFormatter joins first and last name in uppercase with a space.
type UppercaseSpaceFormatter struct{}

// NewUppercaseSpaceFormatter creates a new UppercaseSpaceFormatter.
func NewUppercaseSpaceFormatter() *UppercaseSpaceFormatter {
	return &UppercaseSpaceFormatter{}
}

// Format returns the names in uppercase separated by a space.
func (*UppercaseSpaceFormatter) Format(firstname, surname string) string {
	return strings.ToUpper(firstname) + " " + strings.ToUpper(surname)
}

// CapitalizedSpaceFormatter joins first and last name in capitalized form with a space.
type CapitalizedSpaceFormatter struct {
	caser cases.Caser
}

// NewCapitalizedSpaceFormatter creates a new CapitalizedSpaceFormatter.
// Uses English capitalization rules.
func NewCapitalizedSpaceFormatter() *CapitalizedSpaceFormatter {
	return &CapitalizedSpaceFormatter{
		caser: cases.Title(language.English),
	}
}

// Format returns the names with each word capitalized.
func (f *CapitalizedSpaceFormatter) Format(firstname, surname string) string {
	return f.caser.String(strings.ToLower(firstname)) + " " +
		f.caser.String(strings.ToLower(surname))
}

// LowercaseDashFormatter joins first and last name in lowercase with a dash.
type LowercaseDashFormatter struct{}

// NewLowercaseDashFormatter creates a new LowercaseDashFormatter.
func NewLowercaseDashFormatter() *LowercaseDashFormatter {
	return &LowercaseDashFormatter{}
}

// Format returns the names in lowercase separated by a dash.
func (*LowercaseDashFormatter) Format(firstname, surname string) string {
	return strings.ToLower(firstname) + "-" + strings.ToLower(surname)
}
