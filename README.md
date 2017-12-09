# Seppo
[![Build Status](https://travis-ci.org/koodinikkarit/seppo.svg?branch=master)](https://travis-ci.org/koodinikkarit/seppo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koodinikkarit/seppo)](https://goreportcard.com/report/github.com/koodinikkarit/seppo)

## Environment variables

```
SEPPO_USE_CONFIG_FILE
SEPPO_DB_USERNAME
SEPPO_DB_PASSWORD
SEPPO_DB_IP
SEPPO_DB_PORT
SEPPO_DB_NAME
SEPPO_PORT
```

## Config file

```
dbUsername
dbPasswd
dbIp
dbPort
dbName
port
```

### Database Migrations

Seppo uses migrate(https://github.com/mattes/migrate) for migrations and inside migrations directory sql migration files. Format is yearmonthdayhoursminutesseconds_migrationdescription

### Database models

Seppo uses gorm orm library (https://github.com/jinzhu/gorm) and database models are under db folder.