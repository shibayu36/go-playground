package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	sqldb, err := sql.Open("postgres", "postgres://username:password@localhost:5432/todo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*Todo)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	h := &handlers{e: e, db: db}

	e.GET("/", h.index)
	e.Logger.Fatal(e.Start(":8989"))
}

type Data struct {
	Todos  []Todo
	Errors []error
}

type handlers struct {
	e  *echo.Echo
	db *bun.DB
}

func (h *handlers) index(c echo.Context) error {
	var todos []Todo
	ctx := context.Background()
	err := h.db.NewSelect().Model(&todos).Order("created_at").Scan(ctx)
	if err != nil {
		h.e.Logger.Error(err)
		return c.Render(http.StatusBadRequest, "index", Data{
			Errors: []error{errors.New("Cannot get todos")},
		})
	}

	return c.Render(http.StatusOK, "index", Data{Todos: todos})
}

func (h *handlers) indexPost(c echo.Context) error {
	var todo Todo

	errs := echo.FormFieldBinder(c).
		Int64("id", &todo.ID).
		String("content", &todo.Content).
		Bool("done", &todo.Done).
		CustomFunc("until", customFunc(&todo)).
		BindErrors()

	if errs != nil {
		h.e.Logger.Error(errs)
		return c.Render(http.StatusBadRequest, "index", Data{Errors: errs})
	} else if todo.ID == 0 {
		// IDが0の時は登録
		ctx := context.Background()
		if todo.Content == "" {
			return renderBadRequest(c, "Content is empty")
		} else {
			_, err := h.db.NewInsert().Model(&todo).Exec(ctx)
			if err != nil {
				h.e.Logger.Error(err)
				return renderBadRequest(c, err)
			}
		}
	} else {
		ctx := context.Background()
		if c.FormValue("delete") != "" {
			// 削除
			_, err := h.db.NewDelete().Model(&todo).Where("id = ?", todo.ID).Exec(ctx)
			if err != nil {
				h.e.Logger.Error(err)
				return renderBadRequest(c, err)
			}
		} else {
			// 更新
			var orig Todo
			err := h.db.NewSelect().Model(&orig).Where("id = ?", todo.ID).Scan(ctx)
			if err != nil {
				h.e.Logger.Error(err)
				return renderBadRequest(c, err)
			}

			orig.Done = todo.Done
			_, err = h.db.NewUpdate().Model(&orig).Where("id = ?", todo.ID).Exec(ctx)
			if err != nil {
				h.e.Logger.Error(err)
				return renderBadRequest(c, err)
			}
		}
	}

	return c.Redirect(http.StatusFound, "/")
}

func customFunc(todo *Todo) func([]string) []error {
	return func(values []string) []error {
		if len(values) == 0 || values[0] == "" {
			return nil
		}
		dt, err := time.Parse("2006-01-02T15:04 MST", values[0]+" JST")
		if err != nil {
			return []error{echo.NewBindingError("until", values[0:1], "failed to decode time", err)}
		}
		todo.Until = dt
		return nil
	}
}

func renderBadRequest(c echo.Context, errs any) error {
	var retErrors []error

	switch e := errs.(type) {
	case []error:
		retErrors = e
	case error:
		retErrors = []error{e}
	case string:
		retErrors = []error{errors.New(e)}
	default:
		return fmt.Errorf("Unsupported type: %T", errs)
	}

	return c.Render(http.StatusBadRequest, "index", Data{Errors: retErrors})
}
