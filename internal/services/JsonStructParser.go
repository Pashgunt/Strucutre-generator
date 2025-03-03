package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pashgunt/Strucutre-generator/internal/config"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

type JsonStructParser struct {
	Config ConfigParser
	buffer bytes.Buffer
	ch     chan []byte
}

func NewJsonStructParser(Config ConfigParser) *JsonStructParser {
	return &JsonStructParser{
		buffer: bytes.Buffer{},
		ch:     make(chan []byte),
		Config: Config,
	}
}

func (parser *JsonStructParser) Read() *JsonStructParser {
	file, err := os.Open(parser.Config.File())
	defer func(file *os.File) {
		parser.buffer.WriteString("}")
		close(parser.ch)
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	if err != nil {
		panic(err)
	}

	data := make([]byte, 0)
	chunk := make([]byte, config.ChunkSize4096)

	for {
		n, err := file.Read(chunk)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		data = append(data, chunk[:n]...)
	}

	go func() {
		parser.ch <- data
	}()

	parser.processChunk()

	return parser
}

func (parser *JsonStructParser) Result() (bool, error) {
	if parser.Config.Mode() == config.CliMode {
		fmt.Println(parser.buffer.String())
		return true, nil
	}

	folder := fmt.Sprintf("../structures/%s/%d", time.DateOnly, time.Now().Unix())
	err := os.MkdirAll(folder, os.FileMode(0777))

	if err != nil {
		return false, err
	}

	err = os.WriteFile(
		fmt.Sprintf("./%s/%s.go", folder, parser.Config.Structure()),
		parser.buffer.Bytes(),
		os.FileMode(0777),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (parser *JsonStructParser) processChunk() {
	select {
	case chunk := <-parser.ch:
		var data interface{}
		err := json.Unmarshal(chunk, &data)

		if err != nil {
			panic(err)
		}

		parser.processUnmarshal(data)
	}
}

func (parser *JsonStructParser) processUnmarshal(data interface{}) {
	switch dataStruct := data.(type) {
	case map[string]interface{}:
		parser.writeTypeStructOpen("")
		parser.doProcessObject(dataStruct, 4)
		break
	case []interface{}:
		parser.writeTypeStructOpen("[]")

		m, ok := dataStruct[0].(map[string]interface{})

		if ok {
			parser.doProcessObject(m, 4)
		}

		break
	}
}

func (parser *JsonStructParser) doProcessObject(
	dataStruct map[string]interface{},
	spaceLevel int,
) {
	for key, value := range dataStruct {
		switch data := value.(type) {
		case map[string]interface{}:
			parser.writeAnonymusStructure(spaceLevel, key, "")
			parser.doProcessObject(data, spaceLevel+4)
			parser.writeCloseStructure(spaceLevel, key)
			break
		case []interface{}:
			m, ok := data[0].(map[string]interface{})
			if ok {
				parser.writeAnonymusStructure(spaceLevel, key, "[]")
				parser.doProcessObject(m, spaceLevel+4)
				parser.writeCloseStructure(spaceLevel, key)
			} else {
				parser.writeSimpleKeyValue(spaceLevel, key, value)
			}
			break
		default:
			parser.writeSimpleKeyValue(spaceLevel, key, value)
		}
	}
}

func (parser *JsonStructParser) capitalizeKey(key string) string {
	var re = regexp.MustCompile(`^_+|_+$`)
	var re2 = regexp.MustCompile(`(?i)^id$`)
	formatKey := re.ReplaceAllString(key, "")

	if re2.MatchString(formatKey) {
		return strings.ToTitle(formatKey)
	}

	return strings.Title(formatKey)
}

func (parser *JsonStructParser) writeSimpleKeyValue(spaceLevel int, key string, value interface{}) {
	parser.buffer.WriteString(fmt.Sprintf(
		config.KeyValueSimpleStructure,
		strings.Repeat(" ", spaceLevel),
		parser.capitalizeKey(key),
		value,
		key,
	))
}

func (parser *JsonStructParser) writeCloseStructure(spaceLevel int, key string) {
	parser.buffer.WriteString(fmt.Sprintf(
		config.CloseStructure,
		strings.Repeat(" ", spaceLevel),
		key,
	))
}

func (parser *JsonStructParser) writeAnonymusStructure(
	spaceLevel int,
	key string,
	additionalOpen string,
) {
	parser.buffer.WriteString(fmt.Sprintf(
		config.AnonymusStruct,
		strings.Repeat(" ", spaceLevel),
		parser.capitalizeKey(key),
		additionalOpen,
	))
}

func (parser *JsonStructParser) writeTypeStructOpen(additionalOpen string) {
	parser.buffer.WriteString(fmt.Sprintf(
		config.TypeStructureOpen,
		parser.Config.Structure(),
		additionalOpen,
	))
}
