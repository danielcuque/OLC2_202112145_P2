import Editor from "@monaco-editor/react"
import { ItemsProps } from "../../utils"

export const CompiledContainer = ({ tab }: ItemsProps) => {
  return (
    <>
      <div className="w-1/2 h-full">
        <div className="flex h-full w-full">
          <Editor
            height="100%"
            width="100%"
            theme="vs-dark"
            defaultLanguage="c"
            language="c"
            value={tab?.parser?.compiled || ""}
            options={{ readOnly: true }}
          />
        </div>
      </div>
    </>
  )
}
