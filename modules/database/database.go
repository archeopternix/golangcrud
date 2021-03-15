// Scuffolding for echo framework:  https://echo.labstack.com/
package model

import "github.com/sonyarouje/simdb/db"

func InitializeDb() (err error, driver *db.Driver) {
	driver, err = db.New("data")
	if err != nil {
		panic(err)
	}
	return err, driver
}

var Database *db.Driver

func init() {
	var err error
	err, Database = InitializeDb()

	if err != nil {
		panic(err)
	}
}
