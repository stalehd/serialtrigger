package main
//
//Copyright 2018 Staale Holberg Dahl
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
import (
	"flag"
	"fmt"
	"strings"

	"github.com/jacobsa/go-serial/serial"
)

var port = "/dev/cu.usbmodem14221"
var passphrase = "RELEASE THE KRAKEN!"

func init() {
	flag.StringVar(&port, "port", port, "Port to listen on")
	flag.StringVar(&passphrase, "phrase", passphrase, "Phrase to listen for")
}
func main() {
	options := serial.OpenOptions{
		PortName:        port,
		BaudRate:        9600,
		StopBits:        1,
		DataBits:        8,
		ParityMode:      serial.PARITY_NONE,
		MinimumReadSize: 1,
	}

	reader, err := serial.Open(options)
	if err != nil {
		fmt.Println("got error opening serial port: ", err)
		return
	}

	defer reader.Close()

	readBuffer := make([]byte, 1024)
	pos := 0
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println("Got error reading from serial port: ", err)
			return
		}
		if n > 0 {
			copy(readBuffer[pos:pos+n], buf[0:n])
			pos += n
			if strings.Index(string(readBuffer[0:pos]), passphrase) >= 0 {
				return
			}
		}
	}
}
