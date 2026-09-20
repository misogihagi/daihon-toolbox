package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)


func debugPrintMap(nameMap map[string]string) {
	fmt.Println("nameMap := map[string]string{")
	for k, v := range nameMap {
		fmt.Printf("\t\"%s\": \"%s\",\n", k, v)
	}
	fmt.Println("}")
}

func makeNameMap() map[string]string {
	// テキストファイルを開く（ファイル名は仮）
	file, err := os.Open("characters.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	nameMap := make(map[string]string)
	
	// 「名前（略称）：」の形式にマッチする正規表現
	re := regexp.MustCompile(`^(.+?)（(.+?)）`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) >= 3 {
			name := matches[1]
			key := matches[2]
			nameMap[key] = name
		}
	}
	
	return nameMap
}

func main() {
	nameMap := makeNameMap()
	
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	// 標準入力から脚本を読み込む
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
		os.Exit(1)
	}
	
    re := regexp.MustCompile(`^(.+?)[：:]\s*(.*)$`)

	for _, line := range lines {
		// 空行やコロンを含まない行はそのまま出力
		matches := re.FindStringSubmatch(line)
		if len(matches) < 3 {
			fmt.Println(line)
			continue
		}

		speakersPart := matches[1]
		dialogue := matches[2]

		// カッコの中の略を正しい登場人物名に置換
		// 複数話者の場合（例: "し・み"）を考慮して 「・」 で分割・置換
		speakers := strings.Split(speakersPart, "・")
		var fullSpeakers []string
		for _, s := range speakers {
			s = strings.TrimSpace(s)
			if fullName, exists := nameMap[s]; exists {
				fullSpeakers = append(fullSpeakers, fullName)
			} else {
				fullSpeakers = append(fullSpeakers, s) // マップにない場合はそのまま
			}
		}
		resolvedSpeaker := strings.Join(fullSpeakers, "・")

		// 「…」（三点リーダー）と「―」（ダッシュ）の文字数を2倍にする
		dialogue = doubleCharacters(dialogue)

		// コロンを削除し、カギカッコでくくる
		output := fmt.Sprintf("%s「%s」", resolvedSpeaker, dialogue)
		fmt.Println(output)
	}
}

// 三点リーダー（…）とダッシュ（―）を2倍に増やす関数
func doubleCharacters(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		switch r {
		case '…':
			sb.WriteRune('…')
			sb.WriteRune('…')
		case '―':
			sb.WriteRune('―')
			sb.WriteRune('―')
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
