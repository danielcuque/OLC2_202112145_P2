import { ParserError, TokenSymbol } from "../api";
import { TabI, TabState } from "./tab.types";

type TabsActionType =
  | {
    type: "ADD_TAB";
    payload: TabI;
  }
  | {
    type: "REMOVE_TAB";
  }
  | {
    type: "SET_ACTIVE_TAB";
    payload: TabI | undefined;
  }
  | {
    type: "UPDATE_SELECTED_TAB_CODE";
    payload: string;
  }
  | {
    type: "UPDATE_SELECTED_TAB_PARSER";
    payload: {
      symbols: TokenSymbol[];
      errors: ParserError[];
      compiled: string;
      optimized: string;
    };
  };

export const tabsReducer = (
  state: TabState,
  action: TabsActionType
): TabState => {
  switch (action.type) {
    case "ADD_TAB":
      return {
        ...state,
        tabs: [...state.tabs, action.payload],
      };
    case "SET_ACTIVE_TAB":
      return {
        ...state,
        selectedTab: action.payload,
      };
    case "REMOVE_TAB":
      if (!state.selectedTab) return state;
      return {
        ...state,
        selectedTab: undefined,
      };
    case "UPDATE_SELECTED_TAB_CODE":
      if (!state.selectedTab) return state;
      return {
        ...state,
        tabs: state.tabs.map((tab) => {
          if (tab.id === state.selectedTab?.id) {
            return {
              ...tab,
              code: action.payload,
            };
          }
          return tab;
        }),
        selectedTab: {
          ...state.selectedTab,
          code: action.payload,
        },
      };
    case "UPDATE_SELECTED_TAB_PARSER":
      if (!state.selectedTab) return state;
      return {
        ...state,
        tabs: state.tabs.map((tab) => {
          if (tab.id === state.selectedTab?.id) {
            return {
              ...tab,
              parser: {
                errors: [...action.payload.errors],
                compiled: action.payload.compiled,
                optimized: action.payload.optimized,
                symbols: [...action.payload.symbols],
              },
            };
          }
          return tab;
        }),
        selectedTab: {
          ...state.selectedTab,
          parser: {
            errors: [...action.payload.errors],
            compiled: action.payload.compiled,
            optimized: action.payload.optimized,
            symbols: [...action.payload.symbols],
          },
        },
      };

    default:
      return state;
  }
};
