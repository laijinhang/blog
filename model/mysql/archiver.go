package mysql

import (
	"blog/common"
	"blog/model/model"

	"github.com/sirupsen/logrus"
)

func NewArchiver() *archiver {
	return &archiver{}
}

type archiver struct {
}

func (a *archiver) Get(id string) *model.ArchiverModel {
	var archiver model.ArchiverModel
	if err := DB().Model(&model.ArchiverModel{}).Where("cid=?", id).First(&archiver).Error; err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	return &archiver
}
func (a *archiver) Search(keyword string) []*model.ArchiverModel {
	var archiver []*model.ArchiverModel
	if err := DB().Model(&model.ArchiverModel{}).Where("text like ?", "%"+keyword+"%").Find(&archiver).Error; err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	return archiver
}

func (a *archiver) GetList(filter *common.ListFilter) []*model.ArchiverModel {
	var archiverList []*model.ArchiverModel
	if err := DB().Limit(filter.PageSize).Offset(filter.PageSize * filter.PageIndex).Find(&archiverList).Error; err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	return archiverList
}
