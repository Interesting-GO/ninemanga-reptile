/**
 * @Author: DollarKiller
 * @Description: comic 漫画
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:01 2019-10-21
 */
package datamodels

type PreComic struct {
	Model
	Language         string          `orm:"column(language);size(255);default("")" description:"语言"`
	Name             string          `orm:"column(name);size(255);default("")" description:"漫画名称"`
	Img              string          `orm:"column(img);size(666);default("")" description:"首图img"`
	Year             string          `orm:"column(year);size(255);default("")" description:"年代"`
	State            int             `orm:"column(state);default(0)" description:"状态0 完结 1 更新中"`
	Author           string          `orm:"column(author);size(255);default("")" description:"作者"`
	Describe         string          `orm:"column(describe);default("")" description:"描述"`
	ClassificationId int             `orm:"column(classification_id);default(0)" description:"分类id"`
	Read             int             `orm:"column(read);default(0)" description:"阅读量"`
	PreComicItem     []*PreComicItem `orm:"reverse(many)"`
}
