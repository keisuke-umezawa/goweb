package db

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

type DatabaseInterface interface {
    Connect() *gorm.DB
    DBInstnce(c *gin.Context) *gorm.DB
}
