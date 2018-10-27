
package main 

import (
	"fmt"
	"os"
	"github.com/avekceeb/nfsverificator/codegen"
)

const protoPath = "/home/dima/go/src/github.com/avekceeb/nfsverificator/codegen/in.x"

func main() {
	rdr, err := os.Open(protoPath)
	if err != nil {
		fmt.Printf("failed to open protocol file  %v\n", err)
		os.Exit(1)
	}
	defer rdr.Close()

	if err = codegen.Generate(rdr); err != nil {
		fmt.Println("code generator failed:", err)
		os.Exit(1)
	}
}
