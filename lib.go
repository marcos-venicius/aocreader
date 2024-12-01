package aocreader

import (
	"os"
)

func NewMockReader(lines []string) *MockReader {
	return &MockReader{
		lines: lines,
		index: 0,
	}
}

func NewAocReader(inputPath string) *AocReader {
	data, err := os.ReadFile(inputPath)

	if err != nil {
		panic(err)
	}

	return &AocReader{
		lineIndex:    0,
		currentIndex: 0,
		contentIndex: 0,
		content:      data,
	}
}

func (r *MockReader) Running() bool {
	return r.index < len(r.lines)-1
}

func (r *MockReader) Line() (int, string) {
	defer r.update()

	if r.index+1 >= len(r.lines) {
		panic("invalid range")
	}

	return r.index, r.lines[r.index]
}

func (r *MockReader) update() {
	r.index++
}

func (r *AocReader) Running() bool {
	return r.contentIndex < len(r.content)
}

func (r *AocReader) update() {
	r.lineIndex++
	r.currentIndex++
	r.contentIndex = r.currentIndex
}

func (r *AocReader) Line() (int, string) {
	defer r.update()

	if r.contentIndex >= len(r.content) {
		panic("invalid line range")
	}

	for _, b := range r.content[r.currentIndex:] {
		if b == '\n' {
			line := string(r.content[r.contentIndex:r.currentIndex])

			return r.lineIndex, line
		}

		r.currentIndex++
	}

	return r.lineIndex, string(r.content[r.contentIndex:])
}
