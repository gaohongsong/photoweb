package views

import (
	"io/ioutil"
	"html/template"
	"path"
	"log"

	. "github.com/gmaclinuxer/photoweb/common"
	"net/http"
)

type Templates map[string]*template.Template
type Context map[string]interface{}

const TemplateDir = "./templates"

var T = make(Templates)

func init() {

	fileArr, err := ioutil.ReadDir(TemplateDir)
	CheckError(err)

	var templateName, templatePath string

	for _, f := range fileArr {

		templateName = f.Name()

		if ext := path.Ext(templateName); ext != ".html" {
			log.Printf("ignore files: %s\n", templateName)
			continue
		}

		templatePath = path.Join(TemplateDir, templateName)
		log.Printf("loading template: %s", templatePath)

		T[templatePath] = template.Must(template.ParseFiles(templatePath))
	}

}

func RenderTemplate(w http.ResponseWriter, tpl string, ctx Context) {
	tplPath := path.Join(TemplateDir, tpl+".html")
	err := T[tplPath].Execute(w, ctx)
	CheckError(err)
}
