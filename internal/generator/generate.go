package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/SteveRuble/glodash/internal/fns"
)

func Generate(m *Model, path string) error {
	for _, fn := range m.Functions {

		snakeName := fns.SnakeCase(fn.Name)
		fpath := filepath.Join(path, fmt.Sprintf("fn_%s.go", snakeName))

		err := WriteFunctionFile(m, fn, fpath)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteFunctionFile(m *Model, fn *FunctionInfo, path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o600)
	if err != nil {
		return err
	}

	name := fn.Name
	name = fmt.Sprintf("%s%s", strings.ToUpper(name[0:1]), name[1:])

	var args []string
	for _, arg := range fn.Args {
		args = append(args, fmt.Sprintf("%s %s", arg.Name, arg.GetType()))
	}

	returnType := GetType(fn.Return)

	_, err = fmt.Fprintf(f, `package glodash

%s
func %s[T any](%s) %s {
	var out %s
	return out
}
	`, fn.Comment, name, strings.Join(args, ", "), returnType, returnType)

	return f.Close()
}

func GenerateToDo(m *Model, path string) error {
	var names []string
	for _, fn := range m.Functions {
		names = append(names, fn.Name)
	}
	sort.Strings(names)

	fpath := filepath.Join(path, "todo.md")

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0o600)
	if err != nil {
		return err
	}

	fmt.Fprintln(f, `# Implemented Functions:
	
`)
	for _, name := range names {
		linkName := fmt.Sprintf("./fn_%s.go", fns.SnakeCase(name))
		fmt.Fprintf(f, " - [ ] [%s](%s)\n", name, linkName)
	}

	return f.Close()
}
