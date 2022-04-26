# How to rollback the PROD DATABASE to a previous migration version

With Google's cloud_sql_proxy ([link](https://cloud.google.com/sql/docs/mysql/sql-proxy#install)) and gcloud authenticated with DevOps user credentials:
```bash
export SHERLOCK_DB_INSTANCEPROJECT=$(vault read -field=project secret/suitable/sherlock/prod/postgres/instance)
export SHERLOCK_DB_INSTANCEREGION=$(vault read -field=region secret/suitable/sherlock/prod/postgres/instance)
export SHERLOCK_DB_INSTANCENAME=$(vault read -field=name secret/suitable/sherlock/prod/postgres/instance)
cloud_sql_proxy -instances=$SHERLOCK_DB_INSTANCEPROJECT:$SHERLOCK_DB_INSTANCEREGION:$SHERLOCK_DB_INSTANCENAME=tcp:5432
```

In a new console, with golang-migrate (`brew install golang-migrate`):
```bash
export SHERLOCK_DBNAME=$(vault read -field=db secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DBUSER=$(vault read -field=username secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DBPASSWORD=$(vault read -field=password secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DBURL="postgres://$SHERLOCK_DBUSER:$SHERLOCK_DBPASSWORD@localhost:5432/$SHERLOCK_DBNAME?sslmode=disable"
migrate -path ./db/migrations -database $SHERLOCK_DBURL version
```

You can then use something like `migrate -source ./db/migrations -database $SHERLOCK_DBURL goto` to set the database to a particular numbered migration version.
See the migration files for the available numbered migration versions.