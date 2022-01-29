package model

type User struct {
	ID       uint
	USERNAME string
	PASSWORD string
}

func (u User) TableName() string {
	return "TEST"
}

func AddUser(user User) {
	db := getDB()
	db.Create(&user)
}

func UpdateUser(user User) {
	db := getDB()
	db.Save(&user)
}
