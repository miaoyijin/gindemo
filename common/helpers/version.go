package version

import "fmt"

var (
	AppName    string = "gindemo" // 应用名称
	GoVersion    string // Golang信息
)

func Version() {
	fmt.Printf("App Name:\t%s\n", AppName)
	fmt.Printf("Golang Version: %s\n", GoVersion)
}