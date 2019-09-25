/**
 * @Author: DollarKiller
 * @Description:  解析test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:32 2019-09-24
 */
package test

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"io/ioutil"
	"log"
	"math/rand"
	"ninemanga-reptile/datamodels"
	"ninemanga-reptile/defs"
	"ninemanga-reptile/utils"
	"strconv"
	"strings"
	"testing"
)

func TestParseHome(t *testing.T) {
	url := "http://fr.ninemanga.com/category/index_2.html"
	var homehtml []byte

	// 下载网页
	homehtml = utils.Dow(url)
	if homehtml == nil {
		panic("--")
	}

	//log.Println(string(homehtml))

	document, e := goquery.NewDocumentFromReader(bytes.NewReader(homehtml))
	if e != nil {
		panic(e.Error())
	}

	document.Find("dl.bookinfo").Each(func(i int, selection *goquery.Selection) {
		selection.Find("a.bookname").Each(func(i int, selection *goquery.Selection) {
			val, exists := selection.Attr("href")
			if exists {
				clog.Println(val)
			}
		})
	})
}

func TestParserPHome(t *testing.T) {
	//url := "http://fr.ninemanga.com/manga/The+Prince%27s+Private+Child.html"
	var homehtml []byte

	// 下载网页
	//homehtml = utils.Dow(url)
	//if homehtml == nil {
	//	panic("--")
	//}
	homehtml, err := ioutil.ReadFile("home.html")
	if err != nil {
		panic(err)
	}

	document, e := goquery.NewDocumentFromReader(bytes.NewReader(homehtml))
	if e != nil {
		panic(e.Error())
	}

	// 找到主节点
	mast := document.Find("ul.message")

	data := datamodels.PreCartoon{
		Language:   "fr",
		State:      rand.Intn(2),
		Read:       easyutils.Random(300, 6000),
		CreateTime: easyutils.TimeGetNowTime(),
	}

	dowdata := defs.ParserHoItem{}

	mast.Find("li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("b").Text()
		//clog.Println(text)

		// 名称
		if strings.Index(text, "Nom du livre:") != -1 {
			s := selection.Find("span").Text()
			data.Name = s
		}

		// 分类
		if strings.Index(text, "Genre(s):") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				} else {
					tex += "," + selection.Text()
				}
				ic += 1
			})

			data.Classification = tex
		}

		// 作者
		if strings.Index(text, "Auteur(s):") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				} else {
					tex += "  " + selection.Text()
				}
				ic += 1
			})

			data.Author = tex
		}

		// 年代
		if strings.Index(text, "Année") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				} else {

				}
				ic += 1
			})

			data.Year = tex
		}
	})

	// 描述
	text := document.Find("p[itemprop='description']").Text()
	text = strings.Replace(text, "\nRésumé:\n", "", -1)
	data.Describe = text

	//log.Println(data)

	// 获取需要下载的list
	document.Find("ul.sub_vol_ul").Find("li").Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Find("a").Attr("href")
		var url string
		var num int
		if exists {
			index := strings.LastIndex(val, ".html")
			url = val[:index] + "-10-1" + val[index:]
		}
		s, b := selection.Find("a").Attr("title")
		if b {
			num = utils.GetNum(s)
		}

		item := defs.DowItem{
			Url: url,
			Num: num,
		}

		dowdata.DowUrl = append(dowdata.DowUrl, item)
	})
}

func TestDow(t *testing.T) {
	item := defs.DowItem{SqlId: 1, Url: "http://fr.ninemanga.com/chapter/The%20Prince%27s%20Private%20Child/16548-10-$x$.html", Num: 1}

	url1 := strings.Replace(item.Url, "$x$", "1", -1)

	log.Println(url1)
	chrome := utils.StartChrome(url1)

	if chrome == "" {
		return
	}

	document, e := goquery.NewDocumentFromReader(strings.NewReader(chrome))
	if e != nil {
		panic(e.Error())
	}

	var page int
	document.Find("select#page").Find("option").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		page = utils.GetPage(text)
		return
	})

	data := datamodels.PreCartoonItem{
		Read:easyutils.Random(300,3000),
		CreateTime:easyutils.TimeGetNowTime(),
		Language:"fr",
	}

	data = data

	for i:=1;i<=page;i++ {
		url := strings.Replace(item.Url,"$x$",strconv.Itoa(i),-1)

		i2 := utils.StartChrome(url)

		reader, e := goquery.NewDocumentFromReader(strings.NewReader(i2))
		if e != nil {
			panic(e)
		}

		reader.Find("center").Find("img").Each(func(i int, selection *goquery.Selection) {
			val, exists := selection.Attr("src")
			if exists {
				log.Println(val)
			}
		})

	}

}
