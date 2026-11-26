package okuduke

import "strings"

// ColophonData は奥付テンプレートで差し込めるテキスト一式。
// 空のフィールドはサンプル値に置き換えられる。
type ColophonData struct {
	PageTitle   string         // <title>
	Title       string         // 書名
	Series      string         // 叢書・シリーズ名
	Seal        []string       // 検印風の印章の文字（行ごと）
	PublishedAt string         // 発行日・版数
	Meta        []ColophonMeta // 書誌データ（dl の各行）
	ISBN        string         // ISBN
	Price       string         // 価格表記（Cコードを含む）
	Notice      string         // 乱丁・落丁の案内
	Copyright   string         // 著作権表記
}

// ColophonMeta は奥付の書誌データの1行。
type ColophonMeta struct {
	Term      string // 項目名（例: 著  者）
	Value     string // 内容
	SubDetail string // 補足（住所など。空なら非表示）
}

// DefaultSimple は奥付のサンプルテキスト。
func DefaultSimple() ColophonData {
	return ColophonData{
		PageTitle:   "奥付 — 霧の桟橋で待っている",
		Title:       "霧の桟橋で待っている",
		Series:      "現代文学選",
		Seal:        []string{"朝倉", "之印"},
		PublishedAt: "二〇二六年十一月二十一日 第一刷発行",
		Meta: []ColophonMeta{
			{Term: "著  者", Value: "朝倉 静子"},
			{Term: "発 行 者", Value: "神田 誠一"},
			{Term: "発 行 所", Value: "株式会社 青霧書房", SubDetail: "〒101-0051 東京都千代田区神田神保町二丁目十四番地"},
			{Term: "編  集", Value: "佐伯 恒一"},
			{Term: "装丁設計", Value: "三輪 澄江"},
			{Term: "印刷・製本", Value: "精興堂印刷株式会社"},
		},
		ISBN:      "ISBN978-4-901234-56-7",
		Price:     "C0093 ¥1600E",
		Notice:    "落丁・乱丁本はお手数ですが小社制作部までお送りください。送料小社負担にてお取り替えいたします。",
		Copyright: "© 2026 Shizuko Asakura, Printed in Japan",
	}
}

// withDefaults は空のフィールドだけをサンプル値で埋める。
func withDefaults(data, fallback ColophonData) ColophonData {
	data.PageTitle = valueOr(data.PageTitle, fallback.PageTitle)
	data.Title = valueOr(data.Title, fallback.Title)
	data.Series = valueOr(data.Series, fallback.Series)
	data.Seal = valuesOr(data.Seal, fallback.Seal)
	data.PublishedAt = valueOr(data.PublishedAt, fallback.PublishedAt)
	if len(data.Meta) == 0 {
		data.Meta = fallback.Meta
	}
	data.ISBN = valueOr(data.ISBN, fallback.ISBN)
	data.Price = valueOr(data.Price, fallback.Price)
	data.Notice = valueOr(data.Notice, fallback.Notice)
	data.Copyright = valueOr(data.Copyright, fallback.Copyright)
	return data
}

func valueOr(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func valuesOr(values, fallback []string) []string {
	if len(values) == 0 {
		return fallback
	}
	return values
}
