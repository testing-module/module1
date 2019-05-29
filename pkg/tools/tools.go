package tools

import (
	"fmt"
)

type (
	ToolVar string
)

func Tool(tv ToolVar) {
	fmt.Println(tv)
}
