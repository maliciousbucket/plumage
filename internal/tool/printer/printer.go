package printer

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	infoHeader       = color.New(color.FgHiMagenta, color.Bold).PrintlnFunc()
	infoSubheader    = color.New(color.FgHiMagenta).SprintFunc()
	warningHeader    = color.New(color.FgYellow, color.Bold).PrintlnFunc()
	warningSubHeader = color.New(color.FgYellow).SprintFunc()
	successHeader    = color.New(color.FgGreen, color.Bold).PrintlnFunc()
	successSubheader = color.New(color.FgGreen).SprintFunc()
	errHeader        = color.New(color.FgRed, color.Bold).PrintlnFunc()
	errSubheader     = color.New(color.FgRed).SprintFunc()
)

type Printer struct {
}

func New() *Printer {
	return &Printer{}
}

func (p *Printer) PrintError(title string, err error) {
	errHeader("[Error]")
	fmt.Printf("%s: %s\n", errSubheader(title), err.Error())
}

func (p *Printer) PrintSuccess(title, message string, body interface{}) {
	successHeader("[Success]")
	if message != "" {
		fmt.Printf("%s: %s\n%v\n", successSubheader(title), message, body)
	} else {
		fmt.Printf("%s: %v\n", successSubheader(title), body)
	}

}

func (p *Printer) PrintSuccessWithFields(title, message string, body interface{}) {
	successHeader("[Success]")
	if message != "" {
		fmt.Printf("%s: %s\n%+v\n", successSubheader(title), message, body)
	} else {
		fmt.Printf("%s: %+v\n", successSubheader(title), body)
	}
}

func (p *Printer) PrintWarning(title, message string) {
	warningHeader("[Warning]")
	fmt.Printf("%s: %s\n", warningSubHeader(title), message)
}

func (p *Printer) PrintInfo(title, body string) {
	infoHeader("[Info]")
	fmt.Printf("%s: %s\n", infoSubheader(title), body)
}
