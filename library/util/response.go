package util

import (
	"cld/dao"
	"encoding/json"
	"net/http"
)

type JsonRes struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type JsonUserPage struct {
	Code int        `json:"code"`
	Data []dao.User `json:"data"`
}

const (
	SUCCESS = 1000 + iota
)

func ResponseJson(w http.ResponseWriter, code int, data string) {
	msg, _ := json.Marshal(JsonRes{Code: code, Data: data})
	w.Write(msg)
}

func ResponseSuccess(w http.ResponseWriter, code int, data string) {
	msg, _ := json.Marshal(JsonRes{Code: SUCCESS, Data: ""})
	w.Write(msg)
}

func ResponseUserPage(w http.ResponseWriter, code int, data []dao.User) {
	msg, _ := json.Marshal(JsonUserPage{Code: code, Data: data})
	w.Write(msg)
}
