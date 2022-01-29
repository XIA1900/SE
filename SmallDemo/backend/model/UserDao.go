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
	db.Model(&user).Where("USERNAME", user.USERNAME).Update("PASSWORD", user.PASSWORD)
}

func DeleteUser(user User) {
	db := getDB()
	db.Where("USERNAME", user.USERNAME).Delete(&user)
}
