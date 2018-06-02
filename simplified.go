// 一つのファイルだけでicanhazwordzを解ける例文のプログラム。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	dictionaryFile = flag.String("dictionary", "/usr/share/dict/words", "Dictionary file to use")
	maxLetters     = flag.Int("max_letters", 16, "How many letters are provided on board?")
	minLetters     = flag.Int("min_letters", 3, "Minimum number of letters for an acceptable ansewr.")
)

func main() {
	flag.Parse()                            // コマンドラインのフラグをパース
	dict := loadDictionary(*dictionaryFile) // 辞書を読み込む

	// sortedDictに「文字をソートされた単語→元の単語」のmapを用意する
	sortedDict := make(map[string]string)
	for _, word := range dict {
		sortedDict[sortString(word)] = word
	}

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter available letters: ")
		if !input.Scan() || input.Text() == "" { // 標準入力から1行を読み込んでみる
			return // 入力がなかったら終了する
		}
		letters := sortString(strings.ToUpper(input.Text())) // 大文字にして、アルファベット順に並び直す

		longest := "" // longestに今まで見た最も長い単語を入れておく
		best := ""    // bestに今まで見た最も点数が高い単語を入れておく

		for sorted, word := range sortedDict { // 辞書の全部の言葉に対して
			if !isSubsequence(sorted, letters) { // 与えられたlettersがsorted単語のsubsequenceになってなければ、次の単語に進む。
				continue
			}
			if len(longest) < len(word) { // もしこれが今まで最も長かったら
				longest = word // 結果として保存しておく。
			}
			if Score(best) < Score(word) { // もしこれが今まで最も点数が高かったら
				best = word // 結果として保存しておく。
			}
		}
		fmt.Println("Longest word        : ", Denormalize(longest))
		fmt.Println("Highest scoring word: ", Denormalize(best))
	}
}

func sortString(word string) string {
	letters := strings.Split(word, "") // stringから[]stringにばらして
	sort.Strings(letters)              // アルファベット順にソートして
	return strings.Join(letters, "")   // 一つのstringに戻す
}

// もしneedleがhaystackのsupersequenceになっていればtrueを返す。
// needleとhaystackの両方が既にsortされていなければいけません。
func isSubsequence(needle, haystack string) bool {
	// ni (needle iterator)とhi (haystack iterator)を0にから始まる。
	ni, hi := 0, 0
	for ni < len(needle) && hi < len(haystack) { // どっちかが最後まで進んだら止まる
		switch {
		case needle[ni] == haystack[hi]: // この文字は両方のstringに入っている。
			ni++
			hi++
		case needle[ni] < haystack[hi]:
			// needleにhaystackにはない文字が入っている。
			return false
		default:
			// haystackにneedleにはない文字が入っているので、haystackの1文字を飛ばす。
			hi++
		}
	}
	// needleの最後まで進んだたら、ちゃんとhaystackにneedleがsubstringとして入ってありました。
	return ni == len(needle)
}

var (
	qFixer   = strings.NewReplacer("QU", "Q")
	qUnfixer = strings.NewReplacer("Q", "Qu")
	wordRE   = regexp.MustCompile(`^(?i:(?:[a-pr-z])|qu){3,}$`)
)

func Normalize(word string) string {
	return qFixer.Replace(strings.ToUpper(word)) // 大文字にして、QUをQにする
}

func Denormalize(word string) string {
	return qUnfixer.Replace(Normalize(word)) // QをQUに戻す
}

func loadDictionary(filename string) []string {
	var dict []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open %v: %v", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if !wordRE.MatchString(word) { // もし変な使えない文字があったら飛ばす。
			continue
		}
		norm := Normalize(word) // 大文字にしてQUをQにする。

		if len(norm) > *maxLetters || len(norm) < *minLetters { // もし長過ぎたり短すぎたりしたら飛ばす。
			continue
		}
		dict = append(dict, norm) // 辞書に追加しておく。
	}
	log.Printf("loaded %v words", len(dict))
	return dict
}

// 点数の付け方のためのもの。点数を気にしないなら無視してもいいです。
type PointMap map[rune]int

var LetterPoints PointMap

func init() {
	LetterPoints = make(PointMap)
	// まず全部の文字を1点として初期化
	for c := 'A'; c <= 'Z'; c++ {
		LetterPoints[c] = 1
	}
	// 特別な文字により高い点数を付ける。
	LetterPoints.assignPoints("LCFHMPVWY", 2)
	LetterPoints.assignPoints("JKQXZ", 3)
}

func (m PointMap) assignPoints(letters string, points int) {
	for _, l := range letters {
		m[l] = points
	}
}

func Score(word string) int {
	if word == "" {
		return 0
	}
	score := 1
	for _, l := range word {
		score += LetterPoints[l]
	}
	return score * score
}
