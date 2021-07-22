package session

import (
	"sync"
	"time"
)

//ISession 操作接口
//Session数据结构为散列表kv
type ISession interface {
	//Set 设置
	Set(key, value interface{})
	//Get 获取
	Get(key interface{}) interface{}
	//Delete 删除
	Delete(key interface{})
}

func Session() *session {
	return &s
}

var s session

type session struct {
	m sync.Map
}

func (s *session) Set(key, value interface{}) {
	s.m.Store(key, store{
		data: value,
		time: time.Now(),
	})
}

func (s *session) Get(key interface{}) interface{} {
	val, ok := s.m.Load(key)
	if !ok {
		return nil
	}
	if time.Now().Unix()-val.(store).time.Unix() >= 30*60 {
		s.m.Delete(key)
		return nil
	}
	s.m.Store(key, store{
		data: val.(string),
		time: time.Now(),
	})
	return val.(store).data
}

func (s *session) Delete(key interface{}) {
	s.m.Delete(key)
}

type store struct {
	data interface{}
	time time.Time
}
