package gapi

import (
	db "github.com/CineDeepMatch/Backend-server/db/sqlc"
	"github.com/CineDeepMatch/Backend-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:                user.ID.String(),
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreateAt),
	}
}

func convertActivity(activity db.Activity) *pb.Activity {
	return &pb.Activity{
		Id:            activity.ID.String(),
		UserId:        activity.UserID.String(),
		ViewPage:      activity.ViewPage,
		PageVisitedAt: timestamppb.New(activity.PageVisitedAt),
	}
}
