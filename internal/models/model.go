package models

type User struct {
	Email        string `gorm:"column:email;unique;not null" json:"Email" binding:"required"`
	PasswordHash string `gorm:"column:passwordHash; not null" json:"Password" binding:"required"`
	UserName     string `gorm:"column:userName;not null" json:"UserName" `
	UserId       uint   `gorm:"column:userId;primaryKey" `
	Priority     uint   `gorm:"default:1" json:"Priority"`
}

type Acct struct { // 会计实体
	AcctCode    string     `gorm:"column:acctCode;not null" form:"AcctCode" binding:"required"`
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
	AcctBanks   []AcctBank `gorm:"foreignKey:AcctId" form:"AcctBanks[]"`
	FileId      uint       `gorm:"column:fileId"`
	AcctId      uint       `gorm:"column:acctId;  primaryKey"`
	IsUpload    bool       `gorm:"column:isUpload; not null" form:"IsUpload" binding:"required"`
}

type AcctBank struct {
	AccName    string `gorm:"column:accName; not null" form:"AccName" binding:"required"`
	AccNum     string `gorm:"column:accNum; not null" form:"AccNum" binding:"required"`
	Currency   string `gorm:"column:currency" form:"Currency"`
	BankName   string `gorm:"column:bankName" form:"BankName"`
	BankNum    string `gorm:"column:bankNum" form:"BankNum"`
	SwiftCode  string `gorm:"column:swiftCode" form:"SwiftCode"`
	BankAddr   string `gorm:"column:bankAddr" form:"BankAddr"`
	Notes      string `gorm:"column:notes" form:"Notes"`
	FileId     uint   `gorm:"column:fileId"`
	AcctId     uint   `gorm:"column:AcctId"`
	AcctBankId uint   `gorm:"column:acctBankId; primaryKey; " form:"AcctBankId"`
	IsUpload   bool   `gorm:"column:isUpload; not null" form:"IsUpload" binding:"required"`
}

type File struct {
	Name   string `gorm:"column:Name; not null" json:"Name"`
	MD5    string `gorm:"column:MD5;  not null" json:"MD5"`
	Suffix string `gorm:"column:Suffix; not null" json:"Suffix"`
	FileId uint   `gorm:"column:FileId; primaryKey"`
}

type Message struct {
	Acct       []Acct
	AcctBank   []AcctBank
	RetMessage string
	User       User
}
