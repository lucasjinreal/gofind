package main

import (
	"gitlab.com/jinfagang/colorgo.git"
	"flag"
	"os"
	"bufio"
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


func goFind(rootPath string, mode int, keyWords string) {

	switch mode {
	case DIR_MODE:
		SolveDirMode(rootPath, keyWords)
	case FILE_MODE:
		SolveFileMode(rootPath, keyWords)
	case CONTENT_MODE:
		SolveContentMode(rootPath, keyWords)
	case DIR_FILE_MODE:
		SolveDirMode(rootPath, keyWords)
		SolveFileMode(rootPath, keyWords)
	case DIR_CONTENT_MODE:
		SolveDirMode(rootPath, keyWords)
		SolveContentMode(rootPath, keyWords)
	case FILE_CONTENT_MODE:
		SolveFileMode(rootPath, keyWords)
		SolveContentMode(rootPath, keyWords)
	case ALL_MODE:
		SolveDirMode(rootPath, keyWords)
		SolveContentMode(rootPath, keyWords)
		SolveFileMode(rootPath, keyWords)

	}
}

func main() {
	cg.PrintlnBlue("gofind - find anything you try to find.\n")
	cg.PrintlnGreen(`   __________  ___________   ______
  / ____/ __ \/ ____/  _/ | / / __ \
 / / __/ / / / /_   / //  |/ / / / /
/ /_/ / /_/ / __/ _/ // /|  / /_/ /
\____/\____/_/   /___/_/ |_/_____/
                                    `)
	cg.PrintlnCyan("Author - Jin Tian.")

	searchDir := flag.Bool("d", false, "search dir names.")
	searchFile := flag.Bool("f", false, "search file names.")
	searchContent := flag.Bool("c", false, "search all file content.")
	flag.Parse()
	cg.PrintlnGreen(flag.NArg())

	keyWords := ""
	if flag.NArg() == 0 {
		cg.PrintlnRed("You must specific a path, type: gofind ./ -d -f test")
		flag.Usage()
		os.Exit(1)
	} else if flag.NArg() < 2 {
		reader := bufio.NewReader(os.Stdin)
		cg.PrintlnGreen("What keywords do you want to search? \n" +
			"type: gofind ./ -c Lewis  to search all file content contains Lewis.")
		text, _ := reader.ReadString('\n')
		keyWords = text

	}

	// default mode is all mode
	mode := ALL_MODE

	if !*searchDir && !*searchFile && *searchContent{
		// indicates only search content
		mode = CONTENT_MODE
	} else if !*searchDir && !*searchContent && *searchFile{
		// indicates only search file
		mode = FILE_MODE
	} else if !*searchFile && !*searchContent && *searchDir{
		// indicates only search dir
		mode = DIR_MODE
	} else if *searchDir && *searchFile && !*searchContent{
		// indicates search dir and file
		mode = DIR_FILE_MODE
	} else if *searchDir && !*searchFile && *searchContent {
		// indicates search dir and content
		mode = DIR_CONTENT_MODE
	} else if !*searchDir && *searchFile && *searchContent {
		// indicates search file and content
		mode = FILE_CONTENT_MODE
	} else {
		mode = ALL_MODE
	}

	rootPath := flag.Arg(0)
	goFind(rootPath, mode, keyWords)






}
