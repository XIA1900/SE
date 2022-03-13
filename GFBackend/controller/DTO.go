package controller

type ResponseMsg struct {
	Code    int    `form:"Code" json:"code" example:"200"`
	Message string `form:"Message" json:"message" example:"process successfully"`
}

type CommunityResponseMsg struct {
	Code        int    `json:"code" example:"200"`
	Message     string `json:"message" example:"process successfully"`
	ID          int    `json:"id" example:"1"`
	Creator     string `json:"creator" example:"creator"`
	Name        string `json:"name" example:"name"`
	Description string `json:"description" example:"description"`
	Num_Member  int    `json:"num_member" example:"1"`
	Create_Time string `json:"create_time" example:"create_time"`
}

type UserInfo struct {
	Username    string `form:"Username" json:"Username" example:"jamesbond21" `
	Password    string `form:"Password" json:"Password" example:"f9ae5f68ae1e7f7f3fc06053e9b9b539"`
	NewPassword string `form:"NewPassword" json:"NewPassword" example:"3ecb441b741bcd433288f5557e4b9118"`
	ForAdmin    bool   `form:"ForAdmin" json:"ForAdmin" example:true`
}

type NewUserInfo struct {
	Username   string `form:"Username" json:"Username" example:"jamesbond21"`
	Nickname   string `form:"Nickname" json:"Nickname" example:"Peter Park"`
	Birthday   string `form:"Birthday" json:"Birthday" example:"2022-02-30"`
	Gender     string `form:"Gender" json:"Gender" example:"male/female/unknown"`
	Department string `form:"Department" json:"Department" example:"CS:GO"`
}

type CommunityInfo struct {
	ID          int
	Creator     string `form:"Creator" json:"Creator" example:"test1"`
	Name        string `form:"Name" json:"Name" example:"community1"`
	Description string `form:"Description" json:"Description" example:"this is a test community"`
	Create_Time string `form:"Create_Time" json:"Create_Time" example:"2020-01-01"`
}

type UserFiles struct {
	ResponseMsg
	Filenames []string `form:"Filenames" json:"Filenames" example:"\"xxx.jpg\",\"xxx.png\",\"xxx.gif\""`
}

type UserFollows struct {
	ResponseMsg
	Users []string `form:"Users" json:"Users" example:"\"spriderman\",\"batman\",\"ironman\""`
}

type ArticleTypeInfo struct {
	TypeName    string `form:"TypeName" json:"TypeName" example:"Movie"`
	Description string `form:"Description" json:"Description" example:"Discussion of Movie"`
	Creator     string `form:"Creator" json:"Creator" example:"James Bond"`
	CreateDay   string `form:"CreateDay" json:"CreateDay" example:"2020-02-02"`
}
