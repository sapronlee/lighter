package helpers

import (
	"html/template"
)

func AssetHelpers() template.FuncMap {
	return template.FuncMap{
		"asset_path": func(fileName string) string {
			return "/assets/" + fileName
		},
	}
}
