package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS authors;

CREATE TABLE  authors (
	id serial PRIMARY KEY,
	name text,
	age integer
);

CREATE TABLE posts (
	id serial PRIMARY KEY,
	title text,
	content text,
	author_id integer REFERENCES authors(id)
);`

const (
	host     string = "localhost"
	port     int    = 5432
	user     string = "postgres"
	password string = "password"
	dbname   string = "go-lazy-loading"
)

var db *sqlx.DB

func Connect() {
	var err error

	connectionParams := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sqlx.Connect("postgres", connectionParams)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	db.MustExec(
		"INSERT INTO authors (name, age) VALUES ($1, $2)",
		"John Doe",
		30,
	)

	db.MustExec(
		"INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3)",
		"Lazy loading in Go",
		"Trying to work with lazy loading in Go and SQLX",
		1,
	)
}

type Author struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Post struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	AuthorId int    `db:"author_id"`
	Author   Author
}

func (p *Post) LoadAuthor() {
	if p.Author.Id != 0 {
		fmt.Println("Author already loaded")
		return
	}

	p.Author = GetAuthor(p.AuthorId)
}

func GetAuthor(id int) Author {
	var author Author
	err := db.Get(&author, "SELECT * FROM authors WHERE id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}

	return author
}

func GetPost(id int) Post {
	post := Post{}
	err := db.Get(&post, "SELECT * FROM posts WHERE id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}

	return post
}
