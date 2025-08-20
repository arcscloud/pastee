package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Store interface {
	GetPaste(id string) (Paste, error)
	SavePaste(id string, contents string, encrypted bool, expireAt string) error
}

type db struct {
	ctx *sql.DB
}

type Paste struct {
	Hash    string
	Content string
	Hashed  bool
}

func New() Store {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	ctx, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	init := []string{
		`CREATE TABLE IF NOT EXISTS pastes (
            id INT NOT NULL AUTO_INCREMENT,
            paste_id TEXT,
            content TEXT,
            encrypted TINYINT,
            expire_at DATETIME,
            created_at DATETIME,
            
            PRIMARY KEY (ID)
        )
        `,
	}
	for _, statement := range init {
		_, err = ctx.Exec(statement)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &db{
		ctx: ctx,
	}
}

func (d db) GetPaste(id string) (Paste, error) {
	stmt, err := d.ctx.Prepare("SELECT content, encrypted FROM pastes WHERE paste_id = ?")
	if err != nil {
		return Paste{}, err
	}
	defer stmt.Close()

	var contents string
	var encrypted bool
	err = stmt.QueryRow(id).Scan(&contents, &encrypted)
	if err != nil {
		return Paste{}, err
	}

	return Paste{
		Content: contents,
		Hashed:  encrypted,
	}, nil
}

func (d db) SavePaste(id string, content string, encrypted bool, expireAt string) error {
	_, err := d.ctx.Exec(
		`INSERT INTO pastes (paste_id, content, encrypted, expire_at, created_at) VALUES(?, ?, ?, ?, ?)`,
		id,
		content,
		encrypted,
		expireAt,
		time.Now().UTC().Format("2006-01-02T15:04:05"), // UTC for consistency
	)
	return err
}
