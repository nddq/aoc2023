package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/nddq/aoc2023/common"
)

type mapper struct {
	sourceStart      int64
	destinationStart int64
	rangeLen         int64
}

func mapperFilter(value int64, mappers []*mapper) int64 {
	for _, mapper := range mappers {
		if mapper.withinRange(value) {
			return mapper.mapToDestination(value)
		}
	}
	return value
}

func newMapper(sourceStart, destinationStart, rangeLen int64) *mapper {
	return &mapper{
		sourceStart:      sourceStart,
		destinationStart: destinationStart,
		rangeLen:         rangeLen,
	}
}

func (m *mapper) withinRange(source int64) bool {
	return source >= m.sourceStart && source < m.sourceStart+m.rangeLen
}

func (m *mapper) mapToDestination(source int64) int64 {
	return (source - m.sourceStart) + m.destinationStart
}

func convertStringSliceTOInt64Slice(strSlice []string) []int64 {
	var result []int64
	for _, str := range strSlice {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return result
}

func readFileToMapperSlice(filename string) (mappers []*mapper) {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	strSlices := common.ParseBytesToStringSlices(input)
	for _, slice := range strSlices {
		strSplit := strings.Split(slice, " ")
		sourceStart, err := strconv.ParseInt(strSplit[1], 10, 64)
		if err != nil {
			panic(err)
		}
		destinationStart, err := strconv.ParseInt(strSplit[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rangeLen, err := strconv.ParseInt(strSplit[2], 10, 64)
		if err != nil {
			panic(err)
		}
		mapper := newMapper(sourceStart, destinationStart, rangeLen)
		mappers = append(mappers, mapper)
	}
	return

}

func SolveOne() {
	input, err := os.ReadFile("seeds.txt")
	if err != nil {
		panic(err)
	}
	// convert []byte to string
	seedsStr := string(input)
	seeds := convertStringSliceTOInt64Slice(strings.Split(seedsStr, " "))

	seed2soil := readFileToMapperSlice("seed2soil.txt")
	soil2fertilizer := readFileToMapperSlice("soil2fertilizer.txt")
	fertilizer2water := readFileToMapperSlice("fertilizer2water.txt")
	water2light := readFileToMapperSlice("water2light.txt")
	light2temp := readFileToMapperSlice("light2temp.txt")
	temp2humidity := readFileToMapperSlice("temp2humidity.txt")
	humidity2loc := readFileToMapperSlice("humidity2loc.txt")

	lowestLocation := int64(math.MaxInt64)
	for _, seed := range seeds {
		soil := mapperFilter(seed, seed2soil)
		fertilizer := mapperFilter(soil, soil2fertilizer)
		water := mapperFilter(fertilizer, fertilizer2water)
		light := mapperFilter(water, water2light)
		temp := mapperFilter(light, light2temp)
		humidity := mapperFilter(temp, temp2humidity)
		loc := mapperFilter(humidity, humidity2loc)
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	fmt.Println(lowestLocation)
}

func SolveTwo() {
	input, err := os.ReadFile("seeds.txt")
	if err != nil {
		panic(err)
	}
	// convert []byte to string
	seedsStr := string(input)
	seeds := convertStringSliceTOInt64Slice(strings.Split(seedsStr, " "))

	seed2soil := readFileToMapperSlice("seed2soil.txt")
	soil2fertilizer := readFileToMapperSlice("soil2fertilizer.txt")
	fertilizer2water := readFileToMapperSlice("fertilizer2water.txt")
	water2light := readFileToMapperSlice("water2light.txt")
	light2temp := readFileToMapperSlice("light2temp.txt")
	temp2humidity := readFileToMapperSlice("temp2humidity.txt")
	humidity2loc := readFileToMapperSlice("humidity2loc.txt")

	lowestLocation := int64(math.MaxInt64)
	loc := make(chan int64)
	var wg sync.WaitGroup
	for i := 0; i < len(seeds)-1; i += 2 {
		start := seeds[i]
		end := seeds[i] + seeds[i+1]
		wg.Add(1)
		go func() {
			for seed := start; seed < end; seed++ {
				soil := mapperFilter(seed, seed2soil)
				fertilizer := mapperFilter(soil, soil2fertilizer)
				water := mapperFilter(fertilizer, fertilizer2water)
				light := mapperFilter(water, water2light)
				temp := mapperFilter(light, light2temp)
				humidity := mapperFilter(temp, temp2humidity)
				loc <- mapperFilter(humidity, humidity2loc)
			}
			wg.Done()
		}()
	}
	go func() {
		for {
			res := <-loc
			if res < lowestLocation {
				fmt.Printf("new lowestLocation: %d\n", res)
				lowestLocation = res
			}
		}
	}()
	wg.Wait()
	fmt.Println(lowestLocation)
}

func main() {
	SolveTwo()
}
