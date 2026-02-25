package main

import (
	"bufio"
	"embed"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//go:embed resources/r50k_base.tiktoken
var resource embed.FS

var (
	Vocabulary    = make(map[string]int)
	VacabularyInv = make(map[int]string)
)

func LoadVocabulary() error {
	fd, err := resource.Open("resources/r50k_base.tiktoken")
	if err != nil {
		return fmt.Errorf("fail to open file %w", err)
	}
	defer func() {
		if err := fd.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewReader(fd)
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading a line %w", err)
		}

		split := strings.Split(line, " ")
		split[0] = strings.TrimSpace(split[0])
		split[1] = strings.TrimSpace(split[1])

		j, err := base64.StdEncoding.DecodeString(split[0])
		if err != nil {
			return fmt.Errorf("error decoding string %s, %w", split[0], err)
		}
		i, err := strconv.Atoi(split[1])
		if err != nil {
			return fmt.Errorf("error converting %s to int, %w", split[1], err)
		}
		key := string(j)
		val := i

		Vocabulary[key] = val
		VacabularyInv[val] = key
	}
	return nil
}
