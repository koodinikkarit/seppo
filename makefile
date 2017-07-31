
install:
	go get google.golang.org/grpc
	go get github.com/jinzhu/gorm
	go get github.com/jinzhu/gorm/dialects/mysql
	go get github.com/denisenkom/go-mssqldb
	go get github.com/go-sql-driver/mysql
	go get gopkg.in/yaml.v2

build:
	go build