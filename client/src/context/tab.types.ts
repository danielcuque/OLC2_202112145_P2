import { ParserError } from "../../chore/compiler/error";
import { TokenSymbol } from "../../chore/compiler/symbols";

export interface FileI {
  id: number;
  code: string;
  name: string;
}

export interface TabI extends FileI {
  parser?: {
    errors: ParserError[];
    ast: string;
    symbols: TokenSymbol[];
    logs: unknown[];
  };
}

export interface TabState {
  tabs: TabI[];
  selectedTab: TabI | undefined;
}
