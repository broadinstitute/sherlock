# `role_propagation`

This package contains the logic for propagating different role grants out to their respective cloud providers. Put
differently, this package is the "group sync" part of Sherlock... plus the "manage Firecloud.org accounts" part, plus
the "manage GitHub Org members" part, and so on. This logic is generic so that with a grant stored on a role, and an
engine in `propagation_engines`, we can propagate that grant to the remote system or cloud provider.

There's two sub-packages here:
- `intermediary_user` contains the type definitions for how we understand who does or doesn't have a grant on a 
  particular cloud provider
- `propagation_engines` contains the adapters for actually bridging the logic here to actions on cloud providers

This package itself is split up into three parts:
- `propagator.go`, and the files prefixed with `propagator_`, contain the logic for propagating a single grant from a
  single role to a single cloud provider
- `propagate.go` knows how to run all propagators sequentially (and it knows how to do so non-concurrently)
- `boot.go` is run only during normal full Sherlock boot, and it wires up the set of propagators that we want to have at
  runtime
  - In tests, we can use `test_helpers.go` to wire up whatever we need
