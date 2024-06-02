package weibo

import (
	"fmt"
	"strconv"
	"web_scraping/dataframe"
	"web_scraping/http"
	"web_scraping/utils"
)

var WbSearchBlogUrl = "https://m.weibo.cn/api/container/getIndex?containerid=100103type%s&q%s&page_type=searchall&page=%d"

// FindBlogList helper function to extract blogs from response
// type 4: host search
// type 9: blog
// type 11: group -> may contain type9
func FindBlogList(cardList []interface{}) []interface{} {
	var blogList []interface{}
	for _, c := range cardList {
		// if card type is 11, find out card type 9 if found
		card, _ := c.(map[string]interface{})
		if int(card["card_type"].(float64)) == 11 {
			if group, ok := card["card_group"].([]interface{}); ok {
				for _, g := range group {
					if _card, ok := g.(map[string]interface{}); ok && int(_card["card_type"].(float64)) == 9 {
						card = _card
					}
				}
			}
		}
		// fetch card info
		blog := Blog{}
		if mBlog, ok := card["mblog"].(map[string]interface{}); ok {
			blog.MId = mBlog["mid"].(string)
			blog.Text = mBlog["text"].(string)
			pics := mBlog["pics"]
			if pics != nil {
				for _, p := range pics.([]interface{}) {
					if pic, ok := p.(map[string]interface{}); ok {
						blog.PictureList = append(blog.PictureList, pic["url"].(string))
					}
				}
			}
			blog.LikeCount = int(mBlog["attitudes_count"].(float64))
			blog.CommentCount = int(mBlog["comments_count"].(float64))
			blog.RepostCount = int(mBlog["reposts_count"].(float64))
			user := mBlog["user"].(map[string]interface{})
			blogUser := BlogUser{
				Id:             int(user["id"].(float64)),
				Username:       user["screen_name"].(string),
				Description:    user["description"].(string),
				FollowCount:    user["follow_count"].(float64),
				FollowersCount: user["followers_count"].(string),
				Avatar:         user["avatar_hd"].(string),
				Gender:         user["gender"].(string),
				Verified:       user["verified"].(bool),
				VerifiedReason: user["verified_reason"].(string),
			}
			blog.BlogUser = blogUser
			blog.Source = mBlog["source"].(string)
			if c, ok := mBlog["status_city"].(string); ok {
				blog.StatusCity = c
			}
			if c, ok := mBlog["status_country"].(string); ok {
				blog.StatusCountry = c
			}
			blog.CreatedAt = mBlog["created_at"].(string)
		} else {
			continue
		}
		blogList = append(blogList, blog)
	}
	return blogList
}

func GetWBBlogList(keyword string, typeInt int, pages int) {
	t := utils.UrlEncode("=" + strconv.Itoa(typeInt))
	q := utils.UrlEncode("=" + keyword)

	var blogList []interface{}

	for page := 1; page <= pages; page++ {
		url := fmt.Sprintf(WbSearchBlogUrl, t, q, page)

		resp := http.Get(url, map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		}, nil)

		jsn := resp.Json()
		data := jsn["data"].(map[string]interface{})
		cardList := data["cards"].([]interface{})

		blogList = append(blogList, FindBlogList(cardList)...)
		fmt.Println("page " + strconv.Itoa(page) + " done")
	}

	if len(blogList) > 0 {
		dataframe.ToCSV(blogList, "./weibo/weibo_blog_0_to_"+strconv.Itoa(pages)+".csv")
	}
}
