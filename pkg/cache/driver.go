package cache

type Driver interface {
	// Get returns the value associated with the given key.
	Get(key string) (interface{}, bool)
	// Gets returns the values associated with the given keys.
	Gets(keys []string, prefix string) (map[string]interface{}, []string)
	// Set sets the value associated with the given key.
	Set(key string, value interface{}, ttl int) error
	// Sets sets the value associated with the given keys.
	Sets(values map[string]interface{}, prefix string) error
	// Delete deletes the value associated with the given key.
	Delete(key string) error
	// Deletes deletes the value associated with the given keys.
	Deletes(keys []string, prefix string) error
}
