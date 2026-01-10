package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func status2use(grades map[string]float64, name string) {

	grade, ok := grades[name]
	if !ok {
		fmt.Printf("no grades records for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	} else if grade >= 60 {
		fmt.Printf("%s is ok\n", name)
	}
}

func status2() {
	grades := map[string]float64{"dfsfsf": 28, "fsdafa": 83}
	var name string
	name = "dsfs"
	status2use(grades, name)
	name = "dfsfsf"
	status2use(grades, name)
	name = "fsdafa"
	status2use(grades, name)
}

func Getstring(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func status3() {
	lines ,err:=Getstring("../../data/vote.txt")
	if err != nil{
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _,line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	for name,count := range counts{
		fmt.Println(name,count)
	}
}

func main() {
	// 使用slice来实现这个分数查询系统，参考2026.1.9.map映射
	status2()
	status3()
}
