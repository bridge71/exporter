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
}

type Merchant struct {
	MercCode     string        `gorm:"column:mercCode;not null;unique" form:"MercCode" binding:"required"`
	MercAbbr     string        `gorm:"column:mercAbbr;not null" form:"MercAbbr" binding:"required"`
	ShortMerc    string        `gorm:"column:shortMerc;not null" form:"ShortMerc" binding:"required"`
	Merc         string        `gorm:"column:merc;not null" form:"Merc" `
	EngName      string        `gorm:"column:engName;" form:"EngName"`
	Address      string        `gorm:"column:address;" form:"Address"`
	Nation       string        `gorm:"column:nation;" form:"Nation"`
	PhoneNum     string        `gorm:"column:phoneNum;" form:"PhoneNum"`
	Email        string        `gorm:"column:email;" form:"Email"`
	Fax          string        `gorm:"column:fax;" form:"Fax"`
	Website      string        `gorm:"column:website;" form:"Website"`
	TaxType      string        `gorm:"column:taxType;" form:"TaxType"`
	TaxCode      string        `gorm:"column:taxCode;" form:"TaxCode"`
	MercType     string        `gorm:"column:mercType;" form:"MercType"`
	RegCap       string        `gorm:"column:regCap;" form:"RegCap"`
	Notes        string        `gorm:"column:notes;" form:"Notes"`
	FileName     string        `gorm:"column:fileName" form:"FileName"`
	Custs        []Cust        `gorm:"foreignKey:MercId" form:"Custs[]"`        // 一对多关系
	BankAccounts []BankAccount `gorm:"foreignKey:MercId" form:"BankAccounts[]"` // 一对多关系
	MercId       uint          `gorm:"column:mercId;primaryKey" form:"MercId"`
	FileId       uint          `gorm:"column:fileId" form:"FileId"`
}

type Cust struct {
	Name     string `gorm:"column:name;not null" form:"Name" binding:"required"`
	Gender   string `gorm:"column:gender;" form:"Gender"`
	Nation   string `gorm:"column:nation;" form:"Nation"`
	Addr     string `gorm:"column:addr;" form:"Addr"`
	Email    string `gorm:"column:email;" form:"Email"`
	PhoneNum string `gorm:"column:phoneNum;" form:"PhoneNum"`
	QQ       string `gorm:"column:qq;" form:"QQ"`
	Wechat   string `gorm:"column:wechat;" form:"Wechat"`
	Merc     string `gorm:"column:merc;" form:"Merc"`
	Post     string `gorm:"column:post;" form:"Post"`
	Notes    string `gorm:"column:notes;" form:"Notes"`
	FileName string `gorm:"column:fileName" form:"FileName"`
	CustId   uint   `gorm:"column:custId;primaryKey" form:"CustId"`
	MercId   uint   `gorm:"column:mercId;" form:"MercId"` // 外键，关联到 Merchant 表的 MercId
	FileId   uint   `gorm:"column:fileId" form:"FileId"`
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
	BankAccId   uint   `gorm:"column:bankAccId;primaryKey" form:"BankAccId"`
	FileId      uint   `gorm:"column:fileId" form:"FileId"`
	MercId      uint   `gorm:"column:mercId;" form:"MercId"`
}

type PayMentMethod struct {
	PayMtdName      string `gorm:"column:PayMtdName; unique" form:"PayMtdName"`
	Notes           string `gorm:"column:notes;" form:"Notes"`
	PayMth          string `gorm:"column:payMth;" form:"PayMth"`
	PayLimit        string `gorm:"column:payLimit;" form:"PayLimit"`
	PayMentMethodId uint   `gorm:"column:payMentMethodId; primaryKey" form:"PayMentMethodId"`
	PayRat          uint   `gorm:"column:payRat;" form:"PayRat"`
}
type File struct {
	Name   string `gorm:"column:name; not null" json:"Name"`
	MD5    string `gorm:"column:MD5;  not null" json:"MD5"`
	Suffix string `gorm:"column:suffix; not null" json:"Suffix"`
	FileId uint   `gorm:"column:fileId; primaryKey" form:"FileId" binding:"required"`
}

