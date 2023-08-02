// El navbar solo podrá tener 1 hijo, que será el archivo actual

import { FC, useContext } from "react";
import { FileI, TabContext } from "../../context";

interface NavbarEditorProps {
  tab: FileI;
}

export const NavbarEditor: FC<NavbarEditorProps> = ({ tab: { name } }) => {
  const { closeTab } = useContext(TabContext);
  return (
    <>
      <nav className="w-full h-[5vh] text-white border-b-2 border-b-gray-6">
        <div className="flex gap-4 items-center h-full w-fit border-b-2 border-b-blue-300 border-r-2 border-r-gray-6 px-4">
          <div className="text-sm italic">{name}</div>
          <button
            onClick={() => {
              closeTab();
            }}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-4 h-4"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>
      </nav>
    </>
  );
};
