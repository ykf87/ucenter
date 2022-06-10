package i18n

import (
	"html/template"

	tr "github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

var I18n *tr.I18n

func init() {
	tr.Default = "en"
	I18n = tr.New(
		yaml.New("./i18n"), // load translations from the YAML files in directory `config/locales`
	)
}

func T(lang, key string, args ...interface{}) template.HTML {
	return I18n.T(lang, key, args...)
}
