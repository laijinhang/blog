package model

import (
	"blog/common"
	"blog/model/model"
	"blog/model/mysql"
	"blog/model/redis"
	"fmt"
)

var _ archiverM = mysql.NewArchiver()
var _ archiverM = redis.NewArchiver()
var _ archiverM = NewArchiver()

type archiverM interface {
	Get(id string) *model.ArchiverModel
	GetList(filter *common.ListFilter) []*model.ArchiverModel
}

func NewArchiver() *archiver {
	return &archiver{}
}

type archiver struct {
}

func (a *archiver) Get(id string) *model.ArchiverModel {
	var archiver *model.ArchiverModel
	archiverSql := mysql.NewArchiver()
	archiverRedis := redis.NewArchiver()
	fmt.Println("id", id)
	if archiver = archiverRedis.Get(id); archiver != nil {
		a.SetRanking(id)
		return archiver
	}
	archiver = archiverSql.Get(id)
	if archiver != nil {
		archiverRedis.Set(id, archiver)
		a.SetRanking(id)
	}

	return archiver
}

func (a *archiver) Search(keyword string) []*model.ArchiverModel {
	return mysql.NewArchiver().Search(keyword)
}

func (a *archiver) Ranking() []*model.ArchiverModel {
	return redis.NewArchiver().GetRanking()
}

func (a *archiver) SetRanking(id string) {
	redis.NewArchiver().SetRanking(id)
}

func (a *archiver) GetList(filter *common.ListFilter) []*model.ArchiverModel {
	return mysql.NewArchiver().GetList(filter)
}
