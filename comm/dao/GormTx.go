package dao

import (
	"github.com/jinzhu/gorm"
)

type GormTx struct {
	TxObj    *gorm.DB
	Rollback bool
}

/*
 * @brief 构造事务对象,默认回滚
 *
 * @param db 连接池
 *
 * @return 数据库事务对象
 */
func NewGormTx(db *gorm.DB) (*GormTx) {
	res := new(GormTx)
	res.Rollback = true
	res.TxObj = db.Begin()

	return res
}

/*
 * @brief 设置是否回滚状态
 *
 * @param roll true:回滚；false：提交；
 */
func (this *GormTx) SetRollback(roll bool) {
	this.Rollback = roll
}

/*
 * @brief 事务完成；根据设置回滚或者提交;一般在defer中使用；
 */
func (this *GormTx) Finished() {
	if this.Rollback {
		this.TxObj.Rollback()
	} else {
		this.TxObj.Commit()
	}
}
