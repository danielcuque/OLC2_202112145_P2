import { useContext } from "react";
import { TabContext } from "../context";

export const useSetTab = () => {
  const { setSelectTab } = useContext(TabContext);

  return {
    setSelectTab,
  };
};
