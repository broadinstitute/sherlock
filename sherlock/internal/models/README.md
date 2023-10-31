### Gorm Gotchas

If you do `db.Where(&a).Updates(&b)`, hooks will run on `b`.

What you probably want is `db.Model(&a).Updates(&b)`, where hooks will be run on `a`.

A trick to help remember this is that the argument to `Where` doesn't necessarily refer to one row, so the hooks don't have strong meaning. A non-empty argument to `Model` is interpreted to specifically refer to the model being operated on.