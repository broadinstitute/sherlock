### Gorm Gotchas

#### Hook Receivers

If you do `db.Where(&a).Updates(&b)`, hooks will run on `b`.

What you probably want is `db.Model(&a).Updates(&b)`, where hooks will be run on `a`.

A trick to help remember this is that the argument to `Where` doesn't necessarily refer to one row, so the hooks don't have strong meaning. A non-empty argument to `Model` is interpreted to specifically refer to the model being operated on.

#### Error `invalid value, should be pointer to struct or slice`

You might have something like this:

```go
type MyAssociation struct {}

type MyType struct {
	AssociationID *uint
	Association   *MyAssociation
}

func (m *MyType) myFunction(tx *gorm.DB) error {
	// ...
	if err := tx.Take(m.Association, *m.AssociationID).Error; err != nil {
		return err
	}   
	// ...
}
```

You might think that because `m.Association` is `*MyAssociation`, it is already a pointer you can pass to Gorm.

You'd be wrong, unfortunately. You need to do `&m.Association`, effectively passing `**MyAssociation` to Gorm (!). Gorm will work as expected from here on out, it'll pierce through the pointers to load values into the struct.

#### Gorm functions seem to have no effect

You're probably doing something that's changing the Gorm session from a hook. This breaks Gorm, unfortunately.

This means you can't change the context/user or begin skipping hooks from a hook.
