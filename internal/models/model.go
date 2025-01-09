package models

import (
	"gorm.io/gorm"
)

type Message struct {
	ShouldOut     []ShouldOut
	Buy           []Buy
	Purrec        []Purrec
	Out           []Out
	In            []In
	ShouldIn      []ShouldIn
	Spot          []Spot
	Sale          []Sale
	Send          []Send
	LoadingInfo   []LoadingInfo
	PrdtInfo      []PrdtInfo
	CostInfo      []CostInfo
	Users         []User
	Empl          []Empl
	Acct          []Acct
	AcctBank      []AcctBank
	Brand         []Brand
	PayMentMethod []PayMentMethod
	PackSpec      []PackSpec
	MercType      []MercType
	Cat           []Cat
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
	EduLevel      []EduLevel
	Dept          []Dept
	Position      []Position
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
	Merchant      []Merchant
	Cust          []Cust
	BankAccount   []BankAccount
	RetMessage    string
	User          User
}

type Buy struct {
	OrderNum        string `form:"OrderNum"`  // 订单编号
	OrderDate       string `form:"OrderDate"` // 订单日期
	AcctID          uint   `form:"AcctID"`    // 销售方,绑定id 查找接口 /find/acct
	Acct            Acct
	MerchantID      uint `form:"MerchantID"` // 购买方,绑定id 查找接口 /find/merchant
	Merchant        Merchant
	QualStd         string `form:"QualStd"`                           // 质量标准 从数据字典里选 /find/QualStd 接口返回一个json，里面有QualStd的json数组，选择其中的QualStd
	BillValidity    string `form:"BillValidity"`                      // 账单有效期
	BussOrderSta    string `form:"BussOrderSta"`                      // 单据状态，从数据字典里选 /find/BussOrderSta 接口返回一个json，里面有BussOrderSta的json数组，选择其中的BussOrderSta
	StartShip       string `gorm:"column:startShip" form:"StartShip"` // 发货开始日期
	EndShip         string `gorm:"column:endShip" form:"EndShip"`     // 发货截止日期
	SrcPlace        string `gorm:"column:srcPlace" form:"SrcPlace"`   // 起运地, 从数据字典里选  /find/SrcPlace 同理
	Des             string `gorm:"column:des" form:"Des"`             // 目的地, 从数据字典里选  /find/SrcPlace 同理
	PayMentMethodID uint   `form:"PayMentMethodID"`                   // 付款方式，绑定id，查找接口 /find/PayMentMethod 同理
	PayMentMethod   PayMentMethod
	PrdtInfos       []PrdtInfo `gorm:"many2many:buyPrdtInfo" `          // 产品明细，多表关联，  查找接口 /find/PrdtInfos
	DocReq          []DocReq   `gorm:"many2many:buyDocReq"`             // 单据要求,多选 从数据字典里选 /find/DocReq
	TotAmt          uint       `gorm:"column:totAmt" form:"TotAmt"`     // 总金额
	Currency        string     `gorm:"column:currency" form:"Currency"` // 币种
	TotNum          uint       `gorm:"column:totNum" form:"TotNum"`     // 总件数
	PackSpecID      uint       `gorm:"column:packSpecID"`               // 包装规格 从数据字典里选 /find/PackSpec 同理
	PackSpec        PackSpec
	TotalNetWeight  string `gorm:"column:totalNetWeight" form:"TotalNetWeight"` // 总净重
	UnitMeas        string `gorm:"column:unitMeas" form:"UnitMeas"`             // 单位 从数据字典里选  /find/UnitMeas
	AcctBank        AcctBank
	AcctBankID      uint `gorm:"column:acctBankID" form:"AcctBankID"`       // 我方银行账户  绑定id，查找接口 /find/acctBank
	BankAccountID   uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 对方银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount     BankAccount
	Notes           string      `gorm:"column:notes" form:"Notes"` // 备注
	FileName        string      `gorm:"column:fileName" form:"FileName"`
	FileID          uint        `gorm:"column:fileID" form:"FileID"`
	Purrecs         []Purrec    `gorm:"many2many:buyPurrec" `
	ShouldOuts      []ShouldOut `gorm:"many2many:shouldOutBuy" `
	Outs            []Out       `gorm:"many2many:outBuy" `
	gorm.Model
}
type Sale struct {
	OrderNum        string `gorm:"column:orderNum" form:"OrderNum"`   // 订单编号
	OrderDate       string `gorm:"column:orderDate" form:"OrderDate"` // 订单日期
	AcctID          uint   `form:"AcctID"`                            // 销售方,绑定id 查找接口 /find/acct
	Acct            Acct
	MerchantID      uint `form:"MerchantID"` // 购买方,绑定id 查找接口 /find/merchant
	Merchant        Merchant
	QualStd         string `gorm:"column:qualStd" form:"QualStd"`           // 质量标准 从数据字典里选 /find/QualStd 接口返回一个json，里面有QualStd的json数组，选择其中的QualStd
	BillValidity    string `gorm:"column:billValidity" form:"BillValidity"` // 账单有效期
	BussOrderSta    string `gorm:"column:bussOrderSta" form:"BussOrderSta"` // 单据状态，从数据字典里选 /find/BussOrderSta 接口返回一个json，里面有BussOrderSta的json数组，选择其中的BussOrderSta
	StartShip       string `gorm:"column:startShip" form:"StartShip"`       // 发货开始日期
	EndShip         string `gorm:"column:endShip" form:"EndShip"`           // 发货截止日期
	SrcPlace        string `gorm:"column:srcPlace" form:"SrcPlace"`         // 起运地, 从数据字典里选  /find/SrcPlace 同理
	Des             string `gorm:"column:des" form:"Des"`                   // 目的地, 从数据字典里选  /find/SrcPlace 同理
	PayMentMethodID uint   `form:"PayMentMethodID"`                         // 付款方式，绑定id，查找接口 /find/PayMentMethod 同理
	PayMentMethod   PayMentMethod
	PrdtInfos       []PrdtInfo `gorm:"many2many:salePrdtInfo" `         // 产品明细，多表关联，  查找接口 /find/PrdtInfos
	DocReq          []DocReq   `gorm:"many2many:saleDocReq"`            // 单据要求,多选 从数据字典里选 /find/DocReq
	TotAmt          uint       `gorm:"column:totAmt" form:"TotAmt"`     // 总金额
	Currency        string     `gorm:"column:currency" form:"Currency"` // 币种
	TotNum          uint       `gorm:"column:totNum" form:"TotNum"`     // 总件数
	PackSpecID      uint       `form:"PackSpecID"`                      // 包装规格 从数据字典里选 /find/PackSpec 同理
	PackSpec        PackSpec
	TotalNetWeight  string `gorm:"column:totalNetWeight" form:"TotalNetWeight"` // 总净重
	UnitMeas        string `gorm:"column:unitMeas" form:"UnitMeas"`             // 单位 从数据字典里选  /find/UnitMeas
	AcctBankID      uint   `gorm:"column:acctBankID" form:"AcctBankID"`         // 我方银行账户  绑定id，查找接口 /find/acctBank
	AcctBank        AcctBank
	BankAccountID   uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 对方银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount     BankAccount
	Notes           string     `gorm:"column:notes" form:"Notes"` // 备注
	FileName        string     `gorm:"column:fileName" form:"FileName"`
	FileID          uint       `gorm:"column:fileID" form:"FileID"`
	Sends           []Send     `gorm:"many2many:saleSend" form:"Sends"`         // 销售发货单，多表关联, 查找接口
	ShouldIns       []ShouldIn `gorm:"many2many:shouldInSale" form:"ShouldIns"` // 应收账款单
	Ins             []In       `gorm:"many2many:inSale" form:"ins"`             // 收款单
	gorm.Model
}

