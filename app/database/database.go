package database

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database = &gorm.DB{}
)

// Session session
type Session struct {
	Database *gorm.DB
}

// Configuration config mysql
type Configuration struct {
	Host     string
	Username string
	Password string
	Name     string
}

// New open initialize a new db connection.
func New(config *Configuration) (*Session, error) {
	dns := url.URL{
		User:   url.UserPassword(config.Username, config.Password),
		Scheme: "postgres",
		Host:   config.Host,
		Path:   config.Name,
		RawQuery: (&url.Values{
			"sslmode":  []string{"disable"},
			"TimeZone": []string{"Asia/Bangkok"},
		}).Encode(),
	}

	database, err := gorm.Open(postgres.New(postgres.Config{DSN: dns.String()}))
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return &Session{Database: database}, nil
}

// Get get
func Get(c *gin.Context) *gorm.DB {
	if db, ok := c.Get("db"); ok {
		return db.(*gorm.DB)
	}
	return database
}

// Begin begin
func Begin() *gorm.DB {
	return database.Begin()
}

// Set set
func Set(db *gorm.DB) {
	database = db
}
