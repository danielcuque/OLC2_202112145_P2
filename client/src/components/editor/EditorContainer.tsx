import { FC, useContext, useRef } from "react";
import Editor, { OnMount } from "@monaco-editor/react";
import { TabContext, TabI } from "../../context";

interface EditorContainerProps {
  tab: TabI;
}

export const EditorContainer: FC<EditorContainerProps> = ({ tab }) => {
  const editorRef = useRef<any>(null);
  const { updateSelectedTabCode } = useContext(TabContext);

  const code = tab.code;
  const logs = tab.parser?.logs || [];

  const onChange = (value: string | undefined) => {
    updateSelectedTabCode(value || "");
  };

  const handleEditorDidMount: OnMount = (editor: any) => {
    editorRef.current = editor;
    editor.updateOptions({
      minimap: {
        enabled: false,
      },
    });
  };

  return (
    <>
      <div className="w-full h-full">
        <div className="h-[500px] w-full">
          <Editor
            height="100%"
            theme="vs-dark"
            defaultLanguage="java"
            value={code}
            onChange={onChange}
            onMount={handleEditorDidMount}
          />
        </div>
        <div className="w-full h-[calc(100vh-560px)]">
          <div className="flex gap-2 text-gray-7 font-semibold px-2 items-center">
            Consola
            {logs.length > 0 && (
              <span className="flex bg-blue-500 w-5 h-5 rounded-full text-xs font-normal text-white text-center justify-center items-center">
                {logs.length > 99 ? "99+" : logs.length}
              </span>
            )}
          </div>
          <div className="h-full overflow-y-auto">
            <textarea
              value={logs.join("\n")}
              className="h-full w-full bg-gray-2 text-white p-1"
              readOnly
            ></textarea>
          </div>
        </div>
      </div>
    </>
  );
};
