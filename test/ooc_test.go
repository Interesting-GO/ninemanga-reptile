/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:34 2019-10-11
 */
package test

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"ninemanga-reptile/utils"
	"testing"
)

func TestCol(t *testing.T) {
	url := "http://www.ninemanga.com/category/index_1.html"
	// 下载网页
	homehtml := utils.Dow(url)
	clog.Println(url + "下载完毕")

	if homehtml == nil {
		return
	}

	document, e := goquery.NewDocumentFromReader(bytes.NewReader(homehtml))
	if e != nil {
		panic(e.Error())
	}

	document.Find("dl.bookinfo").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a.bookname").Each(func(i int, selection *goquery.Selection) {
			val, exists := selection.Attr("href")

			if exists {
				log.Println(val)
			}
		})

		text := selection.Find("dd").Find("span").Text()
		log.Print(text)
	})
}
