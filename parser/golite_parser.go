// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoliteParser struct {
	*antlr.BaseParser
}

var goliteparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &goliteparserParserStaticData
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
		"program", "types", "typeDeclaration", "fields", "decl", "type", "declarations",
		"declaration", "ids", "functions", "function", "parameters", "returnType",
		"statements", "statement", "read", "block", "delete", "assignment",
		"print", "conditional", "loop", "return", "invocation", "arguments",
		"lValue", "expression", "boolTerm", "equalTerm", "opRelationTerm", "relationTerm",
		"opSimpleTerm", "simpleTerm", "opTerm", "term", "opUnaryTerm", "unaryTerm",
		"selectorTerm", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 45, 348, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 85,
		8, 1, 10, 1, 12, 1, 88, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 103, 8, 3, 10, 3, 12, 3, 106,
		9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 115, 8, 5, 1, 6,
		5, 6, 118, 8, 6, 10, 6, 12, 6, 121, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 5, 8, 131, 8, 8, 10, 8, 12, 8, 134, 9, 8, 1, 9, 5, 9,
		137, 8, 9, 10, 9, 12, 9, 140, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		146, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 5, 11, 157, 8, 11, 10, 11, 12, 11, 160, 9, 11, 3, 11, 162, 8, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 5, 13, 169, 8, 13, 10, 13, 12, 13, 172,
		9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 183, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 207, 8, 19, 10, 19, 12, 19, 210, 9, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3,
		20, 222, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		3, 22, 232, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 5, 24, 244, 8, 24, 10, 24, 12, 24, 247, 9, 24, 3, 24,
		249, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 5, 25, 256, 8, 25, 10, 25,
		12, 25, 259, 9, 25, 1, 26, 1, 26, 1, 26, 5, 26, 264, 8, 26, 10, 26, 12,
		26, 267, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 272, 8, 27, 10, 27, 12, 27,
		275, 9, 27, 1, 28, 1, 28, 5, 28, 279, 8, 28, 10, 28, 12, 28, 282, 9, 28,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 289, 8, 30, 10, 30, 12, 30, 292,
		9, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 299, 8, 32, 10, 32, 12,
		32, 302, 9, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 5, 34, 309, 8, 34, 10,
		34, 12, 34, 312, 9, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 322, 8, 36, 1, 37, 1, 37, 1, 37, 5, 37, 327, 8, 37, 10, 37,
		12, 37, 330, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 338,
		8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 346, 8, 38, 1,
		38, 0, 0, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 0, 4, 1, 0, 35, 36, 1, 0, 37, 40, 1, 0, 2, 3, 1, 0,
		5, 6, 349, 0, 78, 1, 0, 0, 0, 2, 86, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6,
		97, 1, 0, 0, 0, 8, 107, 1, 0, 0, 0, 10, 114, 1, 0, 0, 0, 12, 119, 1, 0,
		0, 0, 14, 122, 1, 0, 0, 0, 16, 127, 1, 0, 0, 0, 18, 138, 1, 0, 0, 0, 20,
		141, 1, 0, 0, 0, 22, 152, 1, 0, 0, 0, 24, 165, 1, 0, 0, 0, 26, 170, 1,
		0, 0, 0, 28, 182, 1, 0, 0, 0, 30, 184, 1, 0, 0, 0, 32, 188, 1, 0, 0, 0,
		34, 192, 1, 0, 0, 0, 36, 196, 1, 0, 0, 0, 38, 201, 1, 0, 0, 0, 40, 214,
		1, 0, 0, 0, 42, 223, 1, 0, 0, 0, 44, 229, 1, 0, 0, 0, 46, 235, 1, 0, 0,
		0, 48, 239, 1, 0, 0, 0, 50, 252, 1, 0, 0, 0, 52, 260, 1, 0, 0, 0, 54, 268,
		1, 0, 0, 0, 56, 276, 1, 0, 0, 0, 58, 283, 1, 0, 0, 0, 60, 286, 1, 0, 0,
		0, 62, 293, 1, 0, 0, 0, 64, 296, 1, 0, 0, 0, 66, 303, 1, 0, 0, 0, 68, 306,
		1, 0, 0, 0, 70, 313, 1, 0, 0, 0, 72, 321, 1, 0, 0, 0, 74, 323, 1, 0, 0,
		0, 76, 345, 1, 0, 0, 0, 78, 79, 3, 2, 1, 0, 79, 80, 3, 12, 6, 0, 80, 81,
		3, 18, 9, 0, 81, 82, 5, 0, 0, 1, 82, 1, 1, 0, 0, 0, 83, 85, 3, 4, 2, 0,
		84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1,
		0, 0, 0, 87, 3, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 5, 9, 0, 0, 90,
		91, 5, 41, 0, 0, 91, 92, 5, 10, 0, 0, 92, 93, 5, 26, 0, 0, 93, 94, 3, 6,
		3, 0, 94, 95, 5, 27, 0, 0, 95, 96, 5, 42, 0, 0, 96, 5, 1, 0, 0, 0, 97,
		98, 3, 8, 4, 0, 98, 104, 5, 42, 0, 0, 99, 100, 3, 8, 4, 0, 100, 101, 5,
		42, 0, 0, 101, 103, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 103, 106, 1, 0, 0,
		0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 7, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 107, 108, 5, 41, 0, 0, 108, 109, 3, 10, 5, 0, 109, 9,
		1, 0, 0, 0, 110, 115, 5, 11, 0, 0, 111, 115, 5, 12, 0, 0, 112, 113, 5,
		5, 0, 0, 113, 115, 5, 41, 0, 0, 114, 110, 1, 0, 0, 0, 114, 111, 1, 0, 0,
		0, 114, 112, 1, 0, 0, 0, 115, 11, 1, 0, 0, 0, 116, 118, 3, 14, 7, 0, 117,
		116, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 13, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 13,
		0, 0, 123, 124, 3, 16, 8, 0, 124, 125, 3, 10, 5, 0, 125, 126, 5, 42, 0,
		0, 126, 15, 1, 0, 0, 0, 127, 132, 5, 41, 0, 0, 128, 129, 5, 31, 0, 0, 129,
		131, 5, 41, 0, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 17, 1, 0, 0, 0, 134, 132, 1, 0,
		0, 0, 135, 137, 3, 20, 10, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0,
		0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 19, 1, 0, 0, 0, 140,
		138, 1, 0, 0, 0, 141, 142, 5, 14, 0, 0, 142, 143, 5, 41, 0, 0, 143, 145,
		3, 22, 11, 0, 144, 146, 3, 24, 12, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1,
		0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 26, 0, 0, 148, 149, 3, 12,
		6, 0, 149, 150, 3, 26, 13, 0, 150, 151, 5, 27, 0, 0, 151, 21, 1, 0, 0,
		0, 152, 161, 5, 28, 0, 0, 153, 158, 3, 8, 4, 0, 154, 155, 5, 31, 0, 0,
		155, 157, 3, 8, 4, 0, 156, 154, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 161, 153, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 164, 5, 29, 0, 0, 164, 23, 1, 0, 0, 0, 165, 166, 3, 10, 5, 0,
		166, 25, 1, 0, 0, 0, 167, 169, 3, 28, 14, 0, 168, 167, 1, 0, 0, 0, 169,
		172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 27, 1,
		0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 183, 3, 32, 16, 0, 174, 183, 3, 36,
		18, 0, 175, 183, 3, 38, 19, 0, 176, 183, 3, 34, 17, 0, 177, 183, 3, 40,
		20, 0, 178, 183, 3, 42, 21, 0, 179, 183, 3, 44, 22, 0, 180, 183, 3, 30,
		15, 0, 181, 183, 3, 46, 23, 0, 182, 173, 1, 0, 0, 0, 182, 174, 1, 0, 0,
		0, 182, 175, 1, 0, 0, 0, 182, 176, 1, 0, 0, 0, 182, 177, 1, 0, 0, 0, 182,
		178, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 181,
		1, 0, 0, 0, 183, 29, 1, 0, 0, 0, 184, 185, 5, 19, 0, 0, 185, 186, 3, 50,
		25, 0, 186, 187, 5, 42, 0, 0, 187, 31, 1, 0, 0, 0, 188, 189, 5, 26, 0,
		0, 189, 190, 3, 26, 13, 0, 190, 191, 5, 27, 0, 0, 191, 33, 1, 0, 0, 0,
		192, 193, 5, 20, 0, 0, 193, 194, 3, 52, 26, 0, 194, 195, 5, 42, 0, 0, 195,
		35, 1, 0, 0, 0, 196, 197, 3, 50, 25, 0, 197, 198, 5, 4, 0, 0, 198, 199,
		3, 52, 26, 0, 199, 200, 5, 42, 0, 0, 200, 37, 1, 0, 0, 0, 201, 202, 5,
		21, 0, 0, 202, 203, 5, 28, 0, 0, 203, 208, 5, 44, 0, 0, 204, 205, 5, 31,
		0, 0, 205, 207, 3, 52, 26, 0, 206, 204, 1, 0, 0, 0, 207, 210, 1, 0, 0,
		0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0, 210,
		208, 1, 0, 0, 0, 211, 212, 5, 29, 0, 0, 212, 213, 5, 42, 0, 0, 213, 39,
		1, 0, 0, 0, 214, 215, 5, 15, 0, 0, 215, 216, 5, 28, 0, 0, 216, 217, 3,
		52, 26, 0, 217, 218, 5, 29, 0, 0, 218, 221, 3, 32, 16, 0, 219, 220, 5,
		16, 0, 0, 220, 222, 3, 32, 16, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0,
		0, 0, 222, 41, 1, 0, 0, 0, 223, 224, 5, 17, 0, 0, 224, 225, 5, 28, 0, 0,
		225, 226, 3, 52, 26, 0, 226, 227, 5, 29, 0, 0, 227, 228, 3, 32, 16, 0,
		228, 43, 1, 0, 0, 0, 229, 231, 5, 18, 0, 0, 230, 232, 3, 52, 26, 0, 231,
		230, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234,
		5, 42, 0, 0, 234, 45, 1, 0, 0, 0, 235, 236, 5, 41, 0, 0, 236, 237, 3, 48,
		24, 0, 237, 238, 5, 42, 0, 0, 238, 47, 1, 0, 0, 0, 239, 248, 5, 28, 0,
		0, 240, 245, 3, 52, 26, 0, 241, 242, 5, 31, 0, 0, 242, 244, 3, 52, 26,
		0, 243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245,
		246, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 240,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 29,
		0, 0, 251, 49, 1, 0, 0, 0, 252, 257, 5, 41, 0, 0, 253, 254, 5, 30, 0, 0,
		254, 256, 5, 41, 0, 0, 255, 253, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 51, 1, 0, 0, 0, 259, 257, 1,
		0, 0, 0, 260, 265, 3, 54, 27, 0, 261, 262, 5, 33, 0, 0, 262, 264, 3, 54,
		27, 0, 263, 261, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 53, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 268, 273,
		3, 56, 28, 0, 269, 270, 5, 34, 0, 0, 270, 272, 3, 56, 28, 0, 271, 269,
		1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0,
		0, 0, 274, 55, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 280, 3, 60, 30, 0,
		277, 279, 3, 58, 29, 0, 278, 277, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280,
		278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 57, 1, 0, 0, 0, 282, 280, 1,
		0, 0, 0, 283, 284, 7, 0, 0, 0, 284, 285, 3, 60, 30, 0, 285, 59, 1, 0, 0,
		0, 286, 290, 3, 64, 32, 0, 287, 289, 3, 62, 31, 0, 288, 287, 1, 0, 0, 0,
		289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291,
		61, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 7, 1, 0, 0, 294, 295, 3,
		64, 32, 0, 295, 63, 1, 0, 0, 0, 296, 300, 3, 68, 34, 0, 297, 299, 3, 66,
		33, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		300, 301, 1, 0, 0, 0, 301, 65, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 304,
		7, 2, 0, 0, 304, 305, 3, 68, 34, 0, 305, 67, 1, 0, 0, 0, 306, 310, 3, 72,
		36, 0, 307, 309, 3, 70, 35, 0, 308, 307, 1, 0, 0, 0, 309, 312, 1, 0, 0,
		0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 69, 1, 0, 0, 0, 312,
		310, 1, 0, 0, 0, 313, 314, 7, 3, 0, 0, 314, 315, 3, 72, 36, 0, 315, 71,
		1, 0, 0, 0, 316, 317, 5, 32, 0, 0, 317, 322, 3, 74, 37, 0, 318, 319, 5,
		3, 0, 0, 319, 322, 3, 74, 37, 0, 320, 322, 3, 74, 37, 0, 321, 316, 1, 0,
		0, 0, 321, 318, 1, 0, 0, 0, 321, 320, 1, 0, 0, 0, 322, 73, 1, 0, 0, 0,
		323, 328, 3, 76, 38, 0, 324, 325, 5, 30, 0, 0, 325, 327, 5, 41, 0, 0, 326,
		324, 1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329,
		1, 0, 0, 0, 329, 75, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 332, 5, 28,
		0, 0, 332, 333, 3, 52, 26, 0, 333, 334, 5, 29, 0, 0, 334, 346, 1, 0, 0,
		0, 335, 337, 5, 41, 0, 0, 336, 338, 3, 48, 24, 0, 337, 336, 1, 0, 0, 0,
		337, 338, 1, 0, 0, 0, 338, 346, 1, 0, 0, 0, 339, 346, 5, 1, 0, 0, 340,
		341, 5, 22, 0, 0, 341, 346, 5, 41, 0, 0, 342, 346, 5, 23, 0, 0, 343, 346,
		5, 24, 0, 0, 344, 346, 5, 25, 0, 0, 345, 331, 1, 0, 0, 0, 345, 335, 1,
		0, 0, 0, 345, 339, 1, 0, 0, 0, 345, 340, 1, 0, 0, 0, 345, 342, 1, 0, 0,
		0, 345, 343, 1, 0, 0, 0, 345, 344, 1, 0, 0, 0, 346, 77, 1, 0, 0, 0, 27,
		86, 104, 114, 119, 132, 138, 145, 158, 161, 170, 182, 208, 221, 231, 245,
		248, 257, 265, 273, 280, 290, 300, 310, 321, 328, 337, 345,
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

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &goliteparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF         = antlr.TokenEOF
	GoliteParserNUMBER      = 1
	GoliteParserPLUS        = 2
	GoliteParserMINUS       = 3
	GoliteParserASSIGN      = 4
	GoliteParserASTERISK    = 5
	GoliteParserFSLASH      = 6
	GoliteParserPRINT       = 7
	GoliteParserLET         = 8
	GoliteParserTYPE        = 9
	GoliteParserSTRUCT      = 10
	GoliteParserINT_TYPE    = 11
	GoliteParserBOOL_TYPE   = 12
	GoliteParserVAR         = 13
	GoliteParserFUNC        = 14
	GoliteParserIF          = 15
	GoliteParserELSE        = 16
	GoliteParserFOR         = 17
	GoliteParserRETURN      = 18
	GoliteParserSCAN        = 19
	GoliteParserDELETE      = 20
	GoliteParserPRINTF      = 21
	GoliteParserNEW         = 22
	GoliteParserTRUE        = 23
	GoliteParserFALSE       = 24
	GoliteParserNIL         = 25
	GoliteParserLEFTCURL    = 26
	GoliteParserRIGHTCURL   = 27
	GoliteParserLEFTBRAC    = 28
	GoliteParserRIGHTBRAC   = 29
	GoliteParserDOT         = 30
	GoliteParserCOMMA       = 31
	GoliteParserEXCLAMATION = 32
	GoliteParserOR          = 33
	GoliteParserAND         = 34
	GoliteParserEQUALS      = 35
	GoliteParserNOTEQUALS   = 36
	GoliteParserMORETHAN    = 37
	GoliteParserLESSTHAN    = 38
	GoliteParserGEQ         = 39
	GoliteParserLEQ         = 40
	GoliteParserID          = 41
	GoliteParserSEMICOLON   = 42
	GoliteParserCOMMENT     = 43
	GoliteParserSTRING      = 44
	GoliteParserWS          = 45
)

