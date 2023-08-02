import { useContext, useEffect, useRef, useState } from "react";
import { FileI, TabContext } from "../context";
import { validateFilename } from "../utils";
import { toast } from "react-hot-toast";

export const useFileUpload = () => {
  const { addFile, tabs: files } = useContext(TabContext);
  const [listFiles, setListFiles] = useState<FileI[]>(files);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files && event.target.files[0];

    if (file) {
      const reader = new FileReader();
      reader.readAsText(file, "UTF-8");
      reader.onload = () => {
        if (validateFilename(file.name) === false) {
          toast.error("El archivo no tiene la extensión .tw");
        } else {
          const code = reader.result ? reader.result : "";
          addFile(code as string, file.name);
          setListFiles((prev) => [
            ...prev,
            {
              id: prev.length + 1,
              code: code as string,
              name: file.name,
            },
          ]);
        }
      };
    }
    fileInputRef.current?.value && (fileInputRef.current.value = "");
  };

  useEffect(() => {
    setListFiles(files);
  }, [files]);

  return { listFiles, fileInputRef, handleFileChange };
};
