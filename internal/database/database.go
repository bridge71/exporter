package database

import (
	"context"
	"database/sql"
	"exporter/internal/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	CreateUser(user *models.User) error
	FindUser(user *models.User, email string)
	FindFile(file *models.File)
	ModifyUserPass(user *models.User, passwordHash string)
	SaveAcct(acct *models.Acct) error
	SaveAcctBank(acctBank *models.AcctBank) error
	FindAcct(accts *[]models.Acct)
	FindAcctBank(acctBanks *[]models.AcctBank)
	FindAcctBankById(acctBanks *[]models.AcctBank, acctId uint)
	FindAcctBankByAccName(acctBanks *[]models.AcctBank, accName string)
	FindAcctBankByAccNum(acctBanks *[]models.AcctBank, accNum string)
	FindAcctBankByCurrency(acctBanks *[]models.AcctBank, currency string)
	FindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string)
	FzzFindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string)
	FindAcctBankBySwiftCode(acctBanks *[]models.AcctBank, swiftCode string)
	FindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string)
	FzzFindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string)
	FzzFindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string)
	FindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string)
	FindAcctByAcctCode(accts *[]models.Acct, acctCode string)
	FindAcctByAcctAbbr(accts *[]models.Acct, acctAbbr string)
	FindAcctByEtyAbbr(accts *[]models.Acct, EtyAbbr string)
	FindAcctByAcctName(accts *[]models.Acct, acctName string)
	FindAcctByAcctAddr(accts *[]models.Acct, acctAddr string)
	FzzFindAcctByAcctAddr(accts *[]models.Acct, acctAddr string)
	FindAcctByNation(accts *[]models.Acct, nation string)
	FindAcctByTaxType(accts *[]models.Acct, taxType string)
	FindAcctByTaxCode(accts *[]models.Acct, taxCode string)
	FindAcctByPhoneNum(accts *[]models.Acct, phoneNum string)
	FzzFindAcctByPhoneNum(accts *[]models.Acct, phoneNum string)
	FindAcctByEmail(accts *[]models.Acct, email string)
	FzzFindAcctByEmail(accts *[]models.Acct, email string)
	FindAcctByWebsite(accts *[]models.Acct, website string)
	FzzFindAcctByWebsite(accts *[]models.Acct, website string)
	FindAcctByRegDate(accts *[]models.Acct, regDate string)
	FzzFindAcctByRegDate(accts *[]models.Acct, regDate string)
	FindAcctByNotes(accts *[]models.Acct, notes string)
	FzzFindAcctByNotes(accts *[]models.Acct, notes string)

	CreateFile(file *models.File) error
	// FindFile(file *models.File, MD5 string)
}

type service struct {
	db     *sql.DB
	gormDB *gorm.DB
}

var (
	dbname     = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, dbname))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: newLogger,
	})
	gormDB.AutoMigrate(&models.User{}, &models.Acct{}, &models.AcctBank{}, &models.File{})

	dbInstance = &service{
		db:     db,
		gormDB: gormDB,
	}

	return dbInstance
}

func (s *service) CreateFile(file *models.File) error {
	return s.gormDB.Create(file).Error
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
