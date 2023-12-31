package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom type to associate write function to achieve results of "io.Copy(os.Stdout, res.Body)"
type logWriter struct{}

func main() {
	fmt.Println("Hello from Go-http!!")

	res, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println("Response:", res.Body)
	// defer res.Body.Close()
	// body, err := io.ReadAll(res.Body)
	// fmt.Println("Response Body:", string(body))

	// READING: Response Data ▶️
	// Generic syntx for creating byte slice: ➡️ bs := []byte{"data"}
	// New Syntx using make()⬇️: here we define by using the make func and passing the type and data it is to housein.

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// Writer Interface to read the resp data thro Read Interface ⬇️
	// io.Copy(os.Stdout, res.Body)

	// Calling the custom Write Interface
	lw := logWriter{}
	io.Copy(lw, res.Body)

}

// Write functino to read the http resp body data, similat to "io.Copy(os.Stdout, res.Body)"
func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Custom writer called ..", lw)

	fmt.Println(string(bs))
	fmt.Println("No of bytes writte: ", len(bs))

	return len(bs), nil
	// return 1 nil
}
