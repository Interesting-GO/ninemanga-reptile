/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:08 2019-09-24
 */
package reptile

import "fmt"

type GenerateUrl struct {
}

// url 生成器
func (p *GenerateUrl) ParserUrl(url chan interface{}) {
	baseUrl := "http://fr.ninemanga.com/category/index_%v.html"
	for i := 1; i <= 110; i++ {
		spr := fmt.Sprintf(baseUrl, i)
		url <- spr
	}
	// 生成完毕关掉chan
	close(url)
}
