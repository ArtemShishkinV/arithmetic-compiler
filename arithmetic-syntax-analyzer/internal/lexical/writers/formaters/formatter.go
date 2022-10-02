package formaters

import "arithmetic-syntax-analyzer/internal/lexical/models"

type Formatter interface {
	Form([]models.Lexeme) []string
}
