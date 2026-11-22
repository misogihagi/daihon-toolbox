package honbun

import "strings"

// BodyData は本文テンプレートで差し込めるテキスト一式。
// 空のフィールドはサンプル値に置き換えられる。
type BodyData struct {
	Title       string   // 書名（<title> と右ページの柱）
	Author      string   // 著者名（左ページの柱）
	SpreadLabel string   // 見開きの aria-label
	RightPage   PageData // 右ページ（縦書きの1ページ目）
	LeftPage    PageData // 左ページ（縦書きの2ページ目）
}

// PageData は本文の1ページ分のテキスト。
type PageData struct {
	Label          string   // ページの aria-label
	ChapterTitle   string   // 章題（空なら非表示）
	Paragraphs     []string // 本文の段落（場面転換の前）
	SceneBreak     string   // 場面転換の記号（空なら非表示）
	TailParagraphs []string // 場面転換のあとの段落
	Folio          string   // ノンブル（ページ番号）
}

// DefaultSimple は純文学本文のサンプルテキスト。
func DefaultSimple() BodyData {
	return BodyData{
		Title:       "霧の桟橋で待っている",
		Author:      "朝倉 静子",
		SpreadLabel: "純文学の本文見開き",
		RightPage: PageData{
			Label:        "13ページ",
			ChapterTitle: "一 潮の匂い",
			Paragraphs: []string{
				"午前五時の桟橋は、まだ夜の名残を木板の隙間に溜めている。靴底が湿った板を鳴らすたび、足元の暗がりから細い潮の匂いが立ちのぼった。",
				"対岸の灯台は三秒に一度、白く薄い光を水面へ投げかける。光が届くのはほんの一瞬で、そのあとにはまた、どこまで続いているのかわからない濃霧が戻ってくる。",
				"「船は、本当に来るんでしょうか」",
				"コートの襟を立てたまま、そう尋ねたのは昨日から何回目だろう。返事の代わりに聞こえたのは、係留ロープが軋む乾いた音だけだった。",
			},
			Folio: "13",
		},
		LeftPage: PageData{
			Label: "14ページ",
			Paragraphs: []string{
				"ポケットの中で、指先が小さな真鍮の鍵に触れる。誰に返すあてもない鍵を、彼女はずっと手放せずにいた。",
				"あの部屋のドアは、もう別の誰かが開けているかもしれない。それでも、鍵を握っているあいだだけは、自分がまだあの時間の続きにいるような気がした。",
			},
			SceneBreak: "＊ ＊ ＊",
			TailParagraphs: []string{
				"霧の奥で、かすかに汽笛のような音が響いた気がした。顔を上げると、灰色の空と海の境目がわずかに白みはじめている。",
				"待つという行為は、終わりを決めてしまえば案外しずかなものだ。彼女はベンチの端に腰を下ろし、冷えた手のひらをゆっくりとほどいた。",
			},
			Folio: "14",
		},
	}
}

// withDefaults は空のフィールドだけをサンプル値で埋める。
func withDefaults(data, fallback BodyData) BodyData {
	data.Title = valueOr(data.Title, fallback.Title)
	data.Author = valueOr(data.Author, fallback.Author)
	data.SpreadLabel = valueOr(data.SpreadLabel, fallback.SpreadLabel)
	data.RightPage = withPageDefaults(data.RightPage, fallback.RightPage)
	data.LeftPage = withPageDefaults(data.LeftPage, fallback.LeftPage)
	return data
}

func withPageDefaults(data, fallback PageData) PageData {
	data.Label = valueOr(data.Label, fallback.Label)
	data.ChapterTitle = valueOr(data.ChapterTitle, fallback.ChapterTitle)
	data.Paragraphs = valuesOr(data.Paragraphs, fallback.Paragraphs)
	data.SceneBreak = valueOr(data.SceneBreak, fallback.SceneBreak)
	data.TailParagraphs = valuesOr(data.TailParagraphs, fallback.TailParagraphs)
	data.Folio = valueOr(data.Folio, fallback.Folio)
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
