package csv

import (
	"io"
)

type CSVstruct struct {
	current_symbol     byte
	previous_symbol    byte
	current_line       string
	current_field      string
	current_line_slice []string
	quote_enabled      bool
}

func (c CSVstruct) ReadLine(r io.Reader) (string, error) {
	helper_slice := make([]byte, 1)
	c.current_line = ""
	var err error
	for {
		_, err = r.Read(helper_slice)
		c.current_symbol = helper_slice[0]

		if !c.quote_enabled {
			if c.current_line == "" {
				if c.current_symbol == '\n' {
					continue
				} else if c.current_symbol == '"' {
					c.quote_enabled = true
					continue
				}
			} else if c.previous_symbol == ',' {
				if c.current_symbol == '"' {
					c.quote_enabled = true
				}
			} else if c.current_symbol == '\n' || c.current_symbol == '\r' {
				break
			} else if c.current_symbol == '"' {
				err = ErrQuote
				break
			}
		} else {
		}

		if err == ErrQuote {
			return "", err
		} else {
			return c.current_line, err
		}
	}

	return c.current_line, err
}

func (c CSVstruct) GetField(n int) (string, error) {
	return "", nil
}

func (c CSVstruct) GetNumberOfFields() int {
	return 0
}
