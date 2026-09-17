package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func debugPrintMap(nameMap map[string]string) {
	fmt.Println("nameMap := map[string]string{")
	for k, v := range nameMap {
		fmt.Printf("\t\"%s\": \"%s\",\n", k, v)
	}
	fmt.Println("}")
}

func main() {
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

    debugPrintMap(nameMap)
}
