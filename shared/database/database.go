package database

import (
	"fmt"
	"rest-api/entity"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

func Getconnection() *Database {

	db, err := gorm.Open("mysql", "root:Dev1234!@tcp(localhost:3306)/dblocal")
	if err != nil {
		fmt.Println("gagal conn")
		panic(err)

	}
	fmt.Println("seccess CONNNNNN")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	//maksimal idle selama 5 menit setelah 5 menit akan di close
	db.DB().SetConnMaxIdleTime(5 * time.Minute)
	//setelah 60 menit akan di buka koneksi lagi
	db.DB().SetConnMaxLifetime(60 * time.Minute)
	Database{db}.LogMode(true)

	return &Database{db}
}

func Migrate(db *gorm.DB) {
	//mapping entity for autogenerate table
	db.AutoMigrate(&entity.Customer{})
	fmt.Println("Table migrated")
}
