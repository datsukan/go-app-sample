package domain

import "time"

type Users struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Birthday  time.Time
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// この struct はビジネスロジックだと思うので、 usecase で書くべきなのかと思ったけど、
// ここに定義した。
type UsersForGet struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	ImageUrl  string `json:"image_url"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (u *Users) BuildForGet() UsersForGet {
	user := UsersForGet{}
	user.ID = u.ID
	user.Username = u.Username
	user.Email = u.Email
	user.Birthday = u.Birthday.Format("2006-01-02")
	user.ImageUrl = u.ImageUrl
	user.CreatedAt = u.CreatedAt.Format(time.RFC3339)
	user.UpdatedAt = u.UpdatedAt.Format(time.RFC3339)

	return user
}
