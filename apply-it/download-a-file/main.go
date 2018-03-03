// http://localhost:8087/nexus/service/local/repositories/releases/content/com/aramco/AramcoTheme/1.0.3/AramcoTheme-1.0.3.war

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	b, err := downloadFile("aramco-theme.war", "http://localhost:8087/nexus/service/local/repositories/releases/content/com/aramco/AramcoTheme/1.0.3x/AramcoTheme-1.0.3.war")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes file size downloaded\n", b)

}

func downloadFile(filepath, url string) (int64, error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	b, err := io.Copy(out, resp.Body)
	if err != nil {
		return 0, err
	}
	return b, nil
}
