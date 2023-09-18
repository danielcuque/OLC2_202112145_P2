import { useContext, useEffect, useState } from "react";
import { OptimizeContainer, ErrorContainer, SymbolsContainer } from "../reports";
import { RunButton, OptimizeButton } from "../button";
import { TabContext } from "../../context";

export const SidebarMain = () => {
  const { selectedTab } = useContext(TabContext);
  const [tab, setTab] = useState(selectedTab);

  useEffect(() => {
    setTab(selectedTab);
  }, [selectedTab]);

  return (
    <>
      <aside className="flex w-[100px] items-center gap-6 pt-4 flex-col text-white border-r-2 border-gray-4 h-screen">
        <RunButton tab={tab} />
        <OptimizeButton tab={tab} />
        <SymbolsContainer tab={tab} />
        <OptimizeContainer tab={tab} />
        <ErrorContainer tab={tab} />
      </aside>
    </>
  );
};
