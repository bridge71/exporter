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

	FindUserById(user *models.User, userId uint)
	FindUserByEmail(user *models.User, email string)

	DeleteSaleDocReq(sale *models.Sale, DocReq *[]models.DocReq) error
	FindDocReqById(id uint, mercTypes *[]models.DocReq)
	FindById(id uint, model interface{})

	FindSalePrdtInfo(sale *models.Sale)
	FindSalesPrdtInfo(sale *[]models.Sale)
	DeleteSalePrdtInfo(sale *models.Sale, prdtInfo *models.PrdtInfo) error

	Delete(model interface{}) error
	Save(model interface{}) error
	Create(model interface{}) error
	Find(model interface{})

	DeleteEmpl(Empl *models.Empl) error
	SaveEmpl(Empl *models.Empl) error
	FindEmpl(Empls *[]models.Empl)

	DeletePackSpec(PackSpec *models.PackSpec) error
	SavePackSpec(PackSpec *models.PackSpec) error
	FindPackSpec(PackSpecs *[]models.PackSpec)

	DeleteBrand(brand *models.Brand) error
	SaveBrand(brand *models.Brand) error
	FindBrand(brands *[]models.Brand)
	FindBrandById(brand *models.Brand, brandId uint) error

	DeleteCat(cat *models.Cat) error
	SaveCat(cat *models.Cat) error
	FindCat(cats *[]models.Cat)
	FindCatById(cat *models.Cat, catId uint) error

	DeletePayMentMethod(payMentMethod *models.PayMentMethod) error
	SavePayMentMethod(payMentMethod *models.PayMentMethod) error
	FindPayMentMethods(payMentMethods *[]models.PayMentMethod)
	FindPayMentMethodByID(payMentMethod *models.PayMentMethod, id uint) error

	DeleteBankAccount(bankAccount *models.BankAccount) error
	SaveBankAccount(bankAccount *models.BankAccount) error
	FindBankAccount(bankAccounts *[]models.BankAccount)
	FindBankAccountById(bankAccounts *[]models.BankAccount, mercId uint)
	DeleteCust(cust *models.Cust) error
	SaveCust(cust *models.Cust) error
	FindCust(custs *[]models.Cust)
	FindCustById(custs *[]models.Cust, mercId uint)

	DeleteMerchant(merchant *models.Merchant) error
	SaveMerchant(merchant *models.Merchant) error
	FindMerchant(merchants *[]models.Merchant)
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	FindFile(file *models.File)
	SaveAcct(acct *models.Acct) error
	DeleteAcct(acct *models.Acct) error
	SaveAcctBank(acctBank *models.AcctBank) error
	DeleteAcctBank(acctBank *models.AcctBank) error
	FindAcct(accts *[]models.Acct)
	FirstAcct(id uint, acct *[]models.Acct)
	FindAcctBank(acctBanks *[]models.AcctBank)
	FindAcctBankById(acctBanks *[]models.AcctBank, acctId uint)

	CreateFile(file *models.File) error
	// FindFile(file *models.File, MD5 string)

	// Save functions of data directory
	SaveMercType(mercType *models.MercType) error
	SaveSuprType(suprType *models.SuprType) error
	SavePrdtType(prdtType *models.PrdtType) error
	SaveFoodAddType(foodAddType *models.FoodAddType) error
	SaveFeedAddType(feedAddType *models.FeedAddType) error
	SaveUnitMeas(unitMeas *models.UnitMeas) error
	SavePackType(packType *models.PackType) error
	SaveConType(conType *models.ConType) error
	SaveCurrency(currency *models.Currency) error
	SaveTradeTerm(tradeTerm *models.TradeTerm) error
	SaveNation(nation *models.Nation) error
	SavePort(port *models.Port) error
	SaveTaxType(taxType *models.TaxType) error
	SaveBrandType(brandType *models.BrandType) error
	SaveEduLevel(degree *models.EduLevel) error
	SaveDept(dept *models.Dept) error
	SavePosition(post *models.Position) error
	SaveQualStd(qualStd *models.QualStd) error
	SaveInvLoc(invLoc *models.InvLoc) error
	SaveDocReq(docReq *models.DocReq) error
	SavePayMth(payMth *models.PayMth) error
	SavePayLimit(payLimit *models.PayLimit) error
	SaveFinaDocStatus(finaDocStatus *models.FinaDocStatus) error
	SaveFinaDocType(finaDocType *models.FinaDocType) error
	SaveExpType(expType *models.ExpType) error
	SaveRates(rates *models.Rates) error
	SaveBussOrderSta(bussOrderSta *models.BussOrderSta) error

	// Find functions of data directory
	FindMercType(mercTypes *[]models.MercType)
	FindSuprType(suprTypes *[]models.SuprType)
	FindPrdtType(prdtTypes *[]models.PrdtType)
	FindFoodAddType(foodAddTypes *[]models.FoodAddType)
	FindFeedAddType(feedAddTypes *[]models.FeedAddType)
	FindUnitMeas(unitMeas *[]models.UnitMeas)
	FindPackType(packTypes *[]models.PackType)
	FindConType(conTypes *[]models.ConType)
	FindCurrency(currencies *[]models.Currency)
	FindTradeTerm(tradeTerms *[]models.TradeTerm)
	FindNation(nations *[]models.Nation)
	FindPort(ports *[]models.Port)
	FindTaxType(taxTypes *[]models.TaxType)
	FindBrandType(brandTypes *[]models.BrandType)
	FindEduLevel(degrees *[]models.EduLevel)
	FindDept(depts *[]models.Dept)
	FindPosition(posts *[]models.Position)
	FindQualStd(qualStds *[]models.QualStd)
	FindInvLoc(invLocs *[]models.InvLoc)
	FindDocReq(docReqs *[]models.DocReq)
	FindPayMth(payMths *[]models.PayMth)
	FindPayLimit(payLimits *[]models.PayLimit)
	FindFinaDocStatus(finaDocStatuses *[]models.FinaDocStatus)
	FindFinaDocType(finaDocTypes *[]models.FinaDocType)
	FindExpType(expTypes *[]models.ExpType)
	FindRates(rates *[]models.Rates)
	FindBussOrderSta(bussOrderStas *[]models.BussOrderSta)

	FindPayMthByID(payMth *models.PayMth, id uint) error
	FindPayLimitByID(payLimit *models.PayLimit, id uint) error

	// Delete functions of data directory
	DeleteMercType(mercType *models.MercType) error
	DeleteSuprType(suprType *models.SuprType) error
	DeletePrdtType(prdtType *models.PrdtType) error
	DeleteFoodAddType(foodAddType *models.FoodAddType) error
	DeleteFeedAddType(feedAddType *models.FeedAddType) error
	DeleteUnitMeas(unitMeas *models.UnitMeas) error
	DeletePackType(packType *models.PackType) error
	DeleteConType(conType *models.ConType) error
	DeleteCurrency(currency *models.Currency) error
	DeleteTradeTerm(tradeTerm *models.TradeTerm) error
	DeleteNation(nation *models.Nation) error
	DeletePort(port *models.Port) error
	DeleteTaxType(taxType *models.TaxType) error
	DeleteBrandType(brandType *models.BrandType) error
	DeleteEduLevel(degree *models.EduLevel) error
	DeleteDept(dept *models.Dept) error
	DeletePosition(post *models.Position) error
	DeleteQualStd(qualStd *models.QualStd) error
	DeleteInvLoc(invLoc *models.InvLoc) error
	DeleteDocReq(docReq *models.DocReq) error
	DeletePayMth(payMth *models.PayMth) error
	DeletePayLimit(payLimit *models.PayLimit) error
	DeleteFinaDocStatus(finaDocStatus *models.FinaDocStatus) error
	DeleteFinaDocType(finaDocType *models.FinaDocType) error
	DeleteExpType(expType *models.ExpType) error
	DeleteRates(rates *models.Rates) error
	DeleteBussOrderSta(bussOrderSta *models.BussOrderSta) error
	// Merchant 相关接口
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
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local&parseTime=true", username, password, host, port, dbname))
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
	// AutoMigrate all models
	gormDB.AutoMigrate(
		&models.Spot{},
		&models.File{},
		&models.PrdtInfo{},
		&models.DocReq{},
		&models.SuprType{},
		&models.PrdtType{},
		&models.FoodAddType{},
		&models.FeedAddType{},
		&models.UnitMeas{},
		&models.PackType{},
		&models.ConType{},
		&models.Currency{},
		&models.TradeTerm{},
		&models.Nation{},
		&models.Port{},
		&models.TaxType{},
		&models.BrandType{},
		&models.EduLevel{},
		&models.Dept{},
		&models.Position{},
		&models.QualStd{},
		&models.InvLoc{},
		&models.PayMth{},
		&models.PayLimit{},
		&models.FinaDocStatus{},
		&models.FinaDocType{},
		&models.ExpType{},
		&models.Rates{},
		&models.BussOrderSta{},
		&models.MercType{},
		&models.LoadingInfo{},
		&models.CostInfo{},
		&models.Empl{},
		&models.Cat{},
		&models.Brand{},
		&models.User{},
		&models.PayMentMethod{},
		&models.PackSpec{},
		&models.Merchant{},
		&models.BankAccount{},
		&models.Cust{},
		&models.Acct{},
		&models.AcctBank{},
		&models.Sale{},
		&models.Send{},
		&models.ShouldIn{},
		&models.In{},
	)

	dbInstance = &service{
		db:     db,
		gormDB: gormDB,
	}

	return dbInstance
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
