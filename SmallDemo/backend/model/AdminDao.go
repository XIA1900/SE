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

func GetAdminInfo(admin Admin) Admin {
	db := getDB()
	dbAdmin := Admin{}
	db.Take(&dbAdmin, "username = ?", admin.USERNAME)
	return dbAdmin
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
