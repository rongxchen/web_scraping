package repositories

import (
	"web_scraping/database"
	"web_scraping/exceptions"
	"web_scraping/models"
)

var BlogUserDB = database.GetDefaultSqlite3Conn()

type BlogUserRepo struct {
}

func (b *BlogUserRepo) Create(blogUser models.BlogUser) {
	err := BlogUserDB.Create(&blogUser).Error
	exceptions.HandleError(err)
}

func (b *BlogUserRepo) CreateIfNotExists(blogUser models.BlogUser) {
	BlogUserDB.FirstOrCreate(&blogUser, &blogUser)
}

func (b *BlogUserRepo) FindOne(id uint) models.BlogUser {
	var blogUser models.BlogUser
	err := BlogUserDB.First(&blogUser, id).Error
	exceptions.HandleError(err)

	return blogUser
}

func (b *BlogUserRepo) FindPage(page int, size int) []models.BlogUser {
	var blogUsers []models.BlogUser
	BlogUserDB.Limit(size).Offset((page - 1) * size).Find(&blogUsers)
	return blogUsers
}
