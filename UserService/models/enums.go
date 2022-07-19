package models

type Role string

const (
	ADMIN     Role = "ADMIN"
	APPUSER   Role = "APPUSER"
	EMPLOYEE  Role = "EMPLOYEE"
	DELIVERER Role = "DELIVERER"
)

func (s Role) String() string {
	switch s {
	case ADMIN:
		return "ADMIN"
	case APPUSER:
		return "APPUSER"
	case EMPLOYEE:
		return "EMPLOYEE"
	case DELIVERER:
		return "DELIVERER"
	}
	return "unknown"
}
