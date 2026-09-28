package utils

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

const (
	ColorReset   = "\033[0m"
	ColorBold    = "\033[1m"
	ColorDim     = "\033[2m"

	ColorCyan    = "\033[38;2;80;250;123m" 
	ColorBlue    = "\033[38;2;139;233;253m"
	ColorPurple  = "\033[38;2;189;147;249m"
	ColorOrange  = "\033[38;2;255;184;108m"
	ColorRed     = "\033[38;2;255;85;85m"  
	ColorComment = "\033[38;2;98;114;164m" 
)

const (
	// Directories
	IconFolder     = "\ue5fe" 
	IconFolderOpen = "\ue5ff"
	IconRoot       = "\uf015"

	// Categories
	IconCode      = "\uf121" 
	IconTools     = "\uf0ad"
	IconCloud     = "\uf0c2" 
	IconSocial    = "\uf0c0" 

	// Languages & Formats
	IconGo         = "\ue627" 
	IconJS         = "\ue74e" 
	IconBash       = "\uf489"
	IconPython     = "\ue73c" 
	IconYaml       = "\ue6a8" 

	// DevOps & Tools
	IconGit        = "\uf1d3" 
	IconDocker     = "\uf308" 
	IconCI         = "\uf427"
	IconTerraform  = "\ue69a" 
	IconK8s        = "\ufd31" 
	IconNetwork    = "\uf6ff" 

	// Cloud Services
	IconAWS        = "\ue7ad" 

	// Social & Links
	IconLinkedIn   = "\uf08c" 
	IconGitHub     = "\uf09b" 
	IconMedium     = "\uf044" 
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
		Icon:     ColorPurple + IconRoot + ColorReset,
		Children: make(map[string]*Node),
	}

	addFolder := func(parent *Node, name, icon string) *Node {
		node := &Node{
			Name:     name,
			Icon:     icon,
			IsDir:    true,
			Children: make(map[string]*Node),
		}
		parent.Children[name] = node
		return node
	}

	addFile := func(parent *Node, name, icon, desc string) {
		parent.Children[name] = &Node{
			Name:        name,
			Icon:        icon,
			IsDir:       false,
			Description: desc,
			Children:    make(map[string]*Node),
		}
	}

	langs := addFolder(root, "languages", ColorBlue+IconCode+ColorReset)
	addFile(langs, "Golang", ColorCyan+IconGo+ColorReset, "Microservices, CLI tools, Concurrent systems")
	addFile(langs, "Javascript", ColorOrange+IconJS+ColorReset, "Frontend scripting, Node.js tooling")
	addFile(langs, "Bash", ColorComment+IconBash+ColorReset, "Automation scripts, Linux system administration")
	addFile(langs, "Python", ColorBlue+IconPython+ColorReset, "Data processing, scripting, quick prototyping")
	addFile(langs, "Yaml", ColorRed+IconYaml+ColorReset, "Kubernetes manifests, Docker Compose, CI/CD pipelines")

	tools := addFolder(root, "tools", ColorOrange+IconTools+ColorReset)
	addFile(tools, "version-control", ColorRed+IconGit+ColorReset, "Git workflows, branching strategies, rebase")
	addFile(tools, "containerization", ColorBlue+IconDocker+ColorReset, "Docker, multi-stage builds, image optimization")
	addFile(tools, "ci-cd", ColorCyan+IconCI+ColorReset, "GitHub Actions, Jenkins pipelines")
	addFile(tools, "iac", ColorPurple+IconTerraform+ColorReset, "Terraform - Infrastructure as Code")
	addFile(tools, "orchestration", ColorBlue+IconK8s+ColorReset, "Kubernetes - Deployments, Services, Helm")
	addFile(tools, "networking", ColorOrange+IconNetwork+ColorReset, "HTTP/S, TCP/IP, SSH, DNS, VPN, Firewalls")

	cloud := addFolder(root, "cloud", ColorOrange+IconCloud+ColorReset)
	aws := addFolder(cloud, "AWS", ColorOrange+IconAWS+ColorReset)
	addFile(aws, "compute", ColorOrange+IconTools+ColorReset, "EC2, ECS, ECR, EKS")
	addFile(aws, "networking", ColorBlue+IconNetwork+ColorReset, "VPC, Route53, ALB, Gateways")


	socials := addFolder(root, "socials", ColorPurple+IconSocial+ColorReset)
	addFile(socials, "linkedin", ColorBlue+IconLinkedIn+ColorReset, "linkedin.com/in/08abhinav")
	addFile(socials, "github", ColorComment+IconGitHub+ColorReset, "github.com/08abhinav")
	addFile(socials, "medium", ColorCyan+IconMedium+ColorReset, "medium.com/@abhinavnegi101")

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

	if target == "~" || target == "/" {
		fs.Currentpath = "/root/abhinav"
		return nil
	}

	newPath := path.Clean(path.Join(fs.Currentpath, target))

	if !isInsideRoot(newPath) {
		return fmt.Errorf("cd: %s: outside filesystem", target)
	}

	node, err := fs.ResolvePath(newPath)
	if err != nil {
		return fmt.Errorf("cd: %s: no such directory", target)
	}

	if !node.IsDir {
		return fmt.Errorf("cd: %s: Not a directory", target)
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
	var targetPath string
	target = strings.TrimSpace(target)

	if target == "" {
		targetPath = fs.Currentpath
	} else if strings.HasPrefix(target, "/") {
		targetPath = path.Clean(target)
	} else {
		targetPath = path.Join(fs.Currentpath, target)
	}

	dir, err := fs.ResolvePath(targetPath)
	if err != nil {
		fmt.Printf("%sls: cannot access '%s': No such file or directory%s\n", ColorRed, target, ColorReset)
		return
	}

	if !dir.IsDir {
		fmt.Printf("%s  %s %s->%s %s\n", dir.Icon, dir.Name, ColorComment, ColorReset, dir.Description)
		return
	}

	keys := make([]string, 0, len(dir.Children))
	for k := range dir.Children {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		child := dir.Children[k]
		if child.IsDir {
			fmt.Printf("  %s  %s%s%s/\n", child.Icon, ColorBold+ColorBlue, child.Name, ColorReset)
		} else {
			if len(child.Description) > 0 {
				fmt.Printf("  %s  %-20s %s%s%s\n", child.Icon, child.Name, ColorComment, child.Description, ColorReset)
			} else {
				fmt.Printf("  %s  %s\n", child.Icon, child.Name)
			}
		}
	}
}

