package standart

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"strconv"
)

func GetStringValue(param *string, defaultValue *string) string {
	if param != nil && param != defaultValue {
		return *param
	}
	return ""
}

func GetUint64Value(param *uint64, defaultValue *uint64) string {
	if param != nil && param != defaultValue {
		return strconv.FormatUint(*param, 10)
	}
	return ""
}

func GetHtmlFromMarkdown(markdownContent string) string {
	// Convert Markdown to HTML
	htmlContent := markdown.ToHTML([]byte(markdownContent), nil, html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags,
	}))
	return string(htmlContent)
}
