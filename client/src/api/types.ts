export interface ParserError {
    Line: string
    Column: string
    Msg: string
    Type: string
}

// export interface TokenSymbol {
//     name: string;
// }

export interface TokenSymbol {
    Column: number
    IsConst: boolean
    Line: number
    Name: string
    Scope: string
    Type: string
    Value: string

}
