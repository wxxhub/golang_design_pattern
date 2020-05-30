package strategy

type Travel struct {
	Travel_strategy Strategy
}

func (t *Travel) Travel () {
	t.Travel_strategy.Travel()
}