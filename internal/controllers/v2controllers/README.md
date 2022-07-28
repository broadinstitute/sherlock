### Wow, what a lot of struct tags:
#### Do use:
- `json` controls the field name when parsing to/from json (always add)
- `form` controls the field name when parsing from query parameters (always add)
- `swaggertype` can override the type of the field documented on Swagger, useful for anything recursive (only add when Swaggo is parsing the type incorrectly)
- `enums` controls possible values for the field as documented on Swagger (add when reasonable)
- `default` controls (add when reasonable):
  - default values for the field on Swagger (as in, Swagger will fill in the default for you)
  - default values applied internally by Sherlock when it goes to create an entry


I suggest tags be in the order above. [Swaggo](https://github.com/swaggo/swag#available) and [Gin](https://github.com/gin-gonic/gin#model-binding-and-validation) both use the tags on these structs.

#### Don't use
- `validate:"required"` which will document fields as required on Swagger, but in an over-zealous way that interferes with omitting them in query parameters on list calls



