package lexical 

const T_SELECT = 1;
const T_FROM = 2; 
const T_WHERE = 3; 
const T_ORDER = 4; 
const T_BY = 5; 
const T_LIMIT = 6; 
const T_DESC = 7; 
const T_WILD_CARD = 8; 
const T_COMMA = 9;  
const T_SEMICOLON = 10; 
const T_GREATER = 11; 
const T_SMALLER = 12; 
const T_GREATER_OR_EQUAL = 13; 
const T_SMALLER_OR_EQUAL= 14; 
const T_EQUAL = 15; 
const T_NOT_EQUAL = 16; 
const T_LITERAL = 17; 
const T_NUMERIC = 18; 
const T_EOF = 0;

var tokenNameMap map[uint8]string; 

func allocMapTokenNames() {
    if (len(tokenNameMap) == 0 ) {
        tokenNameMap = map[uint8]string{
            T_SELECT : "T_SELECT",
            T_FROM  : "T_FROM",
            T_WHERE  : "T_WHERE",
            T_ORDER : "T_ORDER" ,
            T_BY  : "T_BY",
            T_LIMIT  : "T_LIMIT",
            T_DESC  : "T_DESC",
            T_WILD_CARD  : "T_WILD_CARD",
            T_COMMA   : "T_COMMA",
            T_SEMICOLON : "T_SEMICOLON",
            T_GREATER : "T_GREATER",
            T_SMALLER : "T_SMALLER",
            T_GREATER_OR_EQUAL : "T_GREATER_OR_EQUAL",
            T_SMALLER_OR_EQUAL: "T_SMALLER_OR_EQUAL",
            T_EQUAL : "T_EQUAL",
            T_NOT_EQUAL : "T_NOT_EQUAL", 
            T_LITERAL : "T_LITERAL" ,
            T_NUMERIC : "T_NUMERIC", 
            T_EOF : "T_EOF",
        }   
    }
}

func TokenName(token uint8) string{
    allocMapTokenNames()
    return tokenNameMap[token]
}