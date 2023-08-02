import { FC, ReactNode, useReducer } from "react";
import { TabContext } from "./TabContext";
import { TabI, TabState } from "./tab.types";
import { tabsReducer } from "./tabsReducer";
import { ParserError } from "../../chore/compiler/error";
import { TokenSymbol } from "../../chore/compiler/symbols";

interface TabsProviderProps {
  children: ReactNode;
}

const INITIAL_TAB: TabI = {
  id: 0,
  code: `// Bienvenido a T-Swift
  `,
  name: "example.swift",
};

const TABS_INITIAL_STATE: TabState = {
  tabs: [INITIAL_TAB],
  selectedTab: INITIAL_TAB,
};

export const TabsProvider: FC<TabsProviderProps> = ({ children }) => {
  const [state, dispatch] = useReducer(tabsReducer, TABS_INITIAL_STATE);

  const addFile = (code: string, name: string) => {
    dispatch({
      type: "ADD_TAB",
      payload: {
        id: state.tabs.length + 1,
        code,
        name,
      },
    });
  };
  const setSelectTab = (id: number) => {
    const selectedTab = state.tabs.find((tab) => tab.id === id);
    dispatch({
      type: "SET_ACTIVE_TAB",
      payload: selectedTab,
    });
  };

  const updateSelectedTabCode = (code: string) => {
    dispatch({
      type: "UPDATE_SELECTED_TAB_CODE",
      payload: code,
    });
  };

  const closeTab = () => {
    dispatch({
      type: "REMOVE_TAB",
    });
  };

  const setParserAttributes = (
    symbols: TokenSymbol[],
    errors: ParserError[],
    logs: unknown[],
    ast: any
  ) => {
    dispatch({
      type: "UPDATE_SELECTED_TAB_PARSER",
      payload: {
        errors,
        ast,
        symbols,
        logs,
      },
    });
  };

  return (
    <TabContext.Provider
      value={{
        ...state,
        closeTab,
        updateSelectedTabCode,
        setParserAttributes,
        addFile,
        setSelectTab,
      }}
    >
      {children}
    </TabContext.Provider>
  );
};
