package drivers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySqlDB struct {
	SQL *gorm.DB
}

var MySql = &MySqlDB{}

func Connect(host, user, password, dbname string) (*MySqlDB, error) {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	connStr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbname)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	MySql.SQL = db

	return MySql, nil
}
