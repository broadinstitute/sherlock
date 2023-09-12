# What's a Selector?

## The Problem

You know that thing with bad APIs where you know the name of a thing but the API endpoint is like `/thing/{id}` so instead you have to use `/thing?name={name}`? Or even worse, when the list endpoint doesn't return all the info you need, so you have to get the ID from the list and then make another round trip?

It adds up. There can be differing opinions between the server and client on what uniquely specifies a thing. There can also be a lot of UI complexity to handle creating things with associations when they all need to be specified by ID.

## Sherlock's Solution

Sherlock has a concept of a "selector" to address this problem.

A Sherlock API is like this `/api/charts/v3/:selector` (that colon-syntax is Gin for "consumes rest of URL path", because selectors can contain slashes). Each data type defines how it parses selectors. For charts, either ID or chart name works: if the Leonardo chart had ID 123, both `/api/charts/v3/leonardo` and `/api/charts/v3/123` would behave identically.

Selectors are how data types are universally referred to in the API. If a JSON object referred to a chart, `"chart": "leonardo"` and `"chart": "123"` would both be equivalent for any input to Sherlock. Sherlock's outputs for those fields is standardized, usually to the name (so if you created something with `"chart": "123"`, the response would have `"chart": "leonardo"`).

## The Cool Part

This is a vaguely fun gimmick, but the power of this system comes when data types define associative selectors.

For example, take a chart release. ID and name are both valid selectors. But sometimes, we don't actually know the name of the chart release: we know the chart, say, Leonardo, and we know the environment, say, dev. In the past, you'd have to guess the name, like `leonardo-dev`, which is mostly standardized but not quite.

But chart releases also support `<environment>/<chart>` as a selector. This means that you can pass _any environment selector_ and _any chart selector_, seperated by a slash, and Sherlock will just figure it out for you. To continue with the example, if the dev environment had an ID of 456, then any of the following would be understood as unique references to `leonardo-dev`:

- `dev/leonardo`
- `dev/123`
- `456/leonardo`
- `456/123`

Yeah, `456/123` isn't a very human-readable selector, but who cares -- if you were in a spot where you needed it, you'd be glad that it existed.

This is neat for data types that don't really have names, too. CiIdentifiers are a great example. They have IDs but don't have names, because you can just use `chart-release/dev/leonardo` or `chart-release/leonardo-dev` etc to either get `leonardo-dev`'s CiIdentifier or a 404 (CiIdentifier defines `<resource-type>/<resource-selector...>` as one of its selectors).

This mechanism also means that what would otherwise be bespoke endpoints don't have to be. `/api/users/v3/me` resolves to the calling user because "`me`" is defined as a  selector for users.

## How It Comes Together

If you load [broad.io/beehive/environments/dev/chart-releases/leonardo](https://broad.io/beehive/environments/dev/chart-releases/leonardo), Beehive will make all the following requests in parallel to load everything on the page, just from the URL:

- `/api/users/v3/me`
- `/api/environments/v3`
- `/api/environments/v3/dev`
- `/api/ci-identifiers/v3/environment/dev`
- `/api/chart-releases/v3?environment=dev`
- `/api/chart-releases/v3/dev/leonardo`
- `/api/ci-identifiers/v3/chart-release/dev/leonardo`

Beehive's use of the Remix framework fits perfectly with this model and makes it really ergonomic and performant. All Sherlock has to do is expose solid CRUD API, which is all the better for scripts and Thelma.


