# AocReader

This lib is a helper to my [aoc solutions](https://github.com/marcos-venicius/aoc).

## Mocking

This method allows you to pass an array as the lines of a file

```go
import "github.com/marcos-venicius/aocreader"

lines := []string{
  "line 1",
  "...",
}

reader := aocreader.NewMockReader(lines)

for reader.Running() {
    i, line := reader.Line()

    ...
}
```

## Using

This method allows you to pass the path to a file where you want to read the lines

```go
import "github.com/marcos-venicius/aocreader"

reader := aocreader.NewAocReader("./input.txt")

for reader.Running() {
    i, line := reader.Line()

    ...
}
```