type Purrec struct {
	SaleInvNum      string `gorm:"column:SaleInvNum" form:"SaleInvNum"`   // 销售发票号
	SaleInvDate     string `gorm:"column:SaleInvDate" form:"SaleInvDate"` // 销售发票日期
	Buys            []Buy  `gorm:"many2many:buyPurrec" `                  // 销售订单 多表关联
	Merchant1ID     uint   `form:"Merchant1ID"`
	Merchant2ID     uint   `form:"Merchant2ID"`
	Merchant3ID     uint   `form:"Merchant3ID"`
	Merchant1Name   string `form:"Merchant1Name"`
	Merchant2Name   string `form:"Merchant2Name"`
	Merchant3Name   string `form:"Merchant3Name"`
	AcctID          uint   `gorm:"column:AcctID" form:"AcctID"` // 发货人  通过id绑定，查找接口 /find/acct
	Acct            Acct
	SrcPlace        string `gorm:"column:SrcPlace" form:"SrcPlace"`     // 起运地, 从数据字典里选  /find/SrcPlace
	Des             string `gorm:"column:Des" form:"Des"`               // 目的地, 从数据字典里选  /find/SrcPlace
	ShipName        string `gorm:"column:ShipName" form:"ShipName"`     // 船名
	Voyage          string `gorm:"column:Voyage" form:"Voyage"`         // 航次
	TotNum          uint   `gorm:"column:TotNum" form:"TotNum"`         // 总件数
	PackSpecID      uint   `gorm:"column:packSpecID" form:"PackSpecID"` // 包装规格 从数据字典里选 /find/PackSpec
	PackSpec        PackSpec
	TotalNetWeight  uint   `gorm:"column:TotalNetWeight" form:"TotalNetWeight"`   // 总净重
	UnitMeas1       string `gorm:"column:UnitMeas1" form:"UnitMeas1"`             // 单位 从数据字典里选  /find/UnitMeas
	GrossWt         uint   `gorm:"column:GrossWt" form:"GrossWt"`                 // 总毛重
	UnitMeas2       string `gorm:"column:UnitMeas2" form:"UnitMeas2"`             // 单位 从数据字典里选  /find/UnitMeas
	TotVol          string `gorm:"column:TotVol" form:"TotVol"`                   // 总体积
	UnitMeas3       string `gorm:"column:Size" form:"UnitMeas3"`                  // 尺码
	BillLadNum      string `gorm:"column:BillLadNum" form:"BillLadNum"`           // 提单号
	DateOfShip      string `gorm:"column:DateOfShip" form:"DateOfShip"`           // 发货（开船）日期
	Note1           string `gorm:"column:note1" form:"Note1"`                     // 提单货物描述
	Note2           string `gorm:"column:note2" form:"Note2"`                     // 箱单货物描述
	PayMentMethodID uint   `gorm:"column:payMentMethodID" form:"PayMentMethodID"` // 付款方式，绑定id，查找接口 /find/PayMentMethod
	PayMentMethod   PayMentMethod
	AcctBankID      uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行  绑定id，查找接口 /find/acctBank
	AcctBank        AcctBank
	PrdtInfos       []PrdtInfo    `gorm:"many2many:purrecPrdtInfo" `    // 产品明细，多表关联，  查找接口 /find/PrdtInfos
	LoadingInfos    []LoadingInfo `gorm:"many2many:purrecLoadingInfo" ` // 装货明细，多表关联，查找接口 /find/LoadingInfos
	File1ID         uint          `form:"File1ID"`
	File2ID         uint          `form:"File2ID"`
	File1Name       string        `form:"File1Name"`
	File2Name       string        `form:"File2Name"`
	ShouldOuts      []ShouldOut   `gorm:"many2many:shouldOutPurrec" ` // 应收账款单
	Outs            []Out         `gorm:"many2many:outPurrec"`        // 收款单
	gorm.Model
}
type Send struct {
	SaleInvNum      string `gorm:"column:SaleInvNum" form:"SaleInvNum"`   // 销售发票号
	SaleInvDate     string `gorm:"column:SaleInvDate" form:"SaleInvDate"` // 销售发票日期
	Sales           []Sale `gorm:"many2many:saleSend" form:"Sales"`       // 销售订单 多表关联
	Merchant1ID     uint   `form:"Merchant1ID"`
	Merchant2ID     uint   `form:"Merchant2ID"`
	Merchant3ID     uint   `form:"Merchant3ID"`
	Merchant1Name   string `form:"Merchant1Name"`
	Merchant2Name   string `form:"Merchant2Name"`
	Merchant3Name   string `form:"Merchant3Name"`
	AcctID          uint   `gorm:"column:AcctID" form:"AcctID"` // 发货人  通过id绑定，查找接口 /find/acct
	Acct            Acct
	SrcPlace        string `gorm:"column:SrcPlace" form:"SrcPlace"`     // 起运地, 从数据字典里选  /find/SrcPlace
	Des             string `gorm:"column:Des" form:"Des"`               // 目的地, 从数据字典里选  /find/SrcPlace
	ShipName        string `gorm:"column:ShipName" form:"ShipName"`     // 船名
	Voyage          string `gorm:"column:Voyage" form:"Voyage"`         // 航次
	TotNum          uint   `gorm:"column:TotNum" form:"TotNum"`         // 总件数
	PackSpecID      uint   `gorm:"column:packSpecID" form:"PackSpecID"` // 包装规格 从数据字典里选 /find/PackSpec
	PackSpec        PackSpec
	TotalNetWeight  uint   `gorm:"column:TotalNetWeight" form:"TotalNetWeight"`   // 总净重
	UnitMeas1       string `gorm:"column:UnitMeas1" form:"UnitMeas1"`             // 单位 从数据字典里选  /find/UnitMeas
	GrossWt         uint   `gorm:"column:GrossWt" form:"GrossWt"`                 // 总毛重
	UnitMeas2       string `gorm:"column:UnitMeas2" form:"UnitMeas2"`             // 单位 从数据字典里选  /find/UnitMeas
	TotVol          string `gorm:"column:TotVol" form:"TotVol"`                   // 总体积
	UnitMeas3       string `gorm:"column:Size" form:"UnitMeas3"`                  // 尺码
	BillLadNum      string `gorm:"column:BillLadNum" form:"BillLadNum"`           // 提单号
	DateOfShip      string `gorm:"column:DateOfShip" form:"DateOfShip"`           // 发货（开船）日期
	Note1           string `gorm:"column:note1" form:"Note1"`                     // 提单货物描述
	Note2           string `gorm:"column:note2" form:"Note2"`                     // 箱单货物描述
	PayMentMethodID uint   `gorm:"column:payMentMethodID" form:"PayMentMethodID"` // 付款方式，绑定id，查找接口 /find/PayMentMethod
	PayMentMethod   PayMentMethod
	AcctBankID      uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行  绑定id，查找接口 /find/acctBank
	AcctBank        AcctBank
	PrdtInfos       []PrdtInfo    `gorm:"many2many:sendPrdtInfo" form:"PrdtInfos"`       // 产品明细，多表关联，  查找接口 /find/PrdtInfos
	LoadingInfos    []LoadingInfo `gorm:"many2many:sendLoadingInfo" form:"LoadingInfos"` // 装货明细，多表关联，查找接口 /find/LoadingInfos
	File1ID         uint          `form:"File1ID"`
	File2ID         uint          `form:"File2ID"`
	File1Name       string        `form:"File1Name"`
	File2Name       string        `form:"File2Name"`
	ShouldIns       []ShouldIn    `gorm:"many2many:shouldInSend" form:"ShouldIns"` // 应收账款单
	Ins             []In          `gorm:"many2many:inSend" form:"Ins"`             // 收款单
	gorm.Model
}
type ShouldIn struct {
	BillReceNum   string `gorm:"column:bill_rece_num" form:"BillReceNum" binding:"required"` // 应收账款单号
	DocDate       string `gorm:"column:doc_date" form:"DocDate"`                             // 单据日期
	ExpReceDate   string `gorm:"column:exp_rece_date" form:"ExpReceDate"`                    // 预计收款日期
	FinaDocType   string `gorm:"column:fina_doc_type" form:"FinaDocType"`                    // 财务单据类型
	FinaDocStatus string `gorm:"column:fina_doc_status" form:"FinaDocStatus"`                // 财务单据状态
	MerchantID    uint   `gorm:"column:merchantID" form:"MerchantID"`                        //  付款方,绑定id 查找接口 /find/merchant
	Merchant      Merchant
	AcctID        uint `gorm:"column:acctID" form:"AcctID"` // 收款方,绑定id 查找接口 /find/acct
	Acct          Acct
	BankAccountID uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 付款银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount   BankAccount
	AcctBankID    uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行账户 绑定id，查找接口 /find/acctBank
	AcctBank      AcctBank
	TotAmt        uint   `form:"TotAmt"`   // 总金额
	Currency      string `form:"Currency"` // 币种
	Notes         string `form:"Notes"`    // 描述
	FileID        uint   `form:"FileID"`
	FileName      string `form:"FileName"`
	Sends         []Send `gorm:"many2many:shouldInSend" form:"Sends"` // 销售发货单，多表关联, 查找接口
	Sales         []Sale `gorm:"many2many:shouldInSale" form:"Sales"` // 销售订单，多表关联, 查找接口
	Ins           []In   `gorm:"many2many:inShouldIn" form:"Ins"`     // 收款单
	gorm.Model
}

