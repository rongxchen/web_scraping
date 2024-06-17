package main

import (
	"fmt"
	"web_scraping/repositories"
)

func main() {
	//router := gin.Default()
	//router.Use(cors.Default())
	//
	//// Define a route handler
	//router.GET("/wb/users", api.GetWeiboUserList)
	//router.GET("/wb/update", api.SyncWB)
	//
	//// Start the server
	//exceptions.HandleError(router.Run())

	br := repositories.BlogRepo{}
	blogs := br.FindPage(1, 10)
	for _, blog := range blogs {
		fmt.Println(blog)
	}
}
