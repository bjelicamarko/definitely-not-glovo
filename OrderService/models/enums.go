package models

type OrderStatus string

const (
	ORDERED   OrderStatus = "ORDERED"   // Poruceno
	ACCEPTED  OrderStatus = "ACCEPTED"  // Prihvaceno
	READY     OrderStatus = "READY"     // Spremno
	TAKEN     OrderStatus = "TAKEN"     // Preuzeto
	DELIVERED OrderStatus = "DELIVERED" // Dostavljeno
	CANCELLED OrderStatus = "CANCELLED" // Obustavljeno
)

func (os OrderStatus) String() string {
	switch os {
	case ORDERED:
		return "ORDERED"
	case ACCEPTED:
		return "ACCEPTED"
	case READY:
		return "READY"
	case TAKEN:
		return "TAKEN"
	case DELIVERED:
		return "DELIVERED"
	case CANCELLED:
		return "CANCELLED"
	}
	return "unknown"
}
