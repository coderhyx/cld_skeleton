package controller

import (
	"cld/dao"
	"cld/library/util"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var u dao.User

// 单查
func GoodsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//util.CkToken2(r)
	s := util.CkToken2(r)
	if s == false {
		util.ResponseJson(w, 1003, "token 校验错误")
	}
	//fmt.Println(r.PostFormValue("aa"))
	//fmt.Println(r.PostFormValue("bb"))
	//fmt.Println(r.PostForm)
	//u := goods.GoodListLogic()
	//util.ResponseJson(w, 200, u.Name)
}

// 列表
func GoodsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//接收入参
	pageParam := r.PostFormValue("page")
	limitParam := r.PostFormValue("limit")

	//类型转换
	page, _ := strconv.Atoi(pageParam)
	limit, _ := strconv.Atoi(limitParam)

	//u := goods.GoodListLogic(page, limit, "")
	u := dao.UserList(page, limit, "")
	util.ResponseUserPage(w, 200, u)
}

// 添加
func GoodsAdd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := dao.User{}
	u.Usertel = r.PostFormValue("tel")
	u.Name = r.PostFormValue("name")
	dao.AddUser(u)
	util.ResponseJson(w, 200, "success")
}

// 删除
func GoodsDel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//编辑