// GoliteParser rules.
const (
	GoliteParserRULE_program         = 0
	GoliteParserRULE_types           = 1
	GoliteParserRULE_typeDeclaration = 2
	GoliteParserRULE_fields          = 3
	GoliteParserRULE_decl            = 4
	GoliteParserRULE_type            = 5
	GoliteParserRULE_declarations    = 6
	GoliteParserRULE_declaration     = 7
	GoliteParserRULE_ids             = 8
	GoliteParserRULE_functions       = 9
	GoliteParserRULE_function        = 10
	GoliteParserRULE_parameters      = 11
	GoliteParserRULE_returnType      = 12
	GoliteParserRULE_statements      = 13
	GoliteParserRULE_statement       = 14
	GoliteParserRULE_read            = 15
	GoliteParserRULE_block           = 16
	GoliteParserRULE_delete          = 17
	GoliteParserRULE_assignment      = 18
	GoliteParserRULE_print           = 19
	GoliteParserRULE_conditional     = 20
	GoliteParserRULE_loop            = 21
	GoliteParserRULE_return          = 22
	GoliteParserRULE_invocation      = 23
	GoliteParserRULE_arguments       = 24
	GoliteParserRULE_lValue          = 25
	GoliteParserRULE_expression      = 26
	GoliteParserRULE_boolTerm        = 27
	GoliteParserRULE_equalTerm       = 28
	GoliteParserRULE_opRelationTerm  = 29
	GoliteParserRULE_relationTerm    = 30
	GoliteParserRULE_opSimpleTerm    = 31
	GoliteParserRULE_simpleTerm      = 32
	GoliteParserRULE_opTerm          = 33
	GoliteParserRULE_term            = 34
	GoliteParserRULE_opUnaryTerm     = 35
	GoliteParserRULE_unaryTerm       = 36
	GoliteParserRULE_selectorTerm    = 37
	GoliteParserRULE_factor          = 38
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ProgramContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *ProgramContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Types()
	}
	{
		p.SetState(79)
		p.Declarations()
	}
	{
		p.SetState(80)
		p.Functions()
	}
	{
		p.SetState(81)
		p.Match(GoliteParserEOF)
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypesContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *GoliteParser) Types() (localctx ITypesContext) {
	this := p
	_ = this

	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_types)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(83)
			p.TypeDeclaration()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypeDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeDeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypeDeclarationContext) LEFTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTCURL, 0)
}

