//go:generate gen

package main

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"strconv"
)

const (
	ProfitRatioIrohasu = 0.12
	ProfitRatioVolvic = 0.2
	ProfitRatioCrystalGeyser = 0.05
	ProfitRatioAlcaliIonWater = 0.1
	ProfitRatioPerrier =  0.1
	ProfitRatioContrex = 0.2
	ProfitRatioWilkinson =  0.13
	ProfitRatioAmebaWater = 0
	ProfitRatioOnsensui = 0.4
)

var profitRatio = map[string] float64 {
	"Irohasu" : ProfitRatioIrohasu,
	"Volvic" : ProfitRatioVolvic,
	"CrystalGeyser" : ProfitRatioCrystalGeyser,
	"AlcaliIonWater" : ProfitRatioAlcaliIonWater,
	"Perrier" : ProfitRatioPerrier,
	"Contrex" : ProfitRatioContrex,
	"Wilkinson" : ProfitRatioWilkinson,
	"AmebaWater" : ProfitRatioAmebaWater,
	"Onsensui" : ProfitRatioOnsensui,
}

// +gen slice:"Count,GroupBy[string],Select[int],Select[float64]"
type Water struct {
	Id    int
	Name  string
	Price int
	Size  int
}

type WaterSumalies []WaterSumally
type WaterSumally struct {
	Name string
	Summary Sum
}
type Sum struct {
	sum int
	sizeTotal int
	profit float64
}

func main() {


	var waters []Water = make([]Water, 0)
	waters = readFile("water.csv")

	//fmt.Println(fmt.Sprintf("%+v", waters))

	// GroupByの結果保持用にスライス作っておく
	var waterSlice WaterSlice = waters

	// クロージャを変数に入れておく
	waterNameGroupBy := func (w Water) string {
		return w.Name
	}

	// groupBy実行
	slice := waterSlice.GroupByString(waterNameGroupBy)

	// 出力用にまとめるハッシュ定義
	allSumally := make(map[string]WaterSumalies)

	// 名前をキーにしてWaterSliceをループする
	for waterName, water := range(slice) {
		if waterName == "name" {
			continue
		}

		var sum int = 0 // 価格合計
		var sizeTotal int = 0 // 容量合計

		for _, v := range(water) {
			sum += v.Price
			sizeTotal += v.Size
		}

		// 利益 = 売上 * 利益率
		var profit float64 = float64(sum) * profitRatio[waterName]

		//fmt.Printf("waterName: %s , " , waterName)
		//fmt.Printf("sum : %d , " , sum)
		//fmt.Printf("sizeTotal : %d ," , sizeTotal)
		//fmt.Println("利益:" , profit)

		// sum処理保持用にサマリー構造体を作っておく。 interface{}なのでsumの部分は型はなんでも良い
		var waterSumaly WaterSumally = WaterSumally{
			waterName,
			Sum{
				sum,
				sizeTotal,
				profit,
			},
		}

		allSumally["sum"] = append(allSumally["sum"], waterSumaly)
	}
	//fmt.Println(allSumally)

	printAll(allSumally)
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

func readFile(filename string) []Water {
	file, err := os.Open(filename)
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
	return waters
}

func printAll(allSumally map[string]WaterSumalies) {
	var sum = make(map[string]int)
	var size = make(map[string]int)
	var benefit = make(map[string]int)
	var sumTotal int = 0
	var sizeTotal int = 0
	var benefitTotal int = 0

	for _, waterSumally := range(allSumally["sum"]) {
		sum[waterSumally.Name] = waterSumally.Summary.sum
		size[waterSumally.Name] = waterSumally.Summary.sizeTotal
		benefit[waterSumally.Name] = int(waterSumally.Summary.profit)

		sumTotal += waterSumally.Summary.sum
		sizeTotal += waterSumally.Summary.sizeTotal
		benefitTotal += int(waterSumally.Summary.profit)
	}

	fmt.Println("------------------------------\n金額の合計\n------------------------------")
	for name, s := range sum {
		fmt.Printf("%s : %d\n", name, s)
	}
	fmt.Printf("全ての商品の合計 : %d\n", sumTotal)

	fmt.Println("------------------------------\n容量の合計\n------------------------------")
	for name, sz := range size {
		fmt.Printf("%s : %d\n", name, sz)
	}
	fmt.Printf("全ての商品の合計 : %d\n", sizeTotal)

	fmt.Println("------------------------------\n全商品の利益\n------------------------------")
	for name, bf := range benefit {
		fmt.Printf("%s : %d\n", name, bf)
	}
	fmt.Printf("全ての商品の利益 : %d\n", benefitTotal)

	fmt.Println("------------------------------\n全商品の利率\n------------------------------")
	for name, pf := range profitRatio {
		fmt.Printf("%s : %f\n", name, pf)
	}

}