package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Sample struct {
	ID   string
	Name string
}

var db = OpenConnection()

func TestConnection(t *testing.T) {
	assert.NotNil(t, db)

	/*
		=== RUN   TestConnection
		--- PASS: TestConnection (0.00s)
		PASS
	*/
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "nathan").Error

	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "garzya").Error

	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "3", "santoso").Error

	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values (?, ?)", "4", "canonflow").Error

	assert.Nil(t, err)

	/*
		=== RUN   TestExecuteSQL
		--- PASS: TestExecuteSQL (0.00s)
		PASS
	*/
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error

	assert.Nil(t, err)
	assert.Equal(t, "1", sample.ID)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error

	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))

	/*
		=== RUN   TestRawSQL
		--- PASS: TestRawSQL (0.00s)
		PASS
	*/
}

func TestSQLRow(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)

	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{ID: id, Name: name})
	}

	assert.Equal(t, 4, len(samples))

	/*
		=== RUN   TestSQLRow
		--- PASS: TestSQLRow (0.00s)
		PASS
	*/
}

func TestScanRows(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()

	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 4, len(samples))

	/*
		=== RUN   TestScanRows
		--- PASS: TestScanRows (0.00s)
		PASS
	*/
}
