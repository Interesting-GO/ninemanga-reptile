/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:18 2019-09-24
 */
package reptile

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"io/ioutil"
	"math/rand"
	"ninemanga-reptile/config"
	"ninemanga-reptile/datamodels"
	"ninemanga-reptile/datasources/mysql_conn"
	"ninemanga-reptile/defs"
	"ninemanga-reptile/utils"
	"strings"
	"sync"
)

type ParserHomeIn struct {

}

func (p *ParserHomeIn) ParserUrlItem(ch1 chan interface{},ch2 chan interface{}) {
	numch := make(chan int, 10)
	sy := sync.WaitGroup{}

cc:
	for {
		select {
		case ur,ok := <- ch1 :
			if ok {
				// 开启多协程
				numch <- 1
				sy.Add(1)

				go func(ur interface{}) {
					defer func() {
						<-numch
						sy.Done()
					}()
					url := ur.(string)
					p.logic(url, ch2)

				}(ur)


			}else {
				sy.Wait()
				clog.Println("第三阶段完毕")
				close(ch2)
				break cc
			}
		}
	}
}

func (p *ParserHomeIn) logic(url string,ch chan interface{}) {
	//url := "http://fr.ninemanga.com/manga/The+Prince%27s+Private+Child.html"
	var homehtml []byte

	// 下载网页
	//homehtml = utils.Dow(url)
	chrome := utils.StartChrome(url)
	homehtml = []byte(chrome)
	if homehtml == nil {
		panic("--")
	}


	document, e := goquery.NewDocumentFromReader(bytes.NewReader(homehtml))
	if e != nil {
		panic(e.Error())
	}

	// 找到主节点
	mast := document.Find("ul.message")

	data := datamodels.PreCartoon{
		Language:config.MyConfig.App.Language,
		State:rand.Intn(2),
		Read:easyutils.Random(300,6000),
		CreateTime:easyutils.TimeGetNowTime(),
	}

	dowdata := defs.ParserHoItem{}

	mast.Find("li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("b").Text()
		//clog.Println(text)

		// 名称
		if strings.Index(text,"Book Name:") != -1 {
			s := selection.Find("span").Text()
			data.Name = strings.TrimSpace(s)
		}

		// 分类
		if strings.Index(text,"Genre (s):") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				}else {
					tex += "," + selection.Text()
				}
				ic += 1
			})

			data.Classification = tex
		}

		// 作者
		if strings.Index(text,"Autor (en):") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				}else {
					tex += "  " + selection.Text()
				}
				ic += 1
			})

			data.Author = strings.TrimSpace(tex)
		}

		// 年代
		if strings.Index(text,"Jahr") != -1 {
			tex := ""
			ic := 0
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					tex = selection.Text()
				}else {

				}
				ic += 1
			})

			data.Year = tex
		}
	})

	// 描述
	text := document.Find("p[itemprop='description']").Text()
	text = strings.Replace(text, "\nZusammenfassung:\n", "", -1)
	data.Describe = text

	val, exists := document.Find("img[itemprop='image']").Attr("src")
	if exists {
		data.Img = val
	}

	//log.Println(data)

	// 获取需要下载的list
	document.Find("ul.sub_vol_ul").Find("li").Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Find("a").Attr("href")
		var url string
		var num int
		if exists {
			index := strings.LastIndex(val, ".html")
			url = val[:index] + "-10-$x$" + val[index:]
		}
		s, b := selection.Find("a").Attr("title")
		if b {
			num = utils.GetNum(s)
		}

		item := defs.DowItem{
			Url:url,
			Num:num,
		}

		dowdata.DowUrl = append(dowdata.DowUrl,item)
	})

	// 数据入库
	//data.Language = url
	if data.Name == "" {
		ioutil.WriteFile("err.html",homehtml,00666)
		return
	}
	_, e = mysql_conn.MysqlEngine.InsertOne(&data)
	if e != nil {
		panic(e)
	}

	// 正常入库查询 数据库id
	dat := datamodels.PreCartoon{}
	_, e = mysql_conn.MysqlEngine.Where("name = ? ", data.Name).Get(&dat)
	if e != nil {
		clog.Println(dat)
		panic(e)
	}

	dowdata.SqlId = dat.Id

	for i,k := range dowdata.DowUrl {
		if i < easyutils.Random(15,30) {
			k.SqlId = dowdata.SqlId
			// 写入
			ch <- k
		}else {
			return
		}
	}
}
