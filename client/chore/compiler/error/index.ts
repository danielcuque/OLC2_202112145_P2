import { TokenInfo } from "../utils";

export interface ParserError {
  type: "Léxico" | "Sintáctico" | "Semántico";
  msg: string;
  token: TokenInfo;
}

const errors: ParserError[] = [];

export default errors;
