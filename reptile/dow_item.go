/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:56 2019-09-24
 */
package reptile

import (
	"ninemanga-reptile/defs"
	"sync"
)

type DowItem struct {
}

func (d *DowItem) ParserUrl(ch1 chan interface{}) {
cc:
	for {
		select {
		case val, ok := <-ch1:
			if ok {

				// 开启多协程
				numch := make(chan int, 10)
				gr := sync.WaitGroup{}
				gr.Add(1)
				numch <- 1

				go func(ur interface{}) {
					defer func() {
						gr.Done()
						<-numch
					}()
					url := ur.(string)
					d.logic(url)
				}(val)

				gr.Wait()

			} else {
				break cc
			}
		}
	}
}

func (d *DowItem) logic (dat1 interface{}) {
	data := dat1.(defs.DowItem)



}
