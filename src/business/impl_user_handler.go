package business

import (
	"context"
	"net/http"
	"rosenchat/src/database"
)

// implUserHandler implements IUserHandler.
type implUserHandler struct {
	userInfoDB database.IUserInfoDB
}

func (i *implUserHandler) GetUser(ctx context.Context, userID string) (*ResponseDTO, error) {
	userInfo, err := i.userInfoDB.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &ResponseDTO{
		StatusCode: http.StatusOK,
		Body: &ResponseBodyDTO{
			StatusCode: http.StatusOK,
			CustomCode: "OK",
			Data:       userInfo,
		},
	}, nil
}

func (i *implUserHandler) init() {
	i.userInfoDB = database.Get()
}
