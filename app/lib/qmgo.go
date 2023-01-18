package lib

import (
	"context"
	"github.com/qiniu/qmgo"
)

func GetMgoDb() *qmgo.QmgoClient {
	cli, err := qmgo.Open(context.Background(), &qmgo.Config{Uri: "mongodb://192.168.1.130:27017", Database: "playbill", Coll: "poster"})
	if err != nil {
		return nil
	}
	return cli
}
