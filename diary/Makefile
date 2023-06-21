.PHONY: migrate
migrate:
	mysql -u root -ppassword -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS diary"
	mysql -u root -ppassword -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS diary_test"
	migrate -database 'mysql://root:password@(127.0.0.1:3306)/diary' -path db/migrations up
	migrate -database 'mysql://root:password@(127.0.0.1:3306)/diary_test' -path db/migrations up
