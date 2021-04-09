package main

import "fmt"

var numGenerations = 1024
var x = 31
var y = 919
var s = 42

var cells []int
var ruleset []int
var history [][]int

func init() {
cells = make([]int, numGenerations)
history = make([][]int, 0)
ruleset = []int{0, 1, 1, 0, 0, 1, 0, 1}
for i := 0; i < len(cells); i++ {
cells[i] = 0
}
cells[0] = 1
}

func main() {
for y := 0; y < numGenerations; y++ {
history = append(history, cells)
cells = newGeneration(cells)
}

values := make([]int, s)
q := 0
endIndx := y + s
for i := y; i < endIndx; i++ {
values[q] = history[i][x]
q++
}
fmt.Println(calculateSum(values) - 1)
}

func calculateSum(bits []int) int {
sum := 0
for i := 0; i < len(bits); i++ {
if i == 0 {
sum += (2 << i) * bits[i]
} else {
sum += (2 << (i - 1)) * bits[i]
}
}
return sum
}

func newThreeNeighRule(a, b, c int) int {
if a == 1 && b == 1 && c == 1 {
return ruleset[0]
} else if a == 1 && b == 1 && c == 0 {
return ruleset[1]
} else if a == 1 && b == 0 && c == 1 {
return ruleset[2]
} else if a == 1 && b == 0 && c == 0 {
return ruleset[3]
} else if a == 0 && b == 1 && c == 1 {
return ruleset[4]
} else if a == 0 && b == 1 && c == 0 {
return ruleset[5]
} else if a == 0 && b == 0 && c == 1 {
return ruleset[6]
} else if a == 0 && b == 0 && c == 0 {
return ruleset[7]
}
return 0
}

func newLeftTwoNeighRule(a, b int) int {
if a == 1 && b == 1 {
return 0
} else if a == 1 && b == 0 {
return 1
} else if a == 0 && b == 1 {
return 0
} else if a == 0 && b == 0 {
return 1
}
return 0
}

func newRightTwoNeighRule(a, b int) int {
if a == 1 && b == 1 {
return 1
} else if a == 1 && b == 0 {
return 1
} else if a == 0 && b == 1 {
return 1
} else if a == 0 && b == 0 {
return 1
}
return 0
}

func newGeneration(cells []int) []int {
newcells := make([]int, len(cells))
var left, middle, right int
for i := 0; i < len(cells); i++ {
if i == 0 {
left = cells[i]
middle = cells[i+1]
newcells[i] = newLeftTwoNeighRule(left, middle)
} else if i == len(cells)-1 {
left = cells[i-1]
middle = cells[i]
newcells[i] = newRightTwoNeighRule(left, middle)
} else {
left = cells[i-1]
middle = cells[i]
right = cells[i+1]
newcells[i] = newThreeNeighRule(left, middle, right)
}
}
return newcells
}
