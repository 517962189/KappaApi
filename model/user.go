package model

import (
	"time"
)

type User struct {
	BaseDb
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	CreateTime string `json:"create_time"`
}

func NewUser() *User{
	return &User{}
}

func (user *User) Add() error {

	return user.Db().Create(&User{
		Name:       "张三",
		Password:   "123456",
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}).Error
}
