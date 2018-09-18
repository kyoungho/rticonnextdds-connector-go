/*****************************************************************************
*    (c) 2005-2015 Copyright, Real-Time Innovations, All rights reserved.    *
*                                                                            *
*  RTI grants Licensee a license to use, modify, compile, and create         *
*  derivative works of the Software.  Licensee has the right to distribute   *
*  object form only for use with RTI products. The Software is provided      *
*  "as is", with no warranty of any type, including any warranty for fitness *
*  for any purpose. RTI is under no obligation to maintain or support the    *
*  Software.  RTI shall not be liable for any incidental or consequential    *
*  damages arising out of the use or inability to use the software.          *
*                                                                            *
*****************************************************************************/

package main

import (
	"github.com/rticommunity/rticonnextdds-connector-go"
	"log"
	"path"
	"runtime"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Find the file path to the XML configuration
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("runtime.Caller error")
	}
	filepath := path.Join(path.Dir(filename), "../ShapeExample.xml")

	// Create a channel to receive signals from OS
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Create a connector defined in the XML configuration
	connector, err := rti.NewConnector("MyParticipantLibrary::Zero", filepath)
	if err != nil {
		log.Panic(err)
	}
	// Delete the connector when this main function returns
	defer connector.Delete()

	// Get an input from the connector
	input, err := connector.GetInput("MySubscriber::MySquareReader")
	if err != nil {
		log.Panic(err)
	}

	run := true

	// Get values from a received sample and print them
	for run == true {
		select {
		case sig := <-sigchan:
			log.Printf("Received signal %v: terminating\n", sig)
			run = false
		default:
			input.Take()
			numOfSamples := input.Samples.GetLength()
			for j := 0; j < numOfSamples; j++ {
				if input.Infos.IsValid(j) {
					color := input.Samples.GetString(j, "color")
					x := input.Samples.GetInt(j, "x")
					y := input.Samples.GetInt(j, "y")
					shapesize := input.Samples.GetInt(j, "shapesize")

					log.Println("---Received Sample---")
					log.Printf("color: %s\n", color)
					log.Printf("x: %d\n", x)
					log.Printf("y: %d\n", y)
					log.Printf("shapesize: %d\n", shapesize)
				}
			}
			time.Sleep(time.Second * 1)
		}
	}
}
