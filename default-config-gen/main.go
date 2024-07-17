package main

import (
	"encoding/json"
	"fmt"
	"iter"
	"os"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb "re-sep-config-gen/proto"
)

type Token = struct {
	Value interface{}
	Name  string
	Kind  string
}

func PbRefIter(pbRef protoreflect.Message) iter.Seq2[[3]string, protoreflect.Message] {
	fields := pbRef.Descriptor().Fields()

	return func(yield func([3]string, protoreflect.Message) bool) {
		for i := 0; i < fields.Len(); i++ {
			fd := fields.Get(i)
			v := pbRef.Get(fd)
			name := string(fd.Name())

			var kind string
			var message protoreflect.Message
			if fd.Kind().String() == "message" {
				kind = string(fd.Message().FullName())
				message = v.Message()
			} else {
				kind = string(fd.Kind().String())
				message = nil
			}

			if !yield([3]string{name, kind, v.String()}, message) {
				return
			}
		}
	}
}

func initFlags(pf *PrinterFactory) error {
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"Default config gen of re-sep\n\n"+
				"Usage:\n"+
				"  re-sep-cli [flags] <url>\n\n"+
				"Flags:",
		)
		pflag.PrintDefaults()
	}

	pflag.BoolP("help", "h", false, "Print this help message")
	pflag.StringVarP(&pf.Type, "type", "t", "go", "File output type")
	pflag.Parse()

	return nil
}

func main() {
	pf := PrinterFactory{}

	err := initFlags(&pf)

	fi, err := os.Open("./default_config.json")
	if err != nil {
		panic(err)
	}

	var defaultConfig pb.UserConfig
	err = json.NewDecoder(fi).Decode(&defaultConfig)
	if err != nil {
		panic(err)
	}
	fi.Close()

	ref := defaultConfig.ProtoReflect()

	pkgName := string(ref.Descriptor().ParentFile().Package())
	printer, err := pf.CreatePrinter()
	if err != nil {
		panic(err)
	}

	cfgWriter := NewConfigWriter(printer)
	err = cfgWriter.
		Writer(os.Stdout).
		ProtoReflect(ref).
		Config("pb_pkgname", pkgName).
		Config("go_pkgname", "def").
		Config("go_pb_pkgname", "protoPackage").
		Write()

	if err != nil {
		panic(err)
	}
}
