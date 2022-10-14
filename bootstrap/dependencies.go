package bootstrap

type Dependency struct {
	// Name is the name of the dependency.
	Name string
	// Function is the function that will be called to initialize the dependency.
	Function func()
}

// Dependencies is a list of dependencies that will be initialized in order.
var Dependencies = []Dependency{}
