package model

import (
	"github.com/517962189/Kappa"
	"gorm.io/gorm"
)

type BaseDb struct {
}

// 获取redis db连接池对象
func (this *BaseDb) Db() *gorm.DB {
	return kappa.NewKappaServer().RegisterDbStorage("db3")
}

