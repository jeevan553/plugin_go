package main
import (
    "fmt"
    "os"
)
func main() {
    fmt.Println("Running a custom Drone plugin")
    // Read environment variables
    input := os.Getenv("PLUGIN_INPUT")
 
    // Do something with the input
    if input == "" {
        fmt.Println("No input provided")
    } else {
        fmt.Printf("Input received: %s\n", input)
    }
    fmt.Println("Plugin execution completed")
}
