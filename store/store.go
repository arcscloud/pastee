package store

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "log"
    "os"
    "time"
)

type Store interface {
    GetPaste(id string) (Paste, error)
    SavePaste(id string, contents string, hashed bool) error
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
    _ = godotenv.Load()

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
            hashed TINYINT,
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
    stmt, err := d.ctx.Prepare("SELECT content, hashed FROM pastes WHERE paste_id = ?")
    if err != nil {
        return Paste{}, err
    }
    defer stmt.Close()

    var contents string
    var hashed bool
    err = stmt.QueryRow(id).Scan(&contents, &hashed)
    if err != nil {
        return Paste{}, err
    }

    return Paste{
        Content: contents,
        Hashed:  hashed,
    }, nil
}

func (d db) SavePaste(id string, content string, hashed bool) error {
    _, err := d.ctx.Exec(
        `INSERT INTO pastes (paste_id, content, hashed, created_at) VALUES(?, ?, ?, ?)`,
        id,
        content,
        hashed,
        time.Now().Format("2006-01-02T15:04:05"),
    )
    return err
}
