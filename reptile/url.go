/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:08 2019-09-24
 */
package reptile

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/clog"
)

type GenerateUrl struct {
}

// url 生成器
func (p *GenerateUrl) ParserUrl(url chan interface{}) {
	baseUrl := "http://es.ninemanga.com/category/index_%v.html"
	for i := 2; i <= 100; i++ {
		spr := fmt.Sprintf(baseUrl, i)
		url <- spr
	}
	// 生成完毕关掉chan
	clog.Println("第一阶段完毕")
	close(url)
}
