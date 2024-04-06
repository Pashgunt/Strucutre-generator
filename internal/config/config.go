package config

const (
	ChunkSizeMin  = 1024 * iota
	ChunkSize1024 = 1024 * iota
	ChunkSize2048 = 1024 * iota
	ChunkSize3072 = 1024 * iota
	ChunkSize4096 = 1024 * iota
	CliMode       = "cli"
	FileMode      = "file"
)
