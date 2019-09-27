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
	"ninemanga-reptile/config"
	"ninemanga-reptile/datamodels"
	"ninemanga-reptile/datasources/mysql_conn"
	"ninemanga-reptile/defs"
	"strconv"
	"sync"
	"time"
)

type DowImg struct {
}

func (d *DowImg) End(ch1 chan interface{}, end chan int) {
	numch := make(chan int, 20)
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
					d.logic(ur)
				}(val)

			} else {
				sy.Wait()
				time.Sleep(200 * time.Second)
				clog.Println("第五阶段完毕")
				end <- 1
				break cc
			}
		}
	}
}

func (d *DowImg) logic(ur interface{}) {
	item := ur.(defs.DowImgItem)
	data := datamodels.PreCartoonItem{
		Name:       strconv.Itoa(item.Num),
		Read:       easyutils.Random(300, 6000),
		CreateTime: easyutils.TimeGetNowTime(),
		CartoonId:  item.SqlId,
		Language:   config.MyConfig.App.Language,
		Collection: item.Num,
	}

	//path := "./img/" + strconv.Itoa(item.SqlId) + "/" + strconv.Itoa(item.Num)
	//urlpath := "/img/" + strconv.Itoa(item.SqlId) + "/" + strconv.Itoa(item.Num)
	//easyutils.DirPing(path)
	//for i,k := range item.Context {
	//	s, e := easyutils.FileGetPostfix(k.Img)
	//	if e != nil {
	//		s = "jpg"
	//	}
	//	name := easyutils.SuperRand() + "." + s
	//	imgpath := path + "/" + name
	//	dow := utils.Dow(k.Img)
	//	e = ioutil.WriteFile(imgpath, dow, 000666)
	//	if e != nil {
	//		clog.PrintWa("====================================")
	//		clog.PrintWa(e)
	//		continue
	//	}
	//
	//	item.Context[i].Img = urlpath + imgpath
	//}
	//
	//bytes, e := json.Marshal(item.Context)
	//if e != nil {
	//	clog.PrintWa("====================================")
	//	clog.PrintWa(e)
	//	return
	//}
	//data.Content = string(bytes)

	imgs := item.Context
	bytes, e := json.Marshal(imgs)
	if e != nil {
		clog.PrintWa("====================================")
		clog.PrintWa(e)
		return
	}
	data.Content = string(bytes)
	data.Url = item.Url
	// 进行入库
	_, e = mysql_conn.MysqlEngine.InsertOne(&data)
	if e != nil {
		clog.PrintWa("====================================")
		clog.PrintWa(e)
	}
}
