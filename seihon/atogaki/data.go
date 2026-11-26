package atogaki

import "strings"

// AtogakiData はあとがき・制作コメントのテンプレートで差し込めるテキスト一式。
// 空のフィールドは各テンプレート側でサンプル値に置き換えられる。
type AtogakiData struct {
	PageTitle   string    // <title>
	Title       string    // 書名（柱）
	SpreadLabel string    // 見開きの aria-label
	Essay       EssayData // 作者によるあとがきページ
	Notes       NotesData // 編集・装丁からのコメントページ
}

// EssayData は作者のあとがきページのテキスト。
type EssayData struct {
	Label       string   // ページの aria-label
	RunningHead string   // 柱
	Heading     string   // 見出し
	Author      string   // 署名（横書きでは柱にも表示）
	Paragraphs  []string // あとがきの段落
	Folio       string   // ノンブル
}

// NotesData は編集・装丁コメントのページのテキスト。
type NotesData struct {
	Label        string        // ページの aria-label
	CommentLabel string        // コメント欄の aria-label
	RunningHead  string        // 柱（横書きのみ）
	Heading      string        // 見出し
	Comments     []CommentData // 担当者ごとのコメント
	Folio        string        // ノンブル
}

// CommentData は編集・装丁担当者一人分のコメント。
type CommentData struct {
	Role string // 役職（例: 担当編集）
	Name string // 氏名
	Text string // コメント本文
}

// DefaultSimple は縦組みのあとがきサンプルテキスト。
func DefaultSimple() AtogakiData {
	return AtogakiData{
		PageTitle:   "あとがき・制作コメント — 霧の桟橋で待っている",
		Title:       "霧の桟橋で待っている",
		SpreadLabel: "あとがきと制作陣コメントの見開き",
		Essay: EssayData{
			Label:       "211ページ",
			RunningHead: "あとがき",
			Heading:     "書き終えて",
			Author:      "朝倉 静子",
			Paragraphs: []string{
				"この物語を書きはじめたとき、正直なところ、主人公が桟橋のベンチから立ち上がる場面まで辿り着けるのか自分でもわかりませんでした。霧のなかに立ち止まったままの時間を書くことは、書いている私自身もまた、同じ冷たい空気のなかに座り続けるような作業だったからです。",
				"けれど一行ずつ言葉を置いていくうちに、声高には語られない迷いや、誰にも渡せなかった鍵の重みそのものが、静かに物語を前へ押し出してくれました。大きな出来事は起きなくても、人の心は確かに動いているのだと、登場人物たちに教わった気がします。",
				"書き終えた今、ようやく朝の光が差した桟橋を後ろから見送るような気持ちでいます。この静けさが、読んでくださった方のどこかにも残れば幸いです。",
			},
			Folio: "211",
		},
		Notes: NotesData{
			Label:        "212ページ",
			CommentLabel: "編集・装丁コメント",
			Heading:      "一冊ができるまで",
			Comments: []CommentData{
				{
					Role: "担当編集",
					Name: "佐伯 恒一",
					Text: "最初に第一稿を拝読した夜、読み終えてもしばらく部屋の明かりをつけられずにいました。余白のなかにこれほど豊かな感情を宿せるのかと、改稿を重ねるたびに胸を打たれました。静かな夜に何度も開きたくなる一冊です。",
				},
				{
					Role: "装丁・造本設計",
					Name: "三輪 澄江",
					Text: "原稿から立ちのぼる潮の匂いと朝の冷気を、手触りのある生成り色の紙と最小限の罫線に託しました。タイトルの縦組みが、霧のなかに佇む一本の桟橋のように見えればと思いながら設計しています。",
				},
			},
			Folio: "212",
		},
	}
}

// DefaultHorizontal は横書きのあとがきサンプルテキスト。
func DefaultHorizontal() AtogakiData {
	data := DefaultSimple()
	data.PageTitle = "あとがき・制作コメント（横書き） — 霧の桟橋で待っている"
	data.SpreadLabel = "横書きのあとがきと制作陣コメント見開き"
	data.Notes.RunningHead = "制作ノート"
	data.Notes.Comments[1].Text = "原稿から立ちのぼる潮の匂いと朝の冷気を、手触りのある生成り色の紙と最小限の罫線に託しました。タイトルの組み方が、霧のなかに佇む一本の桟橋のように見えればと思いながら設計しています。"
	return data
}

// withDefaults は空のフィールドだけをサンプル値で埋める。
func withDefaults(data, fallback AtogakiData) AtogakiData {
	data.PageTitle = valueOr(data.PageTitle, fallback.PageTitle)
	data.Title = valueOr(data.Title, fallback.Title)
	data.SpreadLabel = valueOr(data.SpreadLabel, fallback.SpreadLabel)
	data.Essay = withEssayDefaults(data.Essay, fallback.Essay)
	data.Notes = withNotesDefaults(data.Notes, fallback.Notes)
	return data
}

func withEssayDefaults(data, fallback EssayData) EssayData {
	data.Label = valueOr(data.Label, fallback.Label)
	data.RunningHead = valueOr(data.RunningHead, fallback.RunningHead)
	data.Heading = valueOr(data.Heading, fallback.Heading)
	data.Author = valueOr(data.Author, fallback.Author)
	data.Paragraphs = valuesOr(data.Paragraphs, fallback.Paragraphs)
	data.Folio = valueOr(data.Folio, fallback.Folio)
	return data
}

func withNotesDefaults(data, fallback NotesData) NotesData {
	data.Label = valueOr(data.Label, fallback.Label)
	data.CommentLabel = valueOr(data.CommentLabel, fallback.CommentLabel)
	data.RunningHead = valueOr(data.RunningHead, fallback.RunningHead)
	data.Heading = valueOr(data.Heading, fallback.Heading)
	if len(data.Comments) == 0 {
		data.Comments = fallback.Comments
	}
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
