### Why are references always pointers?

Gorm lets us use structs to represent associations--and we can do so with pointer or value types ([link](https://gorm.io/docs/has_one.html#Self-Referential-Has-One)).

So why do we always use pointers, causing us to have to do nil checks in a bunch of places?

The reason why has to do with returning data to the client. When we load something from the database, we load associations too, but only one layer deep.

If we used values for associations, associations at the second/third/fourth layer etc would all _appear_ to be loaded, as zero values. Up at the controller layer, we'd end up returning
a bunch of nonsense data to clients.

As a concrete example, suppose you requested a ChartRelease, and the model and controllers all represented associations with value types:

```json
{
  // Fields at this level would be present because this is the top level
  "environmentInfo": {
    // Fields at this level would be present because this is the first layer of associations
    "defaultClusterInfo": {
      // Fields at this level would be EMPTY because this is the second layer of associations, and it doesn't get loaded
    },
    "templateEnvironmentInfo": {
      // Fields at this level would be EMPTY because this is the second layer of associations, and it doesn't get loaded
      "defaultClusterInfo": {
        // Fields at this level would be EMPTY because this is the third layer of associations, and it doesn't get loaded
      }
    }
  }
}
```

The above is true _even if we use `omitempty`_, because structs in Go don't themselves have an empty value.

The solution is to use pointer types for associations at the controller layer, so `omitempty` works. That gives us this:

```json
{
  // Fields at this level would be present because this is the top level
  "environmentInfo": {
    // Fields at this level would be present because this is the first layer of associations
    // `defaultClusterInfo` and `templateEnvironmentInfo` were nil and so weren't included here
  }
}
```

If we want to use pointer types for associations at the controller, it is very helpful to use them at the model too. This means that a lot of functions inside the model need to handle
pointer types.

### Internal stores

There's a concept here of an "internal" store versus a normal store that gets exported. Here's the differences:

| internalModelStore                        | ModelStore                                     |
|-------------------------------------------|------------------------------------------------|
| Singleton (kept in package-level vars)    | Instanced                                      |
| Stateless                                 | Stateful (keeps a database reference as state) |
| Methods require a database reference      | Methods use database reference from state      |
| Methods accept full model type as queries | Methods accept selector strings as queries     |

(This table references model stores but event stores follow the same principle)

The basic idea is that internal stores can all call each other, even passing transaction references when necessary, while
the exported stores are encapsulated things suitable for a controller to carry around.

