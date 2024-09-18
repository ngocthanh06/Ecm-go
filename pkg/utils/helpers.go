package utils

import (
	"crypto/rand"
	"encoding/hex"
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

func RenderHtmlTemplate(filePath string, dir string) template.HTML {

	content, err := os.ReadFile(fmt.Sprintf("web/templates/%s/%s.%s", dir, filePath, ExtFile))

	if err != nil {
		log.Fatalf("Render template fails", err)
	}

	return template.HTML(string(content))
}

func RenderHtmlTemplateMail(filepath string) template.HTML {
	return RenderHtmlTemplate(filepath, "mail")
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

func CreateRandToken(size int) (string, error) {
	bytes := make([]byte, size)

	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("hash pass error")

		return "", err
	}

	token := hex.EncodeToString(bytes)

	return token, nil
}
