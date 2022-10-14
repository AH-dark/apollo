package bootstrap

func Init() {
	for _, dependency := range Dependencies {
		dependency.Function()
	}
}
