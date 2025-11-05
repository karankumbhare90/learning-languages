package twelve

import "strings"

var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(i int) string {
	verse := "On the " + days[i-1] + " day of Christmas my true love gave to me: "
	var parts []string
	for j := i; j >= 1; j-- {
		if j == 1 && i > 1 {
			parts = append(parts, "and "+gifts[0])
		} else {
			parts = append(parts, gifts[j-1])
		}
	}
	verse += strings.Join(parts, ", ") + "."
	return verse
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	// ✅ Only one newline between verses
	return strings.Join(verses, "\n")
}
