package auth_models

type AuthMethod int

const (
	AuthMethodIAP AuthMethod = iota
	AuthMethodGHA
)
