package datamodels

type PreCartoon struct {
	Id               int    `xorm:"not null pk autoincr INT(10)"`
	Name             string `xorm:"not null default '' VARCHAR(366)" form:"name" binding:"required"`
	Language         string `xorm:"not null comment('语言 简称') index VARCHAR(255)" form:"language"`
	Img              string `xorm:"not null default '' VARCHAR(366)" form:"img" binding:"required"`
	State            int    `xorm:"not null default 0 comment('状态0连载中1完结2停更') TINYINT(3)" form:"state"`
	Author           string `xorm:"not null default '' comment('作者') VARCHAR(366)" form:"author" binding:"required"`
	Describe         string `xorm:"comment('描述') TEXT" form:"describe" binding:"required"`
	ClassificationId int    `xorm:"not null default 0 comment('分类id') index INT(10)" form:"classification_id" binding:"required"`
	Classification   string `xorm:"comment('分类 str') TEXT" form:"classification_id" binding:"required"`
	Hot              string `xorm:"not null comment('hot') index VARCHAR(255)"`
	Year             string `xorm:"not null comment('year') index VARCHAR(255)" form:"year"`
	Read             int    `xorm:"not null default 0 comment('阅读量') INT(10)"`
	CreateTime       int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime       int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	DeleteTime       int    `xorm:"not null default 0 comment('软删除') index INT(10)"`
	Time             string `xorm:"-"`
}
