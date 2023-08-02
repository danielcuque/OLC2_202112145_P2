import { useContext, useEffect, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { TabContext } from "../context";
import { validateFilename } from "../utils";

export const useCreateFile = () => {
  const { addFile } = useContext(TabContext);
  const [isFormNewFileOpen, setIsFormNewFileOpen] = useState(false);
  const [isInputFocused, setIsInputFocused] = useState(false);

  const inputRef = useRef<HTMLInputElement>(null);
  const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const name = inputRef.current?.value || "";
    if (!validateFilename(name)) {
      toast.error("El nombre del archivo no es válido");
      return;
    }
    addFile("", name);
    setIsFormNewFileOpen(false);
  };

  useEffect(() => {
    if (inputRef.current) {
      inputRef.current.focus();
    }
  }, [isInputFocused]);

  return {
    onSubmit,
    setIsFormNewFileOpen,
    setIsInputFocused,
    isFormNewFileOpen,
    inputRef,
  };
};
