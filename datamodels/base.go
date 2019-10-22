/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:13 2019-10-21
 */
package datamodels

type Model struct {
	ID        int `orm:"column(id);pk,auto" description:"id"`
	CreatedAt int `orm:"column(created_at);default(0);index" description:"create time"`
	UpdatedAt int `orm:"column(updated_at);default(0);index" description:"update time"`
	DeletedAt int `orm:"column(deleted_at);default(0);index" description:"del time"`
}
