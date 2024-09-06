package aocreader

import "os"

func NewMockReader(lines []string) *MockReader {
	return &MockReader{
		lines: lines,
	}
}

func NewAocReader(inputPath string) *AocReader {
	return &AocReader{
		inputPath: inputPath,
	}
}

func (r *MockReader) Read(handler LineHandler) {
	for _, line := range r.lines {
		if line != "" && handler(line) {
			break
		}
	}
}

func (r *AocReader) Read(handler LineHandler) {
	data, err := os.ReadFile(r.inputPath)

	if err != nil {
		panic(err)
	}

	lastIndex := 0

	for i, b := range data {
		if b == '\n' && i-lastIndex > 0 {
			line := string(data[lastIndex:i])

			if handler(line) {
				break
			}

			lastIndex = i + 1
		} else if i == len(data)-1 && len(data)-lastIndex > 0 {
			line := string(data[lastIndex:])

			handler(line)
		}
	}
}
