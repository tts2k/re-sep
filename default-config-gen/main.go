package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"iter"
	"log"
	"os"
	"reflect"
	"strconv"
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

const (
	Name  = 0
	Kind  = 1
	Value = 2
)

func printGoAst(pbRef protoreflect.Message) {
	importSpec := &ast.ImportSpec{
		Name: ast.NewIdent("pb"),
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote("path_to_pb"),
		},
	}

	defConfVal := &ast.ValueSpec{
		Names: []*ast.Ident{ast.NewIdent("DefaultUserConfig")},
		Values: []ast.Expr{
			&ast.CompositeLit{
				Lbrace: 8,
				Type: &ast.SelectorExpr{
					X:   ast.NewIdent("pb"),
					Sel: ast.NewIdent("UserConfig"),
				},
			},
		},
	}

	astF := ast.File{
		Name:    ast.NewIdent("something"),
		Imports: []*ast.ImportSpec{importSpec},
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok:    token.IMPORT,
				Lparen: token.NoPos,
				Rparen: token.NoPos,
				Specs:  []ast.Spec{importSpec},
			},
			&ast.GenDecl{
				Tok:   token.VAR,
				Specs: []ast.Spec{defConfVal},
			},
		},
	}

	fset := token.NewFileSet()

	printGoFieldsAst(pbRef, &defConfVal.Values[0].(*ast.CompositeLit).Elts)

	pConfig := printer.Config{
		Mode:     printer.TabIndent,
		Tabwidth: 4,
	}
	err := pConfig.Fprint(os.Stdout, fset, &astF)
	if err != nil {
		log.Fatal(err)
	}
}

func printGoFieldsAst(pbRef protoreflect.Message, elts *[]ast.Expr) {
	for field, message := range PbRefIter(pbRef) {
		switch field[Kind] {
		case "string":
			*elts = append(*elts, &ast.KeyValueExpr{
				Key: ast.NewIdent(field[Name]),
				Value: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(field[Value]),
				},
			})
		case "int32":
			*elts = append(*elts, &ast.KeyValueExpr{
				Key: ast.NewIdent(field[Name]),
				Value: &ast.BasicLit{
					Kind:  token.INT,
					Value: field[Value],
				},
			})
		case "bool":
			*elts = append(*elts, &ast.KeyValueExpr{
				Key:   ast.NewIdent(field[Name]),
				Value: ast.NewIdent(field[Value]),
			})
		default:
			expr := &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.CompositeLit{
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("pb"),
						Sel: ast.NewIdent(strings.Split(field[Kind], ".")[1]),
					},
				},
			}
			printGoFieldsAst(message, &expr.X.(*ast.CompositeLit).Elts)
			*elts = append(*elts, &ast.KeyValueExpr{
				Key:   ast.NewIdent(field[Name]),
				Value: expr,
			})
		}
	}
}

// For playing around with the ast
func buildGoAst() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(
		fset, "../server/user/internal/database/user/database.go",
		nil,
		0,
	)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok != token.VAR {
				return true
			}
			for _, spec := range x.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					if s.Names[0].Name == "DefaultUserConfig" {
						fmt.Println("---")
						fmt.Println(s.Names)
						for _, expr := range s.Values[0].(*ast.CompositeLit).Elts {
							e := expr.(*ast.KeyValueExpr)
							switch val := e.Value.(type) {
							case *ast.UnaryExpr:
								fmt.Println(val.Op)
							default:
								fmt.Println(reflect.TypeOf(e.Value))
								// fmt.Println(e.Value)
							}
						}
					}
				}
			}
		}

		return true
	})
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

	// buildGoAst()
	printGoAst(ref)
}
