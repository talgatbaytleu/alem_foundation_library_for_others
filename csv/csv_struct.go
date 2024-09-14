package csv

import (
	"io"
)

type CSVstruct struct {
	current_symbol []byte
  previous_symbol byte
	current_line   string
  quote_enabled bool
}

func (CSV CSVstruct) ReadLine(r io.Reader) (string, error) {
	CSV.current_symbol = make([]byte, 1)
	CSV.current_line = ""
  var quotes_counter int
	for {
		_, err := r.Read(CSV.current_symbol)
		if CSV.current_line == "" {
      if CSV.current_symbol[0] == '\n' {
        continue
      } else if CSV.current_symbol[0] == '"' {
        CSV.quote_enabled = true
        continue
      }
    } else if {

		} else if !CSV.quote_enabled && (err == io.EOF || CSV.current_symbol[0] == '\n' || CSV.current_symbol[0] == '\r') {
			break
    } else {
			CSV.current_line += string(CSV.current_symbol[0])
		}
    CSV.previous_symbol = CSV.current_symbol[0]
	}
	return CSV.current_line, err
}

func (CSV CSVstruct) GetField(n int) (string, error) {
	return "", nil
}

func (CSV CSVstruct) GetNumberOfFields() int {
	return 0
}
