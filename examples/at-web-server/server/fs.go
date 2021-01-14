package server

import (
	"embed"
	"io/fs"
)

//go:embed static/css/*
//go:embed static/js/*
var staticFS embed.FS

func StaticFS() fs.FS {
	return staticFS
}

//go:embed templates
var templateFS embed.FS

func TemplateFS() fs.FS {
	return templateFS
}
