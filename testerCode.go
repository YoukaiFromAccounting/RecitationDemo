package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Setup type to store a pair of integers
type PythPair [2]int
type scores []int

func LongestCommonSubstring(stringOne, stringTwo string) string {
	//StringOneSubstrings
	substringsOne := []string{}
	substringsTwo := []string{}

	//Get all substrings from stringOne
	for i := 0; i <= len(stringOne); i++ {
		for j := (i + 1); j <= len(stringOne); j++ {
			substringsOne = append(substringsOne, stringOne[i:j])
		}
	}
	//Get all substrings from stringTwo
	for i := 0; i <= len(stringTwo); i++ {
		for j := (i + 1); j <= len(stringTwo); j++ {
			substringsTwo = append(substringsTwo, stringTwo[i:j])
		}
	}

	longestSharedString := "stringtest"
	sum := 0
	for i := range substringsOne {
		for j := range substringsTwo {
			if (substringsOne[i] == substringsTwo[j]) && (len(substringsOne[i]) > sum) {
				longestSharedString = substringsOne[i]
				sum = len(substringsOne[i])
			}
		}
	}

	return longestSharedString
}

func BestAverage(grades [][]int) int {

	averages := []int{}

	for j := range grades {
		//Returns an average for a set
		sum := 0
		for i := range grades[j] {
			sum = sum + grades[j][i]
		}

		averages = append(averages, sum/len(grades[0]))
	}

	//Get Largest Value in Averages
	largest := 0
	for i := range averages {
		if averages[i] > largest {
			largest = averages[i]
		}
	}
	return largest

}

type Body struct {
	name                             string
	mass, radius                     float64
	position, velocity, acceleration PythPair
	red, green, blue                 uint8
}

func Distance(a, b Body) float64 {
	xDist := math.Abs(float64(a.position[0] - b.position[0]))
	yDist := math.Abs(float64(a.position[1] - b.position[1]))

	linDist := math.Sqrt(xDist*xDist + yDist*yDist)
	finDist := linDist - a.radius - b.radius
	return finDist
}

func AltBestAverage(grades map[string][]int) int {
	//Map Data Type Variant
	//mersey := make(map[string][]int)

	//Stores the averages
	averages := []int{}

	//Iterate over the map
	for i := range grades {

		sum := 0
		for j := 0; j < len(grades[i]); j++ {
			sum = sum + grades[i][j]
		}
		averages = append(averages, sum/len(grades[i]))
	}

	//Find largest of the averages
	largest := 0
	for i := range averages {

		if averages[i] > largest {
			largest = averages[i]

		}
	}
	return largest
}

func TwinColorsHouseEdge(numTrials int) int {
	//If both balls are red (chances = 2/10 * 1/9), house loses 10
	//If both balls are blue (chances = 3/10 * 2/9), house loses 1
	//If both balls green (chances = 5/10 * 4/9), house gains 1
	sum := 0
	randomArray := []string{"red", "red", "blue", "blue", "blue", "green", "green", "green", "green", "green"}

	for i := 0; i < numTrials; i++ {
		x := randomArray[rand.Intn(10)]
		y := randomArray[rand.Intn(10)]
		if x == y {
			if x == "red" {
				sum = sum + 10
			} else if x == "blue" {
				sum = sum + 1
			} else if x == "green" {
				sum = sum - 1
			}
		}

	}
	return sum
}

