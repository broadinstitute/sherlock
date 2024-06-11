# `propagation_engines`

This package contains the adapters that bridge between Sherlock's actual propagation logic in `role_propagation` and the
mechanics of each cloud provider. With the `intermediary_user` package, the goal is that each of an engine's methods
should simply delegate to the appropriate client library: repeated conversions to and from cloud provider types should
be unnecessary.

Put differently, an engine can "speak" UUIDs or whatever else is most convenient, so that we don't need to make as many
API calls. Fewer API calls means less code and less chance of hitting rate limits.

Using that UUID example, an engine would be responsible for the following:
- Reading UUIDs that currently have the permission in the remote system
- Determining the UUIDs that should have the permission in the remote system, given a dump of Sherlock's RoleAssignments and Users
- Adding a UUID to the permission in the remote system
- Updating information fields for a UUID in the remote system (if relevant -- think first and last name or profile photo)
- Removing a UUID from the permission in the remote system

The `role_propagation` package has all the logic for actually making propagation happen using these primitives.
