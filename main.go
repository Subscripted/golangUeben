package main

import (
	"fmt"
	"io"
	"os"
)

type TestConfig struct {
	Name   string
	Age    int
	Skills []string
}

func main() {
	/*
		btFile, err := os.ReadFile("person.json")
		if err != nil {
			panic(err)
		}

		reader := bytes.NewReader(btFile)

		data, err := readFile(reader)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))


	*/

	fib()
	aRes, length := sliceFib(100000)
	fmt.Println(aRes, length)
}

func readFile(reader io.Reader) ([]byte, error) {
	if reader == nil {
		panic("reader must not be nil")
	}
	return io.ReadAll(reader)
}

func createNewFile(name string) (*os.File, error) {
	return os.Create(name)
}

func fib() {
	x, y := 0, 1
	for x < 1000 {
		fmt.Println(x)
		x, y = y, x+y
	}
}

func sliceFib(limit int) ([]int, int) {
	aFib := make([]int, 5)
	x, y := 0, 1
	for x < limit {
		aFib = append(aFib, x)
		x, y = y, x+y
	}
	return aFib, len(aFib)
}