func (fs *FileSystem) TreeHelper(target string) {
	var targetPath string
	target = strings.TrimSpace(target)

	if target == "" {
		targetPath = fs.Currentpath
	} else if strings.HasPrefix(target, "/") {
		targetPath = path.Clean(target)
	} else {
		targetPath = path.Join(fs.Currentpath, target)
	}

	node, err := fs.ResolvePath(targetPath)
	if err != nil {
		fmt.Printf("%stree: cannot access '%s': No such file or directory%s\n", ColorRed, target, ColorReset)
		return
	}

	if node.IsDir {
		fmt.Printf("%s  %s%s%s/\n", node.Icon, ColorBold+ColorBlue, node.Name, ColorReset)
	} else {
		fmt.Printf("%s  %s\n", node.Icon, node.Name)
		return
	}

	var printTree func(n *Node, prefix string)
	printTree = func(n *Node, prefix string) {
		keys := make([]string, 0, len(n.Children))
		for k := range n.Children {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for i, k := range keys {
			child := n.Children[k]
			isLast := i == len(keys)-1

			connector := "├── "
			childPrefix := "│   "
			if isLast {
				connector = "└── "
				childPrefix = "    "
			}

			if child.IsDir {
				fmt.Printf("%s%s%s  %s%s%s/\n", prefix, connector, child.Icon, ColorBold+ColorBlue, child.Name, ColorReset)
				printTree(child, prefix+childPrefix)
			} else {
				if len(child.Description) > 0 {
					fmt.Printf("%s%s%s  %-20s %s%s%s\n", prefix, connector, child.Icon, child.Name, ColorComment, child.Description, ColorReset)
				} else {
					fmt.Printf("%s%s%s  %s\n", prefix, connector, child.Icon, child.Name)
				}
			}
		}
	}

	printTree(node, "")
}