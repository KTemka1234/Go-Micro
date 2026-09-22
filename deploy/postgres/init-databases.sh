#!/usr/bin/env bash
# Creates a database and an owner role for every service listed in POSTGRES_DATABASES.
# Role name = database name, password is taken from <NAME>_DB_PASSWORD (e.g. AUTH_DB_PASSWORD).
#
# Runs automatically from /docker-entrypoint-initdb.d on an empty volume.
# The script is idempotent, so it can also be applied to an existing volume:
#   docker compose exec postgres bash /docker-entrypoint-initdb.d/init-databases.sh
set -euo pipefail

if [[ -z "${POSTGRES_DATABASES:-}" ]]; then
  echo "init-databases: POSTGRES_DATABASES is empty, nothing to do"
  exit 0
fi

for db in ${POSTGRES_DATABASES//,/ }; do
  password_var="${db^^}_DB_PASSWORD"
  password="${!password_var:-}"
  if [[ -z "$password" ]]; then
    echo "init-databases: $password_var is not set" >&2
    exit 1
  fi

  echo "init-databases: ensuring database and role '$db'"
  psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d postgres -v db="$db" -v password="$password" <<'SQL'
SELECT format('CREATE ROLE %I LOGIN', :'db')
WHERE NOT EXISTS (SELECT FROM pg_roles WHERE rolname = :'db') \gexec

SELECT format('ALTER ROLE %I WITH LOGIN PASSWORD %L', :'db', :'password') \gexec

SELECT format('CREATE DATABASE %I OWNER %I', :'db', :'db')
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = :'db') \gexec

SELECT format('ALTER DATABASE %I OWNER TO %I', :'db', :'db') \gexec

-- services must not be able to connect to each other's databases
SELECT format('REVOKE ALL ON DATABASE %I FROM PUBLIC', :'db') \gexec
SQL
done
