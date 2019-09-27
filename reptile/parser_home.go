/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:15 2019-09-24
 */
package reptile

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"ninemanga-reptile/utils"
	"sync"
)

// parser home
type ParserHome struct {
}

func (p *ParserHome) ParserUrlItem(ch1 chan interface{}, ch2 chan interface{}) {
	numch := make(chan int, 10)
	sy := sync.WaitGroup{}
cc:
	for {
		select {
		case ur, ok := <-ch1:
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


			} else {
				sy.Wait()
				clog.Println("第二阶段完毕")
				close(ch2)
				break cc
			}
		}
	}
}

func (p *ParserHome) logic(url string, ch chan interface{}) {
	var homehtml []byte

	// 下载网页
	homehtml = utils.Dow(url)
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
				ch <- val
			}
		})
	})
}
