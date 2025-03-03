package contracts

import "github.com/Pashgunt/Strucutre-generator/internal/services"

type JsonStructParserContract interface {
	Read() *services.JsonStructParser
	Result() (bool, error)
}
