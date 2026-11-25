package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"regexp"
)

type CMYKCheckResult struct {
	HasDeviceCMYK   bool
	Has4ChICCBased  bool
	HasCMYKOperator bool
}

func (r CMYKCheckResult) LikelyCMYK() bool {
	return r.HasDeviceCMYK || r.Has4ChICCBased || r.HasCMYKOperator
}

// PDF内の生データおよびFlateDecodeストリームを展開してCMYK関連要素を検査します
func inspectPDFForCMYK(data []byte) CMYKCheckResult {
	var res CMYKCheckResult

	// 1. 辞書・カラースペース指定のチェック（生データ）
	if bytes.Contains(data, []byte("/DeviceCMYK")) {
		res.HasDeviceCMYK = true
	}

	// ICCBased カラースペースで /N 4（4成分 = CMYK）が使われているかチェック
	icc4Re := regexp.MustCompile(`/ICCBased[\s\S]{0,300}?/N\s+4\b`)
	if icc4Re.Match(data) {
		res.Has4ChICCBased = true
	}

	// 2. コンテンツストリーム内のCMYK描画演算子（非ストローク `k` / ストローク `K`）をチェック
	// 例: "0.1 0.2 0.3 0.4 k"
	cmykOpRe := regexp.MustCompile(`(?:^|[\r\n\s])(?:[0-9.]+\s+){4}[kK](?:[\r\n\s]|$)`)
	if cmykOpRe.Match(data) {
		res.HasCMYKOperator = true
	}

	// 3. FlateDecode ストリームを展開して再検査
	streamStartRe := regexp.MustCompile(`stream\r?\n`)
	endStream := []byte("endstream")

	locs := streamStartRe.FindAllIndex(data, -1)
	for _, loc := range locs {
		start := loc[1]
		endIdx := bytes.Index(data[start:], endStream)
		if endIdx == -1 {
			continue
		}
		rawStream := bytes.TrimSuffix(data[start:start+endIdx], []byte("\r\n"))
		rawStream = bytes.TrimSuffix(rawStream, []byte("\n"))

		zr, err := zlib.NewReader(bytes.NewReader(rawStream))
		if err != nil {
			continue
		}
		decompressed, err := io.ReadAll(zr)
		_ = zr.Close()
		if err != nil || len(decompressed) == 0 {
			continue
		}

		if bytes.Contains(decompressed, []byte("/DeviceCMYK")) {
			res.HasDeviceCMYK = true
		}
		if cmykOpRe.Match(decompressed) {
			res.HasCMYKOperator = true
		}
	}

	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "使い方: go run check_cmyk.go <file.pdf>")
		os.Exit(1)
	}

	path := os.Args[1]
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル読み込みエラー: %v\n", err)
		os.Exit(1)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		fmt.Fprintln(os.Stderr, "有効なPDFファイルではありません")
		os.Exit(1)
	}

	res := inspectPDFForCMYK(data)
	fmt.Printf("ファイル: %s\n", path)
	fmt.Printf("  - /DeviceCMYK あり: %v\n", res.HasDeviceCMYK)
	fmt.Printf("  - 4ch ICCBased (/N 4) あり: %v\n", res.Has4ChICCBased)
	fmt.Printf("  - CMYK描画演算子 (k/K) あり: %v\n", res.HasCMYKOperator)
	fmt.Printf("  => CMYK要素を含む可能性: %v\n", res.LikelyCMYK())
}
