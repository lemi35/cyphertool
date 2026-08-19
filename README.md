Cypher Tool
What This Tool Does
This tool encrypts and decrypts text messages. It runs in a terminal. You pick a cypher. The tool changes your message using that cypher.

This tool has 3 cyphers:

Rot13
Reverse
Supasafe
Requirements
You need Go installed on your computer. Get it here: https://go.dev/doc/install

How to Use It
Open a terminal.
Go to the folder with the code.
Run one of the following two commands:
./cyphertool

OR

go run CypherTool.go

Pick Encrypt or Decrypt. Type 1 or 2. Press Enter.
Pick a cypher. Type 1, 2, or 3. Press Enter.
Type your message. Press Enter.
The result shows on the screen.
Example
This example encrypts "Hello World" with Rot13.

Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
Insert number selection > 1

Select operation (1, 2 or 3):
 1.Rot13.
 2.Reverse.
 3.Supasafe.
Insert number selection > 1

 Enter Message: Hello World

Encryptend with Rot13: Uryyb Jbeyq
To get the message back:

Run the tool again.
Pick Decrypt.
Pick Rot13.
Type: Uryyb Jbeyq
Output: Decrypted with Rot13: Hello World

The Cyphers
Rot13
Rot13 shifts each letter 13 places forward in the alphabet. A becomes N. B becomes O. Numbers and symbols do not change.

Encrypt and decrypt use the same steps. 13 + 13 = 26. 26 is the full length of the alphabet. Shifting twice brings each letter back to the start.

Example: Hello World becomes Uryyb Jbeyq

Reverse
Reverse flips each letter's position in the alphabet. A becomes Z. B becomes Y. Numbers and symbols do not change.

Encrypt and decrypt use the same steps. Flipping a letter twice brings it back to the start.

Example: Hello World becomes Svool Dliow

Supasafe
Supasafe changes each letter in 2 steps:

Flip the letter's position in the alphabet.
Shift the result forward 3 places.
Numbers and symbols do not change.

Encrypt and Decrypt do not however follow the same decryption as their encryption.

For encryption we reverse alphabet then we shift by +3. however in reverse we shift by +23 since subtracting can go into negatives which causes an issue with the modulo %26.

Example: Hello World becomes Vyrro Golrz

