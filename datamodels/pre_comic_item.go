/**
 * @Author: DollarKiller
 * @Description: comic item
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:37 2019-10-21
 */
package datamodels

type PreComicItem struct {
	ID         int       `orm:"column(id);pk,auto" description:"id"`
	CreatedAt  int       `orm:"column(created_at);default(0);index" description:"create time"`
	UpdatedAt  int       `orm:"column(updated_at);default(0);index" description:"update time"`
	DeletedAt  int       `orm:"column(deleted_at);default(0);index" description:"del time"`
	Name       string    `orm:"column(name);size(255);default("")" description:"章节名称"`
	Url        string    `orm:"column(url);size(255);default("")" description:"url"`
	Collection int       `orm:"column(collection);default(0)" description:"集"`
	CartoonId  int       `orm:"column(cartoon_id);default(0)" description:"漫画id"`
	Content    string    `orm:"column(content)" description:"图片主体,JSON"`
	Read       int       `orm:"column(read);default(0)" description:"阅读量"`
	PreComic   *PreComic `orm:"rel(fk)"`
}
