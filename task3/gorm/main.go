package main

import (
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
type User struct {
	gorm.Model
	Name      string
	Email     string
	Posts     []Post
	PostCount int // 新增字段：用户的文章数量统计
}
type Post struct {
	gorm.Model
	Title         string
	UserId        uint
	CommentStatus string // 新增字段，用于记录评论状态，默认为 "有评论"
	Comments      []Comment
}
type Comment struct {
	gorm.Model
	PostId  uint
	Content string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(192.168.232.144:3306)/go-test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	/**
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
	*/

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	/*var user User
	db.First(&user, 1) // 查询用户

	var posts []Post
	db.Where("user_id = ?", user.ID).Preload("Comments").Find(&posts) // 查询文章并加载评论

	user.Posts = posts // 手动赋值

	// 打印输出
	for _, post := range user.Posts {
		fmt.Println("文章标题：", post.Title)
		for _, comment := range post.Comments {
			fmt.Println("  评论内容：", comment.Content)
		}
	}*/

	//var post Post
	//db.Table("posts").
	//	Select("posts.*").
	//	Joins("left join comments on comments.post_id = posts.id").
	//	Group("posts.id").
	//	Order("count(comments.id) desc").
	//	Preload("Comments").
	//	Take(&post)
	//
	//fmt.Printf("评论最多的文章: %+v\n", post)
	//fmt.Printf("评论数量: %d\n", len(post.Comments))

	/**
	题目3：钩子函数
	继续使用博客系统的模型。
	要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	*/
	//创建post
	/*db.Create(&Post{
		Title:  "测试",
		UserId: 1,
	})*/

	// 删除某条评论
	var comment Comment
	db.First(&comment, 1)
	db.Delete(&comment)

}

// Post 钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.First(&user, p.UserId)
	tx.Model(&user).UpdateColumn("post_count", user.PostCount+1)
	return
}

// Comment 的 AfterDelete 钩子
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var post Post
	tx.First(&post, c.PostId)

	var count int
	tx.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&count)

	if count == 0 {
		tx.Model(&post).UpdateColumn("comment_status", "无评论")
	}

	return nil
}
