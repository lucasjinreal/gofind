package main

import (
	"fmt"
	"gitlab.com/jinfagang/colorgo"
	"path/filepath"
	"os"
	"strings"
	"strconv"
	"unicode/utf8"
	"bufio"
)


// this structure contains content search result
type ContentResult struct{
	Line string
	LineNumber int
}


func suffixInSupport(suffix string) bool {
	// check a suffix is in support file suffix or not
	supportFiles := []string{
		"txt", "log", "py",
	}
	for _, s := range supportFiles{
		if s == suffix {
			return true
		}
	}
	return false
}

func suffixInIgnored(suffix string) bool {
	// to judge a file suffix is ignored file type or not
	ignoredFiles := []string{
		// executable file
		"exe", "apk", "app",

		// video file
		"mp4", "avi", "mv4", "mp3",

		// audio file
	}
	for _, s := range ignoredFiles{
		if s == suffix {
			return true
		}
	}
	return false
}

// This is a rewrite version of search, do search dir, filename and file content at the same time
func Search(mode int, path string, keyWords string) {


	// this file will search keywords in all file names under path
	// first list all files recursive
	resultDirs := []string{}
	resultFiles := []string{}
	resultContent := map[string][]ContentResult{}

	cg.PrintlnYellow("walking to searching...")
	cg.HighlightPrintln("searching " + keyWords + " from " + path, keyWords, cg.Yellow, cg.Green)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {


		// search file names
		if IsFile(path) {
			if strings.Contains(f.Name(), keyWords) {
				if mode == ALL_MODE || mode == FILE_MODE || mode == FILE_CONTENT_MODE || mode == DIR_FILE_MODE {
					// any mode needs file will call this execute
					resultFiles = append(resultFiles, path)
				}
			}


			// read file and see what inside it
			if mode == ALL_MODE || mode == CONTENT_MODE || mode == DIR_CONTENT_MODE || mode == FILE_CONTENT_MODE {
				f, err := os.Open(path)
				if err != nil {
					fmt.Println(err.Error())
				}

				stat, err := f.Stat()
				// we pass file large than 50M
				if stat.Size() < 50000000 {

					b1 := make([]byte, 40)
					f.Read(b1)
					if utf8.Valid(b1){
						// indicates this is a utf-8 plain text file, we will search this
						fScanner := bufio.NewScanner(f)
						lineNum := 1

						singleResult := []ContentResult{}
						for fScanner.Scan() {
							l := fScanner.Text()
							if strings.Contains(l, keyWords) {
								contentResult := ContentResult{l, lineNum}
								singleResult = append(singleResult, contentResult)
							}
							lineNum += 1
						}
						if len(singleResult) >= 1 {
							resultContent[path] = singleResult
						}

					}
					f.Close()
				}

			}


		}

		// search dir names
		if IsDir(path) {
			if strings.Contains(f.Name(), keyWords) {
				if mode == ALL_MODE || mode == DIR_MODE || mode == DIR_FILE_MODE || mode == DIR_CONTENT_MODE {
					resultDirs = append(resultDirs, path)
				}
			}
		}

		return nil
	})

	if err != nil {
		cg.PrintlnRed("error occurred in walking dir under " + path)
		os.Exit(1)
	}

	switch mode {
	case DIR_MODE:
		printDirResult(path, resultDirs, keyWords)

	case FILE_MODE:
		printFileResult(path, resultFiles, keyWords)

	case CONTENT_MODE:
		printContentResult(path, resultContent, keyWords)

	case DIR_FILE_MODE:
		printDirResult(path, resultDirs, keyWords)
		printFileResult(path, resultFiles, keyWords)

	case DIR_CONTENT_MODE:

		printDirResult(path, resultDirs, keyWords)
		printContentResult(path, resultContent, keyWords)

	case FILE_CONTENT_MODE:
		printFileResult(path, resultFiles, keyWords)
		printContentResult(path, resultContent, keyWords)

	case ALL_MODE:
		printDirResult(path, resultDirs, keyWords)
		printFileResult(path, resultFiles, keyWords)
		printContentResult(path, resultContent, keyWords)
	}
}


func printDirResult(path string, result []string, highlight string){
	cg.Foreground(cg.Blue, false)
	fmt.Print("\033[1m")
	fmt.Println("==== searching result from directory names under " + path)
	fmt.Print("\033[0m")

	if len(result) == 0 {
		// indicates not got result
		s := "not found any directory name contains " + highlight
		cg.HighlightPrintln(s, highlight, cg.Red, cg.Yellow)
	} else {

		// else print result one by one with highlight all
		for _, r := range result {
			cg.HighlightAllPrintln(r, highlight, cg.White, cg.Green)
		}
		s := "Done! found all " + strconv.Itoa(len(result)) + " records contains " + highlight +
			" in directory names under " + path
		cg.HighlightPrintln(s, highlight, cg.Blue, cg.Yellow)
	}
	fmt.Println("")

}

func printFileResult(path string, result []string, highlight string) {
	cg.Foreground(cg.Blue, false)
	fmt.Print("\033[1m")
	fmt.Println("==== searching result from file names under " + path)
	fmt.Print("\033[0m")

	if len(result) == 0 {
		// indicates not got result
		s := "not found any file name contains " + highlight
		cg.HighlightPrintln(s, highlight, cg.Red, cg.Yellow)
	} else {

		// else print result one by one with highlight all
		for _, r := range result {
			cg.HighlightAllPrintln(r, highlight, cg.White, cg.Green)
		}
		s := "Done! found all " + strconv.Itoa(len(result)) + " records contains " + highlight + " in directory names under " + path
		cg.HighlightPrintln(s, highlight, cg.Blue, cg.Yellow)
	}

	fmt.Println("")


}
func printContentResult(path string, result map[string][]ContentResult, highlight string) {
	cg.Foreground(cg.Blue, false)
	fmt.Print("\033[1m")
	fmt.Println("==== searching result from file content names under " + path + "")
	fmt.Print("\033[0m")

	if len(result) == 0 {
		// indicates not got result
		s := "not found any file content include " + highlight
		cg.HighlightPrintln(s, highlight, cg.Red, cg.Yellow)
	} else {

		for s, v := range result {
			fmt.Print("\033[1m")
			s2 := "\nfound file contains " + "\x1b[1;32m" + highlight + "\x1b[0m" + "\033[1m" + ": " + "[" + s + "]"
			fmt.Print(s2)
			fmt.Println("\033[0m")

			// cr is content result
			for _, contentRes := range v {
				cg.PrintYellow("line " + strconv.Itoa(contentRes.LineNumber) + ": ")
				cg.HighlightAllPrintln(contentRes.Line, highlight, cg.White, cg.Green)
			}
		}

		s := "Done! found all " + strconv.Itoa(len(result)) + " records contains " + highlight +
			" in file content under " + path
		cg.HighlightPrintln(s, highlight, cg.Blue, cg.Yellow)
	}
	fmt.Println("")

}

