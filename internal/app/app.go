package app

import (
	"flag"
	"fmt"
	enums "generatorStructure/internal/config"
	"generatorStructure/internal/services"
)

func Init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	file := flag.String("file", "../../testdata/example.json", "Path to file for parsing to struct")
	structure := flag.String("structure", "default", "Name for generated structure")
	mode := flag.String("mode", enums.CliMode, "Type of output result: cli or file")
	flag.Parse()

	if len(*file) == 0 {
		panic("для корректной работы программы передайте путь к файлу")
	}

	config := services.ConfigParser{}
	config.SetFile(*file).SetStructure(*structure).SetMode(*mode)

	_, err := services.NewJsonStructParser(config).Read().Result()

	if err != nil {
		panic(err)
	}
}
