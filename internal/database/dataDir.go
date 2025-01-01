package database

import "exporter/internal/models"

func (s *service) FindMercType(mercTypes *[]models.MercType) {
	s.gormDB.Find(mercTypes)
}

func (s *service) FindSuprType(suprTypes *[]models.SuprType) {
	s.gormDB.Find(suprTypes)
}

func (s *service) FindPrdtType(prdtTypes *[]models.PrdtType) {
	s.gormDB.Find(prdtTypes)
}

func (s *service) FindFoodAddType(foodAddTypes *[]models.FoodAddType) {
	s.gormDB.Find(foodAddTypes)
}

func (s *service) FindFeedAddType(feedAddTypes *[]models.FeedAddType) {
	s.gormDB.Find(feedAddTypes)
}

func (s *service) FindUnitMeas(unitMeas *[]models.UnitMeas) {
	s.gormDB.Find(unitMeas)
}

func (s *service) FindPackType(packTypes *[]models.PackType) {
	s.gormDB.Find(packTypes)
}

func (s *service) FindConType(conTypes *[]models.ConType) {
	s.gormDB.Find(conTypes)
}

func (s *service) FindCurrency(currencies *[]models.Currency) {
	s.gormDB.Find(currencies)
}

func (s *service) FindTradeTerm(tradeTerms *[]models.TradeTerm) {
	s.gormDB.Find(tradeTerms)
}

func (s *service) FindNation(nations *[]models.Nation) {
	s.gormDB.Find(nations)
}

func (s *service) FindPort(ports *[]models.Port) {
	s.gormDB.Find(ports)
}

func (s *service) FindTaxType(taxTypes *[]models.TaxType) {
	s.gormDB.Find(taxTypes)
}

func (s *service) FindBrandType(brandTypes *[]models.BrandType) {
	s.gormDB.Find(brandTypes)
}

func (s *service) FindDegree(degrees *[]models.Degree) {
	s.gormDB.Find(degrees)
}

func (s *service) FindDept(depts *[]models.Dept) {
	s.gormDB.Find(depts)
}

func (s *service) FindPost(posts *[]models.Post) {
	s.gormDB.Find(posts)
}

func (s *service) FindQualStd(qualStds *[]models.QualStd) {
	s.gormDB.Find(qualStds)
}

func (s *service) FindInvLoc(invLocs *[]models.InvLoc) {
	s.gormDB.Find(invLocs)
}

func (s *service) FindDocReq(docReqs *[]models.DocReq) {
	s.gormDB.Find(docReqs)
}

func (s *service) FindPayMth(payMths *[]models.PayMth) {
	s.gormDB.Find(payMths)
}

func (s *service) FindPayLimit(payLimits *[]models.PayLimit) {
	s.gormDB.Find(payLimits)
}

func (s *service) FindFinaDocStatus(finaDocStatuses *[]models.FinaDocStatus) {
	s.gormDB.Find(finaDocStatuses)
}

func (s *service) FindFinaDocType(finaDocTypes *[]models.FinaDocType) {
	s.gormDB.Find(finaDocTypes)
}

func (s *service) FindExpType(expTypes *[]models.ExpType) {
	s.gormDB.Find(expTypes)
}

func (s *service) FindRates(rates *[]models.Rates) {
	s.gormDB.Find(rates)
}

func (s *service) FindBussOrderSta(bussOrderStas *[]models.BussOrderSta) {
	s.gormDB.Find(bussOrderStas)
}

func (s *service) DeleteMercType(mercType *models.MercType) error {
	return s.gormDB.Delete(mercType).Error
}

func (s *service) DeleteSuprType(suprType *models.SuprType) error {
	return s.gormDB.Delete(suprType).Error
}

func (s *service) DeletePrdtType(prdtType *models.PrdtType) error {
	return s.gormDB.Delete(prdtType).Error
}

func (s *service) DeleteFoodAddType(foodAddType *models.FoodAddType) error {
	return s.gormDB.Delete(foodAddType).Error
}

func (s *service) DeleteFeedAddType(feedAddType *models.FeedAddType) error {
	return s.gormDB.Delete(feedAddType).Error
}

func (s *service) DeleteUnitMeas(unitMeas *models.UnitMeas) error {
	return s.gormDB.Delete(unitMeas).Error
}

func (s *service) DeletePackType(packType *models.PackType) error {
	return s.gormDB.Delete(packType).Error
}

func (s *service) DeleteConType(conType *models.ConType) error {
	return s.gormDB.Delete(conType).Error
}

func (s *service) DeleteCurrency(currency *models.Currency) error {
	return s.gormDB.Delete(currency).Error
}

func (s *service) DeleteTradeTerm(tradeTerm *models.TradeTerm) error {
	return s.gormDB.Delete(tradeTerm).Error
}

func (s *service) DeleteNation(nation *models.Nation) error {
	return s.gormDB.Delete(nation).Error
}

func (s *service) DeletePort(port *models.Port) error {
	return s.gormDB.Delete(port).Error
}

func (s *service) DeleteTaxType(taxType *models.TaxType) error {
	return s.gormDB.Delete(taxType).Error
}

func (s *service) DeleteBrandType(brandType *models.BrandType) error {
	return s.gormDB.Delete(brandType).Error
}

func (s *service) DeleteDegree(degree *models.Degree) error {
	return s.gormDB.Delete(degree).Error
}

