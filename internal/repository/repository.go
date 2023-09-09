package repository

import (
	"context"
	"fmt"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/internal/db"
	"os"
	"time"
)

type CrudRepository interface {
	TodoRpInit() *TodoRp
	Close()
	All() []*ent.Todo
	Create(ctx context.Context, todo ent.Todo)
	Update(ctx context.Context, id int, todo ent.Todo)
}

type TodoRp struct {
	Client *ent.Client
}

func NewTodoRp() *TodoRp {
	client, err := db.Client()
	if err != nil {
		fmt.Println("failed to open driver: %v", err)
		os.Exit(1)
	}

	return &TodoRp{
		Client: client,
	}
}

func (tp *TodoRp) Close() {
	tp.Client.Close()
}

func (tp *TodoRp) All(ctx context.Context) []*ent.Todo {
	res, err := tp.Client.Todo.Query().All(ctx)
	if err != nil {
		fmt.Println("failed to query todos: %v", err)
		os.Exit(1)
	}

	return res
}

func (tp *TodoRp) Create(ctx context.Context, todo ent.Todo) {
	todoCreate := tp.Client.Todo.Create()
	todoCreate.SetTitle(todo.Title)
	todoCreate.SetDescription(todo.Description)
	todoCreate.SaveX(ctx)
}

func (tp *TodoRp) Update(ctx context.Context, id int, todo ent.Todo) {
	// get the todo by id
	todoUpdate := tp.Client.Todo.UpdateOneID(id)

	todoUpdate.SetTitle(todo.Title)
	todoUpdate.SetDescription(todo.Description)

	// execute the query
	todoUpdate.SetUpdatedAt(time.Now())
	if err := todoUpdate.Exec(ctx); err != nil {
		fmt.Println("failed to update todo: %v", err)
		os.Exit(1)
	}
}
