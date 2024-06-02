package main

import "web_scraping/weibo"

func main() {
	weibo.GetWBBlogList("以色列", 1, 10)

	//rows := dataframe.ReadCSV("./weibo/weibo_blog_page_0.csv")
	//for i, row := range rows {
	//	fmt.Println(i, row)
	//}
}
