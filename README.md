# Seppo
[![Build Status](https://travis-ci.org/koodinikkarit/seppo.svg?branch=master)](https://travis-ci.org/koodinikkarit/seppo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koodinikkarit/seppo)](https://goreportcard.com/report/github.com/koodinikkarit/seppo)

### Environment variables

```
SEPPO_DB_USERNAME = Database username
SEPPO_DB_PASSWORD = Database password
SEPPO_DB_IP = Database ip
SEPPO_DB_PORT = Database port
SEPPO_DB_NAME = Database name
SEPPO_PORT = Seppo service port
```

### Database Migrations

Seppo uses migrate(https://github.com/mattes/migrate) for migrations and inside migrations directory sql migration files. Format is yearmonthdayhoursminutesseconds_migrationdescription

### Database models

Seppo uses sqlboiler(https://github.com/volatiletech/sqlboiler) for as database orm. Sql boiler uses db access information from sqlboiler.toml config file to connect to database and fetch its schema. Sqlboiler then uses this schema for generating database models

#### Example databaseconfig file

```
blacklist=["schema_migrations"]
schema="schema"
[mysql]
	dbname="name"
	host="host"
	port=port
	user="user"
	pass="pass"
	sslmode="false"
```
Table name in blacklist is used for migrare program so there should not be generated code for it.