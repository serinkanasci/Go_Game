package main

import (
	"math/rand"
	"time"
	"fmt"
	keyboard "github.com/eiannone/keyboard"
)

var (myMessage string)

func main() {
	player1 := newPlayer("Hamidou")
	player2 := newPlayer("Heldou")
	compt := 1

	for compt != 0 {
		if (compt % 2 == 0){
			println("TURN OF : " , player2.name)
			char, _, err := keyboard.GetSingleKey()
			if (err != nil) {
				panic(err)
			}
			switch char {
				case 'z':
					player2.moveVertical(1)
				case 's':
					player2.moveVertical(-1)
				case 'd':
					player2.moveHorinzontal(-1)
				case 'q':
					player2.moveHorinzontal(1)
				case 'f':
					player2.attack(player1)
				default:
					fmt.Printf("FAIL CLICK")
			}
		}else{
			println("TURN OF : " , player1.name)
			char, _, err := keyboard.GetSingleKey()
			if (err != nil) {
				panic(err)
			}
			switch char {
				case 'z':
					player1.moveVertical(1)
				case 's':
					player1.moveVertical(-1)
				case 'd':
					player1.moveHorinzontal(-1)
				case 'q':
					player1.moveHorinzontal(1)
				case 'f':
					player1.attack(player2)
				default:
					fmt.Printf("FAIL CLICK")
			}
		}
		if player1.health <= 0 {
			println(player2.name, " WON")
			compt = 0
			break
		}
		if player2.health <= 0 {
			println(player1.name ," WON")
			compt = 0
			break
		}
		compt += 1 
	}
}
type position struct {
	x int
	y int
}
type player struct {
	name string
	pos position
	health int
}

func newPlayer(myName string) *player{
	p := player{}
	p.name = myName

	rand.Seed(time.Now().UnixNano())
	posX := -10 + rand.Intn(10-(-10)+1)

	rand.Seed(time.Now().UnixNano())
	posY := -10 + rand.Intn(10-(-10)+1)

    p.pos = newPosition(posX, posY)
	p.health = 100
    return &p
}


func newPosition(posX int, posY int) position{
    pos := position{}
    pos.x = posX
	pos.y = posY
    return pos
}

func (attacker *player) attack(attacked *player){
	myDamage := GenerateDamage()
	attacked.health -= myDamage
	println(attacker.name, " picked a damage of ", myDamage)
	println(attacked.name, " has been attacked, remaining life : ", attacked.health)
}

func (moved *player) moveHorinzontal(move int)  {
	moved.pos.x += (move)
	println(moved.name, " moved to  ", moved.pos.x)
}

func (moved *player) moveVertical(move int)  {
	moved.pos.y += (move)
	println(moved.name, " moved to  ", moved.pos.y)
}

func GenerateDamage() int{
	rand.Seed(time.Now().UnixNano())
	damage := 1 + rand.Intn(10-1+1)
	return damage
}



// go get -u github.com/eiannone/keyboard