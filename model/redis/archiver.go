package redis

import (
	"blog/common"
	"blog/model/model"
	"encoding/json"

	"github.com/go-redis/redis"

	"github.com/sirupsen/logrus"
)

func NewArchiver() *archiver {
	return &archiver{}
}

type archiver struct {
}

func (a *archiver) Get(id string) *model.ArchiverModel {
	var archiver model.ArchiverModel
	d := redisClient.Get(id).Val()
	if d == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(d), &archiver); err != nil {
		logrus.Errorln(err.Error)
		return nil
	}
	return &archiver
}

func (a *archiver) Set(id string, archiverModel *model.ArchiverModel) {
	b, _ := json.Marshal(archiverModel)
	redisClient.Set(id, string(b), -1)
}

func (a *archiver) GetRanking() []*model.ArchiverModel {
	var archiver []*model.ArchiverModel
	list := redisClient.ZRevRangeByScore("ranking", redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  0,
	}).Val()
	for id := range list {
		d := redisClient.Get(list[id]).Val()
		if d == "" {
			return nil
		}
		var t model.ArchiverModel
		if err := json.Unmarshal([]byte(d), &t); err != nil {
			logrus.Errorln(err.Error)
			continue
		}
		archiver = append(archiver, &t)
	}

	return archiver

}

func (a *archiver) SetRanking(id string) {
	redisClient.ZIncrBy("ranking", 1, id)
}

func (a *archiver) GetList(filter *common.ListFilter) []*model.ArchiverModel {
	return nil
}
