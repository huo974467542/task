package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
	"github.com/jmoiron/sqlx"
)

type students struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Grade string `db:"grade"`
}

type Accounts struct {
	Id      int `db:"id"`
	Balance int `db:"balance"` // 修改为实际数据库列名
}

type Transactions struct {
	Id            int
	fromAccountId int
	toAccountId   int
	Amount        int
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(192.168.232.144:3306)/go-test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	//r, err := Db.Exec("insert into students(name, age, grade)values(?, ?, ?)", "stu001", 18, "一年级")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//id, err := r.LastInsertId()
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//
	//fmt.Println("insert succ:", id)

	//var student []students
	//err := Db.Select(&student, "select * from students where age>=?", 18)
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//
	//fmt.Println("select succ:", student)

	//res, err := Db.Exec("update students set grade='四年级' where name=?", "stu001")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//row, err := res.RowsAffected()
	//if err != nil {
	//	fmt.Println("rows failed, ", err)
	//}
	//fmt.Println("update succ:", row)

	/*	res, err := Db.Exec("delete from students where age=?", 18)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}

		row, err := res.RowsAffected()
		if err != nil {
			fmt.Println("rows failed, ", err)
		}

		fmt.Println("delete succ: ", row)*/

	//查询accounts表中id为1的信息
	/*	conn, err := Db.Begin()
		if err != nil {
			fmt.Println("begin transaction failed: ", err)
			return
		}

		var account Accounts
		err = conn.QueryRow("SELECT id, balance FROM accounts WHERE id = ?", 1).Scan(&account.Id, &account.Balance)

		if err != nil {
			fmt.Println("query account failed: ", err)
			conn.Rollback()
			return
		}

		if account.Balance > 100 {
			// 账户1减100
			result1, err1 := conn.Exec("UPDATE accounts SET balance = balance - 100 WHERE id = ?", 1)
			if err1 != nil {
				fmt.Println("update account 1 failed: ", err1)
				conn.Rollback()
				return
			}
			rowsAffected1, _ := result1.RowsAffected()
			fmt.Println("account 1 updated, rows affected:", rowsAffected1)

			// 账户2增加100
			result2, err2 := conn.Exec("UPDATE accounts SET balance = balance + 100 WHERE id = ?", 2)
			if err2 != nil {
				fmt.Println("update account 2 failed: ", err2)
				conn.Rollback()
				return
			}
			rowsAffected2, _ := result2.RowsAffected()
			fmt.Println("account 2 updated, rows affected:", rowsAffected2)

			// 保存交易信息
			_, err3 := conn.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)", 1, 2, 100)
			if err3 != nil {
				fmt.Println("insert transaction failed: ", err3)
				conn.Rollback()
				return
			}

		} else {
			fmt.Println("balance not enough")
			return
		}

		err = conn.Commit()
		if err != nil {
			fmt.Println("commit transaction failed: ", err)
			conn.Rollback()
		}*/

	//求 ：
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	type Employees struct {
		Id         int
		Name       string
		Department string
		Salary     float32
	}
	var employees []Employees
	Db.Select(&employees, "select * from employees where department = ?", "技术部")
	for _, v := range employees {
		fmt.Println(v.Id, v.Name, v.Department, v.Salary)
	}
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var emp Employees
	Db.QueryRow("select * from employees where salary = (select max(salary) from employees)").Scan(&emp.Id, &emp.Name, &emp.Department, &emp.Salary)
	fmt.Println(emp.Id, emp.Name, emp.Department, emp.Salary)

	//定义一个 Book 结构体，包含与 books 表对应的字段。
	type Book struct {
		Id     int
		Title  string
		Author string
		Price  float64
	}
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
	books := make([]Book, 0)
	err := Db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		fmt.Println(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}

}
