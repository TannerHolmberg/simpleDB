package main

type Entry struct{ Key, Value string }

type Index struct{ items []Entry }

// constructor for Index to initialize the items slice with a reasonable capacity to avoid too many reallocations.
// returns a pointer to the newly created Index.
func NewIndex() *Index {
	return &Index{items: make([]Entry, 0, 128)}
}

// Get does a linear scan: O(n) and returns the value associated with the key if found,
// along with a boolean indicating whether the key was found.
func (idx *Index) Get(key string) (string, bool) {
	for i := 0; i < len(idx.items); i++ {
		if idx.items[i].Key == key {
			return idx.items[i].Value, true
		}
	}
	return "", false
}

// still performs a linear scan to find the key, but updates the value if found. If the key is not found,
// it appends a new Entry to the items slice. This method also has O(n) time complexity due to the linear scan.
func (idx *Index) Set(key, value string) {
	for i := 0; i < len(idx.items); i++ {
		if idx.items[i].Key == key {
			idx.items[i].Value = value
			return
		}
	}
	idx.items = append(idx.items, Entry{Key: key, Value: value})
}
