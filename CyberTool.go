// Main logic, envoking other functions
package main

import (
	"bufio"
	"fmt"
	"os"
	//	"strconv"
	//"strings"
)
func main() { 

	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Print("\n")

	toEncrypt, encoding, message := getInput()
	
	if toEncrypt == true {
		switch encoding {
		case "Rot13":
			message = encrypt_rot13(message)
		case "Reverse":
			message = encrypt_reverse(message)
		case "Supasafe":
			message = SupaSafe(message)
		}
		fmt.Println("Encryptend with " + encoding + ": " + message)
	} else if toEncrypt == false {
		switch encoding {
			case "Rot13":
				message = decrypt_rot13(message)
			case "Reverse":
				message = decrypt_reverse(message)
			case "Supasafe":
				message = SupaSafeDecrypt(message)
		}
		fmt.Println("Decrypted with " + encoding + ": " + message)
	}
	return 

}

// Get the input data required for the operation
func getInput() (bool, string, string) {
	var toEncrypt bool
	var encoding string
	var message string
	
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.\nInsert number selection > ")
		input.Scan()

		switch input.Text() {
		case "1":
			toEncrypt = true
		case "2":
			toEncrypt = false
		default:
			fmt.Println("\nPlease select a valid input!")
			fmt.Printf("%s is not a valid input\n", input.Text())
			continue
		}

		break
	}

	for {
		fmt.Print("\nSelect operation (1, 2 or 3):\n 1.Rot13. \n 2.Reverse. \n 3.Supasafe. \nInsert number selection >")
		input.Scan()

		switch input.Text() {
		case "1":
			encoding = "Rot13"
		case "2":
			encoding = "Reverse"
		case "3":
			encoding = "Supasafe"
		default:
			fmt.Println("\nPlease select a valid input!")
			fmt.Printf("%s is not a valid input\n", input.Text())
			continue
		}
		break	
	}
	
	for {
		fmt.Print("\n Enter Message: ")
		input.Scan()
		if input.Text() == "" {
			fmt.Println("Please insert a message and not an empty one: ")
			continue
		}
		message = input.Text()
		break
	}
	return toEncrypt, encoding, message
}

// Encrypt the message with rot13
func Shifting(r rune, shift int, ver int) rune{
	if ver == 1 { 
		a := r - 'a'
        	b := ((a + rune(shift)) % 26) + 'a'
        	return b
	}
	if ver == 2 {
		c := r - 'A'
		d := ((c + rune(shift)) % 26 ) + 'A'
		return d
	}
	return r
}

func encrypt_rot13(s string) string {
	str := ""

        for i := 0; i < len(s); i++ {
                if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
                        str += string(Shifting(rune(s[i]), 13, 1))
                } else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
                        str += string(Shifting(rune(s[i]), 13, 2))
                } else {
                        str += string(s[i])
                }
        }
	return str
}

func    ReverseAlphabetValue(r rune, ver int) rune {
        if ver == 1 {
		a := r - 'a'
        	b := 25 - a
        	c := b + 'a'
        	return c
	}
	if ver == 2 {
		d := r - 'A'
		e := 25 - d
		f := e + 'A'
		return f
	}
	return 0
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
			str += string(ReverseAlphabetValue(rune(s[i]), 1))
		} else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			str += string(ReverseAlphabetValue(rune(s[i]), 2))
		} else {
			str += string(s[i])
		}
	}
	return str
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
 str := ""

        for i := 0; i < len(s); i++ {
                if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
                        str += string(Shifting(rune(s[i]), 13, 1))
                } else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
                        str += string(Shifting(rune(s[i]), 13, 2))
                } else {
                        str += string(s[i])
                }
        }
        return str
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
	 str := ""
        for i := 0; i < len(s); i++ {
                if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
                        str += string(ReverseAlphabetValue(rune(s[i]), 1))
                } else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
                        str += string(ReverseAlphabetValue(rune(s[i]), 2))
                } else {
                        str += string(s[i])
                }
        }
	return str
}

func SupaSafe(s string) string {
	str := ""
	for i:= 0; i < len(s); i++ {
		if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
			i := ((25 - (rune(s[i]) - 'a') ) + 'a')
			b := ((((i -'a') + 3 ) % 26 ) + 'a')
			str += string(b)
		} else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			j := ((25 - (rune(s[i]) - 'A') ) + 'A')
			k := ((((j -'A') + 3 ) % 26 ) + 'A')
			str += string(k)
		} else {
                        str += string(s[i])
                }
        }
	return str
}

func SupaSafeDecrypt (s string) string {
	str := ""
        for i:= 0; i < len(s); i++ {
                if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
                        i := (((((rune(s[i])) -'a') + 23 ) % 26 ) + 'a')
                        b := ((25 - (i - 'a') ) + 'a')
                        str += string(b)
                } else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
                        j := (((((rune(s[i])) -'A') + 23 ) % 26 ) + 'A')
                        k :=((25 - (j - 'A') ) + 'A')
                        str += string(k)
                } else {
                        str += string(s[i])
                }
        }
        return str
}

