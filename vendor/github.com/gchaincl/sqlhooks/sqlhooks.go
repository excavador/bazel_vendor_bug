package sqlhooks

import (
	"context"
	"database/sql/driver"
)

// Hook is the hook callback signature
type Hook func(ctx context.Context, query string, args ...interface{}) (context.Context, error)

// Hooks instances may be passed to Wrap() to define an instrumented driver
type Hooks interface {
	Before(ctx context.Context, query string, args ...interface{}) context.Context
	After(ctx context.Context, err error, query string, args ...interface{}) context.Context
}

// Driver implements a database/sql/driver.Driver
type Driver struct {
	driver.Driver
	hooks Hooks
}

// Open opens a connection
func (drv *Driver) Open(name string) (driver.Conn, error) {
	conn, err := drv.Driver.Open(name)
	if err != nil {
		return conn, err
	}

	return &Conn{conn, drv.hooks}, nil
}

// Conn implements a database/sql.driver.Conn
type Conn struct {
	Conn  driver.Conn
	hooks Hooks
}

func (conn *Conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	var (
		stmt driver.Stmt
		err  error
	)

	if c, ok := conn.Conn.(driver.ConnPrepareContext); ok {
		stmt, err = c.PrepareContext(ctx, query)
	} else {
		stmt, err = conn.Prepare(query)
	}

	if err != nil {
		return stmt, err
	}

	return &Stmt{stmt, conn.hooks, query}, nil
}

func (conn *Conn) Prepare(query string) (driver.Stmt, error) { return conn.Conn.Prepare(query) }
func (conn *Conn) Close() error                              { return conn.Conn.Close() }
func (conn *Conn) Begin() (driver.Tx, error)                 { return conn.Conn.Begin() }

// Stmt implements a database/sql/driver.Stmt
type Stmt struct {
	Stmt  driver.Stmt
	hooks Hooks
	query string
}

func (stmt *Stmt) execContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if s, ok := stmt.Stmt.(driver.StmtExecContext); ok {
		return s.ExecContext(ctx, args)
	}

	values := make([]driver.Value, len(args))
	for _, arg := range args {
		values[arg.Ordinal-1] = arg.Value
	}

	return stmt.Exec(values)
}

func (stmt *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	var err error

	list := namedToInterface(args)

	// Exec `Before` Hooks
	if ctx = stmt.hooks.Before(ctx, stmt.query, list...); err != nil {
		return nil, err
	}
	results, err := stmt.execContext(ctx, args)
	stmt.hooks.After(ctx, err, stmt.query, list...)
	return results, err
}

func (stmt *Stmt) queryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if s, ok := stmt.Stmt.(driver.StmtQueryContext); ok {
		return s.QueryContext(ctx, args)
	}

	values := make([]driver.Value, len(args))
	for _, arg := range args {
		values[arg.Ordinal-1] = arg.Value
	}
	return stmt.Query(values)
}

func (stmt *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	var err error

	list := namedToInterface(args)

	// Exec Before Hooks
	if ctx = stmt.hooks.Before(ctx, stmt.query, list...); err != nil {
		return nil, err
	}
	rows, err := stmt.queryContext(ctx, args)
	stmt.hooks.After(ctx, err, stmt.query, list...)
	return rows, err
}

func (stmt *Stmt) Close() error                                    { return stmt.Stmt.Close() }
func (stmt *Stmt) NumInput() int                                   { return stmt.Stmt.NumInput() }
func (stmt *Stmt) Exec(args []driver.Value) (driver.Result, error) { return stmt.Stmt.Exec(args) }
func (stmt *Stmt) Query(args []driver.Value) (driver.Rows, error)  { return stmt.Stmt.Query(args) }

// Wrap is used to create a new instrumented driver, it takes a vendor specific driver, and a Hooks instance to produce a new driver instance.
// It's usually used inside a sql.Register() statement
func Wrap(driver driver.Driver, hooks Hooks) driver.Driver {
	return &Driver{driver, hooks}
}

func namedToInterface(args []driver.NamedValue) []interface{} {
	list := make([]interface{}, len(args))
	for i, a := range args {
		list[i] = a.Value
	}
	return list
}

/*
type hooks struct {
}

func (h *hooks) Before(ctx context.Context, query string, args ...interface{}) error {
	log.Printf("before> ctx = %+v, q=%s, args = %+v\n", ctx, query, args)
	return nil
}

func (h *hooks) After(ctx context.Context, query string, args ...interface{}) error {
	log.Printf("after>  ctx = %+v, q=%s, args = %+v\n", ctx, query, args)
	return nil
}

func main() {
	sql.Register("sqlite3-proxy", Wrap(&sqlite3.SQLiteDriver{}, &hooks{}))
	db, err := sql.Open("sqlite3-proxy", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}

	if _, ok := driver.Stmt(&Stmt{}).(driver.StmtExecContext); !ok {
		panic("NOPE")
	}

	if _, err := db.Exec("CREATE table users(id int)"); err != nil {
		log.Printf("|err| = %+v\n", err)
	}

	if _, err := db.QueryContext(context.Background(), "SELECT * FROM users WHERE id = ?", 1); err != nil {
		log.Printf("err = %+v\n", err)
	}

}
*/
