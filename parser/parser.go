package parser

import (
	"github.com/rocket-internet-berlin/RocketLabsRubberDoc/definition"
	"github.com/rocket-internet-berlin/RocketLabsRubberDoc/parser/transformer"
)

//Parser Interface definition
type Parser interface {
	Parse(filename string, tra transformer.Transformer) (def *definition.Api, err error)
}
