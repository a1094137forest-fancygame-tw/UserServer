package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Player struct {
	IsAdmin        bool   `json:"is_manager" bson:"isAdmin"`
	Account        string `json:"account,omitempty"`
	Password       string `json:"password,omitempty" bson:"password"`
	Avatar         int    `json:"avatar,omitempty" bson:"avatar"`
	Gender         int    `json:"gender,omitempty" bson:"gender"`
	LastLoginTime  int    `json:"lastLoginTime" bson:"lastLoginTime"`
	LastLogoutTime int    `json:"lastLogoutTime" bson:"lastLogoutTime"`
	Balance        int    `json:"balance" bson:"balance"`
}

type PlayerUpdate struct {
	IsManager bool `bson:"is_manager"`
}

func (p *Player) Insert(ctx context.Context, collection *mongo.Collection) error {
	log.Println("player", p)
	_, err := collection.InsertOne(ctx, p)
	return err
}

func FindMemberList(ctx context.Context, collection *mongo.Collection, filter interface{}) (*[]Player, error) {
	var result []Player

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var tmp bson.M
		if err := cur.Decode(&tmp); err != nil {
			return nil, err
		}

		p := Player{
			Account:        tmp["account"].(string),
			Password:       tmp["password"].(string),
			Avatar:         int(tmp["avatar"].(int32)),
			Gender:         int(tmp["gender"].(int32)),
			IsAdmin:        tmp["isAdmin"].(bool),
			LastLoginTime:  int(tmp["lastLoginTime"].(int64)),
			LastLogoutTime: int(tmp["lastLogoutTime"].(int64)),
			Balance:        int(tmp["balance"].(int64)),
		}
		result = append(result, p)
	}
	return &result, nil
}

func Update(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("update err", err)
		return nil, err
	}
	return result, nil
}
