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
func Connect(c consts.MySQLConf) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database))
	if err != nil {
		panic(err)
	}
	DB = db
	go StartAntiTimeout()
}

func StartAntiTimeout() {
	for {
		err := DB.Ping()
		if err != nil {
			logger.Error("Failed to ping, is the Database dead ?")
		}
		time.Sleep(time.Second * 10)
	}
}
