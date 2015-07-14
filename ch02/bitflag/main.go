package main

import (
	"fmt"
	"strings"
)

// BitFlag ...
type BitFlag int

// Enumeration ...
const (
	Active  BitFlag = 1 << iota // 1 << 0 == 1
	Send                        // Implicitly BitFlag = 1 << iota  // 1 << 1 == 2
	Receive                     // Implicitly BitFlag = 1 << iota // 1 << 2 == 4
)

func main() {
	flag := Active | Send
	fmt.Println(BitFlag(0), Active, Send, flag, Receive, flag|Receive)
}

func (flag BitFlag) String() string {
	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}

	if len(flags) > 0 { // int(flag) is vital to avoid infinite recursion!
		return fmt.Sprintf("%d(%b)(%s)", int(flag), int(flag), strings.Join(flags, "|"))
	}
	return "0(0)()"
}
