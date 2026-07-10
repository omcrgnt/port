// Package port defines org-wide outbound ports for application orchestration.
package port

import "context"

// Transactor runs fn inside a local ACID transaction boundary.
// Postgres implementations delegate to conn-sql; in-memory adapters may run fn directly.
type Transactor interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