type Out struct {
	ReceNum       string `gorm:"column:rece_num" form:"ReceNum"`              // 收款单号
	RealReceDate  string `gorm:"column:real_rece_date" form:"RealReceDate"`   // 实际收款日期
	ExpReceDate   string `gorm:"column:exp_rece_date" form:"ExpReceDate"`     // 预计收款日期
	FinaDocType   string `gorm:"column:fina_doc_type" form:"FinaDocType"`     // 财务单据类型
	FinaDocStatus string `gorm:"column:fina_doc_status" form:"FinaDocStatus"` // 财务单据状态
	Merchant      Merchant
	MerchantID    uint `gorm:"column:merchantID" form:"MerchantID"` //  付款方,绑定id 查找接口 /find/merchant
	AcctID        uint `gorm:"column:acctID" form:"AcctID"`         // 收款方,绑定id 查找接口 /find/acct
	Acct          Acct
	BankAccountID uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 付款银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount   BankAccount
	AcctBankID    uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行账户 绑定id，查找接口 /find/acctBank
	AcctBank      AcctBank
	TotAmt        uint        `form:"TotAmt"`   // 总金额
	Currency      string      `form:"Currency"` // 币种
	Notes         string      `form:"Notes"`    // 描述
	Purrecs       []Purrec    `gorm:"many2many:outPurrec"`
	Buys          []Buy       `gorm:"many2many:outBuy"`
	ShouldOuts    []ShouldOut `gorm:"many2many:shouldOutOut"`
	FileID        uint        `form:"FileID"`
	FileName      string      `form:"FileName"`
	gorm.Model
}
type ShouldOut struct {
	BillReceNum   string `gorm:"column:bill_rece_num" form:"BillReceNum" binding:"required"` // 应收账款单号
	DocDate       string `gorm:"column:doc_date" form:"DocDate"`                             // 单据日期
	ExpReceDate   string `gorm:"column:exp_rece_date" form:"ExpReceDate"`                    // 预计收款日期
	FinaDocType   string `gorm:"column:fina_doc_type" form:"FinaDocType"`                    // 财务单据类型
	FinaDocStatus string `gorm:"column:fina_doc_status" form:"FinaDocStatus"`                // 财务单据状态
	MerchantID    uint   `gorm:"column:merchantID" form:"MerchantID"`                        //  付款方,绑定id 查找接口 /find/merchant
	Merchant      Merchant
	AcctID        uint `gorm:"column:acctID" form:"AcctID"` // 收款方,绑定id 查找接口 /find/acct
	Acct          Acct
	BankAccountID uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 付款银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount   BankAccount
	AcctBankID    uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行账户 绑定id，查找接口 /find/acctBank
	AcctBank      AcctBank
	TotAmt        uint     `form:"TotAmt"`   // 总金额
	Currency      string   `form:"Currency"` // 币种
	Notes         string   `form:"Notes"`    // 描述
	FileID        uint     `form:"FileID"`
	FileName      string   `form:"FileName"`
	Purrecs       []Purrec `gorm:"many2many:shouldOutPurrec"`
	Buys          []Buy    `gorm:"many2many:shouldOutBuy" `
	Outs          []Out    `gorm:"many2many:shouldOutOut" `
	gorm.Model
}
type In struct {
	ReceNum       string `gorm:"column:rece_num" form:"ReceNum"`              // 收款单号
	RealReceDate  string `gorm:"column:real_rece_date" form:"RealReceDate"`   // 实际收款日期
	ExpReceDate   string `gorm:"column:exp_rece_date" form:"ExpReceDate"`     // 预计收款日期
	FinaDocType   string `gorm:"column:fina_doc_type" form:"FinaDocType"`     // 财务单据类型
	FinaDocStatus string `gorm:"column:fina_doc_status" form:"FinaDocStatus"` // 财务单据状态
	MerchantID    uint   `gorm:"column:merchantID" form:"MerchantID"`         //  付款方,绑定id 查找接口 /find/merchant
	Merchant      Merchant
	AcctID        uint `gorm:"column:acctID" form:"AcctID"` // 收款方,绑定id 查找接口 /find/acct
	Acct          Acct
	BankAccountID uint `gorm:"column:bankAccountID" form:"BankAccountID"` // 付款银行账户 绑定id，查找接口 /find/bankAccount
	BankAccount   BankAccount
	AcctBankID    uint `gorm:"column:acctBankID" form:"AcctBankID"` // 收款银行账户 绑定id，查找接口 /find/acctBank
	AcctBank      AcctBank
	TotAmt        uint       `form:"TotAmt"`                                // 总金额
	Currency      string     `form:"Currency"`                              // 币种
	Notes         string     `form:"Notes"`                                 // 描述
	Sends         []Send     `gorm:"many2many:inSend" form:"Sends"`         // 销售发货单，多表关联, 查找接口
	Sales         []Sale     `gorm:"many2many:inSale" form:"Sales"`         // 销售订单，多表关联, 查找接口
	ShouldIns     []ShouldIn `gorm:"many2many:inShouldIn" form:"ShouldIns"` // 应收账款单
	FileID        uint       `form:"FileID"`
	FileName      string     `form:"FileName"`
	gorm.Model
}

