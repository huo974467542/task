package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
*/

var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go-test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = db
}
func main() {
	gorm.m
}
