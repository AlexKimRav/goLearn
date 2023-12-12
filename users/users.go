package users

import (
	"fmt"
	"math/rand"
	"time"
)

type person interface {
	getName() string
}

type Admin struct {
	adminId       int
	adminName     string
	adminEmail    string
	adminPw       string
	adminNickname string
	adminAge      int
}

type User struct {
	userId       int
	userName     string
	userEmail    string
	userPw       string
	userNickname string
	userAge      int
}

func (a Admin) getName() string {
	return a.adminName
}

func (u User) getName() string {
	return u.userName
}

func PrintName(p person) {
	fmt.Println(p.getName())
}

func CreateUser(name, email, pw, nickname string, age int) User {
	u := User{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	u.userId = r.Int()
	u.userName, u.userEmail, u.userPw, u.userNickname, u.userAge =
		name, email, pw, nickname, age
	return u
}

//TODO доделать чтение пользователя

func (u User) String() string {
	return fmt.Sprintln("id:", u.userId,
		"\nname:", u.userName,
		"\nmail:", u.userEmail,
		"\nnickname:", u.userNickname,
		"\nage", u.userAge)
}
