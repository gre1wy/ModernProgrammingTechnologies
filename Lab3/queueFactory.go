package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Визначає ліміт на основі введення із io.Reader
func determineLimit(reader io.Reader) (int, error) {
	bufReader := bufio.NewReader(reader)
	fmt.Println("Вкажіть розмір кільцевої черги(додатнє число):")
	s, err := bufReader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	s = strings.TrimSpace(s)
	limit, err := strconv.ParseInt(s, 10, 0)
	if err != nil || limit < 0 {
		return 0, fmt.Errorf("wrong format of size: entry number>=0")
	}
	return int(limit), nil
}

// Отримує чергу залежно від зазначеного типу та введення
func getQueue(queueType string, input io.Reader) (Queue, error) {
	switch queueType {
	case "LinkedListQueue":
		return NewLinkedListQueue(), nil
	case "CircularArrayQueue":
		limit, err := determineLimit(input)
		if err != nil {
			return nil, err
		}
		res, _ := NewCircularArrayQueue(limit)
		return res, nil
	}
	return nil, fmt.Errorf("wrong queue type passed")
}
