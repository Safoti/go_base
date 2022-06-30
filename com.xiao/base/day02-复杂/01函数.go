package main

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}
func greeting()  {
	println("In greeting: Hi!!!!!")
}

//func (st big.Int) Pop() int {
//	v := 0
//	for ix := len(st) - 1; ix >= 0; ix-- {
//		if v = st[ix]; v != 0 {
//			st[ix] = 0
//			return v
//		}
//	}
//}