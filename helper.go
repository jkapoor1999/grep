package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, fl []string, fi []string) []string {
	files := createFiles(fi)	

	flags := make(map[string]int)
	filenames := make(map[string]bool)	

	res := []string{}
	for _, f := range fl {
		flags[f] = 1;
	}
	
	for _, file := range files{
		for ind, line := range file.lines{
			ans, line2 := hit(pattern, line, file.name, ind, len(files), flags)
			if ans {
				filenames[file.name] = true
				res = append(res, line2)
			}			
		}
	}
	
	if flags["-l"] == 1 {		
		res = nil
		for _, file := range fi {
			if filenames[file] {
				res = append(res, file)
			}
		}		
	} 
	return res
}

func hit (pattern, line, fileName string, ind, filesSize int, flags map[string]int) (bool, string) {
	var sum int = flags["-i"] + flags["-v"] + flags["-x"]
	var line2 string

	if (filesSize > 1) {
		line2 = fileName + ":"
	}

	if (flags["-n"] == 1) {
		line2 += fmt.Sprint(ind + 1) + ":"
	}

	switch (sum) {
	case 0:
		if strings.Contains(line, pattern) {			
			return true, line2 + line
		}
	case 1:
		if (flags["-i"] == 1) {
			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {				
				return true, line2 + line
			}
		} else if (flags["-v"] == 1) {
			if !strings.Contains(line, pattern) {				
				return true, line2 + line
			}
		} else if (flags["-x"] == 1) {
			if line == pattern {				
				return true, line2 + line
			}
		}
	case 2:
		if (flags["-i"] == 1 && flags["-v"] == 1) {
			if !strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				return true, line2 + line
			}
		} else if (flags["-i"] == 1 && flags["-x"] == 1) {
			if strings.EqualFold(line, pattern) {
				return true, line2 + line
			}
		} else {
			if (line != pattern) {
				return true, line2 + line
			}
		}
	case 3:
		if !strings.EqualFold(line, pattern){
			return true, line2 + line
		}
	}		
	return false, ""
}

func createFiles(filenames []string) []File {
	lst := []File{}
	for _, filename := range filenames{
		file, err := os.Open(filename)
		var f File
		f.name = filename
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {		
			f.lines = append(f.lines, scanner.Text())
		}
		lst = append(lst, f)
	}
	return lst
}

func StringSlicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

type File struct {
	name  string
	lines []string
}