// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var golitelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golitelexerLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'+'", "'-'", "'='", "'*'", "'/'", "'print'", "'let'", "'type'",
		"'struct'", "'int'", "'bool'", "'var'", "'func'", "'if'", "'else'",
		"'for'", "'return'", "'scan'", "'delete'", "'printf'", "'new'", "'true'",
		"'false'", "'nil'", "'{'", "'}'", "'('", "')'", "'.'", "','", "'!'",
		"'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "", "';'",
	}
	staticData.symbolicNames = []string{
		"", "NUMBER", "PLUS", "MINUS", "ASSIGN", "ASTERISK", "FSLASH", "PRINT",
		"LET", "TYPE", "STRUCT", "INT_TYPE", "BOOL_TYPE", "VAR", "FUNC", "IF",
		"ELSE", "FOR", "RETURN", "SCAN", "DELETE", "PRINTF", "NEW", "TRUE",
		"FALSE", "NIL", "LEFTCURL", "RIGHTCURL", "LEFTBRAC", "RIGHTBRAC", "DOT",
		"COMMA", "EXCLAMATION", "OR", "AND", "EQUALS", "NOTEQUALS", "MORETHAN",
		"LESSTHAN", "GEQ", "LEQ", "ID", "SEMICOLON", "COMMENT", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"NUMBER", "PLUS", "MINUS", "ASSIGN", "ASTERISK", "FSLASH", "PRINT",
		"LET", "TYPE", "STRUCT", "INT_TYPE", "BOOL_TYPE", "VAR", "FUNC", "IF",
		"ELSE", "FOR", "RETURN", "SCAN", "DELETE", "PRINTF", "NEW", "TRUE",
		"FALSE", "NIL", "LEFTCURL", "RIGHTCURL", "LEFTBRAC", "RIGHTBRAC", "DOT",
		"COMMA", "EXCLAMATION", "OR", "AND", "EQUALS", "NOTEQUALS", "MORETHAN",
		"LESSTHAN", "GEQ", "LEQ", "ID", "SEMICOLON", "COMMENT", "STRING", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 282, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 5, 0, 94, 8,
		0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 3, 0, 100, 8, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 5, 40, 247, 8, 40, 10, 40, 12, 40, 250, 9, 40, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 5, 42, 260, 8, 42, 10, 42, 12,
		42, 263, 9, 42, 1, 42, 1, 42, 1, 43, 1, 43, 5, 43, 269, 8, 43, 10, 43,
		12, 43, 272, 9, 43, 1, 43, 1, 43, 1, 44, 4, 44, 277, 8, 44, 11, 44, 12,
		44, 278, 1, 44, 1, 44, 1, 270, 0, 45, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42,
		85, 43, 87, 44, 89, 45, 1, 0, 7, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 65,
		90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 2, 0, 34, 34, 92, 92, 4, 0,
		10, 10, 13, 13, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 288, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 99, 1, 0, 0, 0, 3, 101,
		1, 0, 0, 0, 5, 103, 1, 0, 0, 0, 7, 105, 1, 0, 0, 0, 9, 107, 1, 0, 0, 0,
		11, 109, 1, 0, 0, 0, 13, 111, 1, 0, 0, 0, 15, 117, 1, 0, 0, 0, 17, 121,
		1, 0, 0, 0, 19, 126, 1, 0, 0, 0, 21, 133, 1, 0, 0, 0, 23, 137, 1, 0, 0,
		0, 25, 142, 1, 0, 0, 0, 27, 146, 1, 0, 0, 0, 29, 151, 1, 0, 0, 0, 31, 154,
		1, 0, 0, 0, 33, 159, 1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 37, 170, 1, 0, 0,
		0, 39, 175, 1, 0, 0, 0, 41, 182, 1, 0, 0, 0, 43, 189, 1, 0, 0, 0, 45, 193,
		1, 0, 0, 0, 47, 198, 1, 0, 0, 0, 49, 204, 1, 0, 0, 0, 51, 208, 1, 0, 0,
		0, 53, 210, 1, 0, 0, 0, 55, 212, 1, 0, 0, 0, 57, 214, 1, 0, 0, 0, 59, 216,
		1, 0, 0, 0, 61, 218, 1, 0, 0, 0, 63, 220, 1, 0, 0, 0, 65, 222, 1, 0, 0,
		0, 67, 225, 1, 0, 0, 0, 69, 228, 1, 0, 0, 0, 71, 231, 1, 0, 0, 0, 73, 234,
		1, 0, 0, 0, 75, 236, 1, 0, 0, 0, 77, 238, 1, 0, 0, 0, 79, 241, 1, 0, 0,
		0, 81, 244, 1, 0, 0, 0, 83, 251, 1, 0, 0, 0, 85, 253, 1, 0, 0, 0, 87, 266,
		1, 0, 0, 0, 89, 276, 1, 0, 0, 0, 91, 95, 7, 0, 0, 0, 92, 94, 7, 1, 0, 0,
		93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1,
		0, 0, 0, 96, 100, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 100, 7, 1, 0, 0,
		99, 91, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 2, 1, 0, 0, 0, 101, 102, 5,
		43, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 45, 0, 0, 104, 6, 1, 0, 0, 0,
		105, 106, 5, 61, 0, 0, 106, 8, 1, 0, 0, 0, 107, 108, 5, 42, 0, 0, 108,
		10, 1, 0, 0, 0, 109, 110, 5, 47, 0, 0, 110, 12, 1, 0, 0, 0, 111, 112, 5,
		112, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 105, 0, 0, 114, 115, 5,
		110, 0, 0, 115, 116, 5, 116, 0, 0, 116, 14, 1, 0, 0, 0, 117, 118, 5, 108,
		0, 0, 118, 119, 5, 101, 0, 0, 119, 120, 5, 116, 0, 0, 120, 16, 1, 0, 0,
		0, 121, 122, 5, 116, 0, 0, 122, 123, 5, 121, 0, 0, 123, 124, 5, 112, 0,
		0, 124, 125, 5, 101, 0, 0, 125, 18, 1, 0, 0, 0, 126, 127, 5, 115, 0, 0,
		127, 128, 5, 116, 0, 0, 128, 129, 5, 114, 0, 0, 129, 130, 5, 117, 0, 0,
		130, 131, 5, 99, 0, 0, 131, 132, 5, 116, 0, 0, 132, 20, 1, 0, 0, 0, 133,
		134, 5, 105, 0, 0, 134, 135, 5, 110, 0, 0, 135, 136, 5, 116, 0, 0, 136,
		22, 1, 0, 0, 0, 137, 138, 5, 98, 0, 0, 138, 139, 5, 111, 0, 0, 139, 140,
		5, 111, 0, 0, 140, 141, 5, 108, 0, 0, 141, 24, 1, 0, 0, 0, 142, 143, 5,
		118, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 5, 114, 0, 0, 145, 26, 1, 0,
		0, 0, 146, 147, 5, 102, 0, 0, 147, 148, 5, 117, 0, 0, 148, 149, 5, 110,
		0, 0, 149, 150, 5, 99, 0, 0, 150, 28, 1, 0, 0, 0, 151, 152, 5, 105, 0,
		0, 152, 153, 5, 102, 0, 0, 153, 30, 1, 0, 0, 0, 154, 155, 5, 101, 0, 0,
		155, 156, 5, 108, 0, 0, 156, 157, 5, 115, 0, 0, 157, 158, 5, 101, 0, 0,
		158, 32, 1, 0, 0, 0, 159, 160, 5, 102, 0, 0, 160, 161, 5, 111, 0, 0, 161,
		162, 5, 114, 0, 0, 162, 34, 1, 0, 0, 0, 163, 164, 5, 114, 0, 0, 164, 165,
		5, 101, 0, 0, 165, 166, 5, 116, 0, 0, 166, 167, 5, 117, 0, 0, 167, 168,
		5, 114, 0, 0, 168, 169, 5, 110, 0, 0, 169, 36, 1, 0, 0, 0, 170, 171, 5,
		115, 0, 0, 171, 172, 5, 99, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 110,
		0, 0, 174, 38, 1, 0, 0, 0, 175, 176, 5, 100, 0, 0, 176, 177, 5, 101, 0,
		0, 177, 178, 5, 108, 0, 0, 178, 179, 5, 101, 0, 0, 179, 180, 5, 116, 0,
		0, 180, 181, 5, 101, 0, 0, 181, 40, 1, 0, 0, 0, 182, 183, 5, 112, 0, 0,
		183, 184, 5, 114, 0, 0, 184, 185, 5, 105, 0, 0, 185, 186, 5, 110, 0, 0,
		186, 187, 5, 116, 0, 0, 187, 188, 5, 102, 0, 0, 188, 42, 1, 0, 0, 0, 189,
		190, 5, 110, 0, 0, 190, 191, 5, 101, 0, 0, 191, 192, 5, 119, 0, 0, 192,
		44, 1, 0, 0, 0, 193, 194, 5, 116, 0, 0, 194, 195, 5, 114, 0, 0, 195, 196,
		5, 117, 0, 0, 196, 197, 5, 101, 0, 0, 197, 46, 1, 0, 0, 0, 198, 199, 5,
		102, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201, 5, 108, 0, 0, 201, 202, 5,
		115, 0, 0, 202, 203, 5, 101, 0, 0, 203, 48, 1, 0, 0, 0, 204, 205, 5, 110,
		0, 0, 205, 206, 5, 105, 0, 0, 206, 207, 5, 108, 0, 0, 207, 50, 1, 0, 0,
		0, 208, 209, 5, 123, 0, 0, 209, 52, 1, 0, 0, 0, 210, 211, 5, 125, 0, 0,
		211, 54, 1, 0, 0, 0, 212, 213, 5, 40, 0, 0, 213, 56, 1, 0, 0, 0, 214, 215,
		5, 41, 0, 0, 215, 58, 1, 0, 0, 0, 216, 217, 5, 46, 0, 0, 217, 60, 1, 0,
		0, 0, 218, 219, 5, 44, 0, 0, 219, 62, 1, 0, 0, 0, 220, 221, 5, 33, 0, 0,
		221, 64, 1, 0, 0, 0, 222, 223, 5, 124, 0, 0, 223, 224, 5, 124, 0, 0, 224,
		66, 1, 0, 0, 0, 225, 226, 5, 38, 0, 0, 226, 227, 5, 38, 0, 0, 227, 68,
		1, 0, 0, 0, 228, 229, 5, 61, 0, 0, 229, 230, 5, 61, 0, 0, 230, 70, 1, 0,
		0, 0, 231, 232, 5, 33, 0, 0, 232, 233, 5, 61, 0, 0, 233, 72, 1, 0, 0, 0,
		234, 235, 5, 62, 0, 0, 235, 74, 1, 0, 0, 0, 236, 237, 5, 60, 0, 0, 237,
		76, 1, 0, 0, 0, 238, 239, 5, 62, 0, 0, 239, 240, 5, 61, 0, 0, 240, 78,
		1, 0, 0, 0, 241, 242, 5, 60, 0, 0, 242, 243, 5, 61, 0, 0, 243, 80, 1, 0,
		0, 0, 244, 248, 7, 2, 0, 0, 245, 247, 7, 3, 0, 0, 246, 245, 1, 0, 0, 0,
		247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		82, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252, 5, 59, 0, 0, 252, 84, 1,
		0, 0, 0, 253, 254, 5, 47, 0, 0, 254, 255, 5, 47, 0, 0, 255, 261, 1, 0,
		0, 0, 256, 257, 5, 92, 0, 0, 257, 260, 7, 4, 0, 0, 258, 260, 8, 5, 0, 0,
		259, 256, 1, 0, 0, 0, 259, 258, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 261,
		1, 0, 0, 0, 264, 265, 6, 42, 0, 0, 265, 86, 1, 0, 0, 0, 266, 270, 5, 34,
		0, 0, 267, 269, 9, 0, 0, 0, 268, 267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0,
		270, 271, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272,
		270, 1, 0, 0, 0, 273, 274, 5, 34, 0, 0, 274, 88, 1, 0, 0, 0, 275, 277,
		7, 6, 0, 0, 276, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 276, 1, 0,
		0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 6, 44, 0, 0,
		281, 90, 1, 0, 0, 0, 8, 0, 95, 99, 248, 259, 261, 270, 278, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &golitelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "GoliteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerNUMBER      = 1
	GoliteLexerPLUS        = 2
	GoliteLexerMINUS       = 3
	GoliteLexerASSIGN      = 4
	GoliteLexerASTERISK    = 5
	GoliteLexerFSLASH      = 6
	GoliteLexerPRINT       = 7
	GoliteLexerLET         = 8
	GoliteLexerTYPE        = 9
	GoliteLexerSTRUCT      = 10
	GoliteLexerINT_TYPE    = 11
	GoliteLexerBOOL_TYPE   = 12
	GoliteLexerVAR         = 13
	GoliteLexerFUNC        = 14
	GoliteLexerIF          = 15
	GoliteLexerELSE        = 16
	GoliteLexerFOR         = 17
	GoliteLexerRETURN      = 18
	GoliteLexerSCAN        = 19
	GoliteLexerDELETE      = 20
	GoliteLexerPRINTF      = 21
	GoliteLexerNEW         = 22
	GoliteLexerTRUE        = 23
	GoliteLexerFALSE       = 24
	GoliteLexerNIL         = 25
	GoliteLexerLEFTCURL    = 26
	GoliteLexerRIGHTCURL   = 27
	GoliteLexerLEFTBRAC    = 28
	GoliteLexerRIGHTBRAC   = 29
	GoliteLexerDOT         = 30
	GoliteLexerCOMMA       = 31
	GoliteLexerEXCLAMATION = 32
	GoliteLexerOR          = 33
	GoliteLexerAND         = 34
	GoliteLexerEQUALS      = 35
	GoliteLexerNOTEQUALS   = 36
	GoliteLexerMORETHAN    = 37
	GoliteLexerLESSTHAN    = 38
	GoliteLexerGEQ         = 39
	GoliteLexerLEQ         = 40
	GoliteLexerID          = 41
	GoliteLexerSEMICOLON   = 42
	GoliteLexerCOMMENT     = 43
	GoliteLexerSTRING      = 44
	GoliteLexerWS          = 45
)
