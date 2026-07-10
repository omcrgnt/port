# port

Org-wide outbound ports for omcrgnt applications.

## Transactor

```go
type Transactor interface {
    WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
```

Implementations:

- **conn-sql** — Postgres `BEGIN` / `COMMIT` / `ROLLBACK`
- **memory** (per app) — runs `fn(ctx)` without a real transaction

Saga / TCC orchestrators will be added here separately when needed.
