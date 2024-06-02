package weibo

import "fmt"

type BlogUser struct {
	Id             int
	Username       string
	Description    string
	FollowCount    float64
	FollowersCount string
	Avatar         string
	Gender         string
	Verified       bool
	VerifiedReason string
}

func (b BlogUser) String() string {
	return fmt.Sprintf("Id: %d, Username: %s, Description: %s, FollowCount: %f, "+
		"FollowersCount: %s, Avatar: %s, Gender: %s, Verified: %t, VerifiedReason: %s",
		b.Id, b.Username, b.Description, b.FollowCount, b.FollowersCount, b.Avatar, b.Gender,
		b.Verified, b.VerifiedReason)
}

type Blog struct {
	MId           string
	Text          string
	PictureList   []string
	LikeCount     int
	CommentCount  int
	RepostCount   int
	BlogUser      BlogUser
	Source        string
	StatusCity    string
	StatusCountry string
	CreatedAt     string
}

func (b Blog) String() string {
	return fmt.Sprintf("MId: %s, Text: %s, PictureList: %v, LikeCount: %d, CommentCount: %d, "+
		"RepostCount: %d, BlogUser: %s, Source: %s, StatusCity: %s, StatusCountry: %s, CreatedAt: %s",
		b.MId, b.Text, b.PictureList, b.LikeCount, b.CommentCount, b.RepostCount, &b.BlogUser,
		b.Source, b.StatusCity, b.StatusCountry, b.CreatedAt)
}
