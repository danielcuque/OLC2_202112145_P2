export enum DataType {
  ID = "ID",
  ARRAY = "ARRAY",
  STRING = "STRING",
  BOOLEAN = "BOOLEAN",
  INTEGER = "INTEGER",
  DOUBLE = "DOUBLE",
  DYNAMICLIST = "LIST",
  CHARACTER = "CHAR",
}

export type DataValue = string | number | boolean | DataValue[];

export enum Operator {
  AND = "&&",
  DIVISION = "/",
  EQUALSEQUALS = "==",
  MAJOR = ">",
  MOREOREQUALS = ">=",
  MINOR = "<",
  LESSOREQUALS = "<=",
  MINUS = "-",
  MINUSMINUS = "--",
  MODULE = "%",
  TIMES = "*",
  NONEQUALS = "!=",
  NEGATION = "-*",
  NOT = "!",
  OR = "||",
  PLUS = "+",
  PLUSPLUS = "++",
  POWER = "^",
  TERNARY = "?:",
}

export interface JISONTokenInfo {
  first_line: number;
  last_line: number;
  first_column: number;
  last_column: number;
}

export interface TokenInfo {
  line: number;
  col: number;
}

export default DataType;
