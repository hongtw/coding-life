// https://www.codewars.com/kata/513e08acc600c94f01000001
package main

import (
	"fmt"
	"sync"
)

type colorChannels struct {
	channel map[rune]string
	mux     sync.Mutex
}

func (cc *colorChannels) output() string {
	return cc.channel['r'] + cc.channel['g'] + cc.channel['b']
}

func process(color rune, value int, wg *sync.WaitGroup, cc *colorChannels) {
	if value > 255 {
		value = 255
	} else if value < 0 {
		value = 0
	}
	cc.mux.Lock()
	cc.channel[color] = fmt.Sprintf("%02X", value)
	cc.mux.Unlock()
	wg.Done()
}

func rgb(r, g, b int) string {
	cc := &colorChannels{channel: map[rune]string{}}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go process('r', r, wg, cc)
	go process('g', g, wg, cc)
	go process('b', b, wg, cc)
	wg.Wait()
	return cc.output()
}

func main() {
	fmt.Println(rgb(0, 0, -20))     // '000000'
	fmt.Println(rgb(300, 255, 255)) // 'FFFFFF'
	fmt.Println(rgb(173, 255, 47))  // 'ADFF2F'
	fmt.Println(rgb(258, 206, -4))  // 'FFCE00'
	fmt.Println(rgb(9, 165, 222))   // '09A5DE'
	fmt.Println(rgb(256, 69, 277))  // 'FF45FF'
}
