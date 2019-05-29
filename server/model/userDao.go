package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// UserDao 实例，全局唯一
var CurrentUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

// 初始化一个 UserDao 结构体示例，
func InitUserDao(pool *redis.Pool) (currentUserDao *UserDao) {
	currentUserDao = &UserDao{pool: pool}
	return
}

// 根据用户 id 获取用户信息
// 获取成功返回 user 信息，err nil
// 获取失败返回 err，user 为 nil
func getUsrById(conn redis.Conn, id int) (user User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		err = ERROR_USER_NOT_EXISTS
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Printf("Unmarshal user info error: %v\n", err)
		return
	}
	return
}

func (this *UserDao) Login(id int, password string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = getUsrById(conn, id)
	if err != nil {
		fmt.Printf("get user by id error: %v\n", err)
		return
	}

	if user.Password != password {
		err = ERROR_USER_PWD
		return
	}
	return
}
