package controller

import (
	"blog/model"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Archiver struct {
}

func (a *Archiver) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aModel := model.NewArchiver()
	data := aModel.Get(params.ByName("id"))
	b, _ := json.Marshal(data)
	writer.Write(b)
}

func (a *Archiver) GetList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
}

func (a *Archiver) Ranking(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aModel := model.NewArchiver()
	data := aModel.Ranking()
	b, _ := json.Marshal(data)
	writer.Write(b)
}

func (a *Archiver) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	aModel := model.NewArchiver()
	data := aModel.Search(params.ByName("keyword"))
	b, _ := json.Marshal(data)
	writer.Write(b)
}
