import { TabI } from "../../context";
import { CustomModal } from "../modal";
import { Graphviz } from "graphviz-react";

interface ASTContainerProps {
  tab: TabI | undefined;
}

export const CSTContainer = ({ tab }: ASTContainerProps) => {
  return (
    <>
      <CustomModal
        icon={
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="main-sidebar-icons"
          >
            <path d="M12.378 1.602a.75.75 0 00-.756 0L3 6.632l9 5.25 9-5.25-8.622-5.03zM21.75 7.93l-9 5.25v9l8.628-5.032a.75.75 0 00.372-.648V7.93zM11.25 22.18v-9l-9-5.25v8.57a.75.75 0 00.372.648l8.628 5.033z" />
          </svg>
        }
        title="CST"
      >
        {tab?.parser?.ast ? (
          <div>
            <Graphviz
              dot={tab.parser?.ast}
              options={{
                zoom: true,
              }}
            />
          </div>
        ) : (
          <div>
            <p>No se ha generado CST</p>
          </div>
        )}
      </CustomModal>
    </>
  );
};
