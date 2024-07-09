package utils

import (
	"github.com/valyala/fasttemplate"
	pb "re-sep-content/internal/proto"
)

const (
	H1   = 0
	H2   = 1
	H3   = 2
	H4   = 3
	TEXT = 4
)

var fontSizePresets [][]string = [][]string{
	{"text-5xl", "text-2xl", "text-xl", "text-lg", "text-base"},
	{"text-6xl", "text-3xl", "text-2xl", "text-xl", "text-lg"},
	{"text-7xl", "text-4xl", "text-3xl", "text-2xl", "text-xl"},
	{"text-8xl", "text-5xl", "text-4xl", "text-3xl", "text-2xl"},
	{"text-9xl", "text-6xl", "text-5xl", "text-4xl", "text-3xl"},
}

func ApplyTemplate(html string, userConfig *pb.UserConfig) string {
	t := fasttemplate.New(html, "{{", "}}")

	m := make(map[string]interface{})

	m["h1"] = fontSizePresets[H1][userConfig.FontSize]
	m["h2"] = fontSizePresets[H2][userConfig.FontSize]
	m["h3"] = fontSizePresets[H3][userConfig.FontSize]
	m["h4"] = fontSizePresets[H4][userConfig.FontSize]
	m["text"] = fontSizePresets[TEXT][userConfig.FontSize]

	s := t.ExecuteString(m)

	return s
}
