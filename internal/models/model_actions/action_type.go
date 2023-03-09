package model_actions

type ActionType int

const (
	CREATE ActionType = iota
	EDIT
	DELETE
)

func ActionTypeToString(action ActionType) string {
	switch action {
	case CREATE:
		return "create"
	case EDIT:
		return "edit"
	case DELETE:
		return "delete"
	default:
		return ""
	}
}
