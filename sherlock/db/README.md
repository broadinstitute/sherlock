# How to rollback the PROD DATABASE to a previous migration version

With Google's cloud-sql-proxy ([link](https://cloud.google.com/sql/docs/postgres/sql-proxy#mac-m1)) and gcloud authenticated with DevOps user credentials:
```bash
export SHERLOCK_DB_CONNECTION_NAME=$(gcloud secrets versions access latest --secret=sherlock-cloudsql-maintenance-credentials --project=dsp-devops-super-prod | jq -r '.instance')
cloud-sql-proxy $SHERLOCK_DB_CONNECTION_NAME
```

In a new console, with golang-migrate (`brew install golang-migrate`):
```bash
export SHERLOCK_DB_NAME=$(gcloud secrets versions access latest --secret=sherlock-cloudsql-maintenance-credentials --project=dsp-devops-super-prod | jq -r '.database')
export SHERLOCK_DB_USER=$(gcloud secrets versions access latest --secret=sherlock-cloudsql-maintenance-credentials --project=dsp-devops-super-prod | jq -r '.username')
export SHERLOCK_DB_PASSWORD=$(gcloud secrets versions access latest --secret=sherlock-cloudsql-maintenance-credentials --project=dsp-devops-super-prod | jq -r '.password')
export SHERLOCK_DB_URL="postgres://$SHERLOCK_DB_USER:$SHERLOCK_DB_PASSWORD@localhost:5432/$SHERLOCK_DB_NAME?sslmode=disable"
migrate -path ./db/migrations -database $SHERLOCK_DB_URL version
```

You can then use something like `migrate -source ./db/migrations -database $SHERLOCK_DB_URL goto` to set the database to a particular numbered migration version.
See the migration files for the available numbered migration versions.
