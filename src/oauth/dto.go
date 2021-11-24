package oauth

// UserInfoDTO is the schema of the user info retrieved from an OAuth token.
type UserInfoDTO struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PictureLink string `json:"picture_link"`
}
