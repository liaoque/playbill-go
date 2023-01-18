package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"palybill/app/lib"
)

type Poster struct {
	Id    string                 `bson:"_id"`
	Title string                 `bson:"title"`
	Data  map[string]interface{} `bson:"data"`
}

func (*Poster) GetAll(ctx context.Context, keyword string, page int64, limit int64) []Poster {
	var poster []Poster
	var filter bson.M = make(bson.M, 2)
	if len(keyword) > 0 {
		filter["title"] = bson.M{"$regex": keyword}
	}
	lib.GetMgoDb().
		Find(ctx, filter).
		Skip((page - 1) * limit).
		Limit(limit).
		Select(bson.M{"data": 0}).
		All(&poster)
	return poster
}

func (*Poster) GetOneById(ctx context.Context, id string) Poster {
	var poster Poster
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return poster
	}
	lib.GetMgoDb().
		Find(ctx, bson.M{"_id": hex}).
		Select(bson.M{}).
		One(&poster)
	return poster
}
