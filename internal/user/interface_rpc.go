package user

import (
	"context"
	"ddd_demo2/internal/servers/rpc/protos/in/user"
	user_pb "ddd_demo2/internal/servers/rpc/protos/in/user"
	"ddd_demo2/internal/user/user_model"
)

var _ user_pb.UserServer = &UserRpcServerImpl{}

type UserRpcServerImpl struct {
	UserApp UserAppInterface
	user_pb.UnimplementedUserServer
}

func NewUserServer(userApp UserAppInterface) *UserRpcServerImpl {
	return &UserRpcServerImpl{
		UserApp: userApp,
	}
}

func (u *UserRpcServerImpl) GetUser(ctx context.Context, req *user.G2S_UserInfo) (*user.S2G_UserInfo, error) {
	userID, err := user_model.NewUserID(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := u.UserApp.Get(userID)
	if err != nil {
		return nil, err
	}

	return user.ToS2G_UserInfo(), nil
}
