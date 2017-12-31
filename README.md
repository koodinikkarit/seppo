# Seppo
[![Build Status](https://travis-ci.org/koodinikkarit/seppo.svg?branch=master)](https://travis-ci.org/koodinikkarit/seppo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koodinikkarit/seppo)](https://goreportcard.com/report/github.com/koodinikkarit/seppo)

## Environment variables

```
DB_USERNAME
DB_PASSWORD
DB_IP
DB_PORT
DB_NAME
SEPPO_PORT
MATIAS_PORT
```

### Database Migrations

Seppo uses migrate(https://github.com/mattes/migrate) for migrations and inside migrations directory sql migration files. Format is yearmonthdayhoursminutesseconds_migrationdescription

### Database models

Seppo uses gorm orm library (https://github.com/jinzhu/gorm) and database models are under db folder.