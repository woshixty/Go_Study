package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := "root:Flzx3qcYsyhl9t@tcp(119.45.4.14:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error")
	}

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user) // 通过数据的指针来创建
	// 返回插入数据的主键
	fmt.Println(user.ID)
	// 返回 error
	fmt.Println(result.Error)
	// 返回插入记录的条数
	fmt.Println(result.RowsAffected)

	var find User
	db.First(&find)
	fmt.Println(find)

	res := map[string]interface{}{
		"hello": "Hello",
	}
	db.Model(&User{}).First(&res)
	fmt.Println(res["hello"])
}
