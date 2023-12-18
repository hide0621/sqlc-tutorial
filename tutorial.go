package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"

	"tutorial.sqlc.dev/app/tutorial"
)

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println("Not yet inserted.", authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println("Inserted: ", insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	log.Println("Fetched: ", fetchedAuthor)

	// インサートしてみた
	insertedAuthor2, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "田中",
		Bio:  sql.NullString{String: "", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println("Inserted: ", insertedAuthor2)

	// インサートしたものが参照できるかテスト
	fetchedAuthor2, err := queries.GetAuthor(ctx, insertedAuthor2.ID)
	if err != nil {
		return err
	}
	log.Println("Fetched: ", fetchedAuthor2)

	// テーブル内の対象レコードが全て取れるかテスト
	authors2, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors2)

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	log.Println(reflect.DeepEqual(insertedAuthor2, fetchedAuthor2))

	return nil

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
