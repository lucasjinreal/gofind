package main

import (
	"gitlab.com/jinfagang/colorgo"
	"time"
	"os"
	"fmt"
	flag "github.com/ogier/pflag"
	"strings"
	"os/user"
	"path/filepath"
)

const (
	DIR_MODE = iota
	FILE_MODE
	CONTENT_MODE

	DIR_FILE_MODE
	DIR_CONTENT_MODE
	FILE_CONTENT_MODE

	ALL_MODE
)

// Set Global Variable
var usr, _ = user.Current()
var goFindConfigDir = filepath.Join(usr.HomeDir, ".gofind")


func showInfo() {
	t := time.Now()

	day := t.Day()


	// if day is 18 then show info
	if !Exists(goFindConfigDir) || day == 18 {
		cg.PrintlnBlue("gofind - find anything you try to find.\n")
		cg.PrintlnGreen(`   __________  ___________   ______
  / ____/ __ \/ ____/  _/ | / / __ \
 / / __/ / / / /_   / //  |/ / / / /
/ /_/ / /_/ / __/ _/ // /|  / /_/ /
\____/\____/_/   /___/_/ |_/_____/
                                    `)
		cg.PrintlnCyan("Author - Jin Tian. @WeChat: jintianiloveu")
	}
}

func main() {

	showInfo()
	if !Exists(goFindConfigDir) {
		err := os.Mkdir(goFindConfigDir, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Print("\033[1;33m")
	fmt.Print("gofind was originally written by Jin Tian, welcome fork and star!")
	fmt.Println("\033[0m")


	var searchDir bool
	var searchFile bool
	var searchContent bool

	flag.BoolVarP(&searchDir, "dir", "d",false, "search dir names.")
	flag.BoolVarP(&searchFile,"file", "f", false, "search file names.")
	flag.BoolVarP(&searchContent, "content", "c", false, "search all file content.")
	flag.Parse()

	// posArg are unresolved args
	posArgs := flag.Args()

	keyWords := ""
	rootPath := "./"
	if len(posArgs) == 0 {
		fmt.Print("What do u trying to find: ")
		fmt.Scanf("%s", &keyWords)
		cg.PrintlnYellow("Start searching " + keyWords)
	} else {
		if len(posArgs) == 1 {
			// indicates only input 1 posArg, treat it as keywords
			if IsDir(posArgs[0]) {
				rootPath = posArgs[0]
				fmt.Print("What do u trying to find: ")
				fmt.Scanf("%s", &keyWords)
				cg.PrintlnYellow("Start searching " + keyWords)
			} else {
				keyWords = posArgs[0]
			}
		} else {
			rootPath = posArgs[0]
			keyWords = strings.Join(posArgs[1:], " ")
		}
	}

	if !Exists(rootPath) {
		cg.PrintlnRed("ops! seems your search path is not exist.")
		os.Exit(1)
	} else {
		// default mode is all mode
		mode := ALL_MODE

		if !searchDir && !searchFile && searchContent {
			// indicates only search content
			mode = CONTENT_MODE
		} else if !searchDir && !searchContent && searchFile {
			// indicates only search file
			mode = FILE_MODE
		} else if !searchFile && !searchContent && searchDir {
			// indicates only search dir
			mode = DIR_MODE
		} else if searchDir && searchFile && !searchContent {
			// indicates search dir and file
			mode = DIR_FILE_MODE
		} else if searchDir && !searchFile && searchContent {
			// indicates search dir and content
			mode = DIR_CONTENT_MODE
		} else if !searchDir && searchFile && searchContent {
			// indicates search file and content
			mode = FILE_CONTENT_MODE
		} else {
			mode = ALL_MODE
		}

		Search(mode, rootPath, keyWords)
	}
}
