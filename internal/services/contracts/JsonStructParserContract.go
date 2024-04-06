package contracts

import "generatorStructure/internal/services"

type JsonStructParserContract interface {
	Read() *services.JsonStructParser
	Result() (bool, error)
}
