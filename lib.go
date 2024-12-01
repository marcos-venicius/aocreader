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

func (r *MockReader) Reset() {
	r.index = 0
}

func (r *MockReader) Running() bool {
	return r.index < len(r.lines)
}

func (r *MockReader) Line() (int, string) {
	defer r.update()

	if !r.Running() {
		panic("invalid range")
	}

	return r.index, r.lines[r.index]
}

func (r *MockReader) update() {
	r.index++
}

func (r *AocReader) Reset() {
	r.lineIndex = 0
	r.currentIndex = 0
	r.contentIndex = 0
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
