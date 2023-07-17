package main

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db-range-query-perf")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {
		log.Printf("%d: Creating %d users with %d posts", i, 1000, 100)
		CreateNUsersWithPosts(db, 1000, 100)
		log.Printf("%d: Created %d users with %d posts", i, 1000, 100)
	}
}

// CreateNUsersWithJournals は、uCount人のuserを作成し、各ユーザーにpCount個のpostを作成します。それぞれフォロー状態にします。
func CreateNUsersWithPosts(db *sql.DB, uCount int, pCount int) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	userIDs := make([]int, uCount)
	for i := 0; i < uCount; i++ {
		res, err := db.Exec("INSERT INTO users (name) VALUES (?)", fmt.Sprintf("user%d", i))
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		userIDs[i] = int(id)

		// Create posts with bulk insert using squirrel
		insertBuilder := psql.Insert("posts").Columns("user_id", "body")

		for j := 0; j < pCount; j++ {
			insertBuilder = insertBuilder.Values(userIDs[i], fmt.Sprintf("Post %d of User%d", j, userIDs[i]))
		}

		query, args, err := insertBuilder.ToSql()
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create followings
	for i := 0; i < uCount; i++ {
		insertBuilder := psql.Insert("follows").Columns("follower_id", "followee_id")

		for j := 0; j < uCount; j++ {
			insertBuilder = insertBuilder.Values(userIDs[i], userIDs[j])
		}

		query, args, err := insertBuilder.ToSql()
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(query, args...)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
