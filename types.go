package aocreader

// returns true if want's to stop iterating
type LineHandler func(line string) bool

type LinesReader interface {
	Running() bool
	Line() (int, string)
	Reset()
}

type AocReader struct {
	lineIndex    int
	currentIndex int
	contentIndex int
	content      []byte
}

type MockReader struct {
	lines []string
	index int
}
