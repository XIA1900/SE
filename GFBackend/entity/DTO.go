package entity

type ResponseMsg struct {
	Code    int    `form:"Code" json:"code" example:"200"`
	Message string `form:"Message" json:"message" example:"process successfully"`
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
}

type CommunityNameFuzzyMatch struct {
	Name     string `form:"Name" json:"Name" example:"community1"`
	PageNO   int    `form:"PageNO" json:"PageNO" example:1`
	PageSize int    `form:"PageSize" json:"PageSize" example:5`
}

type CommunitiesInfo struct {
	PageNO      int         `form:"PageNO" json:"PageNO" example:1`
	PageSize    int         `form:"PageSize" json:"PageSize" example:5`
	TotalPageNO int64       `form:"TotalPageNO" json:"TotalPageNO" example:5`
	Communities []Community `form:"Communities" json:"Communities"`
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
}

type ArticleOfES struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Title    string `json:"Title"`
	Content  string `json:"Content"`
}

type ArticleInfo struct {
	Title       string `form:"Title" json:"Title" example:"Gator Forum"`
	TypeID      int    `form:"TypeID" json:"TypeID" example:"12"`
	CommunityID int    `form:"CommunityID" json:"CommunityID" example:"12"`
	Content     string `form:"Content" json:"Content" example:"I love UF"`
}