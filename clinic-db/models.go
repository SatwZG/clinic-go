package clinicDB

import (
	"database/sql"
)

type Account struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	ID       int    `gorm:"column:id"`
}

type Doctor struct {
	ID int             		 	`gorm:"column:id;AUTO_INCREMENT"`
	Name string 				`gorm:"column:name"`
	Sex string 					`gorm:"column:sex"`
	Age int 					`gorm:"column:age"`
	Department string   		`gorm:"column:department"`
	Avatar sql.NullString       `gorm:"column:avatar"`
	Introduction sql.NullString `gorm:"column:introduction"`
}