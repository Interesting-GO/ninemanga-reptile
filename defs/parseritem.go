/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:47 2019-09-24
 */
package defs

type ParserHoItem struct {
	SqlId  int
	DowUrl []DowItem
}

type DowItem struct {
	SqlId int
	Url   string
	Num   int
}

type DowImgItem struct {
	SqlId   int
	Context CartoonItemImg
	Num     int
	Url     string
}

// item

type CartoonItemImg []*CartoonItemImgItem

type CartoonItemImgItem struct {
	Id  int    `json:"id"`
	Img string `json:"img"`
}

type Hc struct {
	Url  string `json:"url"`
	View string `json:"view"`
}
