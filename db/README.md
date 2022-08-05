# How to rollback the PROD DATABASE to a previous migration version

With Google's cloud_sql_proxy ([link](https://cloud.google.com/sql/docs/mysql/sql-proxy#install)) and gcloud authenticated with DevOps user credentials:
```bash
export SHERLOCK_DB_INSTANCE_PROJECT=$(vault read -field=project secret/suitable/sherlock/prod/postgres/instance)
export SHERLOCK_DB_INSTANCE_REGION=$(vault read -field=region secret/suitable/sherlock/prod/postgres/instance)
export SHERLOCK_DB_INSTANCE_NAME=$(vault read -field=name secret/suitable/sherlock/prod/postgres/instance)
cloud_sql_proxy -instances=$SHERLOCK_DB_INSTANCE_PROJECT:$SHERLOCK_DB_INSTANCE_REGION:$SHERLOCK_DB_INSTANCE_NAME=tcp:5432
```

In a new console, with golang-migrate (`brew install golang-migrate`):
```bash
export SHERLOCK_DB_NAME=$(vault read -field=db secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DB_USER=$(vault read -field=username secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DB_PASSWORD=$(vault read -field=password secret/suitable/sherlock/prod/postgres/sherlock-db-creds)
export SHERLOCK_DB_URL="postgres://$SHERLOCK_DB_USER:$SHERLOCK_DB_PASSWORD@localhost:5432/$SHERLOCK_DB_NAME?sslmode=disable"
migrate -path ./db/migrations -database $SHERLOCK_DB_URL version
```

You can then use something like `migrate -source ./db/migrations -database $SHERLOCK_DB_URL goto` to set the database to a particular numbered migration version.
See the migration files for the available numbered migration versions.