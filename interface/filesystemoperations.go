package ui

import (
	"embed"
)

var Files embed.FS
func main(){
	fileSystem, err := fs.Sub(Files, "html")
	if err != nil {
		fmt.Println("Failed to create file system:", err)
	return
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(fileSystem))))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(Files))))

	fmt.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
