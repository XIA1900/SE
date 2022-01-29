package model

type Admin struct {
	ID       uint
	USERNAME string
	PASSWORD string
	SALT     string
}

func (a *Admin) TableName() string {
	return "ADMIN"
}

func AddAdmin(admin Admin) {
	db := getDB()
	db.Create(&admin)
}

func UpdateAdmin(admin Admin) {
	db := getDB()
	db.Model(&admin).Where("USERNAME", admin.USERNAME).Update("PASSWORD", admin.PASSWORD)
}

func DeleteAdmin(admin Admin) {
	db := getDB()
	db.Where("USERNAME", admin.USERNAME).Delete(&admin)
}
