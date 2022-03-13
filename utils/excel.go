package utils

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"regexp"
	"strconv"
)

var ints = []int{0, 0, 0}

type Work struct {
	ID          string
	Dept        string
	Title       string
	Author      string
	School      string
	Description string
	Links       [][]string
}

func GetWorksFromExcelFile() ([]Work, error) {
	f, err := excelize.OpenFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return []Work{}, errors.New("FileOpenError")
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Sheet のすべてのセルを取得
	rows, err := f.GetRows("フォームの回答 1")
	if err != nil {
		fmt.Println(err)
		return []Work{}, errors.New("ExcelFileOpenError")
	}

	Works := make([]Work, 0)
	for i, row := range rows {
		fmt.Printf("Excel: %d / %d\r", i, len(rows))
		tmp := Work{}
		for j, colCell := range row {
			if i < 1 {
				continue
			}

			switch j {
			case 0:
				break
			case 1:
				tmp.School = colCell
				break
			case 2:
				tmp.Dept = colCell
				break
			case 3:
				tmp.Author = colCell
				break
			case 4:
				r := regexp.MustCompile(`(（|\()[ぁ-ん].*(\)|）)$`)
				rc := r.ReplaceAllString(colCell, "")
				tmp.Title = rc
				break
			case 5:
				tmp.Description = colCell
				break
			case 6:
				r := regexp.MustCompile(`https?://[\w!\?/\+\-_~=;\.:,\*&@#\$%\(\)'\[\]]+`)
				rc := r.FindAllStringSubmatch(colCell, -1)
				//fmt.Printf("%v\n", rc)
				tmp.Links = rc
				break
			}
		}
		if i < 1 {
			continue
		}
		tmp.ID = makeId(tmp.Dept)
		Works = append(Works, tmp)
		//fmt.Println()
	}
	return Works, nil
}

func makeId(dept string) string {
	switch dept {
	case "アプリケーション部門":
		ints[0]++
		return "AP" + strconv.Itoa(ints[0])
	case "ゲーム部門":
		ints[1]++
		return "GM" + strconv.Itoa(ints[1])
	case "メディアコンテンツ部門":
		ints[2]++
		return "MC" + strconv.Itoa(ints[2])
	}
	return ""
}