func (s *TypeDeclarationContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypeDeclarationContext) RIGHTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTCURL, 0)
}

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *GoliteParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	this := p
	_ = this

	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(GoliteParserTYPE)
	}
	{
		p.SetState(90)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(91)
		p.Match(GoliteParserSTRUCT)
	}
	{
		p.SetState(92)
		p.Match(GoliteParserLEFTCURL)
	}
	{
		p.SetState(93)
		p.Fields()
	}
	{
		p.SetState(94)
		p.Match(GoliteParserRIGHTCURL)
	}
	{
		p.SetState(95)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserSEMICOLON)
}

func (s *FieldsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, i)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Decl()
	}
	{
		p.SetState(98)
		p.Match(GoliteParserSEMICOLON)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserID {
		{
			p.SetState(99)
			p.Decl()
		}
		{
			p.SetState(100)
			p.Match(GoliteParserSEMICOLON)
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     ITypeContext
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) GetTy() ITypeContext { return s.ty }

func (s *DeclContext) SetTy(v ITypeContext) { s.ty = v }

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(108)

		var _x = p.Type_()

		localctx.(*DeclContext).ty = _x
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT_TYPE, 0)
}

func (s *TypeContext) BOOL_TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL_TYPE, 0)
}

func (s *TypeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(GoliteParserINT_TYPE)
		}

	case GoliteParserBOOL_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(GoliteParserBOOL_TYPE)
		}

	case GoliteParserASTERISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(GoliteParserASTERISK)
		}
		{
			p.SetState(113)
			p.Match(GoliteParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationsContext is an interface to support dynamic dispatch.
type IDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationsContext differentiates from other interfaces.
	IsDeclarationsContext()
}

type DeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationsContext() *DeclarationsContext {
	var p = new(DeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declarations
	return p
}

func (*DeclarationsContext) IsDeclarationsContext() {}

func NewDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationsContext {
	var p = new(DeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarations

	return p
}

func (s *DeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationsContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationsContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarations(s)
	}
}

func (s *DeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarations(s)
	}
}

func (p *GoliteParser) Declarations() (localctx IDeclarationsContext) {
	this := p
	_ = this

	localctx = NewDeclarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_declarations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(116)
			p.Declaration()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     ITypeContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetTy() ITypeContext { return s.ty }

func (s *DeclarationContext) SetTy(v ITypeContext) { s.ty = v }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) Ids() IIdsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdsContext)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(GoliteParserVAR)
	}
	{
		p.SetState(123)
		p.Ids()
	}
	{
		p.SetState(124)

		var _x = p.Type_()

		localctx.(*DeclarationContext).ty = _x
	}
	{
		p.SetState(125)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IIdsContext is an interface to support dynamic dispatch.
type IIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdsContext differentiates from other interfaces.
	IsIdsContext()
}

type IdsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdsContext() *IdsContext {
	var p = new(IdsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_ids
	return p
}

func (*IdsContext) IsIdsContext() {}

func NewIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdsContext {
	var p = new(IdsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_ids

	return p
}

func (s *IdsContext) GetParser() antlr.Parser { return s.parser }

func (s *IdsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserID)
}

func (s *IdsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserID, i)
}

func (s *IdsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *IdsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *IdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterIds(s)
	}
}

func (s *IdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitIds(s)
	}
}

func (p *GoliteParser) Ids() (localctx IIdsContext) {
	this := p
	_ = this

	localctx = NewIdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_ids)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(GoliteParserID)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(128)
			p.Match(GoliteParserCOMMA)
		}
		{
			p.SetState(129)
			p.Match(GoliteParserID)
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionsContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *GoliteParser) Functions() (localctx IFunctionsContext) {
	this := p
	_ = this

	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_functions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNC {
		{
			p.SetState(135)
			p.Function()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) LEFTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTCURL, 0)
}

func (s *FunctionContext) Declarations() IDeclarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationsContext)
}

func (s *FunctionContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionContext) RIGHTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTCURL, 0)
}

func (s *FunctionContext) ReturnType() IReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(GoliteParserFUNC)
	}
	{
		p.SetState(142)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(143)
		p.Parameters()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6176) != 0 {
		{
			p.SetState(144)
			p.ReturnType()
		}

	}
	{
		p.SetState(147)
		p.Match(GoliteParserLEFTCURL)
	}
	{
		p.SetState(148)
		p.Declarations()
	}
	{
		p.SetState(149)
		p.Statements()
	}
	{
		p.SetState(150)
		p.Match(GoliteParserRIGHTCURL)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *ParametersContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *ParametersContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GoliteParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(GoliteParserLEFTBRAC)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserID {
		{
			p.SetState(153)
			p.Decl()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(154)
				p.Match(GoliteParserCOMMA)
			}
			{
				p.SetState(155)
				p.Decl()
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(163)
		p.Match(GoliteParserRIGHTBRAC)
	}

	return localctx
}

// IReturnTypeContext is an interface to support dynamic dispatch.
type IReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsReturnTypeContext differentiates from other interfaces.
	IsReturnTypeContext()
}

type ReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     ITypeContext
}

func NewEmptyReturnTypeContext() *ReturnTypeContext {
	var p = new(ReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_returnType
	return p
}

func (*ReturnTypeContext) IsReturnTypeContext() {}

func NewReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnTypeContext {
	var p = new(ReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_returnType

	return p
}

func (s *ReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnTypeContext) GetTy() ITypeContext { return s.ty }

func (s *ReturnTypeContext) SetTy(v ITypeContext) { s.ty = v }

func (s *ReturnTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturnType(s)
	}
}

func (s *ReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturnType(s)
	}
}

func (p *GoliteParser) ReturnType() (localctx IReturnTypeContext) {
	this := p
	_ = this

	localctx = NewReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_returnType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)

		var _x = p.Type_()

		localctx.(*ReturnTypeContext).ty = _x
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *GoliteParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199094460416) != 0 {
		{
			p.SetState(167)
			p.Statement()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Return_() IReturnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Print_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Delete_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(178)
			p.Loop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(179)
			p.Return_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(180)
			p.Read()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(181)
			p.Invocation()
		}

	}

	return localctx
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *ReadContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
}

func (s *ReadContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *GoliteParser) Read() (localctx IReadContext) {
	this := p
	_ = this

	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_read)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GoliteParserSCAN)
	}
	{
		p.SetState(185)
		p.LValue()
	}
	{
		p.SetState(186)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LEFTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTCURL, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) RIGHTCURL() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTCURL, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(GoliteParserLEFTCURL)
	}
	{
		p.SetState(189)
		p.Statements()
	}
	{
		p.SetState(190)
		p.Match(GoliteParserRIGHTCURL)
	}

	return localctx
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	this := p
	_ = this

	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_delete)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GoliteParserDELETE)
	}
	{
		p.SetState(193)
		p.Expression()
	}
	{
		p.SetState(194)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) LValue() ILValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILValueContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoliteParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.LValue()
	}
	{
		p.SetState(197)
		p.Match(GoliteParserASSIGN)
	}
	{
		p.SetState(198)
		p.Expression()
	}
	{
		p.SetState(199)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINTF, 0)
}

func (s *PrintContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *PrintContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING, 0)
}

func (s *PrintContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *PrintContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *PrintContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	this := p
	_ = this

	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_print)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(GoliteParserPRINTF)
	}
	{
		p.SetState(202)
		p.Match(GoliteParserLEFTBRAC)
	}
	{
		p.SetState(203)
		p.Match(GoliteParserSTRING)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(204)
			p.Match(GoliteParserCOMMA)
		}
		{
			p.SetState(205)
			p.Expression()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(211)
		p.Match(GoliteParserRIGHTBRAC)
	}
	{
		p.SetState(212)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *ConditionalContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_conditional)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(GoliteParserIF)
	}
	{
		p.SetState(215)
		p.Match(GoliteParserLEFTBRAC)
	}
	{
		p.SetState(216)
		p.Expression()
	}
	{
		p.SetState(217)
		p.Match(GoliteParserRIGHTBRAC)
	}
	{
		p.SetState(218)
		p.Block()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(219)
			p.Match(GoliteParserELSE)
		}
		{
			p.SetState(220)
			p.Block()
		}

	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(GoliteParserFOR)
	}
	{
		p.SetState(224)
		p.Match(GoliteParserLEFTBRAC)
	}
	{
		p.SetState(225)
		p.Expression()
	}
	{
		p.SetState(226)
		p.Match(GoliteParserRIGHTBRAC)
	}
	{
		p.SetState(227)
		p.Block()
	}

	return localctx
}

// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_return
	return p
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRETURN, 0)
}

func (s *ReturnContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReturnContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (p *GoliteParser) Return_() (localctx IReturnContext) {
	this := p
	_ = this

	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_return)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(GoliteParserRETURN)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2203649572874) != 0 {
		{
			p.SetState(230)
			p.Expression()
		}

	}
	{
		p.SetState(233)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	this := p
	_ = this

	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_invocation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(236)
		p.Arguments()
	}
	{
		p.SetState(237)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *ArgumentsContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GoliteParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(GoliteParserLEFTBRAC)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2203649572874) != 0 {
		{
			p.SetState(240)
			p.Expression()
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(241)
				p.Match(GoliteParserCOMMA)
			}
			{
				p.SetState(242)
				p.Expression()
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(250)
		p.Match(GoliteParserRIGHTBRAC)
	}

	return localctx
}

// ILValueContext is an interface to support dynamic dispatch.
type ILValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLValueContext differentiates from other interfaces.
	IsLValueContext()
}

type LValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLValueContext() *LValueContext {
	var p = new(LValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_lValue
	return p
}

func (*LValueContext) IsLValueContext() {}

func NewLValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LValueContext {
	var p = new(LValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_lValue

	return p
}

func (s *LValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LValueContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserID)
}

func (s *LValueContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserID, i)
}

func (s *LValueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDOT)
}

func (s *LValueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, i)
}

func (s *LValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLValue(s)
	}
}

func (s *LValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLValue(s)
	}
}

func (p *GoliteParser) LValue() (localctx ILValueContext) {
	this := p
	_ = this

	localctx = NewLValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_lValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(GoliteParserID)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(253)
			p.Match(GoliteParserDOT)
		}
		{
			p.SetState(254)
			p.Match(GoliteParserID)
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllBoolTerm() []IBoolTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolTermContext); ok {
			len++
		}
	}

	tst := make([]IBoolTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolTermContext); ok {
			tst[i] = t.(IBoolTermContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) BoolTerm(i int) IBoolTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTermContext)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.BoolTerm()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(261)
			p.Match(GoliteParserOR)
		}
		{
			p.SetState(262)
			p.BoolTerm()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolTermContext is an interface to support dynamic dispatch.
type IBoolTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolTermContext differentiates from other interfaces.
	IsBoolTermContext()
}

type BoolTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTermContext() *BoolTermContext {
	var p = new(BoolTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolTerm
	return p
}

func (*BoolTermContext) IsBoolTermContext() {}

func NewBoolTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTermContext {
	var p = new(BoolTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolTerm

	return p
}

func (s *BoolTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolTermContext) AllEqualTerm() []IEqualTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualTermContext); ok {
			len++
		}
	}

	tst := make([]IEqualTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualTermContext); ok {
			tst[i] = t.(IEqualTermContext)
			i++
		}
	}

	return tst
}

func (s *BoolTermContext) EqualTerm(i int) IEqualTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualTermContext)
}

func (s *BoolTermContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserAND)
}

func (s *BoolTermContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserAND, i)
}

func (s *BoolTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolTerm(s)
	}
}

func (s *BoolTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolTerm(s)
	}
}

func (p *GoliteParser) BoolTerm() (localctx IBoolTermContext) {
	this := p
	_ = this

	localctx = NewBoolTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_boolTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.EqualTerm()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserAND {
		{
			p.SetState(269)
			p.Match(GoliteParserAND)
		}
		{
			p.SetState(270)
			p.EqualTerm()
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEqualTermContext is an interface to support dynamic dispatch.
type IEqualTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualTermContext differentiates from other interfaces.
	IsEqualTermContext()
}

type EqualTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualTermContext() *EqualTermContext {
	var p = new(EqualTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalTerm
	return p
}

func (*EqualTermContext) IsEqualTermContext() {}

func NewEqualTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualTermContext {
	var p = new(EqualTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalTerm

	return p
}

func (s *EqualTermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualTermContext) RelationTerm() IRelationTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationTermContext)
}

func (s *EqualTermContext) AllOpRelationTerm() []IOpRelationTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpRelationTermContext); ok {
			len++
		}
	}

	tst := make([]IOpRelationTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpRelationTermContext); ok {
			tst[i] = t.(IOpRelationTermContext)
			i++
		}
	}

	return tst
}

func (s *EqualTermContext) OpRelationTerm(i int) IOpRelationTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpRelationTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpRelationTermContext)
}

func (s *EqualTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualTerm(s)
	}
}

func (s *EqualTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualTerm(s)
	}
}

func (p *GoliteParser) EqualTerm() (localctx IEqualTermContext) {
	this := p
	_ = this

	localctx = NewEqualTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_equalTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.RelationTerm()
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserEQUALS || _la == GoliteParserNOTEQUALS {
		{
			p.SetState(277)
			p.OpRelationTerm()
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOpRelationTermContext is an interface to support dynamic dispatch.
type IOpRelationTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRt returns the rt rule contexts.
	GetRt() IRelationTermContext

	// SetRt sets the rt rule contexts.
	SetRt(IRelationTermContext)

	// IsOpRelationTermContext differentiates from other interfaces.
	IsOpRelationTermContext()
}

type OpRelationTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	rt     IRelationTermContext
}

func NewEmptyOpRelationTermContext() *OpRelationTermContext {
	var p = new(OpRelationTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_opRelationTerm
	return p
}

func (*OpRelationTermContext) IsOpRelationTermContext() {}

func NewOpRelationTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpRelationTermContext {
	var p = new(OpRelationTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_opRelationTerm

	return p
}

func (s *OpRelationTermContext) GetParser() antlr.Parser { return s.parser }

func (s *OpRelationTermContext) GetOp() antlr.Token { return s.op }

func (s *OpRelationTermContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpRelationTermContext) GetRt() IRelationTermContext { return s.rt }

func (s *OpRelationTermContext) SetRt(v IRelationTermContext) { s.rt = v }

func (s *OpRelationTermContext) RelationTerm() IRelationTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationTermContext)
}

func (s *OpRelationTermContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQUALS, 0)
}

func (s *OpRelationTermContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOTEQUALS, 0)
}

func (s *OpRelationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpRelationTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpRelationTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterOpRelationTerm(s)
	}
}

func (s *OpRelationTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitOpRelationTerm(s)
	}
}

func (p *GoliteParser) OpRelationTerm() (localctx IOpRelationTermContext) {
	this := p
	_ = this

	localctx = NewOpRelationTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_opRelationTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OpRelationTermContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserEQUALS || _la == GoliteParserNOTEQUALS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OpRelationTermContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(284)

		var _x = p.RelationTerm()

		localctx.(*OpRelationTermContext).rt = _x
	}

	return localctx
}

// IRelationTermContext is an interface to support dynamic dispatch.
type IRelationTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationTermContext differentiates from other interfaces.
	IsRelationTermContext()
}

type RelationTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationTermContext() *RelationTermContext {
	var p = new(RelationTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relationTerm
	return p
}

func (*RelationTermContext) IsRelationTermContext() {}

func NewRelationTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationTermContext {
	var p = new(RelationTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationTerm

	return p
}

func (s *RelationTermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationTermContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *RelationTermContext) AllOpSimpleTerm() []IOpSimpleTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpSimpleTermContext); ok {
			len++
		}
	}

	tst := make([]IOpSimpleTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpSimpleTermContext); ok {
			tst[i] = t.(IOpSimpleTermContext)
			i++
		}
	}

	return tst
}

func (s *RelationTermContext) OpSimpleTerm(i int) IOpSimpleTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpSimpleTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpSimpleTermContext)
}

func (s *RelationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationTerm(s)
	}
}

func (s *RelationTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationTerm(s)
	}
}

