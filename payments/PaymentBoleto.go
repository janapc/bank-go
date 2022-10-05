package payments

type verifyAccount interface {
	WithdrawMoney(value float64) string
}

func PaymentBoleto(account verifyAccount, boletoValue float64) string {
	return account.WithdrawMoney(boletoValue)
}
