//go:generate gen

package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"strconv"
)

type Water struct {
	Id    int
	Name  string
	Price int
	Size  int
}

func main() {
	file, err := os.Open("water.csv")
	if err != nil {
		panic(fmt.Errorf("ファイルがありません"))
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// 保持するWaterの配列
	var waters []Water = make([]Water, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Errorf("読み込みエラーです"))
		}
		waters = append(waters, arrayToWater(record))
	}
	fmt.Println(fmt.Sprintf("%+v", waters))
}

func arrayToWater(arr []string) Water {
	var id int
	id, _ = strconv.Atoi(arr[0])

	var name string
	name = arr[1]

	var price int
	price, _ = strconv.Atoi(arr[2])

	var size int
	size, _ = strconv.Atoi(arr[3])

	return Water{
		id,
		name,
		price,
		size,
	}
}