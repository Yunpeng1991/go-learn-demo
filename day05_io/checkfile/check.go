package checkfile

import (
	"fmt"
	"learn-demo/day05_io/database"
	"learn-demo/day05_io/ops"
)

type CompareFile struct {
	Data []byte

	CompareData []byte

	Differ bool

	DiffCount int64

	DifferData *ops.DiffData
}

func (cf *CompareFile) Compare() (res bool, diffCount int64) {
	cf.Differ = false
	cf.DiffCount = 100
	return cf.Differ, cf.DiffCount
}

func (cf *CompareFile) ReadFile(path string) (data []byte, err error) {
	return nil, nil
}

func (cf *CompareFile) WriteFile(data []byte) error {
	return nil
}

func (cf *CompareFile) QueryDiff(id int64) *ops.DiffData {
	return nil
}

func (cf *CompareFile) SaveDiff() int {
	db := &database.DBFactory{
		DriverName:   "mysql",
		DatabaseName: "root:root@tcp(10.211.55.6:3307)/yp"}
	connection := db.GetConnection()
	defer connection.Close()
	tx, err := connection.Begin()
	if err != nil {
		panic(err)
	}
	sql := fmt.Sprintf("insert DIFF_DATA ( path, diff_count, create_time) values ('%s',%d,%d)", cf.DifferData.Path, cf.DifferData.DiffCount, cf.DifferData.CreateTime)
	fmt.Println(sql)
	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	//exception handler
	errTx := tx.Commit()
	if errTx == nil {
		fmt.Println("transaction commit successfully")
	} else {
		fmt.Println("sql execute error ", errTx.Error())
		panic(errTx)
	}

	return 0
}
