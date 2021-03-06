package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ssaatw/app/controllers"
	"github.com/ssaatw/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Jobtitles struct type
type Jobtitles struct {
	Jobtitle []Jobtitle
}

// Jobtitle struct type
type Jobtitle struct {
	Jobtitlename string
}

// Departments struct type
type Departments struct {
	Department []Department
}

// Department struct type
type Department struct {
	Departmentname string
}

// Genders struct type
type Genders struct {
	Gender []Gender
}

// Gender struct type
type Gender struct {
	Gendername string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewPersonalController(v1, client)
	controllers.NewJobtitleController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewGenderController(v1, client)

	//Set Jobtitle Data
	jobtitles := []string{"แพทย์", "พยาบาล", "เจ้าหน้าที่เภสัชกร", "เจ้าหน้าที่บุคลากร",
		"เจ้าหน้าที่เวชระเบียน", "เจ้าหน้าที่บัญชี", "เจ้าหน้าที่จัดการหอผู้ป่วยใน"}
	for _, jt := range jobtitles {
		client.Jobtitle.
			Create().
			SetJobtitlename(jt).
			Save(context.Background())
	}

	//Set Department Data
	departments := []string{"อายุรกรรม", "ฉุกเฉินและอุบัติเหตุ", "ผู้ป่วยนอก", "ผู้ป่วยใน", "ศัลยกรรม",
		"กุมารเวชกรรม", "สูตินรีเวชกรรม", "จักษุ", "หู คอ จมูก", "ไม่มี"}
	for _, d := range departments {
		client.Department.
			Create().
			SetDepartmentname(d).
			Save(context.Background())
	}

	//Set Gender Data
	genders := []string{"Male", "Female"}
	for _, g := range genders {
		client.Gender.
			Create().
			SetGendername(g).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
