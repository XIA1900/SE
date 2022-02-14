package main

import (
	docs "ComponentTest/docs"
	"ComponentTest/es"
	"ComponentTest/log"
	"ComponentTest/role"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

func main() {
	// logTest()

	// jwt.TestServer()

	// casbinTest()

	// ESTest()

	// FileUploadDownloadTest()

	SwaggoTest()
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
