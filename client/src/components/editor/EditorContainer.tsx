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
      <div className="w-1/2 h-full">
        <div className="flex h-full w-full">
          <Editor
            height="100%"
            width="100%"
            theme="vs-dark"
            defaultLanguage="swift"
            language="swift"
            value={code}
            onChange={onChange}
            onMount={handleEditorDidMount}
          />
          
        </div>
      </div>
    </>
  );
};
