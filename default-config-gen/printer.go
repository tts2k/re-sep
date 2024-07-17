package main

import (
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	Name  = 0
	Kind  = 1
	Value = 2
)

type Printer = interface {
	Print(w io.Writer, pbRef protoreflect.Message, config map[string]string) error
}

type PrinterFactory struct {
	Type string
}

func (pf PrinterFactory) CreatePrinter() (Printer, error) {
	if pf.Type == "go" {
		return &GoPrinter{}, nil

	} else {
		panic("unsupported type")
	}
}

type GoPrinter struct{}

func newGoPrinter() *GoPrinter {
	return &GoPrinter{}
}

func (gp *GoPrinter) Print(w io.Writer, pbRef protoreflect.Message, config map[string]string) error {
	fmt.Fprintf(w, "package %s\n\n", config["go_pkgname"])
	fmt.Fprintf(w, "import pb %s\n\n", config["go_pb_pkgname"])

	fmt.Fprintln(w, "var DefaultConfig = &pb.UserConfig{")
	gp.printGoFields(w, pbRef, config, 1)
	fmt.Fprintln(w, "}")

	return nil
}

func (gp *GoPrinter) printGoFields(w io.Writer, pbRef protoreflect.Message, config map[string]string, indent int) {
	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += "\t"
	}

	for token, message := range PbRefIter(pbRef) {
		switch token[Kind] {
		case "string":
			fmt.Fprintf(w, "%s%s: \"%s\",\n", indentString, token[Name], token[Value])
		case "int32":
			fallthrough
		case "bool":
			fmt.Fprintf(w, "%s%s: %s,\n", indentString, token[Name], token[Value])
		default:
			fmt.Fprintf(w, "%s%s: &%s{\n", indentString, token[Name], strings.Replace(token[Kind], config["pb_pkgname"], "pb", 1))
			gp.printGoFields(w, message, config, indent+1)
			fmt.Fprintf(w, "%s},\n", indentString)
		}
	}
}
