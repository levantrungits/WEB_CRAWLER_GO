## RUN ENV
docker start 50f05238495f       ===> Start laradock_mariadb container
mysql -uroot -proot -h127.0.0.1 ===> Info connect mysql

## RUN PROJECT
cd WEB_CRAWLER_GO
go build
./go-crawler-modules (or go run main.go)

## PROJECT STRUCTURE
...TODO
