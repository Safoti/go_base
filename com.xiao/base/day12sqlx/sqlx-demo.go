package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/19
 * @Description
 **/
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db sql.DB
func main() {
	cfg :=mysql.Config{
	  User:  "root" ,
	  Passwd: "1234",
	  Net:    "tcp",
	  Addr:   "127.0.0.1:3306",
	  DBName: "sqlx",
  }
	var err error
	db, err := sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")



	//查询
	var albums[] Album
   var name="John Coltrane"  //模拟传参
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next(){
		var alb Album
		//赋值操作
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			 fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {

		fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	fmt.Println(albums)



}
