package main

import(
	"fmt"
	"os"
	"io/ioutil"
)

const(
	GT = '>'
	LT = '<'
	PLUS = '+'
	MINUS = '-'
	POINT = '.'
	COMMA = ','
	LBRACKET = '['
	RBRACKET = ']'
	EOF = '0'
)

type Parser struct{
	bfcode string
	currpos int
}

func newParser(bfcode string)(theParser Parser){
	theParser.bfcode = bfcode
	theParser.currpos = 0
	return
}

func (theParser *Parser)nextToken()(currToken byte){
	if theParser.currpos == len(theParser.bfcode){
		currToken = EOF
	}else{
		currToken = theParser.bfcode[theParser.currpos]
		theParser.currpos += 1
	}
	return
}

func (theParser *Parser)backToken()(currToken byte){
	theParser.currpos -= 1
	currToken = theParser.bfcode[theParser.currpos]
	return
}

type Space struct{
	theSpace [1000000]byte
	currPoint int
}

func main(){
	if len(os.Args) == 1{
		fmt.Println("Usage:brainfuck filename.bf")
	}else if len(os.Args) == 2{
		f, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("read fail", err)
		}
		theParser := newParser(string(f))
		var thespace Space
		thespace.currPoint = 500000
		for{
			theToken := theParser.nextToken()
			switch theToken{
			case GT:
				thespace.currPoint += 1
			case LT:
				thespace.currPoint -= 1
			case PLUS:
				thespace.theSpace[thespace.currPoint] += 1
			case MINUS:
				thespace.theSpace[thespace.currPoint] -= 1
			case POINT:
				fmt.Print(string(thespace.theSpace[thespace.currPoint]))
			case COMMA:
				fmt.Scanf("%c",&thespace.theSpace[thespace.currPoint])
			case LBRACKET:
				if thespace.theSpace[thespace.currPoint] == 0{
					bracketStack := 0
					for{
						tempToken := theParser.nextToken()
						if tempToken == LBRACKET{
							bracketStack += 1
						}else if tempToken == RBRACKET{
							if bracketStack == 0{
								break
							}
							bracketStack -= 1
						}
					}
				}
			case RBRACKET:
				if thespace.theSpace[thespace.currPoint] != 0{
					bracketStack := 0
					theParser.backToken()
					for{
						tempToken := theParser.backToken()
						if tempToken == RBRACKET{
							bracketStack += 1
						}else if tempToken == LBRACKET{
							if bracketStack == 0{
								theParser.nextToken()
								break
							}
							bracketStack -= 1
						}
					}
				}
			default:
			}
			if theToken == EOF{
				break
			}
		}
	}else{
		panic("Too many args.")
	}
}