package utils

import (
	"fmt"
	"path"
	"strings"
)

type Node struct {
	Name        string
	Icon		string
	IsDir       bool
	Description string
	Children    map[string]*Node
}

type FileSystem struct {
	Root        *Node
	Currentpath string
}

func NewFileSystem() *FileSystem {
	root := &Node{
		Name:     "abhinav",
		IsDir:    true,
		Children: make(map[string]*Node),
	}

	// Languages
	root.Children["languages"] = &Node{
		Name:        "languages",
		Icon: 		 "</>",
		IsDir:       true,
		Children:    make(map[string]*Node),
	}

	root.Children["languages"].Children["Golang"] = &Node{
		Name:     "Go",
		Icon: 		"",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["languages"].Children["Javascript"] = &Node{
		Name:     "Javascript",
		Icon: 		"",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["languages"].Children["Bash"] = &Node{
		Name:     "Bash",
		Icon: 		"",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["languages"].Children["Python"] = &Node{
		Name:     "Python",
		Icon: 		"",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["languages"].Children["Yaml"] = &Node{
		Name:     "Yaml",
		Icon: 		"",
		IsDir:    false,
		Children: make(map[string]*Node),
	}


	// Tools
	root.Children["tools"] = &Node{
		Name:     "tools",
		Icon: 		"🛠",
		IsDir:    true,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["version control"] = &Node{
		Name:     "version control",
		Icon: 		"",
		Description: "Git",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["containerization"] = &Node{
		Name:     "containerization",
		Description: "Docker",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["CI/Cd"] = &Node{
		Name:     "CI/CD",
		Description: "Github Actions, Jenkins",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["IaC"] = &Node{
		Name:     "IaC",
		Description: "Terraform",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["conatiner orchestration"] = &Node{
		Name:     "conatiner orchestration",
		Description: "Kubernetes",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["tools"].Children["networking"] = &Node{
		Name:     "networking",
		Description: "Https/Http, Tcp/Ip, SSH, DNS, VPN, Firewall",
		IsDir:    false,
		Children: make(map[string]*Node),
	}


	// Socials
	root.Children["socials"] = &Node{
		Name:     "socials",
		Icon: 		"🌐",
		IsDir:    true,
		Children: make(map[string]*Node),
	}

	root.Children["socials"].Children["linkedin"] = &Node{
		Name:     "linkedin",
		Icon: 		"",
		Description: "linkedin.com/in/08abhinav",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["socials"].Children["github"] = &Node{
		Name:     "github",
		Description: "github.com/08abhinav",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	root.Children["socials"].Children["medium"] = &Node{
		Name:     "medium",
		Description: "medium.com/@abhinavnegi101",
		IsDir:    false,
		Children: make(map[string]*Node),
	}

	// Cloud
	root.Children["cloud"] = &Node{
		Name:     "cloud",
		Icon: 	"☁",
		IsDir:    true,
		Children: make(map[string]*Node),
	}

	root.Children["cloud"].Children["AWS"] = &Node{
		Name:     "AWS",
		IsDir:    true,
		Description: "VPC, IAM, EC2, ALB, ECS, ECR, EKS",
		Children: make(map[string]*Node),
	}

	return &FileSystem{
		Root:        root,
		Currentpath: "/root/abhinav",
	}
}

func (fs *FileSystem) ChangeDirectory(target string) error {
	if strings.TrimSpace(target) == "" {
		return fmt.Errorf("cd: missing directory")
	}

	target = strings.TrimSpace(target)

	// Home directory
	if target == "~" || target == "/" {
		fs.Currentpath = "/root/abhinav"
		return nil
	}

	newPath := path.Clean(path.Join(fs.Currentpath, target))

	if !isInsideRoot(newPath) {
		return fmt.Errorf("cd: %s: outside filesystem", target)
	}

	_, err := fs.ResolvePath(newPath)
	if err != nil {
		return fmt.Errorf("cd: %s: no such directory", target)
	}

	fs.Currentpath = newPath
	return nil
}

func isInsideRoot(p string) bool {
	root := "/root/abhinav"

	return p == root || strings.HasPrefix(p, root+"/")
}

func (fs *FileSystem) ResolvePath(p string) (*Node, error) {
	p = path.Clean(p)

	virtualRoot := "/root/abhinav"

	// The filesystem root directly maps to fs.Root
	if p == virtualRoot {
		return fs.Root, nil
	}

	// Every valid path must be inside /root/abhinav
	if !isInsideRoot(p) {
		return nil, fmt.Errorf("path does not exist: %s", p)
	}

	// Remove /root/abhinav from the path.
	relativePath := strings.TrimPrefix(p, virtualRoot)

	// Example:
	//
	// /root/abhinav/languages/Golang
	//
	// becomes:
	//
	// /languages/Golang

	parts := strings.Split(strings.Trim(relativePath, "/"), "/")

	current := fs.Root

	for _, part := range parts {
		if part == "" {
			continue
		}

		next, exists := current.Children[part]

		if !exists {
			return nil, fmt.Errorf("path does not exist: %s", p)
		}

		current = next
	}

	return current, nil
}

func (fs *FileSystem) ListHelper(target string) {
	var dir *Node
	var err error

	if strings.TrimSpace(target) == "" {
		target = fs.Currentpath
	} else {
		target = strings.TrimSpace(target)

		if strings.HasPrefix(target, "/") {
			target = path.Clean(target)
		} else {
			target = path.Join(fs.Currentpath, target)
		}
	}

	dir, err = fs.ResolvePath(target)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !dir.IsDir {
		fmt.Println(dir.Name)
		return
	}

	for _, child := range dir.Children {
		if len(child.Description) != 0{
			fmt.Printf("%s %s -> %s \n", child.Icon, child.Name, child.Description)
		}else{
			fmt.Printf("%s %s\n", child.Icon, child.Name)
		}
	}
}