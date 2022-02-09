package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"unicode/utf8"
)

func main() {
	createPdfFile()
}

func writeText(pdf *gopdf.GoPdf, x float64, y float64, text string, fontSize int) {
	pdf.SetFont("s", "", fontSize)
	pdf.SetX(x)
	pdf.SetY(y)
	err := pdf.Cell(nil, text)
	if err != nil {
		return
	}
}

func convertPdfToJpg() {
	panic("Not implemented!")
}

// 文字列をn文字ずつ改行して書き込む
func split(pdf *gopdf.GoPdf, x int, startY int, fontsize int, space int, len int, text string) {
	if utf8.RuneCountInString(text) > 420 {
		fmt.Println("!表示可能文字数を超えています")
		return
	}
	tmp := 0
	for {
		if utf8.RuneCountInString(text) < len {
			writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), text, fontsize)
			break
		}
		writeText(pdf, float64(x), float64(startY+(fontsize+space)*tmp), string([]rune(text)[0:len]), fontsize)
		text = string([]rune(text)[len:])
		tmp++
	}
}

func createPdfFile() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	_ = pdf.AddTTFFont("s", "./MPLUS1p-Regular.ttf")
	template := pdf.ImportPage("./template.pdf", 1, "/MediaBox")

	pdf.AddPage()
	pdf.UseImportedTemplate(template, 0, 0, 595, 842)
	pdf.Image("./AP_未定-1.jpg", 100, 50, &gopdf.Rect{H: 280, W: 400})
	writeText(&pdf, 100, 350, "ゲーム部門", 15)
	writeText(&pdf, 100, 375, "作品名: ＊＊＊＊＊＊＊＊", 24)
	writeText(&pdf, 100, 410, "製作者: ＊＊＊＊＊＊＊2年 ＊＊ ＊＊", 15)
	split(&pdf, 100, 435, 15, 10, 30, "われらとわれらの子孫のために、諸国民との協和による成果と、わが国全土にわたつて自由のもたらす恵沢を確保し、政府の行為によつて再び戦争の惨禍が起ることのないやうにすることを決意し、ここに主権が国民に存することを宣言し、この憲法を確定する。そもそも国政は、国民の厳粛な信託によるものであつて、その権威は国民に由来し、その権力は国民の代表者がこれを行使し、その福利は国民がこれを享受する。これは人類普遍の原理であり、この憲法は、かかる原理に基くものである。われらは、これに反する一切の憲法、法令及び詔勅を排除する。\n　日本国民は、恒久の平和を念願し、人間相互の関係を支配する崇高な理想を深く自覚するのであつて、平和を愛する諸国民の公正と信義に信頼して、われらの安全と生存を保持しよは、正当に選挙された国会における代表者を通じて行動し、われらとわれらの子孫のためにに選挙された国会における代表者を通じて行動し、われらとわれらの子孫のために、諸00")
	writeText(&pdf, 20, 810, "1", 20)

	pdf.WritePdf("hello.pdf")
}
