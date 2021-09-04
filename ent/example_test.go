package todo

import (
	"context"
	"fmt"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yuchanns/gobyexample/ent/ent"
	"github.com/yuchanns/gobyexample/ent/ent/todo"
)

func Test_Todo(t *testing.T) {
	// open database
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed to open sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// migrate schema
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed to create schema resources: %+v", err)
	}
	// create records
	task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed to creating a todo: %v", err)
	}
	task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
	if err != nil {
		log.Fatalf("faield to create a todo: %v", err)
	}
	// connecting data
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("failed to connecting todo2 to its parent: %v", err)
	}

	// query
	// query all
	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed to querying todos: %v", err)
	}
	fmt.Println("query all:")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// query with condition
	items, err = client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed to querying todos: %v", err)
	}
	fmt.Println("query with condition:")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// query without condition
	items, err = client.Todo.Query().Where(todo.Not(todo.HasParent()), todo.HasChildren()).All(ctx)
	if err != nil {
		log.Fatalf("failed to querying todos: %v", err)
	}
	fmt.Println("query without condition:")
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// query parent through its children
	parent, err := client.Todo.Query().Where(todo.HasParent()).QueryParent().Only(ctx)
	if err != nil {
		log.Fatalf("failed to querying todos: %v", err)
	}
    fmt.Println("query parent throught its children")
	fmt.Printf("%d: %q\n", parent.ID, parent.Text)
}
