package config

import (
	"html/template"
	"log"
)

// AppConfig the aplication config global
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
