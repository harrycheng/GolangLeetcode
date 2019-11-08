package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

import "hc.com/sorter/src/algorithms"

var infile *string = flag.String("i", "infile", "file in...")
var outfile *string = flag.String("o", "outfile", "finle out")
var algorithm *string = flag.String("a", "", "sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println(*infile, *outfile, *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println(values)
		switch *algorithm {
		case "qsort":
			//algorithms.QuickSort(&values)
			algorithms.MySort(&values)
		case "bubblesort":
			algorithms.BubbleSort(values)
			fmt.Println(values)
		default:
			algorithms.BubbleSort(values)
		}
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println(" to long")
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println(" outfile error")
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}
