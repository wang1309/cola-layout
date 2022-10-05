package customer

type customerType int

const (
	POTENTIAL customerType = iota
	INTENTIONAL
	IMPORTANT
	VIP
)
