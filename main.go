package main

import (
	"App/controllers"
	"fmt"
	"plugin"
	"io/fs"
   	"path/filepath"

	"github.com/gofiber/fiber/v2"
)


func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil { return e }
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func main()  {
	fmt.Println("Main APP says hello")
	// Discover and load plugins
	pluginList := find("./plugins", ".so")
	app := fiber.New()

	// Load the data
	for _, pluginPath := range pluginList {
		plugin, err:= plugin.Open(pluginPath)
		if err != nil {
			fmt.Println("Error loading plugin: ", err)
		}
		GetFunctions , err := plugin.Lookup("GetFunctions")
		if err != nil {
			fmt.Println("Error loading plugin: ", err)
		}
		function := GetFunctions.(func() ([]string, []func(c *fiber.Ctx) error))
		paths,functions := function()

		for index, path := range paths{
			app.Get(path, functions[index])
		}
	}

	app.Get("/",controllers.ScienceController)
	app.Listen(":3000")
}