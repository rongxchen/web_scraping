package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web_scraping/models"
	"web_scraping/repositories"
	"web_scraping/weibo"
)

var blogUserRepo = repositories.BlogUserRepo{}

func GetWeiboUserList(c *gin.Context) {
	p := c.Query("page")
	s := c.Query("size")

	page, _ := strconv.Atoi(p)
	size, _ := strconv.Atoi(s)

	userList := blogUserRepo.FindPage(page, size)
	rs := &models.Result{Code: 0, Data: userList, Message: "ok"}

	c.JSON(http.StatusOK, rs)
}

func SyncWB(c *gin.Context) {
	weibo.GetWBBlogList("以色列", 1, 1)
	c.JSON(http.StatusOK, nil)
}