type PrdtInfo struct {
	Brand       Brand
	BrandID     uint
	Cat         Cat
	CatID       uint
	PackSpec    PackSpec
	PackSpecID  uint
	Currency    string `gorm:"column:currency" form:"currency"`       // 币种
	UnitPrice   uint   `gorm:"column:unitPrice" form:"UnitPrice"`     // 单价
	TradeTerm   string `gorm:"column:tradeTerm" form:"TradeTerm"`     // 贸易条款
	DeliveryLoc string `gorm:"column:deliveryLoc" form:"DeliveryLoc"` // 交货地点
	Factory     string `gorm:"column:factory" form:"Factory"`         // 生产工厂
	Unit        string `gorm:"column:unit" form:"Unit"`               // 单位
	Amount      uint   `gorm:"column:amount" form:"Amount"`           // 金额
	ItemNum     uint   `gorm:"column:itemNum" form:"ItemNum"`         // 件数
	Weight      uint   `gorm:"column:weight" form:"Weight"`           // 重量
	WeightUnit  string `gorm:"column:weightUnit" form:"WeightUnit"`   // 重量单位
	gorm.Model
}
type LoadingInfo struct {
	Brand      Brand
	BrandID    uint
	Cat        Cat
	CatID      uint
	PackSpec   PackSpec
	PackSpecID uint
	PrdtPlant  string `gorm:"column:prdt_plant" form:"PrdtPlant"` // 生产工厂
	BatNum     string `gorm:"column:batNum" form:"BatNum"`        // 批次号
	ItemNum    uint   `gorm:"column:itemNum" form:"ItemNum"`      // 件数
	NetWeight  uint   `gorm:"column:netWeight" form:"NetWeight"`  // 净重量
	Unit       string `gorm:"column:unit" form:"Unit"`            // 单位
	CnrNum     string `gorm:"column:cnrNum" form:"CnrNum"`        // 集装箱号
	SealNum    string `gorm:"column:sealNum" form:"SealNum"`      // 铅封号
	VehNum     string `gorm:"column:vehNum" form:"VehNum"`        // 车辆号
	gorm.Model
}

