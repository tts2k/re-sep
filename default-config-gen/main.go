package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/format"
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

func printGoFields(tokens *[]Token, builder *strings.Builder) {
	for _, token := range *tokens {
		switch token.Kind {
		case "string":
			fmt.Fprintf(builder, "%s: \"%s\",\n", token.Name, token.Value)
		case "int32":
			fallthrough
		case "bool":
			fmt.Fprintf(builder, "%s: %s,\n", token.Name, token.Value)
		default:
			fmt.Fprintf(builder, "%s: &%s{\n", token.Name, strings.Replace(token.Kind, pkgName, "pb", 1))
			printGoFields(token.Value.(*[]Token), builder)
			fmt.Fprintln(builder, "},")
		}
	}
}

func printGo(tokens *[]Token) {
	builder := &strings.Builder{}
	fmt.Fprintln(builder, "var DefaultConfig = &pb.UserConfig{")
	printGoFields(tokens, builder)
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

func TokenizePb(ref protoreflect.Message, tokens *[]Token) {
	fields := ref.Descriptor().Fields()

	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		v := ref.Get(fd)

		if fd.Kind().String() == "message" {
			newTokens := []Token{}
			TokenizePb(v.Message(), &newTokens)

			*tokens = append(*tokens, Token{
				Name:  string(fd.Name()),
				Kind:  string(fd.Message().FullName()),
				Value: &newTokens,
			})
			continue
		}

		*tokens = append(*tokens, Token{
			Name:  string(fd.Name()),
			Kind:  fd.Kind().String(),
			Value: v.String(),
		})
	}
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
	tokens := []Token{}

	TokenizePb(ref, &tokens)

	printGo(&tokens)
}
