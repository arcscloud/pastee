package clean

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
)

type Clean interface {
    PerformCleanup() error

    getPastes() ([]int64, error)
    deletePaste(pasteId int64) error
}

type db struct {
    ctx *sql.DB
}

func New() Clean {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
    ctx, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    return &db{ctx: ctx}
}

func (d db) getPastes() ([]int64, error) {
    date := time.Now().Add(-1 * time.Hour).Format("2006-01-02T15:04:05")
    stmt, err := d.ctx.Prepare("SELECT id FROM pastes WHERE expire_at > ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    var pastes []int64
    rows, err := stmt.Query(date)
    if err != nil {
        return nil, err
    }

    defer rows.Close()
    for rows.Next() {
        var pasteId int64
        err := rows.Scan(&pasteId)
        if err != nil {
            return nil, err
        }
        pastes = append(pastes, pasteId)
    }

    return pastes, err
}

func (d db) deletePaste(pasteId int64) error {
    stmt, err := d.ctx.Prepare("DELETE FROM pastes WHERE id = ?")
    if err != nil {
        return err
    }

    defer stmt.Close()
    _, err = stmt.Exec(pasteId)

    return err
}

func (d db) PerformCleanup() error {
    expiredPastes, err := d.getPastes()
    if err != nil {
        return err
    }
    fmt.Printf("%d pastes marked for deletion\n", len(expiredPastes))
    deleted := 0
    for _, paste := range expiredPastes {
        err := d.deletePaste(paste)
        if err != nil {
            fmt.Printf("Paste with ID %s was scheduled for deletion but an error occured: %v\n", paste, err)
        }
        deleted++
    }
    fmt.Printf("Successfully deleted %d/%d pastes\n", deleted, len(expiredPastes))

    return nil
}
