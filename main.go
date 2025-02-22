package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//TODO come and clear the screen
	// exec.Cmd("")
	// cmd.Output()

	fmt.Println("\tTODO app using CLI...")

	fmt.Println("=================================\n")
	fmt.Println("enter command read or write to use DB, anything else to terminate.")

	// a file pointer to open or create a file.
	var file *os.File
	const FILE_PATH = "./DB.txt"

	file, err := os.Open(FILE_PATH)
	if !handleErr(err, "file, doesn't exist") {

		// file doesn't exist but it'll be created
		file, _ = os.Create(FILE_PATH)
		fmt.Println("file created successfully.")
	}

	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	fmt.Println(operation)

	// infinte for loop for db reads and writes
	flag := false
	for !flag {

		switch operation {
		case "read":
			readDB(file)
			break
		case "write":
			data, _ := reader.ReadString('\n')
			writeDB(file, data)
			break
		default:
			fmt.Println("exiting")
			flag = true
			break

		}
	}

}

func handleErr(err error, msg string) bool {

	// this function checks for errors and return flase if there an error
	// and the check statement will revert the return value of this function.

	if err != nil {
		fmt.Println(msg, "\n", err)
		return false
	}
	return true

}

func readDB(file *os.File) {

	// a scanner to read the file as hole
	fileContent := bufio.NewScanner(file)

	// iterate on the file line by line
	for fileContent.Scan() {
		lineNumber := 1
		line := fileContent.Text()
		fmt.Printf("%v. %v\n", lineNumber, line)
		lineNumber++
	}

}

func writeDB(file *os.File, data string) {

	_, err := io.WriteString(file, data)
	if !handleErr(err, "couldn't write to file.") {
		return
	}
}
