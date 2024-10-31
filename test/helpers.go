package test

import (
	"chee-go-backend/health"
	"chee-go-backend/lectures"
	"chee-go-backend/resumes"
	"chee-go-backend/users"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpTestDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"test",        // username
		"test",        // password
		"localhost",   // host
		"3306",        // port
		"cheego_test", // database name
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error initialize database: %s", err)
	}

	db.AutoMigrate(&users.User{}, &resumes.Resume{}, &resumes.Education{}, &resumes.Project{}, &resumes.Keyword{}, &resumes.KeywordResume{}, &resumes.Activity{}, &resumes.Certificate{}, &resumes.WorkExperience{}, &resumes.WorkExperienceDetail{}, &lectures.Subject{}, &lectures.Lecture{})

	ClearTestData(db)

	return db
}

func SetUpTestRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	serverRoute := r.Group("/api")
	users.RegisterUsersRouters(serverRoute.Group("/users"), db)
	resumes.RegisterResumesRouters(serverRoute.Group("/resumes"), db)
	health.RegisterUsersRouters(serverRoute.Group("/health"))
	lectures.RegisterLecturesRouters(serverRoute.Group("/lectures"), db)

	return r
}

func ClearTestData(db *gorm.DB) {
	db.Exec("DELETE FROM lectures")
	db.Exec("DELETE FROM subjects")
	db.Exec("DELETE FROM work_experience_details")
	db.Exec("DELETE FROM work_experiences")
	db.Exec("DELETE FROM certificates")
	db.Exec("DELETE FROM activities")
	db.Exec("DELETE FROM keyword_resumes")
	db.Exec("DELETE FROM keywords")
	db.Exec("DELETE FROM projects")
	db.Exec("DELETE FROM educations")
	db.Exec("DELETE FROM resumes")
	db.Exec("DELETE FROM users")
}

func InitTest() (*gin.Engine, *gorm.DB) {
	db := SetUpTestDB()
	router := SetUpTestRouter(db)

	return router, db
}
