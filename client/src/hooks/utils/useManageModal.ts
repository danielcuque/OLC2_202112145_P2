import { useState } from "react";

export const useManageModal = () => {
  let [isOpen, setIsOpen] = useState<boolean>(false);
  const closeModal = () => setIsOpen(false);
  const openModal = () => setIsOpen(true);
  return { isOpen, closeModal, openModal };
};
