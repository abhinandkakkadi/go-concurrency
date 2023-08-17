package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	clientsChan     chan string // pass clientName to it
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {

	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to waiting room to check for clients.",barber)
		for {
			// if there are no clients, the barber goes to sleep
			// since it is the buffered channel length is same as Number of clients as we have defined
			if len(shop.clientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes  nap.",barber)
				isSleeping = true
			} 

			// if shopOpen is false that means, the channel is closed
			client,shopOpen := <- shop.clientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.",client,barber)
					isSleeping = false
				}

				// cut hair
				shop.cutHair(barber,client)
			} else {
				// shop is closed, so send the barber home and close this go routine
			}
		}
	}()

}

func (shop *BarberShop) cutHair(barber, client string) {

}