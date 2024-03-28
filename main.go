package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const connStr = "postgres://docker:docker@localhost:5432/race_condition"

func main() {
    ctx := context.Background()

    pool, err := pgxpool.New(ctx, connStr)
    if err != nil {
        panic(err)
    }

    defer pool.Close()

    const n = 100
    var wg sync.WaitGroup
    wg.Add(n)
    for range n {
        go func() {
            defer wg.Done()
            comprar(pool)
        } ()
    }
    wg.Wait()

    query := "SELECT quantidade FROM estoque WHERE id = 1"
    var qt int64
    if err := pool.QueryRow(ctx, query).Scan(&qt); err != nil {
        panic(err)
    }
    fmt.Println(qt)
}

func comprar(pool *pgxpool.Pool) {
    ctx := context.Background()

    tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
    if err != nil {
        panic(err)
    }
    defer func() { _ = tx.Rollback(ctx) } ()

    query := "SELECT quantidade, version FROM estoque WHERE id = 1"
    var qt, v int64
    if err := tx.QueryRow(ctx, query).Scan(&qt, &v); err != nil {
        panic(err)
    }

    if qt <= 0 {
        return
    }

    query = "UPDATE estoque set quantidade = quantidade - 1, version = version + 1 WHERE id = 1 AND version = $1"
    if _, err := tx.Exec(ctx, query, v); err != nil {
       panic(err)
    }

    // if wanna throw error
    // if err := tx.QueryRow(ctx, query, v).Scan(); err != nil {
    //     panic(err)
    // }

    if err := tx.Commit(ctx); err != nil {
        panic(err)
    }
}
