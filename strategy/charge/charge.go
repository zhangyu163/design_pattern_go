package charge

import "errors"
import "exercise/design_pattern_go/strategy/chargealgo"

type CashContext struct {
	CashStrategy chargealgo.CashStrategy
}

func (context *CashContext) CashContext(cash string) error {
	var err error
	switch cash {
	case "nor":
		context.CashStrategy = chargealgo.CashNormal{}
	case "dis":
		context.CashStrategy = chargealgo.CashRebate{0.7}
		context.CashStrategy.(chargealgo.CashRebate).SetConfig(0.8)
	case "ret":
		context.CashStrategy = chargealgo.CashReturn{400, 100}
		context.CashStrategy.(chargealgo.CashReturn).SetConfig(400, 100)
	default:
		context.CashStrategy = nil
		err = errors.New("CashContext: the strategy don't exist")
	}
	return err
}

func (context *CashContext) AcceptNum(money float64) float64 {
	return context.CashStrategy.AcceptCash(money)
}
