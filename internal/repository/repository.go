package repository

import (
	"context"
	"fmt"
	"github.com/yugo-ibuki/musdok/ent"
	"github.com/yugo-ibuki/musdok/internal/db"
	"os"
)

type ITodoRp interface {
	Close()
	All() []*ent.Todo
	Create(todo ent.Todo)
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

func (tp *TodoRp) All(ctx context.Context) {
	tp.Client.Todo.Query().All(ctx)
}

func (tp *TodoRp) Create(ctx context.Context, todo ent.Todo) {
	todoCreate := tp.Client.Todo.Create()
	todoCreate.SetTitle(todo.Title)
	todoCreate.SetDescription(todo.Description)
	todoCreate.SaveX(ctx)
}