type CostInfo struct {
	ExpType   string `gorm:"column:expType" form:"ExpType"`     // 费用类型
	Rates     string `gorm:"column:rates" form:"Rates"`         // 收费标准
	UnitPrice uint   `gorm:"column:unitPrice" form:"UnitPrice"` // 单价
	Number    uint   `gorm:"column:number" form:"Number"`       // 数量
	Amount    uint   `gorm:"column:amount" form:"Amount"`       // 金额
	Currency  string `gorm:"column:currency" form:"Currency"`   // 币种
	gorm.Model
}

type Empl struct {
	EmpName          string `gorm:"column:empName; not null" form:"EmpName"`         // 员工姓名
	EmpEngName       string `gorm:"column:empEngName" form:"EmpEngName"`             // 员工英文名
	Dept             string `gorm:"column:dept" form:"Dept"`                         // 部门
	Position         string `gorm:"column:position" form:"Position"`                 // 岗位
	JoinDate         string `gorm:"column:joinDate" form:"JoinDate"`                 // 入职日期
	Gender           string `gorm:"column:gender" form:"Gender"`                     // 性别
	IDCardNum        string `gorm:"column:idCardNum" form:"IDCardNum"`               // 身份证号
	EduLevel         string `gorm:"column:eduLevel" form:"EduLevel"`                 // 学历
	GradSchool       string `gorm:"column:gradSchool" form:"GradSchool"`             // 毕业学校
	Major            string `gorm:"column:major" form:"Major"`                       // 专业
	ResArea          string `gorm:"column:resArea" form:"ResArea"`                   // 居住地区
	AddrDetail       string `gorm:"column:addrDetail" form:"AddrDetail"`             // 详细地址
	MobileNum        string `gorm:"column:mobileNum" form:"MobileNum"`               // 手机号码
	EmailAddr        string `gorm:"column:emailAddr" form:"EmailAddr"`               // 电子邮箱
	EmergContact     string `gorm:"column:emergContact" form:"EmergContact"`         // 紧急联系人
	EmergContactInfo string `gorm:"column:emergContactInfo" form:"EmergContactInfo"` // 紧急联系方式
	PersonalPhoto    string `gorm:"column:personalPhoto" form:"PersonalPhoto"`       // 个人照片
	LaborContract    string `gorm:"column:laborContract" form:"LaborContract"`       // 劳动合同
	DeptComponent    string `gorm:"column:deptComponent" form:"DeptComponent"`       // 部门组件
	OrgRole          string `gorm:"column:orgRole" form:"OrgRole"`                   // 组织角色
	FileName         string `gorm:"column:fileName" form:"FileName"`                 // 文件名
	Notes            string `gorm:"column:notes" form:"Notes"`                       // 备注
	Age              uint   `gorm:"column:age" form:"Age"`                           // 年龄
	FileID           uint   `gorm:"column:fileID" form:"FileID"`                     // 文件ID
	EmplID           uint   `gorm:"column:empID; primaryKey" form:"EmplID"`          // 员工ID
}
type User struct {
	Email    string `gorm:"column:email;unique;not null" form:"Email" `
	Password string `gorm:"column:password; not null" form:"Password"`
	UserName string `gorm:"column:userName;not null" form:"UserName" `
	UserID   uint   `gorm:"column:userID;primaryKey" form:"UserID"`
	Priority uint   `gorm:"default:1" form:"Priority"`
	Empl     Empl
	EmplID   uint
}

