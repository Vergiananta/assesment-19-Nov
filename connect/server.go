package connect

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"superapp/models"
	"time"
)

var (
	cfg *viper.Viper
	db  *gorm.DB
	err error
)

type Connect interface {
	SqlDb() *gorm.DB
	Config() *viper.Viper
	ApiServer(addr []string) string
}

type connect struct{}

func NewConnect() Connect {
	return &connect{}
}

func (i *connect) SqlDb() *gorm.DB {
	dbUser := i.Config().GetString("DB_USER")
	dbPassword := i.Config().GetString("DB_PASSWORD")
	dbHost := i.Config().GetString("DB_HOST")
	dbPort := i.Config().GetString("DB_PORT")
	dbName := i.Config().GetString("DB_NAME")

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", dbHost, dbUser, dbPassword, dbName, dbPort)))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Customer{}, &models.Merchant{}, &models.Transaction{})

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	defer sqlDB.SetMaxIdleConns(1)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	defer sqlDB.SetMaxOpenConns(16)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	defer sqlDB.SetConnMaxLifetime(time.Minute)

	return db
}

func (i *connect) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg
}

func (i *connect) ApiServer(addr []string) string {
	switch len(addr) {
	case 0:
		if port := i.Config().GetString("PORT"); port != "" {
			debugPrint("Environment variable PORT=" + port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8081 by default")
		return ":8081"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func debugPrint(format string, values ...interface{}) {
	fmt.Println(format)
}
