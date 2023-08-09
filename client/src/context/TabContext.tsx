import { createContext } from "react";
import { TabState } from "./";
import { ParserError, TokenSymbol } from "../api";

interface TabContextProps extends TabState {
  addFile: (code: string, name: string) => void;
  setSelectTab: (id: number) => void;
  closeTab: () => void;
  updateSelectedTabCode: (code: string) => void;
  setParserAttributes: (
    symbols: TokenSymbol[],
    errors: ParserError[],
    logs: unknown[],
    ast: string | undefined
  ) => void;
}

export const TabContext = createContext({} as TabContextProps);
