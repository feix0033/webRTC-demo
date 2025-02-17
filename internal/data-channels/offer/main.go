package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"webRTC-demo/internal/helper"

	"github.com/pion/webrtc/v3"
)

func main() {
	// 1. create a new peer connection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Println(err)
		}
	}()
	// 2. create a new data channel
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		return
	}
	dataChannel.OnOpen(func() {
		log.Println("data channel open")
		/* send message every 5 seconds */
		i := -1000
		for range time.NewTicker(time.Second * 2).C {
			i++
			if err := dataChannel.SendText("offer channel message " + strconv.Itoa(i)); err != nil {
				log.Println(err.Error())
			}
		}
	})
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		/*  when the message received, print the message */
		fmt.Println(string(msg.Data))
	})
	// 3. create a new offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		return
	}
	// 4. set the local description
	if err := peerConnection.SetLocalDescription(offer); err != nil {
		return
	}
	// 5. print the offer
	println("Offer: ")
	println(helper.Encoder(offer))
	// 6. input the answer
	println("Input the answer: ")
	var answer webrtc.SessionDescription
	answerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decoder(answerStr, &answer)
	// 7. set the remote description
	if err := peerConnection.SetRemoteDescription(answer); err != nil {
		return
	}
	// 8. block the main thread to wait for the data channel to be opened
	select {}
}
