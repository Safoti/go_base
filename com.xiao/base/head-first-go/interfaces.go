package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/25
 * @Description   接口
	Two different types-1 that have the same methods
 **/

type TapePlayer struct {
	Batteries string
}

func (t *TapePlayer)Play(song string)  {
	fmt.Println("playing ",song)
}
func (t *TapePlayer)stop()  {
	fmt.Println("stopped")
}

type TapeRecode struct {
	MicoPhome int
}

func (t *TapeRecode)Play(song string)  {
	fmt.Println("playing ",song)
}
func (t *TapeRecode)stop()  {
	fmt.Println("stopped")
}
func (t *TapeRecode)Recore()  {
	fmt.Println("Recoring")
}

func main() {
	player:=TapePlayer{}
	//player:=TapeRecode{}  这就是写死了
	music:=[]string{"g to h","9 to 5","1 to 3"}
	PlayList(player,music)
}

/**
  定义播放列表
 */

func  PlayList(decvi TapePlayer,songs[]string)  {
//func  PlayList(decvi TapeRecode,songs[]string)  { //方法重新定义，，有点烦
	for _, song := range songs {
		decvi.Play(song)
	}
	decvi.stop()
}
