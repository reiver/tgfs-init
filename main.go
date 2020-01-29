package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	root = ".tgfs"
)

func main() {

	fmt.Printf("\x1b[1mThe Great File System\x1b[0m (\x1b[1mTGFS\x1b[0m) — \x1b[1m%s\x1b[0m\n\n", os.Args[0])

	// .tgfs/digest/sha-3-512
	{
		var builder strings.Builder

		builder.WriteString(root)
		builder.WriteRune(os.PathSeparator)
		builder.WriteString("digest")
		builder.WriteRune(os.PathSeparator)
		builder.WriteString("sha-3-512")

		var path string = builder.String()

		if err := os.MkdirAll(path, 0755); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Could not create path %q: %s\n", err)
			os.Exit(1)
		}
	}

	// .tgfs/mnt
	{
		var builder strings.Builder

		builder.WriteString(root)
		builder.WriteRune(os.PathSeparator)
		builder.WriteString("mnt")

		var path string = builder.String()

		if err := os.MkdirAll(path, 0755); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Could not create path %q: %s\n", err)
			os.Exit(1)
		}
	}

	{
		path, err := filepath.Abs(root)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Could not determine where the %s is: %s", root, err)
			os.Exit(1)
		}

		fmt.Printf("Created an empty local instance of The Great File System (TGFS) at %s%s\n", path, string(os.PathSeparator))
	}
}
