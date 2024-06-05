# `intermediary_user`

This package contains the type definitions for how other parts of role propagation should understand "principals"
granted some permission in a remote system. 

These type definitions are trivial -- see `./identifier.go` -- but we have them in a separate package because we use 
them in function signatures and type definitions in other packages. Having them in this separate package means that
we won't hit issues with circular dependencies in our mocks.
