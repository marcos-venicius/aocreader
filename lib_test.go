package aocreader

import "testing"

func TestNewMockReader(t *testing.T) {
	data := []string{
		"line 1",
		"line 2",
		"",
	}

	reader := NewMockReader(data)

	if len(reader.lines) != 3 {
		t.Fatalf("%d != %d", len(reader.lines), 3)
	}

	if reader.lines[0] != data[0] {
		t.Fatalf("%s != %s", reader.lines[0], data[0])
	}

	if reader.lines[1] != data[1] {
		t.Fatalf("%s != %s", reader.lines[1], data[1])
	}
}

func TestMockRead(t *testing.T) {
	data := []string{
		"line 1",
		"line 2",
		"",
	}

	iterationsCount := 0

	reader := NewMockReader(data)

	reader.Read(func(line string) bool {
		if line != data[iterationsCount] {
			t.Fatalf(`"%s" != "%s"`, line, data[iterationsCount])
		}

		iterationsCount++

		return false
	})

	if iterationsCount != 2 {
		t.Fatalf("%d != %d", iterationsCount, 2)
	}
}

func TestNewAocReader(t *testing.T) {
	filename := "./test.txt"

	reader := NewAocReader(filename)

	if reader.inputPath != filename {
		t.Fatalf(`"%s" != "%s"`, reader.inputPath, filename)
	}
}

func TestAocRead(t *testing.T) {
	reader := NewAocReader("./test.txt")

	data := []string{
		"line 1",
		"line 2",
	}

	iterationsCount := 0

	reader.Read(func(line string) bool {
		if line != data[iterationsCount] {
			t.Fatalf(`"%s" != "%s"`, line, data[iterationsCount])
		}

		iterationsCount++

		return false
	})

	if iterationsCount != 2 {
		t.Fatalf("%d != %d", iterationsCount, 2)
	}
}
