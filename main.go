package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tj/go-naturaldate"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	if text == "" {
		text = "now"
	}

	t, err := naturaldate.Parse(text, time.Now())
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error parsing text: %s\n", err.Error()))
		os.Exit(1)
	}

	fmt.Println(t.Format("2006-01-02T15:04:05.000Z0700"))
}
