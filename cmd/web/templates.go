package main

import (
	"github.com/zeidlitz/snippetbox/internal/models"
)

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}
