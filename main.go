package main

import "github.com/oktopriima/dimy/config"

func main() {
	i := NewInject()

	// run the seeders
	if err := i.Invoke(func(seed *config.Seeders) {
		seed.RunSeeders()
	}); err != nil {
		panic(err)
	}

	// run the webserver
	if err := i.Invoke(func(instance *config.ServerInstance) {
		instance.RunWithGracefullyShutdown()
	}); err != nil {
		panic(err)
	}
}
