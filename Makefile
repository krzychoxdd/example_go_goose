dbinstall:
	docker run -d --name=exmpl-mysql --publish 6604:3306 -e MYSQL_ROOT_PASSWORD=1234567 -e MYSQL_ROOT_HOST="%" -d mysql/mysql-server:5.7

dbstart:
	docker start exmpl-mysql

dbcreate:
	docker exec -it exmpl-mysql mysql -uroot -p1234567 -e 'CREATE DATABASE exampledb;'

migrate:
	cd cmd && go run migrator.go root:1234567@tcp\(0.0.0.0:6604\)/exampledb?parseTime=true up

.PHONY: dbinstall dbstart dbcreate migrate
