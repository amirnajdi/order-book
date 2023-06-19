package DBconnection

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBconnectionConfig struct {
	Url      string
	Port     string
	Username string
	Password string
	DBname   string
}

var lock = &sync.Mutex{}
var connectionInstance *gorm.DB

func Connect(config DBconnectionConfig) (*gorm.DB, error) {
	if connectionInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		var err error
		connectionInstance, err = createConnection(&config)
		if err != nil {
			return nil, err
		}
	}

	return connectionInstance, nil
}

func createConnection(config *DBconnectionConfig) (*gorm.DB, error) {
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Url, config.Port, config.DBname,
	)

	return gorm.Open(mysql.Open(connection), &gorm.Config{})
}

func GetConnectionInstance() *gorm.DB {
	return connectionInstance
}
