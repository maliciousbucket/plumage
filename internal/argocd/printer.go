package argocd

import "github.com/maliciousbucket/plumage/internal/tool/printer"

type Printer interface {
	PrintError(title string, err error)
	PrintSuccess(title, message string, body interface{})
	PrintSuccessWithFields(title, message string, body interface{})
	PrintWarning(title, message string)
	PrintInfo(title, body string)
}

func NewPrinter() Printer {
	return printer.New()
}
