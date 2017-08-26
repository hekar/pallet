package controllers

import (
	"github.com/spf13/viper"
	"fmt"
	"net/url"
	"io/ioutil"
	"path/filepath"
	"os"
	"strings"
	"github.com/revel/revel"
)

type Application struct {
	*revel.Controller
}

type FileLink struct {
	Self string
	Name string
	FullName string
}

func (c Application) Index() revel.Result {
	root := strings.Replace(viper.GetString("root"), "$HOME", os.Getenv("HOME"), 1)
	path := filepath.Join(root, c.Params.Route.Get("path"))
	
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var links []FileLink
	for _, file := range files {
		name := file.Name()
		fullName := filepath.Join(path, name)
		href := "/" + url.QueryEscape(name)
		link := FileLink{
			Name: file.Name(),
			FullName: fullName,
			Self: href,
		}
		links = append(links, link)
	}
	fmt.Printf("files: %v\n", links)
	c.ViewArgs["links"] = links

	return c.Render()
}
