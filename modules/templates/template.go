package templates

import (
	"fmt"
	"html/template"
	"time"
)

var FuncMap template.FuncMap

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func init() {
	FuncMap = template.FuncMap{
		"formatAsDate": formatAsDate,
	}
}
