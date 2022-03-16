package pdfmake

import (
	"fmt"
	"github.com/signintech/gopdf"
	"strings"
	"unicode/utf8"
)

func writeText(pdf *gopdf.GoPdf, x float64, y float64, text string, fontSize int) {
	err := pdf.SetFont("s", "", fontSize)
	if err != nil {
		return
	}
	pdf.SetX(x)
	pdf.SetY(y)
	err = pdf.Cell(nil, text)
	if err != nil {
		return
	}
}

// 文字列をn文字ずつ改行して書き込む
func split(pdf *gopdf.GoPdf, x int, startY int, fontsize int, space int, length int, text string) int {
	text = strings.Replace(text, "\n", "", -1)

	if utf8.RuneCountInString(text) > 350 || utf8.RuneCountInString(text) < length {
		fmt.Println(utf8.RuneCountInString(text), length)
		writeText(pdf, float64(x), float64(startY), text, fontsize)
		return startY
	}
	tmp := 0
	for {
		// これ以上分割できない場合
		if utf8.RuneCountInString(text) <= length {
			writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), text, fontsize)
			break
		} else {
			if isIllegalChars(length, []rune(text)) != 0 {
				if isIllegalChars(length, []rune(text)) == 2 {
					// 2のとき
					writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), string([]rune(text)[0:length+1]), fontsize)
					fmt.Printf("⤵行末\n")
					text = string([]rune(text)[length+1:])
				} else if isIllegalChars(length, []rune(text)) == 1 {
					// 1のとき
					writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), string([]rune(text)[0:length-1]), fontsize)
					fmt.Printf("⤵行頭\n")
					text = string([]rune(text)[length-1:])

				} else {
					fmt.Printf("%v done\n", string([]rune(text)[0:length+2]))
					writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), string([]rune(text)[0:length+1]), fontsize)
					text = string([]rune(text)[length+1:])
				}
			} else {
				// 0の時
				writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), string([]rune(text)[0:length]), fontsize)
				//fmt.Printf("⤵\n", string([]rune(text)[0:length]))
				text = string([]rune(text)[length:])
			}
			tmp++
		}
	}
	return startY + ((fontsize + space) * tmp)
}

// 禁則文字(Illegal Character)の処理
func isIllegalChars(cutLength int, text []rune) int {
	firstIllegalChars := []rune("、。,.，．・：？！ー」』】)}]") // 先頭に来てはいけない ≒ len(text)+1がこれではいけない -> len(text)で改行
	lastIllegalChars := []rune("({[（〔［｛〈《「『【￥＄")     // 最後に来てはいけない ≒ len(text) がこれらではいけない -> len(text)+1の後に改行

	if 31 >= utf8.RuneCountInString(string(text)) || 61 >= utf8.RuneCountInString(string(text)) {
		fmt.Println(string(text))
		return 3
	}

	for i := 0; i < 17; i++ {
		if i < 14 {
			if text[cutLength+1] == lastIllegalChars[i] {
				fmt.Println(string(text[cutLength-1 : cutLength+2]))
				return 1 // 行頭が禁則文字
			}
		}
		if text[cutLength] == firstIllegalChars[i] {
			fmt.Println(string(text[cutLength-1 : cutLength+2]))
			return 2 // 行末が禁則文字
		}

	}
	return 0 // 禁則処理必要なし
}