func (p *GoliteParser) RelationTerm() (localctx IRelationTermContext) {
	this := p
	_ = this

	localctx = NewRelationTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_relationTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.SimpleTerm()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2061584302080) != 0 {
		{
			p.SetState(287)
			p.OpSimpleTerm()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOpSimpleTermContext is an interface to support dynamic dispatch.
type IOpSimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetSt returns the st rule contexts.
	GetSt() ISimpleTermContext

	// SetSt sets the st rule contexts.
	SetSt(ISimpleTermContext)

	// IsOpSimpleTermContext differentiates from other interfaces.
	IsOpSimpleTermContext()
}

type OpSimpleTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	st     ISimpleTermContext
}

func NewEmptyOpSimpleTermContext() *OpSimpleTermContext {
	var p = new(OpSimpleTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_opSimpleTerm
	return p
}

func (*OpSimpleTermContext) IsOpSimpleTermContext() {}

func NewOpSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpSimpleTermContext {
	var p = new(OpSimpleTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_opSimpleTerm

	return p
}

func (s *OpSimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *OpSimpleTermContext) GetOp() antlr.Token { return s.op }

func (s *OpSimpleTermContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpSimpleTermContext) GetSt() ISimpleTermContext { return s.st }

func (s *OpSimpleTermContext) SetSt(v ISimpleTermContext) { s.st = v }

func (s *OpSimpleTermContext) SimpleTerm() ISimpleTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTermContext)
}

func (s *OpSimpleTermContext) MORETHAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserMORETHAN, 0)
}

func (s *OpSimpleTermContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLESSTHAN, 0)
}

func (s *OpSimpleTermContext) LEQ() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEQ, 0)
}

func (s *OpSimpleTermContext) GEQ() antlr.TerminalNode {
	return s.GetToken(GoliteParserGEQ, 0)
}

func (s *OpSimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpSimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpSimpleTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterOpSimpleTerm(s)
	}
}

func (s *OpSimpleTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitOpSimpleTerm(s)
	}
}

func (p *GoliteParser) OpSimpleTerm() (localctx IOpSimpleTermContext) {
	this := p
	_ = this

	localctx = NewOpSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_opSimpleTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OpSimpleTermContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2061584302080) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OpSimpleTermContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(294)

		var _x = p.SimpleTerm()

		localctx.(*OpSimpleTermContext).st = _x
	}

	return localctx
}

// ISimpleTermContext is an interface to support dynamic dispatch.
type ISimpleTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleTermContext differentiates from other interfaces.
	IsSimpleTermContext()
}

type SimpleTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTermContext() *SimpleTermContext {
	var p = new(SimpleTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleTerm
	return p
}

func (*SimpleTermContext) IsSimpleTermContext() {}

func NewSimpleTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTermContext {
	var p = new(SimpleTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleTerm

	return p
}

func (s *SimpleTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleTermContext) AllOpTerm() []IOpTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpTermContext); ok {
			len++
		}
	}

	tst := make([]IOpTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpTermContext); ok {
			tst[i] = t.(IOpTermContext)
			i++
		}
	}

	return tst
}

func (s *SimpleTermContext) OpTerm(i int) IOpTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpTermContext)
}

func (s *SimpleTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleTerm(s)
	}
}

func (s *SimpleTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleTerm(s)
	}
}

func (p *GoliteParser) SimpleTerm() (localctx ISimpleTermContext) {
	this := p
	_ = this

	localctx = NewSimpleTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_simpleTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Term()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPLUS || _la == GoliteParserMINUS {
		{
			p.SetState(297)
			p.OpTerm()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOpTermContext is an interface to support dynamic dispatch.
type IOpTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetT returns the t rule contexts.
	GetT() ITermContext

	// SetT sets the t rule contexts.
	SetT(ITermContext)

	// IsOpTermContext differentiates from other interfaces.
	IsOpTermContext()
}

type OpTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	t      ITermContext
}

func NewEmptyOpTermContext() *OpTermContext {
	var p = new(OpTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_opTerm
	return p
}

func (*OpTermContext) IsOpTermContext() {}

func NewOpTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpTermContext {
	var p = new(OpTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_opTerm

	return p
}

func (s *OpTermContext) GetParser() antlr.Parser { return s.parser }

func (s *OpTermContext) GetOp() antlr.Token { return s.op }

func (s *OpTermContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpTermContext) GetT() ITermContext { return s.t }

func (s *OpTermContext) SetT(v ITermContext) { s.t = v }

func (s *OpTermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *OpTermContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, 0)
}

func (s *OpTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *OpTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterOpTerm(s)
	}
}

func (s *OpTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitOpTerm(s)
	}
}

func (p *GoliteParser) OpTerm() (localctx IOpTermContext) {
	this := p
	_ = this

	localctx = NewOpTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_opTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OpTermContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserPLUS || _la == GoliteParserMINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OpTermContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(304)

		var _x = p.Term()

		localctx.(*OpTermContext).t = _x
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *TermContext) AllOpUnaryTerm() []IOpUnaryTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpUnaryTermContext); ok {
			len++
		}
	}

	tst := make([]IOpUnaryTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpUnaryTermContext); ok {
			tst[i] = t.(IOpUnaryTermContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) OpUnaryTerm(i int) IOpUnaryTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpUnaryTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpUnaryTermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.UnaryTerm()
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserASTERISK || _la == GoliteParserFSLASH {
		{
			p.SetState(307)
			p.OpUnaryTerm()
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOpUnaryTermContext is an interface to support dynamic dispatch.
type IOpUnaryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetUt returns the ut rule contexts.
	GetUt() IUnaryTermContext

	// SetUt sets the ut rule contexts.
	SetUt(IUnaryTermContext)

	// IsOpUnaryTermContext differentiates from other interfaces.
	IsOpUnaryTermContext()
}

type OpUnaryTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	ut     IUnaryTermContext
}

func NewEmptyOpUnaryTermContext() *OpUnaryTermContext {
	var p = new(OpUnaryTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_opUnaryTerm
	return p
}

func (*OpUnaryTermContext) IsOpUnaryTermContext() {}

func NewOpUnaryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpUnaryTermContext {
	var p = new(OpUnaryTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_opUnaryTerm

	return p
}

func (s *OpUnaryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *OpUnaryTermContext) GetOp() antlr.Token { return s.op }

func (s *OpUnaryTermContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpUnaryTermContext) GetUt() IUnaryTermContext { return s.ut }

func (s *OpUnaryTermContext) SetUt(v IUnaryTermContext) { s.ut = v }

func (s *OpUnaryTermContext) UnaryTerm() IUnaryTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryTermContext)
}

func (s *OpUnaryTermContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *OpUnaryTermContext) FSLASH() antlr.TerminalNode {
	return s.GetToken(GoliteParserFSLASH, 0)
}

func (s *OpUnaryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpUnaryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpUnaryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterOpUnaryTerm(s)
	}
}

