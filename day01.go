package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Map[T, U any](f func(T) (U, error), xs []T) ([]U, error) {
	ys := make([]U, len(xs))
	var err error
	for i, x := range xs {
		ys[i], err = f(x)
		if err != nil {
			return ys, err
		}
	}
	return ys, nil
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	depths, err := Map(strconv.Atoi, strings.Split(strings.TrimSpace(string(b)), "\n"))
	if err != nil {
		panic(err)
	}

	incs := 0
	last := depths[0]
	for _, depth := range depths[1:] {
		if depth > last {
			incs++
		}
		last = depth
	}
	fmt.Println(incs)

	incs = 0
	lasts := depths[:3]
	for _, depth := range depths[3:] {
		if depth > lasts[0] {
			incs++
		}
		lasts[0], lasts[1], lasts[2] = lasts[1], lasts[2], depth
	}
	fmt.Println(incs)
}
