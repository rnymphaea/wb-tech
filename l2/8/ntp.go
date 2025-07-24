package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const (
	UsageExtended = "show extended information from ntp server"

	ErrCode = 1
)

var address string = "0.beevik-ntp.pool.ntp.org"

func main() {
	var extended = flag.Bool("extended", false, UsageExtended)
	flag.BoolVar(extended, "e", false, UsageExtended)
	flag.Parse()

	if !(*extended) {
		currentTime, err := ntp.Time(address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get current time: %v\n", err)
			os.Exit(ErrCode)
		}

		fmt.Printf("Current time is %s\n", currentTime.Format(time.DateTime))
	} else {
		resp, err := ntp.Query(address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get info: %v\n", err)
			os.Exit(ErrCode)
		}

		printNTPResponse(resp)
		printHelp()
	}
}

func printNTPResponse(r *ntp.Response) {
	fmt.Println("Time:", r.Time)
	fmt.Println("ClockOffset:", r.ClockOffset)
	fmt.Println("RTT:", r.RTT)
	fmt.Println("Precision:", r.Precision)
	fmt.Println("Version:", r.Version)
	fmt.Println("Stratum:", r.Stratum)
	fmt.Println("ReferenceID:", r.ReferenceID)
	fmt.Println("ReferenceTime:", r.ReferenceTime)
	fmt.Println("RootDelay:", r.RootDelay)
	fmt.Println("RootDispersion:", r.RootDispersion)
	fmt.Println("RootDistance:", r.RootDistance)
	fmt.Println("Leap:", r.Leap)
	fmt.Println("MinError:", r.MinError)
	fmt.Println("KissCode:", r.KissCode)
	fmt.Println("Poll:", r.Poll)
}

func printHelp() {
	fmt.Println(`
Description:
  Time           Transmit time reported by the server just before it responded to the query
  ClockOffset    Estimated offset of the local system clock relative to the server's clock
  RTT            Measured round-trip-time delay estimate between the client and the server
  Precision      Reported precision of the server's clock
  Version        NTP protocol version number reported by the server
  Stratum        The 'stratum level' of the server. The smaller the number, the closer the server is to the reference clock
  ReferenceID    32-bit integer identifying the server or reference clock
  ReferenceTime  Time when the server's system clock was last set or corrected
  RootDelay      Server's estimated aggregate round-trip-time delay to the stratum 1 server
  RootDispersion Server's estimated maximum measurement error relative to the stratum 1 server
  RootDistance   Estimate of the total synchronization distance between the client and the stratum 1 server
  Leap           Indicate whether a leap second should be added to or removed from the last minute of the current month
  MinError       Lower bound on the error between the client and server clocks
  KissCode       A 4-character string describing the reason for a 'kiss of death' response (stratum=0)
  Poll           Maximum interval between successive NTP query messages to the server`)
}
