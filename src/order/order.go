package order

type Agreement {
	Number string
	Date string
	Issigned int
}

type Order struct {
	Changetimezone int
	Guid string
	InvoiceNumber string
	InvoiceDate string
	InvoiceSubject string
	BaseAgreement Agreement
	ExtendedAgreement Agreement
	Paymenttype int
	Paymentpercent float32
	Paymentdate string
}
