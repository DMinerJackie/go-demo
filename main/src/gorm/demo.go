package main

import (
	db "database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//type DqmUserRole struct {
//	ID        int64     `gorm:"column:id;primary_key" json:"id"`
//	UserId    string    `gorm:"column:user_id" json:"user_id"`
//	RoleId    string    `gorm:"column:role_id" json:"role_id"`
//	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
//	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
//}
//
////func (Test) TableName() string {
////	return "test"
////}
//

func test() {
	for {

		msg := make(chan int)
		defer func() {
			if recover() != nil {
				// the return result can be altered
				// in a defer function call
				return
			}
		}()
		close(msg)
		msg <- 1

		time.Sleep(time.Second * 2)
	}
}
func main() {

	//test()
	//fmt.Println("hhh")

	var rows *db.Rows
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(0)
	var name string
	//j:=1
	//for i := 0; i<3; i++  {
	rows, _ = db.Query("select name from test where name = 'jackie'")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name)
		//j++
		fmt.Println("close")
		//if j==4 {
		//	return
		//}
	}
	fmt.Println("=========")
	//rows.Close()
	//}

	rows, _ = db.Query("select name from test where name = 'hello'")
	for rows.Next() {
		fmt.Println("last close")
		rows.Scan(&name)
	}

	fmt.Println("=========")

	rows, _ = db.Query("select name from test where name = 'hello'")
	for rows.Next() {
		fmt.Println("last close")
		rows.Scan(&name)
	}

	//row,_ := db.Query("select * from test") //此操作将一直阻塞
	//
	////正确
	//db.SetMaxOpenConns(1)
	//r,_ := db.Query("select * from test")
	//r.Close() //将连接的所属权归还，释放连接
	//row,_ = db.Query("select * from test")
	////other op
	//row.Close()

	//if err != nil {
	//	panic("failed to connect database")
	//}
	//defer db.Close()
	//
	//// 如果不设置则gorm默认表名会在后面加s（注意struct的名字要和表名一致，否则会找不到）
	//db.SingularTable(true)
	//
	////test := &Test{
	////	ID:3,
	////	Name:"jackie",
	////	Age:18,
	////}
	////db.Create(test)
	//
	////db.AutoMigrate(&DqmUserRole{})
	//
	//// 判定表是否存在
	//var result bool
	//result = db.HasTable(&DqmUserRole{})
	//fmt.Println(result)
	//
	////db.CreateTable(Product{})
	//
	//user := &DqmUserRole{ID: 1}
	//db.Model(user).Update("role_id", "1,2").Update("created_at", time.Now())

	//userRole := UserRole{UserId: 3, RoleId: 3, CreatedAt: time.Now(), UpdatedAt:time.Now()}
	//db.NewRecord(userRole)
	//db.Create(&userRole)
	//db.NewRecord(userRole)
	//db.Where("user_id = ?", 1).Find(&UserRole{})
	//fmt.Printf("==%v", db.Last(&userRole))

	//var useRole DqmUserRole
	//db.Where("user_id = ?", "1").First(&useRole)
	//fmt.Println(useRole.RoleId)

}
