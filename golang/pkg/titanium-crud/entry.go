package titaniumcrud

// Entry definition
type Entry struct {
	fields map[string]string
}

// NewEntry create new entry
func NewEntry() *Entry {
	return &Entry{fields: make(map[string]string)}
}

// GetFields get fields
func (e *Entry) GetFields() map[string]string {
	return e.fields
}

// SetFields set fields
func (e *Entry) SetFields(newFields map[string]string) {
	e.fields = newFields
}

// Put put one field
func (e *Entry) Put(key string, value string) {
	e.fields[key] = value
}

// Get get one field
func (e *Entry) Get(key string) string {
	return e.fields[key]
}
