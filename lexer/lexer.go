package lexer

type Lexer struct {
	input        string // stream of input as a string
	position     int    // current position in the input
	readPosition int    // current reading position in the input (meaning char after the current char; char + 1)
	ch           byte   // current character under examination
}

// New A kind of constructor for the Lexer
func New(input string) *Lexer {
	return &Lexer{input: input}
}

// readChar reads the character and moves forward
func (lexer *Lexer) readChar() {
	if lexer.readPosition > len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}
