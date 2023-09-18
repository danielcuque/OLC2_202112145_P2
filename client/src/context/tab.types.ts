import { ParserError, TokenSymbol } from "../api";

export interface FileI {
  id: number;
  code: string;
  name: string;
}

export interface TabI extends FileI {
  parser?: {
    errors: ParserError[];
    symbols: TokenSymbol[];
    logs: unknown[];
    compiled: string;
    optimized: string;
  };
}

export interface TabState {
  tabs: TabI[];
  selectedTab: TabI | undefined;
}
