package todo

import (
	"encoding/json"
	"errors"
	"os"
)

const defaultStorePath = "todos.json"

// Store handles loading and saving todos to a JSON file.
type Store struct {
	path string
}

// NewStore creates a Store that persists to the given file path.
// Pass an empty string to use the default path ("todos.json").
func NewStore(path string) *Store {
	if path == "" {
		path = defaultStorePath
	}
	return &Store{path: path}
}

// feat: optimize store loading
// Load reads todos from the JSON file. If the file does not exist it returns
// an empty slice — the file will be created on the first Save.
func (s *Store) Load() ([]*Todo, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []*Todo{}, nil
		}
		return nil, err
	}

	var items []*Todo
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// Save marshals the provided todos and overwrites the JSON file.
func (s *Store) Save(items []*Todo) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}

func (s *Store) Add(title string) (*Todo, error) {
	items, err := s.Load()
	if err != nil {
		return nil, err
	}

	list := NewList(items)
	todo := list.Add(title)

	if err := s.Save(list.Items()); err != nil {
		return nil, err
	}
	return todo, nil
}
