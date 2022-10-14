package model

// Init initializes the model package.
func Init(force bool) {
	client := BuildClient(force)
	Global = client
}
