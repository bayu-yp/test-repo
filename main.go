package main

import (
	"fmt"
	"os"
	"strconv"

	"test-repo/todo"
)

// chain-3: dummy change
func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	store := todo.NewStore("")

	items, err := store.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading todos: %v\n", err)
		os.Exit(1)
	}

	list := todo.NewList(items)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo add <title>")
			os.Exit(1)
		}
		title := os.Args[2]
		t := list.Add(title)
		if err := store.Save(list.Items()); err != nil {
			fmt.Fprintf(os.Stderr, "error saving todos: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added todo #%d: %s\n", t.ID, t.Title)

	case "list":
		todos := list.ListAll()
		if len(todos) == 0 {
			fmt.Println("No todos yet. Add one with: todo add <title>")
			return
		}
		for _, t := range todos {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("[%s] #%d  %s\n", mark, t.ID, t.Title)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid id %q: must be an integer\n", os.Args[2])
			os.Exit(1)
		}
		if err := list.MarkDone(id); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := store.Save(list.Items()); err != nil {
			fmt.Fprintf(os.Stderr, "error saving todos: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Marked #%d as done.\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: todo delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid id %q: must be an integer\n", os.Args[2])
			os.Exit(1)
		}
		if err := list.Delete(id); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if err := store.Save(list.Items()); err != nil {
			fmt.Fprintf(os.Stderr, "error saving todos: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted todo #%d.\n", id)

	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  todo add <title>   — add a new todo")
	fmt.Fprintln(os.Stderr, "  todo list          — list all todos")
	fmt.Fprintln(os.Stderr, "  todo done <id>     — mark a todo as done")
	fmt.Fprintln(os.Stderr, "  todo delete <id>   — delete a todo")
}
