### What's the overall layout? What's different between the v1 and v2 APIs?

The primary thing to note is that v1 was built as a listener and v2 was built to itself be the authority. v2 isn't a re-do of v1, it is a different design with an expanded scope.

- `models` provides database types and functions
  - v2 also has data validation and permissions handling here
- `controllers` exposes functions with API-like behavior
  - v1 does so in terms of model types themselves, while v2 provides its own API types
  - v2 also has all default-value logic here
- `serializers` is for v1 only, and it converts model types to API types it defines
  - v2's controller does the conversions itself because the API types are tightly coupled to the signatures of the API-like functions
- `handlers` wires Gin up to actually call controller functions
  - v1 has logic for error messages here, while v2 generally relies on the `errors` package
  - v2 also has Swagger page documentation here

The other notable overall difference between v1 and v2 is that v2 was written after Go's generics was introduced--it makes use of them to help provide a wider API across its types, since v2 plans to have new clients.

### What about the other packages?
Briefly:
- `auth` is for v2 and includes code for turning IAP headers into actionable "is this user suitable" info
- `cli` is for v1 right now; it is the Cobra commands used when Sherlock is run as a CLI client for itself
- `config` is for Koanf-powered configuration across Sherlock
  - v1 used to use Viper but has been converted to Koanf to follow the lead of Thelma and Monkey
- `db` is for direct database work--namely migrations
- `metrics` is for v1 right now; it exposes the acclerate metrics Sherlock's v1 was originally built for
- `sherlock` defines the top-level application instance and configures Gin
- `testutils` includes helpers for running functional tests within Go itself
- `version` records and embeds Sherlock's own version at build-time
