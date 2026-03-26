package todo

import (
	"fmt"
	"time"
)

// Todo represents a single todo item.
type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Done       bool      `json:"done"`
	CreatedAt  time.Time `json:"created_at"`
	AssignedTo string    `json:"assigned_to,omitempty"`
}

// List holds a collection of todos and provides mutation operations.
type List struct {
	items []*Todo
}

// NewList constructs a List from an existing slice of todos, typically
// loaded from persistent storage.
func NewList(items []*Todo) *List {
	return &List{items: items}
}

// Items returns the underlying slice so the store can persist it.
func (l *List) Items() []*Todo {
	return l.items
}

// Add creates a new todo with an auto-incremented ID and appends it to the list.
// It returns the newly created todo.
func (l *List) Add(title string) *Todo {
	nextID := 1
	for _, t := range l.items {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
	todo := &Todo{
		ID:        nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	l.items = append(l.items, todo)
	return todo
}

// ListAll returns all todos in insertion order.
func (l *List) ListAll() []*Todo {
	return l.items
}

// MarkDone sets Done = true on the todo with the given ID.
// It returns an error if no todo with that ID exists.
func (l *List) MarkDone(id int) error {
	for _, t := range l.items {
		if t.ID == id {
			t.Done = true
			return nil
		}
	}
	fmt.Println("mark pr 1")
	return fmt.Errorf("todo #%d not found", id)
}

// Delete removes the todo with the given ID from the list.
// It returns an error if no todo with that ID exists.
func (l *List) Delete(id int) error {
	for i, t := range l.items {
		if t.ID == id {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo #%d not found", id)
}

func (l *List) Assign(id int, assignee string) error {
	for _, t := range l.items {
		if t.ID == id {
			t.AssignedTo = assignee
			return nil
		}
	}
	return fmt.Errorf("todo #%d not found", id)
}
