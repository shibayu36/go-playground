.PHONY: migrate
migrate:
	mysql -u root -ppassword -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS diary"
	mysql -u root -ppassword -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS diary_test"
	goose -dir db/migrations mysql "root:password@(127.0.0.1:3306)/diary" up
	goose -dir db/migrations mysql "root:password@(127.0.0.1:3306)/diary_test" up

.PHONY: run
run: migrate
	DATABASE_DSN="root:password@(127.0.0.1:3306)/diary?parseTime=true" go run ./cmd/diary
