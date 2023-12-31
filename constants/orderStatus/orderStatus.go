package orderStatus

type ORDER_STATUS string

const (
	ToBeCreated ORDER_STATUS = "TO_BE_CREATED"
	Created     ORDER_STATUS = "CREATED"
	Processing  ORDER_STATUS = "PROCESSING"
	Completed   ORDER_STATUS = "COMPLETED"
	Canceled    ORDER_STATUS = "CANCELED"
	Failed      ORDER_STATUS = "FAILED"
)
