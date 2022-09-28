package state

type AppointmentStatus string

const (
	PENDING          AppointmentStatus = "PENDING"
	ACCEPTED         AppointmentStatus = "ACCEPTED"
	REJECTED         AppointmentStatus = "REJECTED"
	PAYMENT_APPROVED AppointmentStatus = "PAYMENT_APPROVED"
	PAYMENT_DECLINED AppointmentStatus = "PAYMENT_DECLINED"
	CLOSED           AppointmentStatus = "CLOSED"
)
