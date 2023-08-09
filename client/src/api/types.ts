export interface ParserError {
    line: string
    column: string
    message: string
    type: string
}

export interface TokenSymbol {
    name: string;
}

// export interface TokenSymbol {
//     scope: string;
//     type: string;
//     name: string;
//     dataType: string;
//     value: string;
//     params?: string[];
//   }
