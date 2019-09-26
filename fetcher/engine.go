/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:39 2019-09-24
 */
package fetcher

import (
	"ninemanga-reptile/reptile"
)

// 分发url
type ParserUrl interface {
	ParserUrl(chan interface{})
}

// 解析url
type ParserUrlItem interface {
	ParserUrlItem(chan interface{},chan interface{})
}

type End interface {
	End(chan interface{},chan int)
}

type Reptile struct {
	churl chan interface{}   // url 主页url任务   // 用于生产任务页面
	chin1 chan interface{}   // item 解析
	chin2 chan interface{}   // 二级 内容页面      // 用于解析动漫页面  获取 动漫 单个动漫的list 和动漫的详情 然后入库
	chin3 chan interface{}   // 三级 用于下载动漫   // 下载动漫然后 入库
	end chan interface{}

	ParserUrl ParserUrl
	ParserUrlHome ParserUrlItem
	ParserItem ParserUrlItem
	ParserItemIn ParserUrlItem
	ParserImg End
}

// 中央控制
func ReptileEngine() {
	url := reptile.GenerateUrl{}    // url 生成器
	urlHome := reptile.ParserHome{} // 解析二级url
	in := reptile.ParserHomeIn{}    // 三级解析item
	item := reptile.DowItem{}       // 四级下载 chrome解析流量
	img := reptile.DowImg{}         // 下载图片

	end := make(chan int)

	i := Reptile{
		churl:make(chan interface{},15),
		chin1:make(chan interface{},15),
		chin2:make(chan interface{},20),
		chin3:make(chan interface{},30),
		ParserUrl:&url,
		ParserUrlHome:&urlHome,
		ParserItem:&in,
		ParserItemIn:&item,
		ParserImg:&img,
	}


	go i.ParserUrl.ParserUrl(i.churl)
	go i.ParserUrlHome.ParserUrlItem(i.churl,i.chin1)
	go i.ParserItem.ParserUrlItem(i.chin1,i.chin2)
	go i.ParserItemIn.ParserUrlItem(i.chin2,i.chin3)
	go i.ParserImg.End(i.chin3,end)

	<-end

}