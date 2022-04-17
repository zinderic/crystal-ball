package orb

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ResetColor = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Purple     = "\033[35m"
	Cyan       = "\033[36m"
	Gray       = "\033[37m"
	White      = "\033[97m"
)

const Orb = `
         _...._
       .'      '.
      / ***      \
     : **         :
     :            :
      \          /
       '-.,,,,.-'
        _(    )_
       )        (
      (          )
       '-......-'
`

func ReadInput() (string, error) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print(Red + "🔮 > ")
	_, err := fmt.Fscan(reader)
	if err != nil {
		return "", err
	}
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	fmt.Println(ResetColor)
	return text, nil
}
