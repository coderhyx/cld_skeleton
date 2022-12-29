package goods

import (
	"cld/dao"
)

func GoodListLogic(page int, limit int, title string) []dao.User {
	u := dao.UserList(page, limit, title)
	return u
}
