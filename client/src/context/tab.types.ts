import { ParserError, TokenSymbol } from "../api";

export interface FileI {
  id: number;
  code: string;
  name: string;
}

export interface TabI extends FileI {
  parser?: {
    errors: ParserError[];
    cst: string;
    symbols: TokenSymbol[];
    logs: unknown[];
  };
}

export interface TabState {
  tabs: TabI[];
  selectedTab: TabI | undefined;
}
