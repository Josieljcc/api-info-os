package service

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := ctx.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}
		pageSizeStr := ctx.DefaultQuery("pageSize", "10")
		pageSize, _ := strconv.Atoi(pageSizeStr)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 1:
			pageSize = 10
		}

		dbClone := db.Session(&gorm.Session{})
		var total int64
		dbClone.Count(&total)

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

		ctx.Set("totalPages", totalPages)
		ctx.Set("page", page)
		ctx.Set("pageSize", pageSize)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
