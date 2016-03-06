package item

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type repositor interface {
	Add(itemType *ItemType)
	GetAll() []*ItemType
	Drop(itemType *ItemType)
	FindByCode(code string) *ItemType
}

// in memory repository
type inMemRepository struct {
	list []*ItemType
}

func (repo *inMemRepository) Add(itemType *ItemType) {
	repo.list = append(repo.list, itemType)
}
func (repo *inMemRepository) Drop(itemType *ItemType) {

}
func (repo *inMemRepository) GetAll() []*ItemType {
	result := make([]*ItemType, 0, len(repo.list))
	for _, itemType := range repo.list {
		result = append(result, itemType)
	}
	return result
}
func (repo *inMemRepository) FindByCode(code string) *ItemType {
	return nil
}

// MySQL repository

type mySqlRepository struct{}

func (repo *mySqlRepository) connect() *sql.DB {
	db, err := sql.Open("mysql", "root:coppermine@tcp(192.168.21.129:3306)/pale_rabbit?charset=utf8&parseTime=true")
	checkErr(err)
	return db
}

func (repo *mySqlRepository) Drop(itemType *ItemType) {
	log.Println("drop")
	db := repo.connect()
	stmt, err := db.Prepare("DELETE FROM `pale_rabbit_master`.`item_type` WHERE code=?")
	checkErr(err)
	result, err := stmt.Exec(itemType.Code)
	checkErr(err)
	log.Println(result)
	db.Close()
}
func (repo *mySqlRepository) Add(itemType *ItemType) {
	log.Println("add")
	db := repo.connect()
	stmt, err := db.Prepare("INSERT INTO `pale_rabbit_master`.`item_type` (code, name, updated_by) values(?, ?, ?)")
	checkErr(err)
	result, err := stmt.Exec(itemType.Code, itemType.Name, "system")
	checkErr(err)
	log.Println(result)
	db.Close()
}

func (repo *mySqlRepository) GetAll() []*ItemType {
	db := repo.connect()
	rows, err := db.Query("SELECT code, name, updated_by, updated_at FROM pale_rabbit_master.item_type limit 100")
	checkErr(err)
	result := make([]*ItemType, 0, 100)

	for rows.Next() {
		itemType := &ItemType{}
		err = rows.Scan(&itemType.Code, &itemType.Name, &itemType.UpdatedBy, &itemType.UpdatedAt)
		
		attr_rows, err := db.Query("SELECT attr_code, attr_name, rank, updated_by, updated_at FROM pale_rabbit_master.item_type_attr where item_type=? order by rank", itemType.Code)
		checkErr(err)

		for attr_rows.Next() {
			attr := &itemTypeAttr{}
			err = attr_rows.Scan(&attr.AttrCode, &attr.AttrName, &attr.Rank, &attr.UpdatedBy, &attr.UpdatedAt)
			itemType.Attrs[attr.AttrCode] = attr

			log.Println(attr)
		}
		result = append(result, itemType)
	}

	db.Close()
	return result
}
func (repo *mySqlRepository) FindByCode(code string) *ItemType {
	db := repo.connect()

	rows, err := db.Query("SELECT code, name, updated_by, updated_at FROM pale_rabbit_master.item_type where code=?", code)
	checkErr(err)

	itemType := &ItemType{}
	if rows.Next() {
		err = rows.Scan(&itemType.Code, &itemType.Name, &itemType.UpdatedBy, &itemType.UpdatedAt)
		checkErr(err)
	}
	attr_rows, err := db.Query("SELECT attr_code, attr_name, rank, updated_by, updated_at FROM pale_rabbit_master.item_type_attr where item_type=? order by rank", code)
	checkErr(err)

	for attr_rows.Next() {
		attr := &itemTypeAttr{}
		err = attr_rows.Scan(&attr.AttrCode, &attr.AttrName, &attr.Rank, &attr.UpdatedBy, &attr.UpdatedAt)
		itemType.Attrs[attr.AttrCode] = attr

		log.Println(attr)
	}

	log.Println(itemType)
	db.Close()
	return itemType
}

var memRepository *inMemRepository = nil
var myRepository *mySqlRepository = nil

func GetRepository(repoType string) repositor {
	if repoType == "mysql" {
		if myRepository == nil {
			myRepository = &mySqlRepository{}
		}
		return myRepository
	} else {
		if memRepository == nil {
			memRepository = &inMemRepository{
				list: make([]*ItemType, 0, 32),
			}
		}
		return memRepository
	}
}

// -------
// private
// -------

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
