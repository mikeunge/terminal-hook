package store

import (
	"fmt"
	"strings"

	fileHelper "github.com/mikeunge/go/pkg/file-helper"
	pathHelper "github.com/mikeunge/go/pkg/path-helper"
)

type Store struct {
	Path string
}

func New(path string) (Store, error) {
	store := Store{
		Path: path,
	}
	return store, nil
}

func (s *Store) GetPath(key string) (string, error) {
	var path string

	if path, err := find(key, s.Path); err != nil {
		return path, err
	}
	return path, nil
}

func (s *Store) WritePath(key string, path string) error {
	key = replaceInvalidChars(key)
	path = replaceInvalidChars(path)

	if !pathHelper.FileExists(s.Path) {
		if err := fileHelper.WriteFile(s.Path, fmt.Sprintf("%s:%s", key, path), 0640); err != nil {
			return err
		}
		return nil
	}

	store, err := loadStore(s.Path)
	if err != nil {
		return err
	}

	// if we have 0 entries, we want to directly write to the store, no splitting needed
	if len(store) == 0 {
		if err := fileHelper.WriteFile(s.Path, fmt.Sprintf("%s:%s", key, path), 0640); err != nil {
			return err
		}
		return nil
	}

	if _, err := find(key, store); err == nil {
		return fmt.Errorf("hook with key '%s' already exists", key)
	}

	data := strings.Split(store, ";")
	data = append(data, fmt.Sprintf("%s:%s", key, path))

	if err := fileHelper.WriteFile(s.Path, strings.Join(data, ";"), 0640); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeletePath(key string) error {
	key = reverseInvalidChars(key)
	store, err := loadStore(s.Path)
	if err != nil {
		return err
	}

	if _, err := find(key, store); err != nil {
		return err
	}

	var newStore []string
	data := strings.Split(store, ";")
	for i := 0; i < len(data); i++ {
		hook := strings.Split(data[i], ":")
		vKey := reverseInvalidChars(hook[0])
		if vKey != key {
			newStore = append(newStore, fmt.Sprintf("%s:%s", hook[0], hook[1]))
		}
	}

	if err := fileHelper.WriteFile(s.Path, strings.Join(newStore, ";"), 0640); err != nil {
		return err
	}
	return nil
}

func find(key string, store string) (string, error) {
	data := strings.Split(store, ";")
	for i := 0; i < len(data); i++ {
		hook := strings.Split(data[i], ":")
		vKey := reverseInvalidChars(hook[0])
		if vKey == key {
			return hook[i], nil
		}
	}
	return "", fmt.Errorf("key '%s' not found in store", key)
}

func loadStore(path string) (string, error) {
	var store []byte

	store, err := fileHelper.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(store), nil
}

// We need to get rid of some special characters we use, mainly ; and :
func replaceInvalidChars(input string) string {
	replacer := strings.NewReplacer(";", "-*-", ":", "-+-")
	return replacer.Replace(input)
}

// Reverse the special/invalid characters
func reverseInvalidChars(input string) string {
	replacer := strings.NewReplacer("-*-", ";", "-+-", ":")
	return replacer.Replace(input)
}
