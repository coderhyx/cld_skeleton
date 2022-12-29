package goods

import (
	"cld/dao"
)

func GoodAddLogic() dao.User {
	u := dao.GetById(1)
	return u
}
