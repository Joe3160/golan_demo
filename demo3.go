package main

import "fmt"

func main() {
	const ( 						//		iota  说明
		a       = iota             //0   	0
		b                          //1		1     隐式使用iota关键字，实际等同于 b = iota
		c                          //2		2	  实际等同于 c = iota
		d, e, f = iota, iota, iota //3,3,3	3     同一行值相同，此处不能只写一个iota
		g       = iota             //4		4
		h       = "h"              //h  	5     单独赋佳，iota依旧递增为 5
		i                          //h 		6	  默认使用上面的赋值，iota依旧递增为6
		j       = iota             //7		7
	)
	const z = iota //每个单独定义的 const 常量中， iota 都会重豆，此时 z=θ
	fmt.Println(a, b, c, d, e, f, g, h, i, j, z)
}
