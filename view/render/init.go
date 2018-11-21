package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rsniezynski/go-asset-helper"
	"github.com/tribes/logging"
)

var renderApp *template.Template
var cfg Config

type Config struct {
	Render RenderConfig
}

type RenderConfig struct {
	TemplatePath string
	ImagePath    string
}

func init() {
	ok := logging.ReadModuleConfig(&cfg, "config/render", "render")
	if !ok {
		log.Fatal("Failed to read render config")
	}
	static, err := asset.NewStatic("/static/", "view/asset/manifest.json")
	if err != nil {
		// Manifest file doesn't exist or is not a valid JSON
	}

	funcMap := static.FuncMap()
	templatePath := fmt.Sprintf("%s*.html", cfg.Render.TemplatePath)
	renderApp = template.Must(template.New("").Funcs(funcMap).ParseGlob(templatePath))
}

func GetImagePath() string {
	return cfg.Render.ImagePath
}

func RenderPage(w http.ResponseWriter, renderName string, context map[string]interface{}) {
	context["include"] = renderName
	context["ImageDir"] = GetImagePath()
	renderName = renderName + ".html"
	err := renderApp.ExecuteTemplate(w, renderName, context)
	if err != nil {
		fmt.Println(err)
	}
}
