package utils

import (
	"fmt"
	"path"
	"strings"
)

type Node struct{
	Name 		string
	IsDir 		bool
	Children 	map[string]*Node
}

type FileSystem struct{
	Root 		*Node
	Currentpath string
}

func NewFileSystem() *FileSystem{
	root := &Node{
		Name: "root",
		IsDir: true,
		Children: make(map[string]*Node),
	}

	root.Children["languages"] = &Node{
		Name: "languages",
		IsDir: true,
		Children: make(map[string]*Node),
	}
	
	root.Children["tools"] = &Node{
		Name: "tools",
		IsDir: true,
		Children: make(map[string]*Node),
	}

	root.Children["socials"] = &Node{
		Name: "socials",
		IsDir: true,
		Children: make(map[string]*Node),
	}

	root.Children["projects"] = &Node{
		Name: "projects",
		IsDir: true,
		Children: make(map[string]*Node),
	}

	return &FileSystem{
		Root: root,
		Currentpath: "/root",
	}
}

func (fs *FileSystem) ChangeDirectory(target string) error{
	if target == ""{
		return fmt.Errorf("cd: missing directory")
	}

	if target == "~" || target == "/" || (target == ".." && fs.Currentpath == "/root"){
		fs.Currentpath = "/root"
		return nil
	}

	newPath := path.Join(fs.Currentpath, target)

	if newPath != "/root" && !isInsideroot(newPath){
		fs.Currentpath = "/root"
		return nil
	}

	fs.Currentpath = newPath
	return nil
}

func isInsideroot(p string) bool{
	return p == "/root" || len(p) > len("/root") && p[:len("/root")+1] == "/root/"
}