// file handling -> it means performing operations like Create, Read, Write, Delete on files using packages like os and io.
package main

// "os"
// "fmt"

// 1. Create a file
// func main() {
// 	file, err := os.Create("example.txt")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer file.Close()

// 	fmt.Println("File created successfully")
// }

// 2. Write to a file
// func main() {
// 	file, err := os.Create("example.txt")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("Hello, this is a sample text written to the file.")
// 	fmt.Println("File created and written to successfully")
// }

// 3. Read from a file
// func main() {
// 	file, err := os.ReadFile("example.txt")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(file))
// }

// 4. Append to a file
// func main() {
// 	file, _ := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
// 	defer file.Close()

// 	file.WriteString("\nNew line added")

// }

// 5. Delete a file
// func main() {
// 	os.Remove("example.txt")
// }

// NOTE: defer -> it delays a function call until the surrounding function finishes execution.
