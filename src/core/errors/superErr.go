package errors

import (
	"runtime"
)

type superErr struct {
	line int
	file string
}

func newSuperErr() *superErr {
	_, file, line, _ := runtime.Caller(2)
	return &superErr{line, file}
}

func (e superErr) File() string {
	return e.file
}

func (e superErr) Line() int {
	return e.line
}

func (e superErr) LogLevel() int {
	return InfoLevel
}
