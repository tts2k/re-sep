package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
	"iter"
	"os"
	"strings"
	"text/template"

	"google.golang.org/protobuf/reflect/protoreflect"

	pb "re-sep-config-gen/proto"
)

type Token = struct {
	Value interface{}
	Name  string
	Kind  string
}

//go:embed go.tmpl
var f embed.FS
var pkgName string

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

const (
	Name  = 0
	Kind  = 1
	Value = 2
)

func printGoFields(pbRef protoreflect.Message, builder *strings.Builder) {
	for token, message := range PbRefIter(pbRef) {
		switch token[Kind] {
		case "string":
			fmt.Fprintf(builder, "%s: \"%s\",\n", token[Name], token[Value])
		case "int32":
			fallthrough
		case "bool":
			fmt.Fprintf(builder, "%s: %s,\n", token[Name], token[Value])
		default:
			fmt.Fprintf(builder, "%s: &%s{\n", token[Name], strings.Replace(token[Kind], pkgName, "pb", 1))
			printGoFields(message, builder)
			fmt.Fprintln(builder, "},")
		}
	}
}

func printGo(pbRef protoreflect.Message) {
	builder := &strings.Builder{}
	fmt.Fprintln(builder, "var DefaultConfig = &pb.UserConfig{")
	printGoFields(pbRef, builder)
	fmt.Fprintln(builder, "}")

	tmpl, err := template.ParseFS(f, "go.tmpl")
	if err != nil {
		panic(err)
	}

	type GoTemplate = struct {
		Package       string
		ProtoPackage  string
		DefaultConfig string
	}

	buf := bytes.NewBufferString("")
	// Use hardcoded values for now
	err = tmpl.Execute(buf, GoTemplate{
		Package:       "def",
		ProtoPackage:  "protoPackage",
		DefaultConfig: builder.String(),
	})
	if err != nil {
		panic(err)
	}

	res, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}

func main() {
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
	pkgName = string(ref.Descriptor().ParentFile().Package())

	printGo(ref)
}
