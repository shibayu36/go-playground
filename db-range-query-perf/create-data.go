package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

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
		log.Printf("%d: Creating %d users with %d posts", i, 1000, 1000)
		CreateNUsersWithPosts(db, 1000, 1000)
		log.Printf("%d: Created %d users with %d posts", i, 1000, 1000)
	}
}

// CreateNUsersWithJournals は、uCount人のuserを作成し、各ユーザーにpCount個のpostを作成します。それぞれフォロー状態にします。
func CreateNUsersWithPosts(db *sql.DB, uCount int, pCount int) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Question)

	userIDs := make([]int, uCount)
	for i := 0; i < uCount; i++ {
		res, err := db.Exec("INSERT INTO users (name) VALUES (?)", fmt.Sprintf("user%d", i+1))
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		userIDs[i] = int(id)

		// Create posts with bulk insert using squirrel
		insertBuilder := psql.Insert("posts").Columns("user_id", "body", "posted_at", "deleted")

		for j := 0; j < pCount; j++ {
			insertBuilder = insertBuilder.Values(
				userIDs[i],
				fmt.Sprintf("Post %d of User%d", j+1, userIDs[i]),
				randomPostedAt(),
				randomDeleted(),
			)
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

// 2ヶ月前から現在までの間でランダムに時刻を生成します。
func randomPostedAt() time.Time {
	now := time.Now()
	twoMonthAgo := now.AddDate(0, -2, 0)

	diff := now.Unix() - twoMonthAgo.Unix()

	// ランダムな差分を生成
	randomDiff := rand.Int63n(diff)

	// ランダムな時刻を生成
	randomTime := twoMonthAgo.Add(time.Duration(randomDiff) * time.Second)

	return randomTime
}

// 0.1%の確率でtrueを返す
func randomDeleted() bool {
	return rand.Intn(1000) == 0
}
