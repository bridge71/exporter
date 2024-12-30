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

type Message struct {
	Acct       []Acct
	AcctBank   []AcctBank
	RetMessage string
	User       User
}
