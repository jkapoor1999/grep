package main

func main() {
	for _, line := range Search("hello", []string{"-i","-n","-x"}, []string{"input.txt", "greeting.txt"}) {
		println(line)
	}
}