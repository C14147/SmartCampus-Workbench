package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/your-org/smartcampus/internal/models"
    "github.com/your-org/smartcampus/internal/utils"
    "github.com/your-org/smartcampus/pkg/response"
)

func ListSchools(c *gin.Context) {
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var list []models.School
    gdb.Find(&list)
    response.Success(c, list)
}

func CreateSchool(c *gin.Context) {
    var req models.School
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
        return
    }
    if err := utils.ValidateStruct(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "validation failed", err.Error())
        return
    }
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    if err := gdb.Create(&req).Error; err != nil {
        response.Error(c, http.StatusBadRequest, "create failed", err.Error())
        return
    }
    response.Success(c, req)
}

func GetSchool(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var s models.School
    if err := gdb.First(&s, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusNotFound, "not found", nil)
        return
    }
    response.Success(c, s)
}

func UpdateSchool(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var s models.School
    if err := gdb.First(&s, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusNotFound, "not found", nil)
        return
    }
    if err := c.ShouldBindJSON(&s); err != nil {
        response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
        return
    }
    if err := utils.ValidateStruct(&s); err != nil {
        response.Error(c, http.StatusBadRequest, "validation failed", err.Error())
        return
    }
    if err := gdb.Save(&s).Error; err != nil {
        response.Error(c, http.StatusInternalServerError, "update failed", err.Error())
        return
    }
    response.Success(c, s)
}

func DeleteSchool(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    if err := gdb.Delete(&models.School{}, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusInternalServerError, "delete failed", err.Error())
        return
    }
    c.Status(http.StatusNoContent)
}
