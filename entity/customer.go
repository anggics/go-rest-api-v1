package entity

type Customer struct {
	CustomerID int    `gorm:"primary_key;auto_increment" json:"CustomerID"`
	FirstName  string `gorm:"size:100;" json:"FirstName"`
	LastName   string `gorm:"size:100;" json:"LastName"`
}
