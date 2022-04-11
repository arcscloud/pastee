package store

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
    "time"
)

type Store interface {
    GetPaste(id string) (Paste, error)
    SavePaste(id string, hash string, contents string) error
}

type db struct {
    ctx *sql.DB
}

type Paste struct {
    Hash    string
    Content string
}

func New() Store {
    dbDir := "data"
    if _, err := os.Stat(dbDir); os.IsNotExist(err) {
        os.Mkdir(dbDir, os.ModePerm)
    }

    ctx, err := sql.Open("sqlite3", dbDir+"/app.db")
    if err != nil {
        log.Fatal(err)
    }

    init := []string{
        `CREATE TABLE IF NOT EXISTS pastes (
            id TEXT PRIMARY KEY,
            hash TEXT NULL,
            content TEXT,
            created_at DATETIME
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
    stmt, err := d.ctx.Prepare("SELECT content, hash FROM pastes WHERE id = ?")
    if err != nil {
        return Paste{}, err
    }
    defer stmt.Close()

    var contents string
    var hash string
    err = stmt.QueryRow(id).Scan(&contents, &hash)
    if err != nil {
        return Paste{}, err
    }

    return Paste{
        Hash:    hash,
        Content: contents,
    }, nil
}

func (d db) SavePaste(id string, hash string, content string) error {
    _, err := d.ctx.Exec(
        `INSERT INTO pastes (id, hash, content, created_at) VALUES(?, ?, ?, ?)`,
        id,
        hash,
        content,
        time.Now().Format("2006-01-02T15:04:05"),
    )
    return err
}
