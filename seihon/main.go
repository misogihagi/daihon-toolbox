//go:build ignore

// seihon ディレクトリは package seihon なので、通常のビルドからは除外する。
// 実行例:
//
//	go run seihon/main.go > seihon/b6_output.html
//	go run seihon/main.go seihon/b6_output.html
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/misogihagi/daihon-toolbox/seihon"
)

func main() {
	out := os.Stdout
	if len(os.Args) > 1 {
		f, err := os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		out = f
	}

	if err := seihon.B6().Render(context.Background(), out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
