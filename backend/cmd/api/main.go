package main

import (
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/C14147/SmartCampus-Workbench/internal/config"
    "github.com/C14147/SmartCampus-Workbench/internal/db"
    "github.com/C14147/SmartCampus-Workbench/internal/handlers"
    authpkg "github.com/C14147/SmartCampus-Workbench/internal/auth"
    "github.com/C14147/SmartCampus-Workbench/internal/models"
    "github.com/C14147/SmartCampus-Workbench/internal/middleware"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    cfg, err := config.LoadConfig()
    if err != nil {
        logger.Fatal("failed to load config", zap.Error(err))
    }

    // Initialize DB if DSN provided via env (DATABASE_DSN)
    dsn := cfgRawDSN(cfg)

    r := gin.Default()
    // register prometheus middleware
    r.Use(middleware.PrometheusMiddleware())

    if dsn != "" {
        gdb, err := db.Connect(dsn)
        if err != nil {
            logger.Fatal("db connect failed", zap.Error(err))
        }
        // auto migrate (keep minimal set)
        if err := gdb.AutoMigrate(&models.User{}); err != nil {
            logger.Fatal("auto migrate failed", zap.Error(err))
        }
        // register middleware to provide db to handlers
        r.Use(func(c *gin.Context) {
            c.Set("db", gdb)
            c.Next()
        })
    // gdb is a *gorm.DB returned from db.Connect and stored in the request context above
    }

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "healthy",
            "timestamp": time.Now().Format(time.RFC3339),
            "version":   "0.1.0",
        })
    })

    // Prometheus metrics endpoint
    r.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // initialize casbin enforcer
    enforcer := authpkg.NewEnforcer("./backend/config/rbac_model.conf", "./backend/config/rbac_policy.csv")

    api := r.Group("/api/v1")
    {
        api.POST("/auth/register", handlers.RegisterHandler)
        api.POST("/auth/login", handlers.LoginHandler)
        api.GET("/auth/me", handlers.MeHandler)
    }

    // Protected routes: require auth and RBAC checks
    protected := r.Group("/api/v1")
    protected.Use(handlers.AuthMiddleware())
    protected.Use(authpkg.RequirePermission(enforcer))
    {
        // schools
        protected.GET("/schools", handlers.ListSchools)
        protected.POST("/schools", handlers.CreateSchool)
        protected.GET("/schools/:id", handlers.GetSchool)
        protected.PUT("/schools/:id", handlers.UpdateSchool)
        protected.DELETE("/schools/:id", handlers.DeleteSchool)

        // assignments
        protected.GET("/assignments", handlers.ListAssignments)
        protected.POST("/assignments", handlers.CreateAssignment)
        protected.GET("/assignments/:id", handlers.GetAssignment)
        protected.PUT("/assignments/:id", handlers.UpdateAssignment)
        protected.DELETE("/assignments/:id", handlers.DeleteAssignment)
    }

    addr := ":" + cfg.Server.Port
    logger.Info("starting server", zap.String("addr", addr))
    if err := r.Run(addr); err != nil {
        logger.Fatal("server exited", zap.Error(err))
    }
}

// cfgRawDSN will try to read DATABASE_DSN from env or build one from config if present.
func cfgRawDSN(cfg *config.Config) string {
    if v := os.Getenv("DATABASE_DSN"); v != "" {
        return v
    }
    // no DB configured
    return ""
}
