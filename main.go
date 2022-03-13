package main

/*
	Comfes-pamphlet-maker
	(C) 2022 Tatsuto Yamamoto [National institute of technology, Matsue College]
	This source code is under licensed under MIT license.
	See LICENSE file.
*/

import (
	"comfes-pamphlet-maker/pdfmake"
	"comfes-pamphlet-maker/utils"
)

func main() {
	file, err := utils.GetWorksFromExcelFile()
	if err != nil {
		return
	}
	//marshal, _ := json.Marshal(file)
	//fmt.Printf("%s", string(marshal))
	pdfmake.CreatePdfFile(file)
}
