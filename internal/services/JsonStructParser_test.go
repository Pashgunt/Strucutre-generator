package services

import (
	"bytes"
	"testing"
)

func TestJsonStructParser_writeTypeStructOpen(t *testing.T) {
	type fields struct {
		Config ConfigParser
		buffer bytes.Buffer
		ch     chan []byte
	}
	type args struct {
		additionalOpen string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test",
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			args: args{additionalOpen: "[]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &JsonStructParser{
				Config: tt.fields.Config,
				buffer: tt.fields.buffer,
				ch:     tt.fields.ch,
			}
			parser.writeTypeStructOpen(tt.args.additionalOpen)
		})
	}
}

func TestJsonStructParser_writeAnonymusStructure(t *testing.T) {
	type fields struct {
		Config ConfigParser
		buffer bytes.Buffer
		ch     chan []byte
	}
	type args struct {
		spaceLevel     int
		key            string
		additionalOpen string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test",
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			args: args{
				spaceLevel:     4,
				key:            "name",
				additionalOpen: "[]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &JsonStructParser{
				Config: tt.fields.Config,
				buffer: tt.fields.buffer,
				ch:     tt.fields.ch,
			}
			parser.writeAnonymusStructure(tt.args.spaceLevel, tt.args.key, tt.args.additionalOpen)
		})
	}
}

func TestJsonStructParser_writeCloseStructure(t *testing.T) {
	type fields struct {
		Config ConfigParser
		buffer bytes.Buffer
		ch     chan []byte
	}
	type args struct {
		spaceLevel int
		key        string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test",
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			args: args{
				spaceLevel: 4,
				key:        "name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &JsonStructParser{
				Config: tt.fields.Config,
				buffer: tt.fields.buffer,
				ch:     tt.fields.ch,
			}
			parser.writeCloseStructure(tt.args.spaceLevel, tt.args.key)
		})
	}
}

func TestJsonStructParser_writeSimpleKeyValue(t *testing.T) {
	type fields struct {
		Config ConfigParser
		buffer bytes.Buffer
		ch     chan []byte
	}
	type args struct {
		spaceLevel int
		key        string
		value      interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test",
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			args: args{
				spaceLevel: 4,
				key:        "name",
				value:      "Pavel",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &JsonStructParser{
				Config: tt.fields.Config,
				buffer: tt.fields.buffer,
				ch:     tt.fields.ch,
			}
			parser.writeSimpleKeyValue(tt.args.spaceLevel, tt.args.key, tt.args.value)
		})
	}
}

func TestJsonStructParser_capitalizeKey(t *testing.T) {
	type fields struct {
		Config ConfigParser
		buffer bytes.Buffer
		ch     chan []byte
	}
	type args struct {
		key string
	}
	tests := []struct {
		fields fields
		name   string
		args   args
		want   string
	}{
		{
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			name: "Test1",
			args: args{key: "id"},
			want: "ID",
		},
		{
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			name: "Test2",
			args: args{key: "Test"},
			want: "Test",
		},
		{
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			name: "Test3",
			args: args{key: "test"},
			want: "Test",
		},
		{
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			name: "Test4",
			args: args{key: "___id__"},
			want: "ID",
		},
		{
			fields: fields{
				Config: ConfigParser{
					file:      "Tets",
					structure: "Test",
					mode:      "cli",
				},
				buffer: bytes.Buffer{},
				ch:     make(chan []byte),
			},
			name: "Test5",
			args: args{key: "___name__"},
			want: "Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &JsonStructParser{
				Config: tt.fields.Config,
				buffer: tt.fields.buffer,
				ch:     tt.fields.ch,
			}
			if got := parser.capitalizeKey(tt.args.key); got != tt.want {
				t.Errorf("capitalizeKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
