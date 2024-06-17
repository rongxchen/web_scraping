package repositories

import (
	"web_scraping/database"
	"web_scraping/exceptions"
	"web_scraping/models"
)

var BlogDB = database.GetDefaultSqlite3Conn()

type BlogRepo struct {
}

func (b *BlogRepo) Create(blog models.Blog) {
	err := BlogDB.Create(&blog).Error
	exceptions.HandleError(err)
}

func (b *BlogRepo) CreateIfNotExists(blog models.Blog) {
	BlogDB.FirstOrCreate(&blog, &blog)
}

func (b *BlogRepo) FindOne(id uint) models.Blog {
	var blog models.Blog
	err := BlogDB.First(&blog, id).Error
	exceptions.HandleError(err)

	return blog
}

func (b *BlogRepo) FindPage(page int, size int) []models.Blog {
	var blogs []models.Blog
	err := BlogDB.Offset((page - 1) * size).Limit(size).Find(&blogs).Error
	exceptions.HandleError(err)

	return blogs
}
