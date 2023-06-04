package v1

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"UserServer/mongodb"
	"UserServer/mongodb/model"
	fancygame "UserServer/proto/fancygame"
)

type UserServe struct {
	fancygame.UnimplementedUserServer
}

func (s *UserServe) GetMemberList(ctx context.Context, req *fancygame.EmptyReq) (*fancygame.GetMemberListRes, error) {
	log.Println("user server get memberList req")
	var resp fancygame.GetMemberListRes
	ps, err := model.FindMemberList(ctx, mongodb.PoolWr.Read("player"), bson.M{})
	if err != nil {
		return nil, err
	}
	log.Println("here", ps)

	memberList := make([]*fancygame.MemberList, len(*ps))

	for i, item := range *ps {
		memberList[i] = &fancygame.MemberList{
			Account:        item.Account,
			Password:       item.Password,
			Avatar:         fancygame.AvatarEnum(item.Avatar),
			Gender:         fancygame.GenderEnum(item.Gender),
			Balance:        int64(item.Balance),
			LastLoginTime:  int64(item.LastLoginTime),
			LastLogoutTime: int64(item.LastLogoutTime),
		}
	}

	resp = fancygame.GetMemberListRes{
		StatusCode: 200,
		Message:    "Ok",
		Data: &fancygame.MemberListInfo{
			MemberList: memberList,
		},
	}
	log.Println("memberList req success")
	return &resp, nil
}

func (s *UserServe) SetMemberData(ctx context.Context, req *fancygame.SetMember) (*fancygame.SetMemberRes, error) {
	var resp fancygame.SetMemberRes

	filter := bson.M{
		"account": req.Account,
	}

	update := bson.M{
		"$set": bson.M{
			"password": req.Password,
			"gender":   req.Gender,
			"avatar":   req.Avatar,
			"Balance":  req.Balance,
		},
	}

	updateResult, err := model.Update(ctx, mongodb.PoolWr.Write("player"), filter, update)
	if err != nil {
		return nil, err
	}

	if updateResult.ModifiedCount == 0 {
		errorMsg := errors.New("OriginData is same as UpdateData")
		return nil, errorMsg
	}

	resp = fancygame.SetMemberRes{
		StatusCode: 200,
		Message:    "Ok",
	}

	return &resp, nil
}

func (s *UserServe) KickOutMember(ctx context.Context, req *fancygame.KickOutMemberInfo) (*fancygame.KickOutMemberRes, error) {
	var resp fancygame.KickOutMemberRes

	filter := bson.M{
		"account": req.Account,
	}

	update := bson.M{
		"$set": bson.M{
			"Balance": req.Balance,
		},
	}

	_, err := model.Update(ctx, mongodb.PoolWr.Write("player"), filter, update)
	if err != nil {
		return nil, err
	}

	resp = fancygame.KickOutMemberRes{
		StatusCode: 200,
		Message:    "Ok",
	}

	return &resp, nil
}