type Acct struct { // 会计实体
	AcctCode    string `gorm:"column:acctCode;not null;" form:"AcctCode" binding:"required"`
	AcctAbbr    string `gorm:"column:acctAbbr;not null" form:"AcctAbbr" binding:"required"`
	EtyAbbr     string `gorm:"column:etyAbbr;not null" form:"EtyAbbr" binding:"required"`
	AcctName    string `gorm:"column:acctName;not null" form:"AcctName" binding:"required"`
	AcctEngName string `gorm:"column:acctEngName;" form:"AcctEngName"`
	AcctAddr    string `gorm:"column:acctAddr;" form:"AcctAddr"`
	Nation      string `gorm:"column:nation;" form:"Nation"`
	TaxType     string `gorm:"column:taxType;" form:"TaxType"`
	TaxCode     string `gorm:"column:taxCode;" form:"TaxCode"`
	PhoneNum    string `gorm:"column:phoneNum;" form:"PhoneNum"`
	Email       string `gorm:"column:email;" form:"Email"`
	Website     string `gorm:"column:website;" form:"Website"`
	RegDate     string `gorm:"column:regDate;" form:"RegDate"`
	Notes       string `gorm:"column:notes;" form:"Notes"`
	FileName    string `gorm:"column:fileName" form:"FileName"`
	FileID      uint   `gorm:"column:fileID" form:"FileID"`
	gorm.Model
}
type Spot struct {
	InvLocName string `form:"InvLocName" binding:"required"`
	InvLocAbbr string `form:"InvLocAbbr"`
	InvAddr    string `form:"InvAddr"`
	Notes      string `form:"Notes"`
	gorm.Model
}

type AcctBank struct {
	AccName   string `gorm:"column:accName; not null" form:"AccName" binding:"required"`
	AccNum    string `gorm:"column:accNum; not null; " form:"AccNum" binding:"required"`
	Currency  string `gorm:"column:currency" form:"Currency"`
	BankName  string `gorm:"column:bankName" form:"BankName"`
	BankNum   string `gorm:"column:bankNum" form:"BankNum"`
	SwiftCode string `gorm:"column:swiftCode" form:"SwiftCode"`
	BankAddr  string `gorm:"column:bankAddr" form:"BankAddr"`
	Notes     string `gorm:"column:notes" form:"Notes"`
	FileName  string `gorm:"column:fileName" form:"FileName"`
	FileID    uint   `gorm:"column:fileID" form:"FileID"`
	AcctID    uint   `gorm:"column:acctID" form:"AcctID"`
	Acct      Acct
	gorm.Model
}

type Merchant struct {
	MercCode  string `gorm:"column:mercCode;not null;" form:"MercCode" binding:"required"`
	MercAbbr  string `gorm:"column:mercAbbr;not null" form:"MercAbbr" `
	ShortMerc string `gorm:"column:shortMerc;not null" form:"ShortMerc" binding:"required"`
	Merc      string `gorm:"column:merc;not null" form:"Merc" `
	EngName   string `gorm:"column:engName;" form:"EngName"`
	Address   string `gorm:"column:address;" form:"Address"`
	Nation    string `gorm:"column:nation;" form:"Nation"`
	PhoneNum  string `gorm:"column:phoneNum;" form:"PhoneNum"`
	Email     string `gorm:"column:email;" form:"Email"`
	Fax       string `gorm:"column:fax;" form:"Fax"`
	Website   string `gorm:"column:website;" form:"Website"`
	TaxType   string `gorm:"column:taxType;" form:"TaxType"`
	TaxCode   string `gorm:"column:taxCode;" form:"TaxCode"`
	MercType  string `gorm:"column:mercType;" form:"MercType"`
	RegCap    string `gorm:"column:regCap;" form:"RegCap"`
	Notes     string `gorm:"column:notes;" form:"Notes"`
	FileName  string `gorm:"column:fileName" form:"FileName"`
	FileID    uint   `gorm:"column:fileID" form:"FileID"`
	gorm.Model
}

type Cust struct {
	Name       string `gorm:"column:name;not null" form:"Name" binding:"required"`
	Gender     string `gorm:"column:gender;" form:"Gender"`
	Nation     string `gorm:"column:nation;" form:"Nation"`
	Addr       string `gorm:"column:addr;" form:"Addr"`
	Email      string `gorm:"column:email;" form:"Email"`
	PhoneNum   string `gorm:"column:phoneNum;" form:"PhoneNum"`
	QQ         string `gorm:"column:qq;" form:"QQ"`
	Wechat     string `gorm:"column:wechat;" form:"Wechat"`
	Merc       string `gorm:"column:merc;" form:"Merc"`
	Post       string `gorm:"column:post;" form:"Post"`
	Notes      string `gorm:"column:notes;" form:"Notes"`
	FileName   string `gorm:"column:fileName" form:"FileName"`
	MerchantID uint   `gorm:"column:merchantID;" form:"MerchantID" binding:"required"` // 外键，关联到 Merchant 表的 MercID
	Merchant   Merchant
	FileID     uint `gorm:"column:fileID" form:"FileID"`
	gorm.Model
}

