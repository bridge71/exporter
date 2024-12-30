package models

type User struct {
	Email        string `gorm:"column:email;unique;not null" json:"Email" `
	PasswordHash string `gorm:"column:passwordHash; not null" json:"Password"`
	UserName     string `gorm:"column:userName;not null" json:"UserName" `
	UserId       uint   `gorm:"column:userId;primaryKey" json:"UserId"`
	Priority     uint   `gorm:"default:1" json:"Priority"`
}

type Acct struct { // 会计实体
	AcctCode    string     `gorm:"column:acctCode;not null; unique" form:"AcctCode" binding:"required"`
	AcctAbbr    string     `gorm:"column:acctAbbr;not null" form:"AcctAbbr" binding:"required"`
	EtyAbbr     string     `gorm:"column:etyAbbr;not null" form:"EtyAbbr" binding:"required"`
	AcctName    string     `gorm:"column:acctName;not null" form:"AcctName" binding:"required"`
	AcctEngName string     `gorm:"column:acctEngName;" form:"AcctEngName"`
	AcctAddr    string     `gorm:"column:acctAddr;" form:"AcctAddr"`
	Nation      string     `gorm:"column:nation;" form:"Nation"`
	TaxType     string     `gorm:"column:taxType;" form:"TaxType"`
	TaxCode     string     `gorm:"column:taxCode;" form:"TaxCode"`
	PhoneNum    string     `gorm:"column:phoneNum;" form:"PhoneNum"`
	Email       string     `gorm:"column:email;" form:"Email"`
	Website     string     `gorm:"column:website;" form:"Website"`
	RegDate     string     `gorm:"column:regDate;" form:"RegDate"`
	Notes       string     `gorm:"column:notes;" form:"Notes"`
	FileName    string     `gorm:"column:fileName" form:"FileName"`
	AcctBanks   []AcctBank `gorm:"foreignKey:AcctId" form:"AcctBanks[]"`
	FileId      uint       `gorm:"column:fileId" form:"FileId"`
	AcctId      uint       `gorm:"column:acctId;  primaryKey" form:"AcctId"`
	// IsUpload    bool       `gorm:"column:isUpload; not null" form:"IsUpload" binding:"required"`
}

type AcctBank struct {
	AccName    string `gorm:"column:accName; not null" form:"AccName" binding:"required"`
	AccNum     string `gorm:"column:accNum; not null; unique" form:"AccNum" binding:"required"`
	Currency   string `gorm:"column:currency" form:"Currency"`
	BankName   string `gorm:"column:bankName" form:"BankName"`
	BankNum    string `gorm:"column:bankNum" form:"BankNum"`
	SwiftCode  string `gorm:"column:swiftCode" form:"SwiftCode"`
	BankAddr   string `gorm:"column:bankAddr" form:"BankAddr"`
	Notes      string `gorm:"column:notes" form:"Notes"`
	AcctName   string `gorm:"column:acctName" form:"AcctName"`
	FileName   string `gorm:"column:fileName" form:"FileName"`
	FileId     uint   `gorm:"column:fileId" form:"FileId"`
	AcctId     uint   `gorm:"column:acctId" form:"AcctId"`
	AcctBankId uint   `gorm:"column:acctBankId; primaryKey; " form:"AcctBankId"`
	// IsUpload   bool   `gorm:"column:isUpload; not null" form:"IsUpload" binding:"required"`
}

type File struct {
	Name   string `gorm:"column:name; not null" json:"Name"`
	MD5    string `gorm:"column:MD5;  not null" json:"MD5"`
	Suffix string `gorm:"column:suffix; not null" json:"Suffix"`
	FileId uint   `gorm:"column:fileId; primaryKey" form:"FileId" binding:"required"`
}
type SuprType struct {
	SuprType   string `gorm:"column:suprType; not null; unique" json:"SuprType"`
	SuprTypeId uint   `gorm:"column:suprTypeId; primaryKey; " form:"SuprTypeId"`
}

type PrdtType struct {
	PrdtType   string `gorm:"column:prdtType; not null; unique" json:"PrdtType"`
	PrdtTypeId uint   `gorm:"column:prdtTypeId; primaryKey; " form:"PrdtTypeId"`
}

type FoodAddType struct {
	FoodAddType   string `gorm:"column:foodAddType; not null; unique" json:"FoodAddType"`
	FoodAddTypeId uint   `gorm:"column:foodAddTypeId; primaryKey; " form:"FoodAddTypeId"`
}

type FeedAddType struct {
	FeedAddType   string `gorm:"column:feedAddType; not null; unique" json:"FeedAddType"`
	FeedAddTypeId uint   `gorm:"column:feedAddTypeId; primaryKey; " form:"FeedAddTypeId"`
}

type UnitMeas struct {
	UnitMeas   string `gorm:"column:unitMeas; not null; unique" json:"UnitMeas"`
	UnitMeasId uint   `gorm:"column:unitMeasId; primaryKey; " form:"UnitMeasId"`
}

type PackType struct {
	PackType   string `gorm:"column:packType; not null; unique" json:"PackType"`
	PackTypeId uint   `gorm:"column:packTypeId; primaryKey; " form:"PackTypeId"`
}

type ConType struct {
	ConType   string `gorm:"column:conType; not null; unique" json:"ConType"`
	ConTypeId uint   `gorm:"column:conTypeId; primaryKey; " form:"ConTypeId"`
}

