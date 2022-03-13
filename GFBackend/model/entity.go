package model

type User struct {
	ID         int
	Username   string
	Password   string
	Salt       string
	Nickname   string
	Birthday   string
	Gender     string
	Department string
}

func (u User) TableName() string {
	return "User"
}

type Follow struct {
	Followee   string
	Follower   string
	Create_Day string
}

func (follow Follow) TableName() string {
	return "Follow"
}

type Community struct {
	ID          int
	Creator     string
	Name        string
	Description string
	Num_Member  int
	Create_Time string
}

func (c Community) TableName() string {
	return "Community"
}

type Space struct {
	ID       int
	Username string
	Capacity float64
	Used     float64
}

func (space Space) TableName() string {
	return "Space"
}

type ArticleType struct {
	ID          int
	TypeName    string
	Description string
	Creator     string
	Create_Day  string
}

func (articleType ArticleType) TableName() string {
	return "Article_Type"
}
