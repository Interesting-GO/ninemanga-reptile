/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:56 2019-09-24
 */
package reptile

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"ninemanga-reptile/defs"
	"ninemanga-reptile/utils"
	"strconv"
	"strings"
	"sync"
)

type DowItem struct {
}

func (d *DowItem) ParserUrlItem(ch1 chan interface{},ch2 chan interface{}) {
	numch := make(chan int, 10)
	sy := sync.WaitGroup{}

cc:
	for {
		select {
		case val, ok := <-ch1:
			if ok {

				// 开启多协程
				numch <- 1
				sy.Add(1)

				go func(ur interface{}) {
					defer func() {
						<-numch
						sy.Done()
					}()
					d.logic(ur,ch2)
				}(val)

			} else {
				sy.Wait()
				clog.Println("第四阶段完毕")
				close(ch2)
				break cc
			}
		}
	}
}

func (d *DowItem) logic (dat1 interface{},ch2 chan interface{}) {
	data1 := dat1.(defs.DowItem)

	url1 := strings.Replace(data1.Url, "$x$", "1", -1)

	//log.Println(url1)
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

	// #cx
	data2 := defs.DowImgItem{
		SqlId:data1.SqlId,
		Num:data1.Num,
	}

	data3 := defs.CartoonItemImg{}

	for i:=1;i<=page;i++ {
		url := strings.Replace(data1.Url,"$x$",strconv.Itoa(i),-1)

		i2 := utils.StartChrome(url)

		reader, e := goquery.NewDocumentFromReader(strings.NewReader(i2))
		if e != nil {
			panic(e)
		}

		reader.Find("center").Find("img").Each(func(i int, selection *goquery.Selection) {
			val, exists := selection.Attr("src")
			if exists {
				dataitem := defs.CartoonItemImgItem{
					Id:i,
					Img:val,
				}
				data3 = append(data3,&dataitem)
			}
		})
	}



	data2.Url = data1.Url
	data2.Context = data3

	ch2 <- data2
}
