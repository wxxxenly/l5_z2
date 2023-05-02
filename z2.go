package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "lab_array.dat"

	data := generateArray(5, 5)
	fmt.Printf("Собран массив значений %s", fmt.Sprintln(data))

	err := writeArrayFile(path, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	source, err := readArrayFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	processAndWriteData(source)
}

func processAndWriteData(source [][]int) {
	//TODO: Add logic here
	
    fmt.Printf("Массив в начале: %v\n", source)
	
	//source - mn := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    maxSum := 0
    maxIndex := 0
    minSum := 100
    minIndex := 100
    
    // находим строки с максимальной суммой
    for i := 0; i < len(source); i++ {
        sum := 0
        for j := 0; j < len(source[0]); j++ {
            sum += source[i][j]
        }
        if sum > maxSum {
            maxSum = sum
            maxIndex = i
        }
    }
    
/*    
    for i, row := range source {
        sum := 0
        for _, val := range row {
            sum += val
        }
        if sum > maxSum {
            maxSum = sum
            maxIndex = i
        }
    }
*/
    fmt.Printf("Массив после первой итерации: %v\n", source)
    // находим столбцы с минимальной суммой
    for j := 0; j < len(source[0]); j++ {
        sum := 0
        for i := 0; i < len(source); i++ {
            sum += source[i][j]
        }
        if sum < minSum {
            minSum = sum
            minIndex = j
        }
    }
	fmt.Sprintln(source)
	
	
    // выводим результат
    fmt.Printf("Строка с максимальной суммой: %v\n", source[maxIndex])
    fmt.Printf("minIndex: %v\n", minIndex)
    fmt.Printf("Массив: %v\n", source)
/*    fmt.Printf("Столбец с минимальной суммой: ")
    for i := 0; i < len(source); i++ {
        fmt.Printf("%v ", source[i][minIndex])
    }
*/
    fmt.Println()
	
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func readArrayFile(path string) ([][]int, error) {
	var file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, errors.New("не удалось прочитать данные из файла")
	}

	var result [][]int
	for scanner.Scan() {
		line, err := sliceAtoi(strings.Split(scanner.Text(), ","))
		if err != nil {
			return nil, err
		}

		result = append(result, line)
	}

	return result, nil
}

func getValueFromUser(valueType string) int {
	var n int

	fmt.Printf("Введите %v: ", valueType)
	fmt.Scan(&n)
	return n
}

func generateArray(n int, m int) [][]int {
	var result [][]int
	vMin := getValueFromUser("v_min")
	vMax := getValueFromUser("v_max")

	for i := 0; i < n; i++ {
		line := make([]int, 0, m)
		for j := 0; j < m; j++ {
			value := rand.Intn(vMax-vMin+1) + vMin
			line = append(line, value)
		}
		result = append(result, line)
	}

	return result
}

func writeArrayFile(path string, data [][]int) error {
	var file, err = os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, line := range data {
		st := fmt.Sprintln(strings.Join(strings.Fields(strings.Trim(fmt.Sprint(line), "[]")), ","))
		_, err = file.WriteString(st)
		if err != nil {
			return err
		}
	}

	return nil
}
