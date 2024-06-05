package authentication_method

type Method uint

const (
	// UNKNOWN exists so that zero values of Method don't masquerade as IAP.
	UNKNOWN Method = iota
	IAP
	GHA
	TEST
	LOCAL
	// SHERLOCK_INTERNAL means Sherlock authenticating as itself internally
	// to take some action inside the database. An example would be updating
	// a table based on a cronjob or something, where the table has RBAC that
	// checks the authenticated User and Sherlock wants to attribute the
	// action to itself.
	SHERLOCK_INTERNAL
)
