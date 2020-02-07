# capture
Parse regular expression named capturing groups into structs

## Installation

```
go get github.com/enrico5b1b4/capture
```

## Usage

```go
type reminder struct {
    Who     string `regexpGroup:"who"`
    Day     int    `regexpGroup:"day"`
    Month   string `regexpGroup:"month"`
    Year    int    `regexpGroup:"year"`
    Message string `regexpGroup:"message"`
}

myReminder := &reminder{}
err := capture.Parse(
    `remind (?P<who>\w+) on the (?P<day>\d{1,2})(?:(st|nd|rd|th))? of (?P<month>october|november|december) (?P<year>\d{4}) to (?P<message>.*)`,
    "remind John on the 31st of october 2030 to buy milk",
    myReminder,
)

fmt.Println(myReminder.Who) // John
fmt.Println(myReminder.Day) // 31
fmt.Println(myReminder.Month) // october
fmt.Println(myReminder.Year) // 2030
fmt.Println(myReminder.Message) // buy milk
```