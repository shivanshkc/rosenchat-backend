package database

const (
	userInfoTableName = "user_info"
)

// UserInfoDTO is the schema of the user document.
type UserInfoDTO struct {
	ID          string `json:"_id" bson:"_id"`
	Email       string `json:"email" bson:"email"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PictureLink string `json:"picture_link" bson:"picture_link"`
}
