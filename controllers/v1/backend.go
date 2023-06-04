package v1

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"UserServer/mongodb"
	"UserServer/mongodb/model"
	fancygame "UserServer/proto/fancygame"
)

func (s *UserServe) Login(ctx context.Context, req *fancygame.LoginInfo) (*fancygame.LoginRes, error) {
	log.Println("get login")
	var resp fancygame.LoginRes
	log.Println("a p", req.Account, req.Password)
	filter := bson.M{
		"account": req.Account,
	}

	ps, err := model.FindMemberList(ctx, mongodb.PoolWr.Read("player"), filter)
	if err != nil {
		return nil, err
	}

	if len(*ps) == 0 {
		return nil, errors.New("no exists account")
	} else {
		password := (*ps)[0].Password
		if password != req.Password {
			return nil, errors.New("wromg password")
		} else {
			resp = fancygame.LoginRes{
				StatusCode: 200,
				Message:    "Ok",
				Data: &fancygame.LoginResInfo{
					Avatar:  fancygame.AvatarEnum((*ps)[0].Avatar),
					Gender:  fancygame.GenderEnum((*ps)[0].Gender),
					IsAdmin: (*ps)[0].IsAdmin,
					Balance: int64((*ps)[0].Balance),
				},
			}
		}
	}

	log.Println("resp", resp.Data.IsAdmin, resp.Data.Balance)

	return &resp, nil
}

func (s *UserServe) Logout(ctx context.Context, req *fancygame.LogoutInfo) (*fancygame.LogoutRes, error) {
	var resp fancygame.LogoutRes

	filter := bson.M{
		"account": req.Account,
	}

	update := bson.M{
		"$set": bson.M{
			"Balance": req.Balance,
		},
	}

	_, err := model.Update(ctx, mongodb.PoolWr.Read("player"), filter, update)
	if err != nil {
		return nil, err
	}

	resp = fancygame.LogoutRes{
		StatusCode: 200,
		Message:    "Ok",
	}

	return &resp, nil
}

func (s *UserServe) Register(ctx context.Context, req *fancygame.RegisterInfo) (*fancygame.RegisterRes, error) {
	log.Println("user get register")
	var resp fancygame.RegisterRes
	log.Println("req", req)
	var player = model.Player{
		IsAdmin:        false,
		Account:        req.Account,
		Password:       req.Password,
		Avatar:         int(req.Avatar),
		Gender:         int(req.Gender),
		LastLoginTime:  int(time.Now().Unix()),
		LastLogoutTime: int(time.Now().Unix()),
		Balance:        0,
	}
	log.Println("start insert")
	err := player.Insert(ctx, mongodb.PoolWr.Write("player"))
	if err != nil {
		return nil, err
	}
	log.Println("insert success")
	resp = fancygame.RegisterRes{
		StatusCode: 200,
		Message:    "Ok",
		Data: &fancygame.RegisterResInfo{
			Avatar:  fancygame.AvatarEnum(player.Avatar),
			Gender:  fancygame.GenderEnum(player.Gender),
			IsAdmin: player.IsAdmin,
			Balance: int64(player.Balance),
		},
	}

	log.Println("resp", resp)
	return &resp, nil
}
