package config

const (
	TypeStructureOpen       = "type %s %sstruct {\n"
	KeyValueSimpleStructure = "%s%s %T `json:\"%s\"`\n"
	AnonymusStruct          = "%s%s %sstruct {\n"
	CloseStructure          = "%s} `json:\"%s\"`\n"
)