type Currency struct {
	Currency   string `gorm:"column:currency; not null; unique" json:"Currency"`
	CurrencyId uint   `gorm:"column:currencyId; primaryKey; " form:"CurrencyId"`
}

type TradeTerm struct {
	TradeTerm   string `gorm:"column:tradeTerm; not null; unique" json:"TradeTerm"`
	TradeTermId uint   `gorm:"column:tradeTermId; primaryKey; " form:"TradeTermId"`
}

type Nation struct {
	Nation   string `gorm:"column:nation; not null; unique" json:"Nation"`
	NationId uint   `gorm:"column:nationId; primaryKey; " form:"NationId"`
}

type Port struct {
	Port   string `gorm:"column:port; not null; unique" json:"Port"`
	PortId uint   `gorm:"column:portId; primaryKey; " form:"PortId"`
}

type TaxType struct {
	TaxType   string `gorm:"column:taxType; not null; unique" json:"TaxType"`
	TaxTypeId uint   `gorm:"column:taxTypeId; primaryKey; " form:"TaxTypeId"`
}

type BrandType struct {
	BrandType   string `gorm:"column:brandType; not null; unique" json:"BrandType"`
	BrandTypeId uint   `gorm:"column:brandTypeId; primaryKey; " form:"BrandTypeId"`
}

type Degree struct {
	Degree   string `gorm:"column:degree; not null; unique" json:"Degree"`
	DegreeId uint   `gorm:"column:degreeId; primaryKey; " form:"DegreeId"`
}

type Dept struct {
	Dept   string `gorm:"column:dept; not null; unique" json:"Dept"`
	DeptId uint   `gorm:"column:deptId; primaryKey; " form:"DeptId"`
}

type Post struct {
	Post   string `gorm:"column:post; not null; unique" json:"Post"`
	PostId uint   `gorm:"column:postId; primaryKey; " form:"PostId"`
}

type QualStd struct {
	QualStd   string `gorm:"column:qualStd; not null; unique" json:"QualStd"`
	QualStdId uint   `gorm:"column:qualStdId; primaryKey; " form:"QualStdId"`
}

type InvLoc struct {
	InvLoc   string `gorm:"column:invLoc; not null; unique" json:"InvLoc"`
	InvLocId uint   `gorm:"column:invLocId; primaryKey; " form:"InvLocId"`
}

type DocReq struct {
	DocReq   string `gorm:"column:docReq; not null; unique" json:"DocReq"`
	DocReqId uint   `gorm:"column:docReqId; primaryKey; " form:"DocReqId"`
}

type PayMth struct {
	PayMth   string `gorm:"column:payMth; not null; unique" json:"PayMth"`
	PayMthId uint   `gorm:"column:payMthId; primaryKey; " form:"PayMthId"`
}

type PayLimit struct {
	PayLimit   string `gorm:"column:payLimit; not null; unique" json:"PayLimit"`
	PayLimitId uint   `gorm:"column:payLimitId; primaryKey; " form:"PayLimitId"`
}

type FinaDocStatus struct {
	FinaDocStatus   string `gorm:"column:finaDocStatus; not null; unique" json:"FinaDocStatus"`
	FinaDocStatusId uint   `gorm:"column:finaDocStatusId; primaryKey; " form:"FinaDocStatusId"`
}

type FinaDocType struct {
	FinaDocType   string `gorm:"column:finaDocType; not null; unique" json:"FinaDocType"`
	FinaDocTypeId uint   `gorm:"column:finaDocTypeId; primaryKey; " form:"FinaDocTypeId"`
}

type ExpType struct {
	ExpType   string `gorm:"column:expType; not null; unique" json:"ExpType"`
	ExpTypeId uint   `gorm:"column:expTypeId; primaryKey; " form:"ExpTypeId"`
}

type Rates struct {
	Rates   string `gorm:"column:rates; not null; unique" json:"Rates"`
	RatesId uint   `gorm:"column:ratesId; primaryKey; " form:"RatesId"`
}

type BussOrderSta struct {
	BussOrderSta   string `gorm:"column:bussOrderSta; not null; unique" json:"BussOrderSta"`
	BussOrderStaId uint   `gorm:"column:bussOrderStaId; primaryKey; " form:"BussOrderStaId"`
}

type MercType struct {
	MercType   string `gorm:"column:mercType; not null; unique" json:"MercType"`
	MercTypeId uint   `gorm:"column:mercTypeId; primaryKey; " form:"MercTypeId"`
}
type Message struct {
	Acct          []Acct
	AcctBank      []AcctBank
	MercType      []MercType
	SuprType      []SuprType
	PrdtType      []PrdtType
	FoodAddType   []FoodAddType
	FeedAddType   []FeedAddType
	UnitMeas      []UnitMeas
	PackType      []PackType
	ConType       []ConType
	Currency      []Currency
	TradeTerm     []TradeTerm
	Nation        []Nation
	Port          []Port
	TaxType       []TaxType
	BrandType     []BrandType
	Degree        []Degree
	Dept          []Dept
	Post          []Post
	QualStd       []QualStd
	InvLoc        []InvLoc
	DocReq        []DocReq
	PayMth        []PayMth
	PayLimit      []PayLimit
	FinaDocStatus []FinaDocStatus
	FinaDocType   []FinaDocType
	ExpType       []ExpType
	Rates         []Rates
	BussOrderSta  []BussOrderSta
	RetMessage    string
	User          User
}
