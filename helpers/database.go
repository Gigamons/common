package helpers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/logger"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// DB Database
var DB *sql.DB

// Connect to MySQL Database
func Connect(c *consts.MySQLConf) {
	go startAntiTimeout(c)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database))
	if err != nil {
		logger.Errorln("Could not connect to Database.", err)
		return
	}
	DB = db
}

func startAntiTimeout(c *consts.MySQLConf) {
	var err error
	for {
		err = DB.Ping()
		if err != nil {
			logger.Errorln("Failed to ping, is the Database dead ?")
			Connect(c)
			break
		}
		time.Sleep(time.Second * 10)
	}
}
