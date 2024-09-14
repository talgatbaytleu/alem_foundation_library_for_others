package csv

import (
	"io"
)

type CSVstruct struct {
	current_symbol []byte
	current_line   string
}

func (CSV CSVstruct) ReadLine(r io.Reader) (string, error) {
	CSV.current_symbol = make([]byte, 1)
	CSV.current_line = ""
	for {
		_, err := r.Read(CSV.current_symbol)
		if CSV.current_line == "" && CSV.current_symbol[0] == '\n' {
			continue
		} else if err == io.EOF || CSV.current_symbol[0] == '\n' || CSV.current_symbol[0] == '\r' {
			break
		} else {
			CSV.current_line += string(CSV.current_symbol[0])
		}
	}
	return CSV.current_line, nil
}

func (CSV CSVstruct) GetField(n int) (string, error) {
	return "", nil
}

func (CSV CSVstruct) GetNumberOfFields() int {
	return 0
}
