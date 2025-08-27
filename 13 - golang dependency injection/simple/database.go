package simple

type Database struct {
	Name string
}

/*
Harus bikin type alias, karena WIRE tidak mengizinkan provider memiliki tipe data yg sama lebih dari 1
*/
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	database := &Database{
		Name: "postgres",
	}

	return (*DatabasePostgreSQL)(database)
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "mongodb",
	}

	return (*DatabaseMongoDB)(database)
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMongoDB:    mongoDB,
	}
}
