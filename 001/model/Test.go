package models

import (
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO ドライバの指定の仕方は考える必要がある。

// Test test
type Test struct {
	gorm.Model
	Pass string
}
