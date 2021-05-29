package mila_view

import "github.com/aymerick/raymond"

type View struct {
	DirTemplate string
	Layout string
	Extension string
	TemplateMaster *raymond.Template
}