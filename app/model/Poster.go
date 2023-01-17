package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"palybill/app/lib"
)

type Poster struct {
	_id string
}

func (*Poster) GetAll(ctx context.Context) []Poster {
	var poster []Poster
	lib.GetMgoDb().Find(ctx, bson.M{"age": 10}).Select(bson.M{"age": 1}).All(&poster)
	return poster
}

func (*Poster) GetOne(ctx context.Context) Poster {
	var poster Poster
	lib.GetMgoDb().Find(ctx, bson.M{"age": 10}).Select(bson.M{"age": 1}).One(&poster)
	return poster
}
