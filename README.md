# grep

Important Points to Focus before Attempting the Problem 
1.	Create a public GitHub repo and commit your code 
2.	Use modules while writing the code
3.	Create a test file and write the test cases as well. 
4.	The code should be properly structured and naming conventions should be followed.  
Implement GREP in Go
Search a file for lines matching a regular expression pattern. Return the line number and contents of each matching line.
The Unix grep command can be used to search for lines in one or more files that match a user-provided search query (known as the pattern).
1.	The grep command takes three arguments:
2.	The pattern which is used to match lines in a file.
3.	Zero or more flags to customize the matching behavior.
One or more files in which to search for matching lines.
Your task is to implement the grep function given a list of files, find all lines that match the specified pattern. Return the lines in the order they appear in the files. You'll also have to handle options (given as flags), which control how matching is done and how the results are to be reported.
As an example, suppose there is a file named "input.txt" with the following contents:
•	hello
•	world
•	hello again

If we were to call grep "hello" input.txt, the result should be:
•	hello
•	hello again
If given multiple files, grep should prefix each found line with the file it was found in. As an example:
•	input.txt:hello
•	input.txt:hello again
•	greeting.txt:hello world
If given just one file, this prefix is not present.



Flags
As said earlier, the grep command should also support the following flags:
1.	-n Prefix each matching line with its line number within its file. When multiple files are present, this prefix goes after the filename prefix.
2.	-l Print only the names of files that contain at least one matching line.
3.	-i Match line using a case-insensitive comparison.
4.	-v Invert the program -- collect all lines that fail to match the pattern.
5.	-x Only match entire lines, instead of lines that contain a match.
If we run grep -n "hello" input.txt, the -n flag will require the matching lines to be prefixed with its line number:
1: hello
3: hello again
And if we run grep -i "HELLO" input.txt, we'll do a case-insensitive match, and the output will be:
hello
hello again
The grep command should support multiple flags at once.
For example, running grep -l -v "hello" file1.txt file2.txt should print the names of files that do not contain the string "hello".
In package grep, Define a single Go func, Search, which accepts a pattern string, a slice of flags which are strings, and a slice of filename strings. Search should return a slice of strings of the output for the given flags and filenames.
Use the following signature for func Search:
func Search(pattern string, flags, files []string) []string {}
