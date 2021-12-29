package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type DBUserInfo struct {
	ID        uint   `gorm:"primary_key"`
	Mail      string `gorm:"type:varchar(127);unique_index"`
	Nick      string `gorm:"type:varchar(63)"`
	Face      string `gorm:"size:255"`
	Role      string `gorm:"size:31;default:'guest'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DBUserInfo) TableName() string {
	return "user_info"
}

type DBTetsInfo struct {
}

func (DBTetsInfo) TableName() string {
	return "test_info"
}

func DBInitialize(engine *gorm.DB) {
	// 使用gorm自动创建表结构时,定义的结构的字段一定要大写
	engine.AutoMigrate(&DBUserInfo{})
}