// 国家
type Nation struct {
	Nation   string `gorm:"column:nation; not null; unique" form:"Nation"`
	NationId uint   `gorm:"column:nationId; primaryKey; " form:"NationId"`
}
type Brand struct {
	BrandName    string `gorm:"column:brandName;not null"  form:"BrandName" binding:"required"` // 品牌名称
	BrandEngName string `gorm:"column:brandEngName;"  form:"BrandEngName"`                      // 品牌英文名称
	BrandType    string `gorm:"column:brandType;"  form:"BrandType"`                            // 品牌类型
	Notes        string `gorm:"column:notes;"  form:"Notes"`                                    // 备注
	FileName     string `gorm:"column:fileName;"  form:"FileName"`                              // 文件名字
	Nation       string `gorm:"column:Nation" form:"Nation" `                                   // 品牌国别，绑定 Nation 表项
	FileId       uint   `gorm:"column:fileId;"  form:"FileId"`                                  // 文件 ID
	BrandYear    uint   `gorm:"column:brandYear;"  form:"BrandYear"`                            // 品牌创建年份
	BrandId      uint   `gorm:"column:brandId;primaryKey" form:"BrandId"`                       // 品牌 ID（主键）
}
type Cat struct {
	CatAbbr    string `gorm:"column:catAbbr;not null;unique" form:"CatAbbr"` // 品类缩写
	CatName    string `gorm:"column:catName;not null" form:"CatName"`        // 品类名称
	CatEngName string `gorm:"column:catEngName" form:"CatEngName"`           // 品类英文名称
	DeclHS     string `gorm:"column:declHS" form:"DeclHS"`                   // 报关HS编码
	DocHS      string `gorm:"column:docHS" form:"DocHS"`                     // 单据HS编码
	CAS        string `gorm:"column:CAS" form:"CAS"`                         // CAS编码
	GB         string `gorm:"column:GB" form:"GB"`                           // 国标编码
	StdDoc     string `gorm:"column:stdDoc" form:"StdDoc"`                   // 国标文件
	Notes      string `gorm:"column:notes" form:"Notes"`                     // 备注
	FileName   string `gorm:"column:fileName" form:"FileName"`               // 文件名字
	FileId     uint   `gorm:"column:fileId" form:"FileId"`                   // 文件ID
	CatId      uint   `gorm:"column:catId;primaryKey" form:"CatId"`          // 品类ID
}
type PackSpec struct {
	SpecName    string `gorm:"column:specName; unique" form:"SpecName"`
	SpecEngName string `gorm:"column:specEngName" form:"SpecEngName"`
	UnitMeas    string `gorm:"column:unitMeas" form:"UnitMeas"`
	PackType    string `gorm:"column:packType" form:"PackType"`
	FileName    string `gorm:"column:fileName" form:"FileName"`
	Notes       string `gorm:"column:notes" form:"Notes"` // 备注
	FileId      uint   `gorm:"column:fileId" form:"FileId"`
	PackSpecId  uint   `gorm:"column:packSpecId; primaryKey" form:"PackSpecId"`
	NetWt       uint   `gorm:"column:netWt" form:"NetWt"`
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
	FileId           uint   `gorm:"column:fileId" form:"FileId"`                     // 文件ID
	EmpID            uint   `gorm:"column:empId; primaryKey" form:"EmpId"`           // 员工ID
}
type Message struct {
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

// 客商（户）类型
type MercType struct {
	MercType   string `gorm:"column:mercType; not null; unique" form:"MercType"`
	MercTypeId uint   `gorm:"column:mercTypeId; primaryKey; " form:"MercTypeId"`
}

// 供应商类型
type SuprType struct {
	SuprType   string `gorm:"column:suprType; not null; unique" form:"SuprType"`
	SuprTypeId uint   `gorm:"column:suprTypeId; primaryKey; " form:"SuprTypeId"`
}

// 产品类型
type PrdtType struct {
	PrdtType   string `gorm:"column:prdtType; not null; unique" form:"PrdtType"`
	PrdtTypeId uint   `gorm:"column:prdtTypeId; primaryKey; " form:"PrdtTypeId"`
}

// 食品添加剂类型
type FoodAddType struct {
	FoodAddType   string `gorm:"column:foodAddType; not null; unique" form:"FoodAddType"`
	FoodAddTypeId uint   `gorm:"column:foodAddTypeId; primaryKey; " form:"FoodAddTypeId"`
}

// 饲料添加剂类型
type FeedAddType struct {
	FeedAddType   string `gorm:"column:feedAddType; not null; unique" form:"FeedAddType"`
	FeedAddTypeId uint   `gorm:"column:feedAddTypeId; primaryKey; " form:"FeedAddTypeId"`
}

// 计量单位
type UnitMeas struct {
	UnitMeas   string `gorm:"column:unitMeas; not null; unique" form:"UnitMeas"`
	UnitMeasId uint   `gorm:"column:unitMeasId; primaryKey; " form:"UnitMeasId"`
}

// 包装类型
type PackType struct {
	PackType   string `gorm:"column:packType; not null; unique" form:"PackType"`
	PackTypeId uint   `gorm:"column:packTypeId; primaryKey; " form:"PackTypeId"`
}

// 集装箱类型
type ConType struct {
	ConType   string `gorm:"column:conType; not null; unique" form:"ConType"`
	ConTypeId uint   `gorm:"column:conTypeId; primaryKey; " form:"ConTypeId"`
}

// 币种
type Currency struct {
	Currency   string `gorm:"column:currency; not null; unique" form:"Currency"`
	CurrencyId uint   `gorm:"column:currencyId; primaryKey; " form:"CurrencyId"`
}

// 贸易条款
type TradeTerm struct {
	TradeTerm   string `gorm:"column:tradeTerm; not null; unique" form:"TradeTerm"`
	TradeTermId uint   `gorm:"column:tradeTermId; primaryKey; " form:"TradeTermId"`
}

// 港口
type Port struct {
	Port   string `gorm:"column:port; not null; unique" form:"Port"`
	PortId uint   `gorm:"column:portId; primaryKey; " form:"PortId"`
}

// 税号类型
type TaxType struct {
	TaxType   string `gorm:"column:taxType; not null; unique" form:"TaxType"`
	TaxTypeId uint   `gorm:"column:taxTypeId; primaryKey; " form:"TaxTypeId"`
}

// 品牌类型
type BrandType struct {
	BrandType   string `gorm:"column:brandType; not null; unique" form:"BrandType"`
	BrandTypeId uint   `gorm:"column:brandTypeId; primaryKey; " form:"BrandTypeId"`
}

// 学历
type EduLevel struct {
	EduLevel   string `gorm:"column:degree; not null; unique" form:"EduLevel"`
	EduLevelId uint   `gorm:"column:degreeId; primaryKey; " form:"EduLevelId"`
}

// 公司部门
type Dept struct {
	Dept   string `gorm:"column:dept; not null; unique" form:"Dept"`
	DeptId uint   `gorm:"column:deptId; primaryKey; " form:"DeptId"`
}

// 公司岗位
type Position struct {
	Position   string `gorm:"column:postion; not null; unique" form:"Position"`
	PositionId uint   `gorm:"column:postId; primaryKey; " form:"PositionId"`
}

// 质量标准
type QualStd struct {
	QualStd   string `gorm:"column:qualStd; not null; unique" form:"QualStd"`
	QualStdId uint   `gorm:"column:qualStdId; primaryKey; " form:"QualStdId"`
}

// 库存地点位置
type InvLoc struct {
	InvLoc   string `gorm:"column:invLoc; not null; unique" form:"InvLoc"`
	InvLocId uint   `gorm:"column:invLocId; primaryKey; " form:"InvLocId"`
}

// 单据要求
type DocReq struct {
	DocReq   string `gorm:"column:docReq; not null; unique" form:"DocReq"`
	DocReqId uint   `gorm:"column:docReqId; primaryKey; " form:"DocReqId"`
}

// 转账方式
type PayMth struct {
	PayMth   string `gorm:"column:payMth; not null; unique" form:"PayMth"`
	PayMthId uint   `gorm:"column:payMthId; primaryKey; " form:"PayMthId"`
}

// 后付款转账期限
type PayLimit struct {
	PayLimit   string `gorm:"column:payLimit; not null; unique" form:"PayLimit"`
	PayLimitId uint   `gorm:"column:payLimitId; primaryKey; " form:"PayLimitId"`
}

// 财务单据状态
type FinaDocStatus struct {
	FinaDocStatus   string `gorm:"column:finaDocStatus; not null; unique" form:"FinaDocStatus"`
	FinaDocStatusId uint   `gorm:"column:finaDocStatusId; primaryKey; " form:"FinaDocStatusId"`
}

// 财务单据类型
type FinaDocType struct {
	FinaDocType   string `gorm:"column:finaDocType; not null; unique" form:"FinaDocType"`
	FinaDocTypeId uint   `gorm:"column:finaDocTypeId; primaryKey; " form:"FinaDocTypeId"`
}

// 费用类型
type ExpType struct {
	ExpType   string `gorm:"column:expType; not null; unique" form:"ExpType"`
	ExpTypeId uint   `gorm:"column:expTypeId; primaryKey; " form:"ExpTypeId"`
}

// 收费标准
type Rates struct {
	Rates   string `gorm:"column:rates; not null; unique" form:"Rates"`
	RatesId uint   `gorm:"column:ratesId; primaryKey; " form:"RatesId"`
}

// 业务单据状态
type BussOrderSta struct {
	BussOrderSta   string `gorm:"column:bussOrderSta; not null; unique" form:"BussOrderSta"`
	BussOrderStaId uint   `gorm:"column:bussOrderStaId; primaryKey; " form:"BussOrderStaId"`
}
