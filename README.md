More about magic numbers [here](https://en.wikipedia.org/wiki/Magic_number_(programming)#Magic_numbers_in_files) .  

### \# package
```
$ go get -u github.com/michalswi/magic-numbers/magic
```

### \# example
```
import (
	"fmt"
	"log"

	"github.com/michalswi/magic-numbers/magic"
)

func main() {
	magic, err := magic.MagicNumbers("/bin/bash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(magic)
}
```