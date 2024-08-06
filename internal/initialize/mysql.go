package initialize

import (
	"fmt"

	"github.com/marktran77/go-ecomerce-backend-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func SetPool() {

}

func migrateTables() {

}
