// https://www.codewars.com/kata/5279f6fe5ab7f447890006a7
package main

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	output := PosPeaks{[]int{}, []int{}}
	total := len(array)
	pt := 0
	for i := 1; i < total; i++ {
		if array[i] > array[i-1] {
			pt = i
		} else if pt > 0 && array[i] < array[pt] {
			output.Pos = append(output.Pos, pt)
			output.Peaks = append(output.Peaks, array[pt])
			pt = 0
		}
	}
	return output
}
