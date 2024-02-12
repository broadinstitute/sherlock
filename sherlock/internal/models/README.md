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

#### Human error with `Select`

`Select` lets you selectively load columns from the database.

That's great, but you specify arguments with strings. If you're expecting to, say, read a ChartRelease's AppVersionBranch field, but `Select` some columns and forget to specify `app_version_branch`, you'll get a zero value for the field and no error. That can be tough to track down.

A slightly better mechanism in some circumstances is to use what Gorm calls "smart select." The idea is basically this:

```go
func myFunction(db *gorm.DB) {
	var chartRelease struct {
		AppVersionBranch string
	}
	err := db.
		Model(&ChartRelease{}).
		Where(&ChartRelease{ID: 123}).
		First(&chartRelease).
		Error
	println(chartRelease.AppVersionBranch)
}
```

Essentially, rather than explicitly adding a `Select` to the chain, you simply parse the result into a struct with just the fields you need. Gorm will automatically only select those fields. 

You'll get a runtime error if your custom struct includes a field that isn't a column in the table you're looking for (which would happen for `Select` too) but the big thing is that you'll get compile-time errors if you access any fields that aren't on the struct. This prevents you from unexpectedly getting zero values for fields that are actually present in the table.