/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:39 2019-09-24
 */
package fetcher

import "ninemanga-reptile/reptile"

// 分发url
type ParserUrl interface {
	ParserUrl(chan interface{})
}

// 解析url
type ParserUrlItem interface {
	ParserUrlItem(chan interface{},chan interface{})
}

type Reptile struct {
	churl chan interface{}   // url 主页url任务   // 用于生产任务页面
	chin1 chan interface{}   // item 解析
	chin2 chan interface{}   // 二级 内容页面      // 用于解析动漫页面  获取 动漫 单个动漫的list 和动漫的详情 然后入库
	chin3 chan interface{}   // 三级 用于下载动漫   // 下载动漫然后 入库

	ParserUrl ParserUrl
	ParserUrlHome ParserUrlItem
	ParserItem ParserUrlItem
}

// 中央控制
func ReptileEngine() {
	url := reptile.GenerateUrl{}    // url 生成器
	urlhome := reptile.ParserHome{} // 解析二级url
	in := reptile.ParserHomeIn{}    // 三级解析item
	// 四级下载

	i := Reptile{
		churl:make(chan interface{},1000),
		chin1:make(chan interface{},1000),
		chin2:make(chan interface{},1000),
		chin3:make(chan interface{},1000),
		ParserUrl:&url,
		ParserUrlHome:&urlhome,
		ParserItem:&in,

	}


}