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
	// 2. on data channel
	peerConnection.OnDataChannel(func(dataChannel *webrtc.DataChannel) {
		dataChannel.OnOpen(func() {
			i := -1000
			for range time.NewTicker(time.Second * 2).C {
				i++
				if err := dataChannel.SendText("answer channel message " + strconv.Itoa(i)); err != nil {
					log.Println(err.Error())
				}
			}
		})
		dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
			dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
				fmt.Println("targetMessage" + string(msg.Data))
			})
		})
	})
	// 3. input new offer
	println("Input the offer: ")
	var offer webrtc.SessionDescription
	offerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decoder(offerStr, &offer)
	// 4. set the remote description
	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		panic(err)
	}
	// 5. create a new answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}
	// 6. set the local description
	if err := peerConnection.SetLocalDescription(answer); err != nil {
		panic(err)
	}
	// 7. gather the completed
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)
	<-gatherComplete
	// 8. print answer
	fmt.Println("Answer: ")
	fmt.Println(helper.Encoder(peerConnection.LocalDescription()))

	select {}
}
