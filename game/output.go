package game

import (
	"bufio"
	"os"
)

type Output interface {
	Write(text string)
	Read() (string, error)
}

type DefaultOutput struct {
}

func (output *DefaultOutput) Write(text string) {
	println(text)
}

func (output *DefaultOutput) Read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func NewOutput() *DefaultOutput {
	return &DefaultOutput{}
}