func (s *OpUnaryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitOpUnaryTerm(s)
	}
}

func (p *GoliteParser) OpUnaryTerm() (localctx IOpUnaryTermContext) {
	this := p
	_ = this

	localctx = NewOpUnaryTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoliteParserRULE_opUnaryTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OpUnaryTermContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserASTERISK || _la == GoliteParserFSLASH) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OpUnaryTermContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(314)

		var _x = p.UnaryTerm()

		localctx.(*OpUnaryTermContext).ut = _x
	}

	return localctx
}

// IUnaryTermContext is an interface to support dynamic dispatch.
type IUnaryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryTermContext differentiates from other interfaces.
	IsUnaryTermContext()
}

type UnaryTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryTermContext() *UnaryTermContext {
	var p = new(UnaryTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryTerm
	return p
}

func (*UnaryTermContext) IsUnaryTermContext() {}

func NewUnaryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryTermContext {
	var p = new(UnaryTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryTerm

	return p
}

func (s *UnaryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryTermContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(GoliteParserEXCLAMATION, 0)
}

func (s *UnaryTermContext) SelectorTerm() ISelectorTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorTermContext)
}

func (s *UnaryTermContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserMINUS, 0)
}

func (s *UnaryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryTerm(s)
	}
}

func (s *UnaryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryTerm(s)
	}
}

func (p *GoliteParser) UnaryTerm() (localctx IUnaryTermContext) {
	this := p
	_ = this

	localctx = NewUnaryTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoliteParserRULE_unaryTerm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(321)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserEXCLAMATION:
		{
			p.SetState(316)
			p.Match(GoliteParserEXCLAMATION)
		}
		{
			p.SetState(317)
			p.SelectorTerm()
		}

	case GoliteParserMINUS:
		{
			p.SetState(318)
			p.Match(GoliteParserMINUS)
		}
		{
			p.SetState(319)
			p.SelectorTerm()
		}

	case GoliteParserNUMBER, GoliteParserNEW, GoliteParserTRUE, GoliteParserFALSE, GoliteParserNIL, GoliteParserLEFTBRAC, GoliteParserID:
		{
			p.SetState(320)
			p.SelectorTerm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectorTermContext is an interface to support dynamic dispatch.
type ISelectorTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorTermContext differentiates from other interfaces.
	IsSelectorTermContext()
}

type SelectorTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorTermContext() *SelectorTermContext {
	var p = new(SelectorTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorTerm
	return p
}

func (*SelectorTermContext) IsSelectorTermContext() {}

func NewSelectorTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorTermContext {
	var p = new(SelectorTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorTerm

	return p
}

func (s *SelectorTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorTermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectorTermContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserDOT)
}

func (s *SelectorTermContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserDOT, i)
}

func (s *SelectorTermContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserID)
}

func (s *SelectorTermContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserID, i)
}

func (s *SelectorTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorTerm(s)
	}
}

func (s *SelectorTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorTerm(s)
	}
}

func (p *GoliteParser) SelectorTerm() (localctx ISelectorTermContext) {
	this := p
	_ = this

	localctx = NewSelectorTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoliteParserRULE_selectorTerm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Factor()
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserDOT {
		{
			p.SetState(324)
			p.Match(GoliteParserDOT)
		}
		{
			p.SetState(325)
			p.Match(GoliteParserID)
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) LEFTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserLEFTBRAC, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RIGHTBRAC() antlr.TerminalNode {
	return s.GetToken(GoliteParserRIGHTBRAC, 0)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GoliteParserNUMBER, 0)
}

func (s *FactorContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserFALSE, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL, 0)
}

func (s *FactorContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoliteParserRULE_factor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLEFTBRAC:
		{
			p.SetState(331)
			p.Match(GoliteParserLEFTBRAC)
		}
		{
			p.SetState(332)
			p.Expression()
		}
		{
			p.SetState(333)
			p.Match(GoliteParserRIGHTBRAC)
		}

	case GoliteParserID:
		{
			p.SetState(335)
			p.Match(GoliteParserID)
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoliteParserLEFTBRAC {
			{
				p.SetState(336)
				p.Arguments()
			}

		}

	case GoliteParserNUMBER:
		{
			p.SetState(339)
			p.Match(GoliteParserNUMBER)
		}

	case GoliteParserNEW:
		{
			p.SetState(340)
			p.Match(GoliteParserNEW)
		}
		{
			p.SetState(341)
			p.Match(GoliteParserID)
		}

	case GoliteParserTRUE:
		{
			p.SetState(342)
			p.Match(GoliteParserTRUE)
		}

	case GoliteParserFALSE:
		{
			p.SetState(343)
			p.Match(GoliteParserFALSE)
		}

	case GoliteParserNIL:
		{
			p.SetState(344)
			p.Match(GoliteParserNIL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
