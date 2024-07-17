package main

import (
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type ConfigWriter struct {
	writer  io.Writer
	printer Printer
	pbRef   protoreflect.Message
	config  map[string]string
}

func NewConfigWriter(printer Printer) *ConfigWriter {
	return &ConfigWriter{
		printer: printer,
		config:  make(map[string]string),
	}
}

func (cw *ConfigWriter) Writer(w io.Writer) *ConfigWriter {
	cw.writer = w
	return cw
}

func (cw *ConfigWriter) Config(k, v string) *ConfigWriter {
	cw.config[k] = v
	return cw
}

func (cw *ConfigWriter) ProtoReflect(pbRef protoreflect.Message) *ConfigWriter {
	cw.pbRef = pbRef
	return cw
}

func (cw *ConfigWriter) Write() error {
	return cw.printer.Print(cw.writer, cw.pbRef, cw.config)
}