func (s *service) DeleteDept(dept *models.Dept) error {
	return s.gormDB.Delete(dept).Error
}

func (s *service) DeletePost(post *models.Post) error {
	return s.gormDB.Delete(post).Error
}

func (s *service) DeleteQualStd(qualStd *models.QualStd) error {
	return s.gormDB.Delete(qualStd).Error
}

func (s *service) DeleteInvLoc(invLoc *models.InvLoc) error {
	return s.gormDB.Delete(invLoc).Error
}

func (s *service) DeleteDocReq(docReq *models.DocReq) error {
	return s.gormDB.Delete(docReq).Error
}

func (s *service) DeletePayMth(payMth *models.PayMth) error {
	return s.gormDB.Delete(payMth).Error
}

func (s *service) DeletePayLimit(payLimit *models.PayLimit) error {
	return s.gormDB.Delete(payLimit).Error
}

func (s *service) DeleteFinaDocStatus(finaDocStatus *models.FinaDocStatus) error {
	return s.gormDB.Delete(finaDocStatus).Error
}

func (s *service) DeleteFinaDocType(finaDocType *models.FinaDocType) error {
	return s.gormDB.Delete(finaDocType).Error
}

func (s *service) DeleteExpType(expType *models.ExpType) error {
	return s.gormDB.Delete(expType).Error
}

func (s *service) DeleteRates(rates *models.Rates) error {
	return s.gormDB.Delete(rates).Error
}

func (s *service) DeleteBussOrderSta(bussOrderSta *models.BussOrderSta) error {
	return s.gormDB.Delete(bussOrderSta).Error
}

func (s *service) SaveMercType(mercType *models.MercType) error {
	return s.gormDB.Save(mercType).Error
}

func (s *service) SaveSuprType(suprType *models.SuprType) error {
	return s.gormDB.Save(suprType).Error
}

func (s *service) SavePrdtType(prdtType *models.PrdtType) error {
	return s.gormDB.Save(prdtType).Error
}

func (s *service) SaveFoodAddType(foodAddType *models.FoodAddType) error {
	return s.gormDB.Save(foodAddType).Error
}

func (s *service) SaveFeedAddType(feedAddType *models.FeedAddType) error {
	return s.gormDB.Save(feedAddType).Error
}

func (s *service) SaveUnitMeas(unitMeas *models.UnitMeas) error {
	return s.gormDB.Save(unitMeas).Error
}

func (s *service) SavePackType(packType *models.PackType) error {
	return s.gormDB.Save(packType).Error
}

func (s *service) SaveConType(conType *models.ConType) error {
	return s.gormDB.Save(conType).Error
}

func (s *service) SaveCurrency(currency *models.Currency) error {
	return s.gormDB.Save(currency).Error
}

func (s *service) SaveTradeTerm(tradeTerm *models.TradeTerm) error {
	return s.gormDB.Save(tradeTerm).Error
}

func (s *service) SaveNation(nation *models.Nation) error {
	return s.gormDB.Save(nation).Error
}

func (s *service) SavePort(port *models.Port) error {
	return s.gormDB.Save(port).Error
}

func (s *service) SaveTaxType(taxType *models.TaxType) error {
	return s.gormDB.Save(taxType).Error
}

func (s *service) SaveBrandType(brandType *models.BrandType) error {
	return s.gormDB.Save(brandType).Error
}

func (s *service) SaveDegree(degree *models.Degree) error {
	return s.gormDB.Save(degree).Error
}

func (s *service) SaveDept(dept *models.Dept) error {
	return s.gormDB.Save(dept).Error
}

func (s *service) SavePost(post *models.Post) error {
	return s.gormDB.Save(post).Error
}

func (s *service) SaveQualStd(qualStd *models.QualStd) error {
	return s.gormDB.Save(qualStd).Error
}

func (s *service) SaveInvLoc(invLoc *models.InvLoc) error {
	return s.gormDB.Save(invLoc).Error
}

func (s *service) SaveDocReq(docReq *models.DocReq) error {
	return s.gormDB.Save(docReq).Error
}

func (s *service) SavePayMth(payMth *models.PayMth) error {
	return s.gormDB.Save(payMth).Error
}

func (s *service) SavePayLimit(payLimit *models.PayLimit) error {
	return s.gormDB.Save(payLimit).Error
}

func (s *service) SaveFinaDocStatus(finaDocStatus *models.FinaDocStatus) error {
	return s.gormDB.Save(finaDocStatus).Error
}

func (s *service) SaveFinaDocType(finaDocType *models.FinaDocType) error {
	return s.gormDB.Save(finaDocType).Error
}

func (s *service) SaveExpType(expType *models.ExpType) error {
	return s.gormDB.Save(expType).Error
}

func (s *service) SaveRates(rates *models.Rates) error {
	return s.gormDB.Save(rates).Error
}

func (s *service) SaveBussOrderSta(bussOrderSta *models.BussOrderSta) error {
	return s.gormDB.Save(bussOrderSta).Error
}

// FindPayMthByID 根据 ID 查询 PayMth 记录
func (s *service) FindPayMthByID(payMth *models.PayMth, id uint) error {
	return s.gormDB.First(payMth, id).Error
}

// FindPayLimitByID 根据 ID 查询 PayLimit 记录
func (s *service) FindPayLimitByID(payLimit *models.PayLimit, id uint) error {
	return s.gormDB.First(payLimit, id).Error
}

