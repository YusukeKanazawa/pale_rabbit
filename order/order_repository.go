package order

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
//	"math"
)

// Object -> Store
// Object <- Store 
//

type Repositor interface {
	Add(*Order) error
	Query() error
	GetAll() ([]*Order, error)
}

type MemRepository struct {
	orderList []*Order
}


func GetRepository(repoType string) Repositor {
	if repoType == "memory" {
		return &MemRepository{
				orderList: make([]*Order, 0, 32),
			}
	} 
	//else {
	//	return &SqlRepository{}
	//}
	return nil
}

func (this *MemRepository) Add(order *Order) error {
	log.Println(order)
	this.orderList = append(this.orderList, order)
	log.Println(this.orderList)
	return nil
}

func (this *MemRepository) GetAll() ([]*Order, error) {
	return this.orderList, nil
}

func (this *MemRepository) Query() error {
	return nil
}



// 追加
//func (this *SqlRepository) Add(order *Order) error{
//	log.Println("add")
//	db := connect()
//
//	stmt, err := db.Prepare("insert into pale_rabbit.order (order_id) values(?)")
//	checkErr(err)
//	result, err := stmt.Exec(order.Id)
//	checkErr(err)
//	log.Println(result)
//	stmt, err = db.Prepare("insert into pale_rabbit.order_detail (order_id, detail_no) values(?, ?)")
//	checkErr(err)
//	for _, detail := range order.Details {
//		result, err := stmt.Exec(order.Id, detail.DetailNo)
//		
//		log.Println(result)
//		checkErr(err)
//	}
//	db.Close()
//	return nil
//}


//func (this *SqlRepository) GetAll() ([]*Order, error) {
//	var ROWS_LIMIT = 100
//	var ROWS_PER_PAGE = 20
//	count := 99
//	pages := count / ROWS_PER_PAGE
//	log.Println(pages)
//	
//	if (count == ROWS_LIMIT){
//		log.Println("more")
//	}
//	return nil, nil
//}

//func (this *Repository) GetAll(offset int, limit int) *[]Order {
//	db := connect()
//	// 最大100件でカウント
//	rows, err := db.Query("select count(*) as count from `pale_rabbit`.`order` limit 100")
//	
//	checkErr(err)
//	result := make([]Order, 0, 100)
//	
//	for rows.Next() {
//		order := NewOrder();
//		err = rows.Scan(&order.Id, &order.UpdatedAt)
//		details , err := db.Query("select order_id, detail_no, updated_at from `pale_rabbit`.`order_detail` where order_id='1' order by detail_no limit 100")
//		
//		for details.Next() {
//			detail := &Detail{}
//			err = details.Scan(detail.DetailNo, detail.UpdatedAt)
//			checkErr(err)
//			order.Details = append (order.Details, detail)
//		}
//		result = append(result, *order)
//	}
//
//	log.Println(result)
//	db.Close()
//	return &result
//}
//func (this *SqlRepository) Query() error {
//	return nil
//	}

/* 削除
 *
 */
//func (this *Repository) Remove(item *Item) {
//	log.Println("delete")
//	db := connect()
//	stmt, err := db.Prepare("delete from rabbit_house.mst_item where item_code = ?")
//	checkErr(err)
//	result, err := stmt.Exec(item.ItemCode)
//	checkErr(err)
//	log.Println(result)
//	db.Close()
//}

// -------
// private
// -------
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:coppermine@tcp(192.168.21.129:3306)/pale_rabbit?charset=utf8")
	checkErr(err)
	return db
}
