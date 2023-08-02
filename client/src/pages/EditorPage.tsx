import { useContext, useEffect } from "react";
import { EditorContainer, NavbarEditor } from "../components";
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
        <EditorContainer tab={selectedTab} />
      </div>
    </>
  );
};
