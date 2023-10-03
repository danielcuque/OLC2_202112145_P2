import { useContext, useEffect } from "react";
import { CompiledContainer, EditorContainer, NavbarEditor } from "../components";
import { TabContext } from "../context";

export const EditorPage = () => {
  const { selectedTab } = useContext(TabContext);

  useEffect(() => {}, [selectedTab]);

  if (!selectedTab)
    return (
      <>
        <div className="flex flex-1 text-3xl text-white justify-center align-middle">
          <h1 className="text-center">
            Abrir un archivo o crear uno nuevo para comenzar a editar
          </h1>
        </div>
      </>
    );
  return (
    <>
      <div className="flex flex-1 flex-col h-screen">
        <NavbarEditor tab={selectedTab} />
        <div className="flex w-full h-full">
          <EditorContainer tab={selectedTab} />
          <CompiledContainer tab={selectedTab} />
        </div>
      </div>
    </>
  );
};
