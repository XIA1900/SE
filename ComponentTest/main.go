package main

import (
	docs "ComponentTest/docs"
	"ComponentTest/es"
	"ComponentTest/log"
	"ComponentTest/resources"
	"ComponentTest/role"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	// logTest()

	// jwt.TestServer()

	// casbinTest()

	// ESTest()

	// FileUploadDownloadTest()

	// SwaggoTest()

	// fmt.Println(GetRandomString(6))

	// fmt.Println(EncodeInMD5("jake16"))

	// RedisTest()

	// UserTest()

	// FollowTest()

	//go TestSingleton()
	//go TestSingleton()
	//time.Sleep(time.Second)

	//StaticTest()

	// SpaceTest()

	FileRelatedOperation()
}

func FileRelatedOperation() {
	files, err := resources.GetFilesInDir("admin")
	if err != nil {
		return
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

type Space struct {
	ID        int
	Username  string
	Capacity  float32
	Remaining float32
}

func (space Space) TableName() string {
	return "Space"
}

func SpaceTest() {
	space := Space{
		Username: "jake16",
	}

	dsn := "root:admin@tcp(127.0.0.1:3306)/gf_test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Select("Username").Create(&space)

}

func StaticTest() {
	router := gin.Default()
	router.Static("/resources", "./resources")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

type SingleDog struct{}

var lock sync.Mutex
var singleDog *SingleDog

func NewSingleDog() *SingleDog {
	lock.Lock()
	if singleDog == nil {
		singleDog = new(SingleDog)
	}
	lock.Unlock()
	return singleDog
}

func TestSingleton() {
	dog := NewSingleDog()
	fmt.Printf("%p\n", &*dog)
}

type Follow struct {
	Followee   string
	Follower   string
	Create_Day string
}

func (follow Follow) TableName() string {
	return "Follow"
}

func FollowTest() {
	follow := Follow{
		Followee:   "joey",
		Follower:   "jerry",
		Create_Day: "2022-02-22",
	}

	dsn := "root:admin@tcp(127.0.0.1:3306)/gf_test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&follow)

	orgFollow := Follow{}
	db.Where("followee = ?", "tom").First(&orgFollow)
	fmt.Println(orgFollow)
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

func EmbeddedStructTest() {
	var circle Circle
	circle.X = 1
	circle.Y = 2
	circle.Radius = 3
	fmt.Println(circle)
}

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

func UserTest() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gf_test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	//date, _ := time.Parse("2006-01-02", "1997-12-16")
	//user := User{
	//	Username: "test123",
	//	Password: "test1234",
	//	Salt:     "23452",
	//	Birthday: date.Format("2006-01-02"),
	//}
	//
	//db.Select("Username", "Password", "Salt", "Birthday").Create(&user)

	orgUser := User{}
	db.Where("Username = ?", "test123").Find(&orgUser)
	fmt.Println()
}

func logTest() {
	log.InitLog(zapcore.InfoLevel)
	defer log.Logger.Sync()

	log.Logger.Info("logger", zap.String("name", "修华师6"))
	log.Logger.Error("logger", zap.String("name", "修华师7"))
	log.Logger.Debug("logger", zap.String("name", "修华师8"))
}

func casbinTest() {
	role.InitCasbin()
	e := role.GetCasbinEnforcer()
	sub := "data2_admin"
	obj := "data2"
	act := "read"
	e.AddPolicy(sub, obj, act)
	user := "jake16"
	e.AddGroupingPolicy(user, sub)
	flag, _ := e.Enforce(sub, obj, act)
	if flag == true {
		fmt.Println("pass")
	} else {
		fmt.Println("block")
	}
}

func ESTest() {
	es.InitES()
	elasticSearch := es.GetES()
	res, err := elasticSearch.Info()
	if err != nil {
		fmt.Printf("Error getting response: %s\n", err)
	}
	defer res.Body.Close()
	fmt.Println(res)
}

func FileUploadDownloadTest() {
	router := gin.Default()
	router.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("pf")
		if err != nil {
			context.String(http.StatusBadRequest, "A Bad Request")
			return
		}
		filename := file.Filename
		fmt.Println("Upload: " + filename)
		if err := context.SaveUploadedFile(file, file.Filename); err != nil {
			context.String(http.StatusBadRequest, "Upload File Error: %s", err.Error())
			return
		}
		context.String(http.StatusCreated, "upload successfully")
	})
	router.GET("/download", func(context *gin.Context) {
		filename := "haha.txt"
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.File("./test_files/cert.txt")
	})
	router.Run(":10016")
}

// @BasePath /api/v1

// HelloWorld godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func SwaggoTest() {
	r := gin.Default()

	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", HelloWorld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":10016")
}

func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func EncodeInMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func RedisTest() {
	RDB := redis.NewClient(&redis.Options{
		Addr:     "167.71.166.120" + ":" + strconv.Itoa(6379),
		Password: "rocendis",
		DB:       0,
	})

	sign, err := RDB.Del(context.Background(), "jake16").Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(sign)
	}
}
