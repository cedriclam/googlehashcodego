package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"bufio"
	"container/list"
	"github.com/cedriclam/googlehashcodego/src/utils"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	bio := bufio.NewReader(os.Stdin)
	
	var line []byte
	hasMoreInLine := true
	var err error = nil

	for hasMoreInLine  {
		line, hasMoreInLine, err = bio.ReadLine()
		if err == nil {
			fmt.Println(line)
		}
	}
		
	utils.GetInput("foo")

}

func ReadFromStdinToBytes() ([]byte, error){
	bytes, err := ioutil.ReadAll(os.Stdin)
	return bytes, err;
}

func ReadFromStdinToLines(iList *list.List) {
	reader := bufio.NewReader(os.Stdin)

    for {
        line, err := reader.ReadString('\n')

        if err != nil {
            // You may check here if err == io.EOF
            break
        }

		iList.PushBack(line)
        log.Println(line)
    }
}