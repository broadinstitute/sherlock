package authentication_method

type Method uint

const (
	// UNKNOWN exists so that zero values of Method don't masquerade as IAP.
	UNKNOWN Method = iota
	IAP
	GHA
	TEST
	LOCAL
)
