package hyoushi

import "strings"

// CoverData は5種類の表紙テンプレートで差し込めるテキスト一式。
// 空のフィールドは各テンプレート側でサンプル値に置き換えられる。
type CoverData struct {
	// 共通の基本情報
	Title     []string // タイトル（行ごと）
	Author    string   // 著者名（「著：」などの接頭辞はテンプレート側で付く）
	Publisher string   // 出版社・レーベル名
	Subtitle  string   // サブタイトル・副題
	Catch     string   // キャッチコピー・煽り文
	Obi       []string // 帯の文言（行ごと）

	// レーベル / 版表示
	Label   string // レーベル表示
	Series  string // 叢書・シリーズ名
	Volume  string // 巻数
	Edition string // 版表示
	Number  string // 番号

	// デザイン固有のテキスト
	Genre       string        // ジャンル帯（comedy）
	Caution     string        // 注意書き（comedy）
	Speech      string        // ツッコミ吹き出し（comedy）
	GagLeft     string        // ギャグパネル見出し左（comedy）
	GagRight    string        // ギャグパネル見出し右（comedy）
	Receipt     []ReceiptLine // 領収書風の明細（comedy）
	ObiBadge    string        // 帯のバッジ（lightnovel）
	Illustrator string        // イラストレーター名（lightnovel）
	ArtCaption  string        // イラスト枠のキャプション（lightnovel）
	Epigraph    []string      // 詞書（literary）
	Hanko       string        // 印章の文字（literary）
	Pills       []string      // 上部のピル（pop）
	Kicker      string        // 煽り文言（pop）
	Badges      []string      // ミニバッジ（pop）
}

// ReceiptLine はコメディ表紙の領収書風パネルの1行。
type ReceiptLine struct {
	Text     string
	Emphasis string // 赤く強調する部分
}

// DefaultComedy はコメディ表紙のサンプルテキスト。
func DefaultComedy() CoverData {
	return CoverData{
		Title:    []string{"魔王軍の経理課は", "今日も残業が止まらない！"},
		Author:   "電卓 太郎",
		Label:    "WARAI BUNKO",
		Genre:    "お仕事ドタバタコメディ",
		Caution:  "※笑いすぎて作業中断注意",
		Speech:   "「勇者撃退の領収書、宛名が『世界平和』ってなんですか!?」",
		GagLeft:  "【本日の差し戻し案件】",
		GagRight: "承認待ち：99件",
		Receipt: []ReceiptLine{
			{Text: "・火の玉魔法の燃料費：計上不可"},
			{Text: "・城の壁を壊した損害賠償：", Emphasis: "月3回目"},
			{Text: "・部下の「ちょっと世界征服してきます」：", Emphasis: "理由不備"},
		},
		Obi: []string{
			"世界征服より、まずは経費精算！",
			"胃薬片手に奮闘する中間管理職コメディ第1弾！",
		},
	}
}

// DefaultLightNovel はライトノベル表紙のサンプルテキスト。
func DefaultLightNovel() CoverData {
	return CoverData{
		Title:      []string{"放課後、時計塔の鍵をなくした僕らは", "世界を巻き戻す"},
		Author:     "神楽 悠真",
		Label:      "STARLIGHT BUNKO",
		Volume:     "第1巻",
		ArtCaption: "ILLUSTRATION AREA（画像・SVG差し替え可）",
		Catch:      "「あと五分だけ、君とやり直したい」——停まった時間が動き出す青春ファンタジー！",
		ObiBadge:   "発売即重版決定！",
		Obi: []string{
			"無口な図書委員と落ちこぼれ鍵師が挑む、七度目の夏休み。",
			"ラスト三ページの伏線に、きっと息をのむ。",
		},
		Illustrator: "白銀 こなつ",
	}
}

// DefaultLiterary は純文学表紙のサンプルテキスト。
func DefaultLiterary() CoverData {
	return CoverData{
		Title:     []string{"霧の桟橋で", "待っている"},
		Author:    "朝倉 静子",
		Publisher: "青霧書房",
		Series:    "現代文学選",
		Epigraph: []string{
			"潮の引く音を数えながら、",
			"誰も来ない船着き場に立つ。",
			"忘れてしまった言葉だけが、",
			"冷たい空気のなかで澄んでいる。",
		},
		Hanko:   "栞",
		Edition: "第一刷",
	}
}

// DefaultPop はポップ表紙のサンプルテキスト。
func DefaultPop() CoverData {
	return CoverData{
		Title:     []string{"毎日がちょっと", "楽しくなる", "工夫ノート"},
		Author:    "ひなた まゆ",
		Publisher: "POP BOOKS",
		Kicker:    "5分でできる！",
		Pills:     []string{"初心者OK", "イラスト付き"},
		Badges:    []string{"#朝活", "#片付け", "#気分転換"},
		Obi: []string{
			"がんばりすぎないのに、ちゃんと変わる。",
			"小さなアイデア100選！",
		},
	}
}

// DefaultSimple はシンプル表紙のサンプルテキスト。
func DefaultSimple() CoverData {
	return CoverData{
		Title:    []string{"静かな夜の", "地図"},
		Author:   "湊 蒼一郎",
		Label:    "KAZE BUNKO",
		Number:   "042",
		Subtitle: "短編連作集",
		Obi: []string{
			"終電を逃した街角で、まだ名前のない約束を交わす。",
			"読むたびに行き先が変わる、十一の夜咄。",
		},
	}
}

// withDefaults は空のフィールドだけをデザイン別サンプルで埋める。
func withDefaults(data, fallback CoverData) CoverData {
	data.Title = valuesOr(data.Title, fallback.Title)
	data.Author = valueOr(data.Author, fallback.Author)
	data.Publisher = valueOr(data.Publisher, fallback.Publisher)
	data.Subtitle = valueOr(data.Subtitle, fallback.Subtitle)
	data.Catch = valueOr(data.Catch, fallback.Catch)
	data.Obi = valuesOr(data.Obi, fallback.Obi)

	data.Label = valueOr(data.Label, fallback.Label)
	data.Series = valueOr(data.Series, fallback.Series)
	data.Volume = valueOr(data.Volume, fallback.Volume)
	data.Edition = valueOr(data.Edition, fallback.Edition)
	data.Number = valueOr(data.Number, fallback.Number)

	data.Genre = valueOr(data.Genre, fallback.Genre)
	data.Caution = valueOr(data.Caution, fallback.Caution)
	data.Speech = valueOr(data.Speech, fallback.Speech)
	data.GagLeft = valueOr(data.GagLeft, fallback.GagLeft)
	data.GagRight = valueOr(data.GagRight, fallback.GagRight)
	if len(data.Receipt) == 0 {
		data.Receipt = fallback.Receipt
	}
	data.ObiBadge = valueOr(data.ObiBadge, fallback.ObiBadge)
	data.Illustrator = valueOr(data.Illustrator, fallback.Illustrator)
	data.ArtCaption = valueOr(data.ArtCaption, fallback.ArtCaption)
	data.Epigraph = valuesOr(data.Epigraph, fallback.Epigraph)
	data.Hanko = valueOr(data.Hanko, fallback.Hanko)
	data.Pills = valuesOr(data.Pills, fallback.Pills)
	data.Kicker = valueOr(data.Kicker, fallback.Kicker)
	data.Badges = valuesOr(data.Badges, fallback.Badges)

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

// joinLines はタイトルの後続行などを区切りなしで連結する。
func joinLines(values []string) string {
	return strings.Join(values, "")
}