type BankAccount struct {
	BankAccName string `gorm:"column:bankAccName;not null" form:"BankAccName" binding:"required"`
	CompName    string `gorm:"column:compName;" form:"CompName" binding:"required"`
	AcctNum     string `gorm:"column:acctNum;not null" form:"AcctNum" binding:"required"`
	Currency    string `gorm:"column:currency;" form:"Currency"`
	BankName    string `gorm:"column:bankName;" form:"BankName" binding:"required"`
	BankEng     string `gorm:"column:bankEng;" form:"BankEng"`
	BankNum     string `gorm:"column:bankNum;" form:"BankNum"`
	SwiftCode   string `gorm:"column:swiftCode;" form:"SwiftCode"`
	BankAddr    string `gorm:"column:bankAddr;" form:"BankAddr"`
	Notes       string `gorm:"column:notes;" form:"Notes"`
	FileName    string `gorm:"column:fileName" form:"FileName"`
	FileID      uint   `gorm:"column:fileID" form:"FileID"`
	MerchantID  uint   `gorm:"column:merchantID;" form:"MerchantID"`
	Merchant    Merchant
	gorm.Model
}

type PayMentMethod struct {
	PayMtdName string `gorm:"column:PayMtdName;" form:"PayMtdName" binding:"required"`
	Notes      string `gorm:"column:notes;" form:"Notes"`
	PayMth     string `gorm:"column:payMth;" form:"PayMth"`
	PayLimit   string `gorm:"column:payLimit;" form:"PayLimit"`
	PayRat     uint   `gorm:"column:payRat;" form:"PayRat"`
	gorm.Model
}

type PackSpec struct {
	SpecName    string `gorm:"column:specName; " form:"SpecName" binding:"required"`
	SpecEngName string `gorm:"column:specEngName" form:"SpecEngName"`
	UnitMeas    string `gorm:"column:unitMeas" form:"UnitMeas"`
	PackType    string `gorm:"column:packType" form:"PackType"`
	FileName    string `gorm:"column:fileName" form:"FileName"`
	Notes       string `gorm:"column:notes" form:"Notes"` // 备注
	FileID      uint   `gorm:"column:fileID" form:"FileID"`
	NetWt       uint   `gorm:"column:netWt" form:"NetWt"`
	gorm.Model
}
type File struct {
	Name   string `gorm:"column:name; not null" form:"Name"`
	MD5    string `gorm:"column:MD5;  not null" form:"MD5"`
	Suffix string `gorm:"column:suffix; not null" form:"Suffix"`
	gorm.Model
}

type Brand struct {
	BrandName    string `gorm:"column:brandName;not null"  form:"BrandName" binding:"required"` // 品牌名称
	BrandEngName string `gorm:"column:brandEngName;"  form:"BrandEngName"`                      // 品牌英文名称
	BrandType    string `gorm:"column:brandType;"  form:"BrandType"`                            // 品牌类型
	Notes        string `gorm:"column:notes;"  form:"Notes"`                                    // 备注
	FileName     string `gorm:"column:fileName;"  form:"FileName"`                              // 文件名字
	Nation       string `gorm:"column:Nation" form:"Nation" `                                   // 品牌国别，绑定 Nation 表项
	FileID       uint   `gorm:"column:fileID;"  form:"FileID"`                                  // 文件 ID
	BrandYear    uint   `gorm:"column:brandYear;"  form:"BrandYear"`                            // 品牌创建年份
	gorm.Model
}
type Cat struct {
	CatAbbr    string `gorm:"column:catAbbr;not null;" form:"CatAbbr"` // 品类缩写
	CatName    string `gorm:"column:catName;not null" form:"CatName"`  // 品类名称
	CatEngName string `gorm:"column:catEngName" form:"CatEngName"`     // 品类英文名称
	DeclHS     string `gorm:"column:declHS" form:"DeclHS"`             // 报关HS编码
	DocHS      string `gorm:"column:docHS" form:"DocHS"`               // 单据HS编码
	CAS        string `gorm:"column:CAS" form:"CAS"`                   // CAS编码
	GB         string `gorm:"column:GB" form:"GB"`                     // 国标编码
	StdDoc     string `gorm:"column:stdDoc" form:"StdDoc"`             // 国标文件
	Notes      string `gorm:"column:notes" form:"Notes"`               // 备注
	FileName   string `gorm:"column:fileName" form:"FileName"`         // 文件名字
	FileID     uint   `gorm:"column:fileID" form:"FileID"`             // 文件ID
	gorm.Model
}

// 国家
type Nation struct {
	Nation   string `gorm:"column:nation; not null; unique" form:"Nation"`
	NationID uint   `gorm:"column:nationID; primaryKey; " form:"NationID"`
}

// 客商（户）类型
type MercType struct {
	MercType   string `gorm:"column:mercType; not null; unique" form:"MercType"`
	MercTypeID uint   `gorm:"column:mercTypeID; primaryKey; " form:"MercTypeID"`
}

// 供应商类型
type SuprType struct {
	SuprType   string `gorm:"column:suprType; not null; unique" form:"SuprType"`
	SuprTypeID uint   `gorm:"column:suprTypeID; primaryKey; " form:"SuprTypeID"`
}

// 产品类型
type PrdtType struct {
	PrdtType   string `gorm:"column:prdtType; not null; unique" form:"PrdtType"`
	PrdtTypeID uint   `gorm:"column:prdtTypeID; primaryKey; " form:"PrdtTypeID"`
}

// 食品添加剂类型
type FoodAddType struct {
	FoodAddType   string `gorm:"column:foodAddType; not null; unique" form:"FoodAddType"`
	FoodAddTypeID uint   `gorm:"column:foodAddTypeID; primaryKey; " form:"FoodAddTypeID"`
}

