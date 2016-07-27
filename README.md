# SKK dictionary reader

Example

```go
import skkdict "github.com/koron/go-skkdict"

f, err := os.Open("SKK-JISYO.utf-8.S")
if err != nil {
	panic(err)
}
defer f.Close()

r := skkdict.NewReader(f)
for {
	entry, err := r.Read()
	if err != nil {
		break
	}
	fmt.Printf("%+v\n", entry)
}
```
