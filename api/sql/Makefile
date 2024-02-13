mysql:
	docker run -p 3306:3306 --name devbook-mysql -e MYSQL_ROOT_PASSWORD=examplepasswd -e MYSQL_DATABASE=devbook -e MYSQL_USER=golang -e MYSQL_PASSWORD=user-passwd -d mysql

create-table:
	docker exec -i devbook-mysql mysql -u golang -puser-passwd devbook < sql.sql
	
bash:
	docker exec -it devbook-mysql bash