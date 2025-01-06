package utils

import (
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}

func GetTotalPage(model interface{}, db *gorm.DB, pageSize int) (totalPage float64) {
	var count int64

	// Execute a count query on the specified model to get the total number of records
	err := db.Model(&model).Count(&count).Error
	if err != nil {
		log.Printf("err : %v\n", err.Error())
		return -1
	}

	/* Calculate the total number of pages by dividing the total count of records by the page size
	   i.e. rounds the given number up to the nearest greatest integer i.e. 3.33 will equivalent to 4 */

	totalPage = math.Ceil(float64(count) / float64(pageSize))

	return totalPage
}

func CheckPageInQuery(c *fiber.Ctx) int {
	page := 1

	if c.Query("page") != "" {
		pageRequest, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			log.Panicln(err)
		}

		if pageRequest <= 0 {
			page = 1
		} else {
			page = pageRequest
		}
	}
	return page
}

func CheckPageSizeInQuery(c *fiber.Ctx) int {
	var err error
	pageRequest := 0
	pageSize := 0

	if c.Query("pageSize") != "" {
		pageRequest, err = strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Panicln(err)
		}
	}

	if pageRequest <= 0 {
		pageSize = viper.GetInt("pagination.page_size")
	} else {
		pageSize = pageRequest
	}

	return pageSize
}
