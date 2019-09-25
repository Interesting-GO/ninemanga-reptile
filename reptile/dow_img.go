/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:24 2019-09-25
 */
package reptile

import (
	"encoding/json"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"io/ioutil"
	"ninemanga-reptile/datamodels"
	"ninemanga-reptile/datasources/mysql_conn"
	"ninemanga-reptile/defs"
	"ninemanga-reptile/utils"
	"strconv"
	"sync"
)

type DowImg struct {
}

func (d *DowImg) ParserUrlItem(ch1 chan interface{},end chan interface{}) {
cc:
	for {
		select {
		case val,ok := <- ch1:
			if ok {

				// 开启多协程
				numch := make(chan int, 20)
				gr := sync.WaitGroup{}
				gr.Add(1)
				numch <- 1

				go func(ur interface{}) {
					defer func() {
						gr.Done()
						<-numch
					}()
					d.logic(ur)
				}(val)

				gr.Wait()

			}else {
				end <- 1
				break cc
			}
		}
	}
}

func (d *DowImg) logic(ur interface{}) {
	item := ur.(defs.DowImgItem)
	data := datamodels.PreCartoonItem{
		Name:strconv.Itoa(item.Num),
		Read:easyutils.Random(300,6000),
		CreateTime:easyutils.TimeGetNowTime(),
		CartoonId:item.SqlId,
		Language:"fr",
		Collection:item.Num,
	}

	path := "./img/" + strconv.Itoa(item.SqlId) + "/" + strconv.Itoa(item.Num)
	urlpath := "/img/" + strconv.Itoa(item.SqlId) + "/" + strconv.Itoa(item.Num)
	easyutils.DirPing(path)
	for i,k := range item.Context {
		s, e := easyutils.FileGetPostfix(k.Img)
		if e != nil {
			s = "jpg"
		}
		name := easyutils.SuperRand() + "." + s
		imgpath := path + "/" + name
		dow := utils.Dow(k.Img)
		e = ioutil.WriteFile(imgpath, dow, 000666)
		if e != nil {
			clog.PrintWa("====================================")
			clog.PrintWa(e)
			continue
		}

		item.Context[i].Img = urlpath + imgpath
	}

	bytes, e := json.Marshal(item.Context)
	if e != nil {
		clog.PrintWa("====================================")
		clog.PrintWa(e)
		return
	}
	data.Content = string(bytes)

	// 进行入库
	_, e = mysql_conn.MysqlEngine.InsertOne(&data)
	if e != nil {
		clog.PrintWa("====================================")
		clog.PrintWa(e)
	}
}
