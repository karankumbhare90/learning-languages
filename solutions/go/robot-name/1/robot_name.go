package robotname

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Exercise defines maxNames in the test, so we must use a different name.
const totalPossibleNames = 26 * 26 * 1000 // 676,000 possible names

type Robot struct {
	name string
}

var (
	usedNames = make(map[string]bool, totalPossibleNames)
	mu        sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	newName, err := generateUniqueName()
	if err != nil {
		return "", err
	}

	r.name = newName
	return r.name, nil
}

func (r *Robot) Reset() {
	mu.Lock()
	defer mu.Unlock()

	delete(usedNames, r.name)
	r.name = ""
}

func generateUniqueName() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	if len(usedNames) >= totalPossibleNames {
		return "", errors.New("no available robot names")
	}

	for {
		name := randomName()
		if !usedNames[name] {
			usedNames[name] = true
			return name, nil
		}
	}
}

func randomName() string {
	return string([]byte{
		'A' + byte(rand.Intn(26)),
		'A' + byte(rand.Intn(26)),
		'0' + byte(rand.Intn(10)),
		'0' + byte(rand.Intn(10)),
		'0' + byte(rand.Intn(10)),
	})
}
