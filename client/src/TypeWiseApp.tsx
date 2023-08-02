import { SidebarMain, SidebarFiles } from "./components/sidebar";
import { EditorPage } from "./pages";
import { Toaster } from "react-hot-toast";

export const TypeWiseApp = () => {
  return (
    <>
      <Toaster position="top-right" />
      <div className="flex justify-start items-center w-full min-h-screen">
        <SidebarMain />
        <SidebarFiles />
        <EditorPage />
      </div>
    </>
  );
};
