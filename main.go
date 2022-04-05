package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//CasbinDB()
	CasbinCSV()
}

func CasbinDB() {
	adapter, err := gormadapter.NewAdapter("mysql", "root:123456@tcp(0.0.0.0:3306)/go-casbin", true)

	if err != nil {
		fmt.Println("NewAdapter Error:", err)
		return
	}

	e, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		fmt.Println("NewEnforcer Error:", err)
		return
	}

	sub := "jasonzhou"
	obj := "data1"
	act := "read"

	//e.AddPolicy(sub, obj, act)添加策略到数据库
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		fmt.Println("Enforce Error:", err)
		return
	}

	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func addPolicy(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.AddPolicy(sub, obj, act)
	if err != nil {
		fmt.Printf("AddPolicy err: %v\n", err)
		return
	}
	if ok {
		fmt.Println("添加成功")
	} else {
		fmt.Printf("\"添加失败\": %v\n", "添加失败")
	}
}

func CasbinCSV() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println("NewENforcer Error:", err)
		return
	}
	sub := "jasonzhou"
	obj := "data13"
	act := "read"
	addPolicy(e, sub, obj, act)
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("enforce err: %v\n", err)
		return
	}
	if ok {
		fmt.Println("通过")
	} else {
		fmt.Println("未通过")
	}
}
