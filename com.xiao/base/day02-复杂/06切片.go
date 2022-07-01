package main

import "fmt"

/**
动态数组
 */
func main() {
	var arr2 [6]int
	var sli1 []int=arr2[2:5]
	for i := 0; i <len(arr2) ; i++ {
		arr2[i]=i
	}

	for i := 0; i <len(sli1) ; i++ {
		fmt.Printf("Slice at %d is %d\n", i, sli1[i])
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr2))
	fmt.Printf("The length of slice1 is %d\n", len(sli1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(sli1))

	// grow the slice
	sli1 = sli1[0:4]

	for i := 0; i < len(sli1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, sli1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(sli1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(sli1))



  //切片创建make

   s:= make([] int ,3)
	fmt.Println("emp:", s)

   s[0]=1
   s[1]=2
   s[2]=3

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
   fmt.Println("len",len(s))

	 s=  append(s,4)
	s = append(s, 5, 6)
	fmt.Println("apd:", s)

    //创建切片的长度为50 ，容量为500
	y:=	make([]int,50,500)
	fmt.Println(y)
	//根据切片C创建
	c:=make([]int,len(s))
    fmt.Println(c)
	copy(c,s)
	fmt.Println(c)


	l:=s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[:4]
	fmt.Println("sl3:", l)
	l = s[2:]
	fmt.Println("sl4:", l)

	is:=make([]int,5 )
	for i := 0; i < 5; i++ {
		is[i]=i
	}
	fmt.Println(is)
	lis:=is[2:]
	fmt.Println(lis)
	lis = is[2:3]
	fmt.Println(lis)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)


	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}

	seasons:=[]string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	for i := range seasons {
		fmt.Printf("%d", i)
	}


}
