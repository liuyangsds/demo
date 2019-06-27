package main

func main()  {

	//fmt.Print("aaa");

	//var total int = 0
	//total := max(6,8)
	//print(total)

	//fmt.Print("hello world")

	var pp Person
	pp.age = 23
	pp.name = "yang"
	pp.sex = "n"
	pp.height = 175.05
	pp.weight = 136.52

	//print(pp.name)

	var(
		i int = 0
		j int = 1
	)
	for i = 1; i <= 9; i++  {
		for j = 1; j <= i; j++ {
			print(j, " * ", i, " = ", j*i,"\t")
		}
		println()
	}
}

func max(x int, y int) int {

	if(x > y) {
		return x
	}
	return y

}

type Person struct {
	name string
	age int
	sex string
	height float64
	weight float64
}