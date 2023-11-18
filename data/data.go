package data

import (
	"embed"
	"encoding/csv"
	"strings"
)

//go:embed csv
var data embed.FS

type csvEmbed struct {
	files embed.FS
}

func LoadCsvData() (csv *csvEmbed) {
	return &csvEmbed{files: data}
}

func (c *csvEmbed) Read(filename string) (record [][]string, err error) {
	var text []byte
	if text, err = c.files.ReadFile("csv/" + filename + ".csv"); err != nil {
		return
	}
	r := csv.NewReader(strings.NewReader(string(text)))
	r.Comment = '#'
	return r.ReadAll()
}
