package payment

type Payment interface {
	GetClient() Payment
}

func Get(pm string) Payment {

	return nil
}
