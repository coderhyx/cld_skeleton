package dao

import (
	"cld/bootstrap"
	"fmt"
)

type User struct {
	Usertel string
	Img     string
	Name    string
}

func GetById(id int) User {
	var u User
	sqlStr := "select usertel,name,img from user where id=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err2 := bootstrap.DB.QueryRow(sqlStr, id).Scan(&u.Usertel, &u.Name, &u.Img)
	if err2 != nil {
		fmt.Printf("scan failed, err:%v\n", err2)
		return u
	}
	return u
	//fmt.Printf("id:%d name:%s age:%d\n", u.usertel, u.name, u.img)
}

func AddUser(u User) int64 {
	sqlStr := "insert into user(name, usertel) values (?,?)"
	ret, err := bootstrap.DB.Exec(sqlStr, u.Name, u.Usertel)
	var theID int64
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return theID
	}
	theID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return theID
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return theID
}

// 查询多条数据示例
func UserList(page int, limit int, name string) []User {
	//收集多行列表信息的容器
	users := []User{}
	//定义一条查询语句
	sqlStr := "select id, name, usertel from user where id > ?"
	rows, err := bootstrap.DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return users
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		//空车
		var u User
		//scan 将每一行的数据扫描到user结构体里， 数据装到车里了
		err := rows.Scan(&u.Img, &u.Name, &u.Usertel)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return users
		}
		//加到车队里
		users = append(users, u)
	}
	return users
}
