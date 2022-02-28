package entitas

type Person struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id" `
	FirstName string `gorm:"type:varchar(200)" json:"first_name"`
	LastName  string `gorm:"type:varchar(200)" json:"last_name"`
	Username  string `gorm:"type:varchar(200)" json:"username"`
	Password  string `gorm:"type:varchar(200)" json:"password"`
}
