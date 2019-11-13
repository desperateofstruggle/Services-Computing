package server

import "github.com/go-martini/martini"

// NewServer - new server
func NewServer(port string) {
	// instantiation
	m := martini.Classic()

	// set methods to handle the request to the /helloworld/name
	m.Get("/helloworld/:name", func(params martini.Params) string {
		return "hello world " + params["name"] + "!\n"
	})

	// set methods to handle the request to the /helloNum/num
	m.Get("/helloNum/:num", func(params martini.Params) string {
		return "hello NO. " + params["num"] + "...\n"
	})

	// set methods to handle the request to the /nothing
	m.Get("/:nothing", func(params martini.Params) string {
		return "nothing hanppen: " + params["nothing"] + "?\n"
	})

	// set port correspondence to the main func
	m.RunOnAddr(":" + port)
}
