package database


var defaultSchema = Schema{
	create: `
`,
	drop: `
`,
}

func init() {
 
var db *sqlx.DB
 
// exactly the same as the built-in
db = sqlx.Open("sqlite3", ":memory:")
 
// from a pre-existing sql.DB; note the required driverName
db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")
 
// force a connection and test that it worked
err = db.Ping()


// execute a query on the server
result, err := db.Exec(schema)


p := Place{}
pp := []Place{}
 
// this will pull the first place directly into p
err = db.Get(&p, "SELECT * FROM place LIMIT 1")
 
// this will pull places with telcode > 50 into the slice pp
err = db.Select(&pp, "SELECT * FROM place WHERE telcode > ?", 50)
 
// they work with regular types as well
var id int
err = db.Get(&id, "SELECT count(*) FROM place")
 
// fetch at most 10 place names
var names []string
err = db.Select(&names, "SELECT name FROM place LIMIT 10")

}

