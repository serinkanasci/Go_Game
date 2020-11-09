package main

import (
	"math/rand"
	"time"
    "fmt"
	keyboard "github.com/eiannone/keyboard"
)

var (myMessage string)

func main() {
	println("HELLO PLAYERS !!")
	
	println("ENTER PLAYER 1 : ")
	var player1Name string
	fmt.Scanln(&player1Name)

	println("ENTER PLAYER 2 : ")
	var player2Name string
	fmt.Scanln(&player2Name)

	player1 := newPlayer(player1Name)
	player2 := newPlayer(player2Name)
	compt := 1

	for compt != 0 {
		if (compt % 2 == 0){
			println("TURN OF : " , player2.name)
			println("POS X : ", player2.pos.x , "POS Y : ", player2.pos.y)
			println(GenerateMap(player1, player2))
			char, _, err := keyboard.GetSingleKey()
			if (err != nil) {
				panic(err)
			}
			switch char {
				case 'z':
					player2.moveVertical(-1)
				case 's':
					player2.moveVertical(1)
				case 'd':
					player2.moveHorinzontal(1)
				case 'q':
					player2.moveHorinzontal(-1)
				case 'f':
					player2.attack(player1)
				default:
					fmt.Printf("FAIL CLICK")
			}
		}else{
			println("TURN OF : " , player1.name)
			println("POS X : ", player1.pos.x , "POS Y : ", player1.pos.y)
			println(GenerateMap(player2, player1))
			char, _, err := keyboard.GetSingleKey()
			if (err != nil) {
				panic(err)
			}
			switch char {
				case 'z':
					player1.moveVertical(-1)
				case 's':
					player1.moveVertical(1)
				case 'd':
					player1.moveHorinzontal(1)
				case 'q':
					player1.moveHorinzontal(-1)
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
	posX := -5 + rand.Intn(5-(-5)+1)

	rand.Seed(time.Now().UnixNano())
	posY := -2 + rand.Intn(2-(-2)+1)

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
	diffX := attacker.pos.x - attacked.pos.x
	diffY := attacker.pos.y - attacked.pos.y
	if(diffX >= -1 && diffX <= 1 && diffY >= -1 && diffY <= 1) {
		attacked.health -= myDamage
		println(attacker.name, " picked a damage of ", myDamage)
		println(attacked.name, " has been attacked, remaining life : ", attacked.health)
	}else{
		println(attacker.name, " cannot attack ", attacked.name, " is too far")
	}
}

func (moved *player) moveHorinzontal(move int)  {
	if move == 1 {
		if moved.pos.x == 5{
			println("Can't go any more to the right ... ")
		}else{
			moved.pos.x += (move)
		}
	}else{
		if moved.pos.x == -5{
			println("Can't go any more to the left ... ")
		}else{
			moved.pos.x += (move)
		}
	}
	
}

func (moved *player) moveVertical(move int)  {
	if move == 1 {
		if moved.pos.y == 2{
			println("Can't go any lower ... ")
		}else{
			moved.pos.y += (move)
			println(moved.name, " moved to  ", moved.pos.y)
		}
	}else{
		if moved.pos.y == -2{
			println("Can't go any higher ... ")
		}else{
			moved.pos.y += (move)
			println(moved.name, " moved to  ", moved.pos.y)
		}
	}
}

func GenerateDamage() int{
	rand.Seed(time.Now().UnixNano())
	damage := 1 + rand.Intn(10-1+1)
	return damage
}

func GenerateMap(attacked *player, attacker *player) string{
	startX := -5
	startY := -2
	genMap := "_______________________________________ \n"
	i := 0
	for i < 60 {
		if attacked.pos.x == startX && attacked.pos.y == startY || attacker.pos.x == startX && attacker.pos.y == startY {
			if attacked.pos.x == startX && attacked.pos.y == startY {
				genMap += " " + attacked.name + " "
				if startX == 6{
					startX = -5
					startY ++
				}else{
					startX ++
				}
			}else{
				genMap += " " + attacker.name + " "
				if startX == 6{
					startX = -5
					startY ++
				}else{
					startX ++
				}
			}
		}else{
			if startX == 6 {
				startX = -5
				startY ++
				genMap += "\n"
				
			}else{
				startX ++
				genMap += " _ "
			}
		}
		i++
	}
	genMap += "\n_______________________________________"
	return genMap
}



// go get -u github.com/eiannone/keyboard