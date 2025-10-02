package main

import (
    "github.com/gin-gonic/gin"
    "smartcampus/internal/handlers"
    "smartcampus/internal/middleware"
)

func main() {
    r := gin.Default()

    // Middlewares
    r.Use(middleware.AuthMiddleware(nil)) // Replace nil with actual JWTUtil
    r.Use(middleware.RBAC("admin", "teacher", "student"))

    // Auth routes
    r.POST("/api/auth/login", handlers.AuthHandler{}.Login)

    // User routes
    r.GET("/api/users", handlers.GetUsers)
    r.GET("/api/users/:id", handlers.GetUserByID)
    r.POST("/api/users", handlers.CreateUser)
    r.PUT("/api/users/:id", handlers.UpdateUser)
    r.DELETE("/api/users/:id", handlers.DeleteUser)

    // Course routes
    r.GET("/api/courses", handlers.GetCourses)
    r.GET("/api/courses/:id", handlers.GetCourseByID)
    r.POST("/api/courses", handlers.CreateCourse)
    r.PUT("/api/courses/:id", handlers.UpdateCourse)
    r.DELETE("/api/courses/:id", handlers.DeleteCourse)
    r.GET("/api/courses/:id/students", handlers.GetCourseStudents)

    // Assignment routes
    r.GET("/api/assignments", handlers.GetAssignments)
    r.GET("/api/assignments/:id", handlers.GetAssignmentByID)
    r.POST("/api/assignments", handlers.CreateAssignment)
    r.PUT("/api/assignments/:id", handlers.UpdateAssignment)
    r.DELETE("/api/assignments/:id", handlers.DeleteAssignment)

    // Submission routes
    r.GET("/api/submissions", handlers.GetSubmissions)
    r.GET("/api/submissions/:id", handlers.GetSubmissionByID)
    r.POST("/api/submissions", handlers.CreateSubmission)
    r.PUT("/api/submissions/:id", handlers.UpdateSubmission)
    r.DELETE("/api/submissions/:id", handlers.DeleteSubmission)

    // Message routes
    r.GET("/api/messages", handlers.GetMessages)
    r.GET("/api/messages/:id", handlers.GetMessageByID)
    r.POST("/api/messages", handlers.CreateMessage)
    r.PUT("/api/messages/:id", handlers.UpdateMessage)
    r.DELETE("/api/messages/:id", handlers.DeleteMessage)

    // File routes
    r.POST("/api/files/upload", handlers.UploadFile)
    r.GET("/api/files/:id", handlers.GetFileByID)
    r.DELETE("/api/files/:id", handlers.DeleteFile)

    r.Run(":8080")
}
