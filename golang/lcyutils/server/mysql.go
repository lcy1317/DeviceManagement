package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBConnect() {
	DB, err := sql.Open("mysql", "root:SEUjyh1317@tcp(127.0.0.1:3306)/device?charset=utf8mb4&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail", err)
		return
	}
	fmt.Println("connnect success")
}

type device struct {
	categorie string
	dtype     string
	time      string
	belong    string
	ifmark    string
	status    string
	username  string
	yikatong  string
	position  string
}

func QueryDevice(deviceID int) {
	var d device
	//单行查询操作
	queryURL := "select type from devices where device_id = ?"
	fmt.Println(queryURL)
	//DB.QueryRow(queryURL).Scan(d.categorie, d.dtype, d.time, d.belong, d.ifmark, d.status, d.username, d.yikatong, d.position)
	DB.QueryRow(queryURL, deviceID).Scan(&d.dtype)
	fmt.Println(d.dtype)
}
