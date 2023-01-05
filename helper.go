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
			ans, line2 := hit(pattern, line, ind, flags, len(files), file.name)
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

func hit (pattern string, line string, ind int, flags map[string]int, filesSize int, fileName string) (bool, string) {
	var sum int = flags["-i"] + flags["-v"] + flags["-x"]
	var line2 string

	if filesSize > 1 {
		line2 = fileName + ":"
	}

	if flags["-n"] == 1 {
		line2 += fmt.Sprint(ind + 1) + ":"
	}

	if sum == 0 {
		if strings.Contains(line, pattern) {			
			return true, line2 + line
		}
	} else if sum == 1 {
		if flags["-i"] == 1 {
			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {				
				return true, line2 + line
			}
		} else if flags["-v"] == 1 {
			if !strings.Contains(line, pattern) {				
				return true, line2 + line
			}
		} else if flags["-x"] == 1 {
			if line == pattern {				
				return true, line2 + line
			}
		}
	} else if sum == 2 {

		if flags["-i"] == 1 && flags["-v"] == 1 {
			if !strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				return true, line2 + line
			}
		} else if flags["-i"] == 1 && flags["-x"] == 1 {
			if strings.EqualFold(line, pattern) {
				return true, line2 + line
			}
		} else {
			if line != pattern {
				return true, line2 + line
			}
		}
	} else {
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

// if line == pattern{
// 	filenames[file.name] = true
// 	if flags["-n"]{					
// 		if len(files) > 1 {
// 			res = append(res, file.name + ":" + fmt.Sprint(ind + 1) + ":" + line)
// 		} else {
// 			res = append(res, fmt.Sprint(ind + 1) + ":" + line)
// 		}
// 	} else{
// 		if len(files) > 1 {
// 			res = append(res, file.name + ":" + line)
// 		} else {
// 			res = append(res, line)
// 		}
// 	}
// }