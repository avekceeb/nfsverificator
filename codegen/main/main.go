
package main 

import (
	"fmt"
	"os"
	"github.com/avekceeb/nfsverificator/codegen"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		panic("File names not provided!")
	}
    gopath := os.Getenv("GOPATH")
    p := filepath.Join(gopath,
        "src/github.com/avekceeb/nfsverificator/codegen",
        os.Args[1])
	rdr, err := os.Open(p)
	if err != nil {
		fmt.Printf("failed to open protocol file %v\n", err)
		os.Exit(1)
	}
	defer rdr.Close()

	if err = codegen.Generate(rdr, os.Args[2]); err != nil {
		fmt.Println("code generator failed:", err)
		os.Exit(1)
	}
}
