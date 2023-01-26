dbrun:
	docker start example-mysql || docker run -d --name=example-mysql --publish 6603:3306 -e MYSQL_ROOT_PASSWORD=1234567 -e MYSQL_ROOT_HOST="%" -d mysql/mysql-server:5.7

migrate:
	cd cmd && go run migrator.go root:1234567@tcp\(0.0.0.0:6603\)/exampledb?parseTime=true up

.PHONY: dbrun migrate
