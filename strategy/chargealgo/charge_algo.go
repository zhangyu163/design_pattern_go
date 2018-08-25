package chargealgo

type CashStrategy interface {
	//SetConfig(base, condition float64)
	AcceptCash(money float64) float64
}

type CashNormal struct {
}

func (cash CashNormal) AcceptCash(money float64) float64 {
	return money
}

type CashRebate struct {
	Discount float64
}

func (cash CashRebate) SetConfig(discount float64) {
	cash.Discount = discount
}

func (cash CashRebate) AcceptCash(money float64) float64 {
	return money * cash.Discount
}

type CashReturn struct {
	Base, RetNum float64
}

func (cash CashReturn) SetConfig(base, retNum float64) {
	cash.Base = base
	cash.RetNum = retNum
}

func (cash CashReturn) AcceptCash(money float64) float64 {
	if money < cash.Base {
		return money
	}
	multiple := int(money / cash.Base)
	return money - float64(multiple)*cash.RetNum
}
