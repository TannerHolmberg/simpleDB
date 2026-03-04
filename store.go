package main

import (
	"bufio"
	"fmt"
	"os"
)

type Store struct {
	path  string
	file  *os.File
	index *Index
}

func openStore(path string) (*Store, error) {
	//create a new store with the given path and initialize the index
	s := &Store{
		path:  path,
		index: NewIndex(),
	}

	// Ensure file exists; open read-only first to replay.
	// If file doesn't exist, create it.
	// ignore metadata file if it exists, as we don't need it for this simple implementation.
	if _, err := os.Stat(path); err != nil {
		// if the error is because the file does not exist, create it
		if os.IsNotExist(err) {
			f, cerr := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
			if cerr != nil {
				return nil, cerr
			}
			// close the file immediately since we just wanted to create it
			_ = f.Close()
		} else {
			return nil, err
		}
	}

	// Rebuild the in-memory index by replaying the log file.
	if err := s.replay(); err != nil {
		return nil, err
	}

	// Open the file for appending new entries.
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Store the file handle for appending new entries.
	s.file = f

	return s, nil
}

func (s *Store) replay() error {
	f, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for sc.Scan() {
		line := sc.Text()
		cmd, key, value, ok := parseInput(line)
		if !ok {
			// Ignore malformed lines in the log.
			continue
		}
		if cmd == "SET" {
			// last write wins
			s.index.Set(key, value)
		}
		// GET/EXIT lines should not appear in log; ignore if they do.
	}

	return sc.Err()
}

// Close the store by closing the file handle if it's open.
func (s *Store) Close() error {
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}

func (s *Store) Set(key, value string) error {
	// Append-only write: one line per SET
	// Format: SET <key> <value>\n
	_, err := fmt.Fprintf(s.file, "SET %s %s\n", key, value)
	if err != nil {
		return err
	}

	// Persist immediately
	if err := s.file.Sync(); err != nil {
		return err
	}

	// Update in-memory index
	s.index.Set(key, value)
	return nil
}

func (s *Store) Get(key string) (string, bool) {
	return s.index.Get(key)
}