func main() {
	//FindPythagoreanLegs(25)
	//fmt.Println(LongestCommonSubstring("applePiesandclamChowder", "clamChowder"))

	/*
		a := make(map[string][]int)
		b := []int{2, 3, 4, 5, 6, 7}
		c := []int{100, 90, 80, 80, 80, 80}
		d := []int{90, 80, 90, 90, 90, 90, 90, 90, 90}
		a["Jack"] = b
		a["Taylor"] = c
		a["Jeremy"] = d
	*/
	/*
		test := []int{5, 6, 7, 4}
		testtwo := []int{40, 50, 500}

		b := make(map[string][]int)
		b["Jack"] = test
		b["Ben"] = testtwo

		fmt.Println(AltBestAverage(b))

	*/
	//arr := []int{1, 2, 3, 4, 6, 6, 6, 6, 6, 6, 6, 6, 2, 3, 3, 4, 2, 1, 1, 1, 2, 3, 4, 5, 6, 3, 1, 1, 2, 3, 4, 2}
	//fmt.Println(FreqMap(arr))
	//fmt.Println(TwinColorsHouseEdge(100))
	//fmt.Println(EulerGCD(9000, 2000))
	//fmt.Println(SubstringParse("SuperCarnegieMellonUniversity", "AndrewSuperCarnegieandAndrewMellon"))
	//a := []string{"albus", "severus", "lily", "james", "RemoveThis"}

	for i := 0; i < 1000; i++ {
		fmt.Println((rand.Intn(1000-(-90)) - 90))
	}
}

func FindPythagoreanLegs(n int) []PythPair {

	arrayPairs := []PythPair{}

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			for k := 0; k < 99999; k++ {
				if (i*i)+(j*j) == (k * k) {
					tempArray := PythPair{}
					tempArray[0] = i
					tempArray[1] = j
					arrayPairs = append(arrayPairs, tempArray)
				}
			}
		}
	}
	return arrayPairs
}

func ReturnLonger(intOne, intTwo int) int {
	if intOne > intTwo {
		return intOne
	} else {
		return intTwo
	}
}

func FreqMap(arrayInts []int) map[int]int {
	freq := make(map[int]int)
	for i := 0; i < len(arrayInts); i++ {
		for num := range arrayInts {
			freq[num] = freq[num] + 1
		}
	}
	return freq
}

func EulerGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if b > a {
			b = b - a
		}
	}
	return a
}

func BestAverageTwo(grades map[string][]int) float64 {
	highestAverage := 0.0
	for i := range grades {
		averager := 0.0
		sum := 0
		for _, j := range grades[i] {
			sum = sum + j
		}

		averager = float64(sum) / float64(len(grades[i]))
		if highestAverage < averager {
			highestAverage = averager
		}
	}
	return highestAverage
}

func BallReturn(trials int) int {
	houseEdge := 0
	jar := []string{"red", "red", "blue", "blue", "blue", "green", "green", "green", "green", "green"}

	for i := 0; i < trials; i++ {
		//Draw matching balls
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Println("X is", x, "Y is ", y)
		if jar[x] == jar[y] {
			if jar[x] == "red" {
				houseEdge -= 10
			} else if jar[x] == "blue" {
				houseEdge -= 1
			} else if jar[x] == "green" {
				houseEdge += 1
			}
		}
	}

	return houseEdge
}

func SubstringParse(text0, text1 string) string {
	substringsOne := []string{}
	substringsTwo := []string{}

	for i := 0; i < len(text0); i++ {
		for j := i; j < len(text0); j++ {
			substringsOne = append(substringsOne, text0[i:j])
		}
	}

	for i := 0; i < len(text1); i++ {
		for j := i; j < len(text1); j++ {
			substringsTwo = append(substringsTwo, text1[i:j])
		}
	}

	length := 0
	longestString := ""
	for _, i := range substringsOne {
		for _, j := range substringsTwo {
			if i == j {
				if len(i) > length {
					longestString = i
					length = len(i)
				}
			}
		}
	}

	return longestString
}

func RemovingFromSlice(tester []string) {
	for x, i := range tester {
		if i == "RemoveThis" {
			tester = append(tester[:x], tester[x+1:]...)
		}
	}

	fmt.Println(tester)
}

func SumSlice(u []int) int {
	sum := 0
	for _, i := range u {
		sum = sum + i
	}
	return sum
}
