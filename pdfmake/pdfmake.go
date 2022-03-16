package pdfmake

import (
	"comfes-pamphlet-maker/utils"
	"fmt"
	"github.com/signintech/gopdf"
	"strconv"
)

type PDF struct {
	pdf      *gopdf.GoPdf
	template int
}

func CreatePdfFile(Works []utils.Work) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	err := pdf.AddTTFFont("s", "./fonts/MPLUS1p-Regular.ttf")
	if err != nil {
		fmt.Printf("Err: failed to load font file")
		return
	}
	template := pdf.ImportPage("./template.pdf", 1, "/MediaBox")

	createWorkPage(PDF{&pdf, template}, Works)

	pdf.WritePdf("hello.pdf")
}

func createWorkPage(pdf PDF, Works []utils.Work) {
	//last := len(Works) + 1
	for j := range Works {
		work := Works[j]
		pdf.pdf.AddPage()
		if j == 0 {
			t := pdf.pdf.ImportPage("./top.pdf", 1, "/MediaBox")
			pdf.pdf.UseImportedTemplate(t, 0, 0, 595, 842)
			continue
		}

		//fmt.Printf("Page: %d / %d\n", j, len(Works))

		pdf.pdf.UseImportedTemplate(pdf.template, 0, 0, 595, 842)
		err := pdf.pdf.Image("./data/"+work.ID+".jpg", 100, 30, &gopdf.Rect{H: 280, W: 400})
		if err != nil {

		}
		writeText(pdf.pdf, 70, 320, work.Dept+" (作品ID: "+work.ID+")", 15)
		y := split(pdf.pdf, 70, 355, 22, 10, 25, "作品名: "+work.Title)
		y = split(pdf.pdf, 70, int(float64(y)+22+10), 10, 10, 60, "製作者: "+work.Author)
		y = split(pdf.pdf, 70, int(float64(y)+20), 15, 10, 30, work.Description)
		writeText(pdf.pdf, 20, 810, strconv.Itoa(j), 20)

		//writeText(pdf.pdf, 20, 0, work.ID, 15)

		for i := range work.Links {
			writeText(pdf.pdf, 70, float64(y+20+(i*20)), "リンク"+strconv.Itoa(i+1), 15)
			pdf.pdf.AddExternalLink(work.Links[i][0], 70, float64(y+20+(i*20)), 53, 15)
		}
	}
	pdf.pdf.AddPage()
	pdf.pdf.UseImportedTemplate(pdf.template, 0, 0, 595, 842)
	writeText(pdf.pdf, 50, 700, "中国地区高専コンピュータフェスティバル2022 作品集", 20)
	writeText(pdf.pdf, 50, 730, "2022年3月13日 発行", 18)
	writeText(pdf.pdf, 50, 753, "発行者: 松江工業高等専門学校 情報科学研究部", 18)

}
