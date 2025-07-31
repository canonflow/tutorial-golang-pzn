package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-mysql/db"
	"golang-mysql/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(db.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "test repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(db.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 124)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
	/*
		=== RUN   TestFindById
		{124 repository@test.com test repository}
		--- PASS: TestFindById (0.00s)
		PASS
	*/

	comment, err = commentRepository.FindById(context.Background(), 999)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
	/*
		--- FAIL: TestFindById (0.00s)
		panic: ID: 999 Not Found [recovered]
			panic: ID: 999 Not Found
	*/
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(db.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
	/*
		=== RUN   TestFindAll
		{1 nathan@gmail.com test comment}
		{2 nathan@gmail.com test comment lagi}
		{3 nathan0@gmail.com Comment ke 0}
		{4 nathan1@gmail.com Comment ke 1}
		{5 nathan2@gmail.com Comment ke 2}
		{6 nathan3@gmail.com Comment ke 3}
		--- PASS: TestFindAll (0.00s)
		PASS
	*/
}
