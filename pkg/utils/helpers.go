package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var ExtFile string
var Route string

func LoadTemplateDir(pattern string) (*template.Template, error) {
	tmpl := template.New("")

	err := filepath.WalkDir(pattern, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == "."+ExtFile {
			//log.Println("Loading template:", path)
			if _, err := tmpl.ParseFiles(path); err != nil {
				return err
			}
		}

		return nil
	})

	return tmpl, err
}

func TemplateRender(context *gin.Context, httpStatus int, contentTemplate string, data gin.H, isAdminTemplate ...bool) {
	dirTmpl := "userLayout"

	if isAdminTemplate != nil {
		dirTmpl = "adminLayout"
	}

	data["templates"] = contentTemplate
	//data["content"] = renderHtmlTemplate(contentTemplate)
	context.HTML(httpStatus, dirTmpl, data)
}

func renderHtmlTemplate(filePath string) template.HTML {
	content, err := os.ReadFile(fmt.Sprintf("web/templates/%s/%s.%s", Route, filePath, ExtFile))

	if err != nil {
		log.Fatalf("Render template fails", err)
	}

	return template.HTML(string(content))
}

func ConvertFieldName(field string) string {
	var result []string
	for index, val := range field {
		if index > 0 && unicode.IsUpper(val) {
			result = append(result, " ")
		}

		result = append(result, strings.ToLower(string(val)))
	}

	return strings.Join(result, "")
}
