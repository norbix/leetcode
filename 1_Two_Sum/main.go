package main

import (
	log "github.com/sirupsen/logrus"
)

func TwoSum(one, two int) int {
	return one + two
}

func main() {
	i := 1
	j := 1

	result := TwoSum(i, j)

	log.Infof("%v", result)
}
