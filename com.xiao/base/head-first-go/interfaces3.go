package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/25
 * @Description   接口
	Two different types-1 that have the same methods
 **/
type Players interface {
	Play( string)
	Stop()
}

type TapePlayers struct {
	Batteries string
}



func (t *TapePlayers)Play(song string)  {
	fmt.Println("playing ",song)
}
func (t *TapePlayers)Stop()  {
	fmt.Println("stopped")
}

type TapeRecodes struct {
	MicoPhome int
}

func (t *TapeRecodes)Play(song string)  {
	fmt.Println("playing ",song)
}
func (t *TapeRecodes)Stop()  {
	fmt.Println("stopped")
}
func (t *TapeRecodes)Recore()  {
	fmt.Println("Recoring")
}

func main() {
	music:=[]string{"g to h","9 to 5","1 to 3"}
    var player Players=&TapeRecodes{}

	PlayLists(player,music)

    player=&TapePlayers{}
	PlayLists(player,music)
}

/**
  定义播放列表
 */

func  PlayLists(decvi Players,songs[]string)  {
//func  PlayList(decvi TapeRecode,songs[]string)  { //方法重新定义，，有点烦
	for _, song := range songs {
		decvi.Play(song)
	}
	decvi.Stop()

}