// 饲料添加剂类型
type FeedAddType struct {
	FeedAddType   string `gorm:"column:feedAddType; not null; unique" form:"FeedAddType"`
	FeedAddTypeID uint   `gorm:"column:feedAddTypeID; primaryKey; " form:"FeedAddTypeID"`
}

// 计量单位
type UnitMeas struct {
	UnitMeas   string `gorm:"column:unitMeas; not null; unique" form:"UnitMeas"`
	UnitMeasID uint   `gorm:"column:unitMeasID; primaryKey; " form:"UnitMeasID"`
}

// 包装类型
type PackType struct {
	PackType   string `gorm:"column:packType; not null; unique" form:"PackType"`
	PackTypeID uint   `gorm:"column:packTypeID; primaryKey; " form:"PackTypeID"`
}

// 集装箱类型
type ConType struct {
	ConType   string `gorm:"column:conType; not null; unique" form:"ConType"`
	ConTypeID uint   `gorm:"column:conTypeID; primaryKey; " form:"ConTypeID"`
}

// 币种
type Currency struct {
	Currency   string `gorm:"column:currency; not null; unique" form:"Currency"`
	CurrencyID uint   `gorm:"column:currencyID; primaryKey; " form:"CurrencyID"`
}

// 贸易条款
type TradeTerm struct {
	TradeTerm   string `gorm:"column:tradeTerm; not null; unique" form:"TradeTerm"`
	TradeTermID uint   `gorm:"column:tradeTermID; primaryKey; " form:"TradeTermID"`
}

// 港口
type Port struct {
	Port   string `gorm:"column:port; not null; unique" form:"Port"`
	PortID uint   `gorm:"column:portID; primaryKey; " form:"PortID"`
}

// 税号类型
type TaxType struct {
	TaxType   string `gorm:"column:taxType; not null; unique" form:"TaxType"`
	TaxTypeID uint   `gorm:"column:taxTypeID; primaryKey; " form:"TaxTypeID"`
}

// 品牌类型
type BrandType struct {
	BrandType   string `gorm:"column:brandType; not null; unique" form:"BrandType"`
	BrandTypeID uint   `gorm:"column:brandTypeID; primaryKey; " form:"BrandTypeID"`
}

// 学历
type EduLevel struct {
	EduLevel   string `gorm:"column:degree; not null; unique" form:"EduLevel"`
	EduLevelID uint   `gorm:"column:degreeID; primaryKey; " form:"EduLevelID"`
}

// 公司部门
type Dept struct {
	Dept   string `gorm:"column:dept; not null; unique" form:"Dept"`
	DeptID uint   `gorm:"column:deptID; primaryKey; " form:"DeptID"`
}

// 公司岗位
type Position struct {
	Position   string `gorm:"column:postion; not null; unique" form:"Position"`
	PositionID uint   `gorm:"column:postID; primaryKey; " form:"PositionID"`
}

// 质量标准
type QualStd struct {
	QualStd   string `gorm:"column:qualStd; not null; unique" form:"QualStd"`
	QualStdID uint   `gorm:"column:qualStdID; primaryKey; " form:"QualStdID"`
}

// 库存地点位置
type InvLoc struct {
	InvLoc   string `gorm:"column:invLoc; not null; unique" form:"InvLoc"`
	InvLocID uint   `gorm:"column:invLocID; primaryKey; " form:"InvLocID"`
}

// 单据要求
type DocReq struct {
	DocReq   string `gorm:"column:docReq; not null; unique" form:"DocReq"`
	DocReqID uint   `gorm:"column:docReqID; primaryKey; " form:"DocReqID"`
}

// 转账方式
type PayMth struct {
	PayMth   string `gorm:"column:payMth; not null; unique" form:"PayMth"`
	PayMthID uint   `gorm:"column:payMthID; primaryKey; " form:"PayMthID"`
}

// 后付款转账期限
type PayLimit struct {
	PayLimit   string `gorm:"column:payLimit; not null; unique" form:"PayLimit"`
	PayLimitID uint   `gorm:"column:payLimitID; primaryKey; " form:"PayLimitID"`
}

// 财务单据状态
type FinaDocStatus struct {
	FinaDocStatus   string `gorm:"column:finaDocStatus; not null; unique" form:"FinaDocStatus"`
	FinaDocStatusID uint   `gorm:"column:finaDocStatusID; primaryKey; " form:"FinaDocStatusID"`
}

// 财务单据类型
type FinaDocType struct {
	FinaDocType   string `gorm:"column:finaDocType; not null; unique" form:"FinaDocType"`
	FinaDocTypeID uint   `gorm:"column:finaDocTypeID; primaryKey; " form:"FinaDocTypeID"`
}

// 费用类型
type ExpType struct {
	ExpType   string `gorm:"column:expType; not null; unique" form:"ExpType"`
	ExpTypeID uint   `gorm:"column:expTypeID; primaryKey; " form:"ExpTypeID"`
}

// 收费标准
type Rates struct {
	Rates   string `gorm:"column:rates; not null; unique" form:"Rates"`
	RatesID uint   `gorm:"column:ratesID; primaryKey; " form:"RatesID"`
}

// 业务单据状态
type BussOrderSta struct {
	BussOrderSta   string `gorm:"column:bussOrderSta; not null; unique" form:"BussOrderSta"`
	BussOrderStaID uint   `gorm:"column:bussOrderStaID; primaryKey; " form:"BussOrderStaID"`
}
