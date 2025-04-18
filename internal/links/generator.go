package links

import (
	"errors"
	"math/rand"
)

const MAX_LENGTH = 3_000

type Generator struct {
	hashOptions []int
}

func getHashOptions() []int {
	results := make([]int, 64)
	results[0] = 36 // $
	results[1] = 42 // *
	index := 2

	//numbers
	for i := 0; i < 10; i++ {
		results[index] = 48 + i
		index++
	}
	//capital letters
	for i := 0; i < 26; i++ {
		results[index] = 65 + i
		index++
	}
	//lower case letters
	for i := 0; i < 26; i++ {
		results[index] = 97 + i
		index++
	}
	return results
}

func NewGenerator() *Generator {
	return &Generator{
		hashOptions: getHashOptions(),
	}
}

/*
* Characters that can make up links are as follows: [A-Za-z0-9$*]
* This means that we have 26+26+10+2=64 options to create our characters
* How long can we make this that it's still pretty short, but also that it's going to be
* hard to accidentally duplicate?
* x^64 options. 10 characters long is 10^64. That's like a decillion, squared.
* We'll never reach that amount, we just need to have something that's deterministic and creates
* fast redirects for users
*
* 36, 42, 48-57, 65-90, 97-122
 */
func (g *Generator) GeneratePath(original string) (string, error) {
	if len(original) >= MAX_LENGTH {
		return "", errors.New("link must have fewer than 3000 characters")
	}
	hash := ""
	for i := 0; i < 10; i++ {
		index := rand.Intn(len(g.hashOptions))
		hash += string(rune(g.hashOptions[index]))
	}
	return hash, nil
}
