package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Arjit801/TheBloggies/dao/models"
	"github.com/Arjit801/TheBloggies/database"
	"github.com/Arjit801/TheBloggies/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatPost(c *fiber.Ctx) error {
	var blogPost models.Blog
	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("unable to parse body")
	}
	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message":"Success!! Your post is live",
	})
}

func GetAllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page-1)*limit
	var totalPages int64
	var getBlog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getBlog)
	database.DB.Model(&models.Blog{}).Count(&totalPages)
	return c.JSON(fiber.Map{
		"data":getBlog,
		"meta":fiber.Map{
			"totalPages": totalPages,
			"page":page,
			"last_page":math.Ceil(float64(int(totalPages)/limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogPost models.Blog
	database.DB.Where("id=?", id).Preload("user").First(&blogPost)
	return c.JSON(fiber.Map{
		"data": blogPost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("unable to parse body")
	}
	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "post updated successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var blogs []models.Blog
	database.DB.Model(&blogs).Where("user_id=?", id).Preload("User").Find(&blogs)
	return c.JSON(blogs)
}
func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"Cannot delete the log post! Record not found",
		})
	}
	return c.JSON(fiber.Map{
		"message":"Blog post deleted successfully",
	})

}