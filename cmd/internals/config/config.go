package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// App config hold the config at application level
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
