package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/C14147/SmartCampus-Workbench/internal/models"
    "github.com/C14147/SmartCampus-Workbench/internal/utils"
    "github.com/C14147/SmartCampus-Workbench/pkg/response"
)

func ListAssignments(c *gin.Context) {
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var list []models.Assignment
    gdb.Find(&list)
    response.Success(c, list)
}

func CreateAssignment(c *gin.Context) {
    var req models.Assignment
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

func GetAssignment(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var a models.Assignment
    if err := gdb.First(&a, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusNotFound, "not found", nil)
        return
    }
    response.Success(c, a)
}

func UpdateAssignment(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    var a models.Assignment
    if err := gdb.First(&a, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusNotFound, "not found", nil)
        return
    }
    if err := c.ShouldBindJSON(&a); err != nil {
        response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
        return
    }
    if err := utils.ValidateStruct(&a); err != nil {
        response.Error(c, http.StatusBadRequest, "validation failed", err.Error())
        return
    }
    if err := gdb.Save(&a).Error; err != nil {
        response.Error(c, http.StatusInternalServerError, "update failed", err.Error())
        return
    }
    response.Success(c, a)
}

func DeleteAssignment(c *gin.Context) {
    id := c.Param("id")
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)
    if err := gdb.Delete(&models.Assignment{}, "id = ?", id).Error; err != nil {
        response.Error(c, http.StatusInternalServerError, "delete failed", err.Error())
        return
    }
    c.Status(http.StatusNoContent)
}
