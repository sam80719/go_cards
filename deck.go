package main

import (
	"fmt"
	"math/rand"
	"time"

	//"io/ioutil"
	"os"
	"strings" // 引用 pkg strings
)

// creat a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "diamonds", "Hear", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// 與上面的迴圈相同，指示index的i與j沒使用到的話，go會噴錯誤，所以在這裡使用底線，避免噴錯誤
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // 匿名函數
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666) // 系統權限表示任何人都可以讀
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil { // handle error
		fmt.Println("err: ", err)
		os.Exit(1) // quit program
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//// 教程的程式碼
	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)
	//for i := range d {
	//	newPosition := r.Intn(len(d) - 1)
	//	d[i], d[newPosition] = d[newPosition], d[i]
	//}

	// 較新的版本有rand,Seed
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // swap
	}
}
