package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


func main() {

	db, err := sql.Open("mysql", "root:Xiaomi_123@/test2?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("root", "工程", time.Now().Unix())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)



	//更新数据
	//stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("rootupdate", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo WHERE username ='root' order by department desc limit 10 ")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}