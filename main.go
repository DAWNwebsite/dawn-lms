package main

import (
	"aida/auth"
	"aida/database"
	"aida/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

// "github.com/joechristophers/GoEasyJWT"
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB := database.ConnectDB()

	DB.Logger.LogMode(logger.Info)

	r := gin.Default()

	r.POST("auth/signup", auth.SignUp)
	r.POST("/auth/login", auth.Login)

	User := r.Group("user")
	Student := r.Group("student")
	Course := r.Group("course")
	Teacher := r.Group("/teacher")
	User.GET("/profile", routers.UserProfile)

	Student.POST("new", routers.CreateStudentProfile)
	Student.GET("/all", routers.GetStudents)
	Student.GET("/:id", routers.GetASudent)
	Student.GET("/myprofile", routers.GetStudentProfile)
	Student.DELETE("/delete", routers.DeleteStudentProfile)

	Course.POST("/new", routers.CreateCourse)
	Course.GET("/all", routers.GetAllCourses)
	Course.GET("/:id", routers.GetOneCourse)
	Course.DELETE("/:id/delete", routers.DeleteCourse)
	Course.PUT("/:id/students/new", routers.AddCourseToStudent)
	Course.DELETE("/:id/students/delete", routers.RemoveStudentFromCourse)
	Course.DELETE("/:id/teachers/delete", routers.RemoveTeacherFromCourse)
	Course.PUT("/:id/teachers/new", routers.AddTeacherToCourse)
	Course.GET("/:id/teachers", routers.GetAllCourseTeachers)

	Teacher.POST("/new", routers.CreateTeachers)
	Teacher.GET("/all", routers.GetAllTeachers)
	Teacher.GET("/:id", routers.GetOneTeacher)

	Course.GET("/:id/students", routers.GetCourseStudents)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

	//
	///go build -tags netgo -ldflags '-s -w' -o app
}
