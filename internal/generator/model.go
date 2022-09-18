package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Model struct {
	Lines     []string
	Functions map[string]*FunctionInfo
}

type FunctionInfo struct {
	Name          string
	WrappedResult bool
	Args          []ArgInfo
	Aliases       []string
	Comment       string
	Return        string
}

type ArgInfo struct {
	Name string
	Type string
}

func (a ArgInfo) GetType() string {
	return GetType(a.Type)
}

func GetType(s string) string {
	switch s {
	case "Array":
		return "[]T"
	case "Function":
		return "func(T) bool"
	case "Object":
		return "T"
	case "string":
		return "string"
	case "number":
		return "int"
	default:
		return fmt.Sprintf("any /*%s*/", s)
	}
}

func NewModel(path string) (*Model, error) {
	m := &Model{
		Functions: map[string]*FunctionInfo{},
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return m, err
	}

	m.Lines = strings.Split(string(b), "\n")
	log.Printf("got %d lines", len(m.Lines))

	err = m.init()

	return m, err
}

func (m *Model) init() error {
	m.initFunctionInfo()
	return nil
}

func (m *Model) initFunctionInfo() {
	src := &srclines{lines: m.Lines}

	for _, comment := range []string{
		"// Add methods that return wrapped values in chain sequences.",
		"// Add methods that return unwrapped values in chain sequences.",
	} {
		_ = src.mustSeek(comment)

		log.Printf("adding functions after comment %q at index %d", comment, src.index)
		for _, tokens := range src.sequence(`lodash.(\w+) = .*;`) {
			name := tokens[1]

			m.Functions[name] = &FunctionInfo{
				Name:          makeSafe(name),
				WrappedResult: true,
			}
		}

		log.Println("adding aliases")

		_ = src.mustSeek("// Add aliases.")
		for _, tokens := range src.sequence(`lodash.(\w+) = (\w+);`) {
			alias := tokens[1]
			name := tokens[2]
			m.Functions[name].Aliases = append(m.Functions[name].Aliases, alias)
		}
	}

	src.reset()

	for {
		block := src.block(
			`      /\*\*`,
			`      (function|var)`,
		)
		if len(block) == 0 {
			break
		}
		last := block[len(block)-1]
		parts := regexp.MustCompile(`(function|var) (\w+)`).FindStringSubmatch(last)
		name := makeSafe(parts[2])
		if name == "result" {
			continue
		}
		fn, known := m.Functions[name]
		if !known {
			continue
		}

		blockLines := &srclines{lines: block}
		for _, paramInfo := range blockLines.sequence(`@param \{([^}]+)} \[?(\w+)`) {
			fn.Args = append(fn.Args, ArgInfo{
				Name: makeSafe(paramInfo[2]),
				Type: paramInfo[1],
			})
		}
		blockLines.reset()

		returnInfo, ok := blockLines.seek(`@returns \{(.*)\}`)
		if ok {
			fn.Return = returnInfo[1]
		}

		for i := 0; i < len(block)-1; i++ {
			block[i] = strings.TrimSpace(block[i])
		}
		body := strings.Join(block[:len(block)-1], "\n")

		fn.Comment = body
	}
}

var reservedMap = map[string]string{
	"func":   "fn",
	"string": "str",
}

func makeSafe(name string) string {
	if replacement, replace := reservedMap[name]; replace {
		return replacement
	}
	return name
}

type srclines struct {
	lines []string
	index int
}

func (s *srclines) reset() {
	s.index = 0
}

func (s *srclines) seek(re string) (match []string, ok bool) {
	exp := regexp.MustCompile(re)

	for i := s.index; i < len(s.lines); i++ {
		line := s.lines[i]
		m := exp.FindStringSubmatch(line)
		if len(m) > 0 {
			s.index = i + 1
			return m, true
		}
	}

	return nil, false
}

func (s *srclines) block(first, last string) []string {
	var out []string

	firstExp := regexp.MustCompile(first)
	lastExp := regexp.MustCompile(last)

	for ; s.index < len(s.lines); s.index++ {
		line := s.lines[s.index]
		if len(out) == 0 {
			if firstExp.MatchString(line) {
				out = []string{line}
			}
			continue
		}
		out = append(out, line)
		if lastExp.MatchString(line) {
			s.index++
			break
		}
	}

	return out
}

func (s *srclines) sequence(re string) [][]string {
	var out [][]string

	exp := regexp.MustCompile(re)

	for ; s.index < len(s.lines); s.index++ {
		line := s.lines[s.index]
		m := exp.FindStringSubmatch(line)
		if len(m) == 0 {
			if len(out) == 0 {
				continue
			}
			s.index++
			break
		}
		out = append(out, m)
	}

	return out
}

func (s *srclines) mustSeek(re string) []string {
	i := s.index
	m, ok := s.seek(re)
	if !ok {
		panic(fmt.Errorf("no next matched %q: checked lines from %d to %d",
			re, i, s.index,
		))
	}
	return m
}
