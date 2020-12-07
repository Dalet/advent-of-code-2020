package util

import "bytes"

// SplitOnBlankLine is a split function for a Scanner that splits on blank lines ("\n\n")
func SplitOnBlankLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[:i+1], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
