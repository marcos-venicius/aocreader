package aocreader

// returns true if want's to stop iterating
type LineHandler func(line string) bool

type LinesReader interface {
	Read(handler LineHandler)
}

type AocReader struct {
	inputPath string
}

type MockReader struct {
	lines []string
}
