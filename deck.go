package main

import (
	"fmt"
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
		// option #1 - log the error and return call to newDeck() => 在這裡不合用，因為在從新拿取card 會造成 有重複的牌
		// option #2 - log the error and entirely quit the program
		fmt.Println("err: ", err)
		os.Exit(1) // quit program
	}
	// string(bs) // ce of Spades,Two of Spades,Three of Spades....
	s := strings.Split(string(bs), ",")
	return deck(s)
}
