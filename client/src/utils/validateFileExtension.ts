import { TabI } from "../context";

export const validateFilename = (filename: string) => {
  // Validamos que el nombre no sea una cadena vacía y que tenga extensión .tw
  const regex = /^([a-zA-Z0-9\s_\\.\-\(\):])+(.swift)$/i;
  return regex.test(filename);
};


export interface ItemsProps {
  tab: TabI | undefined;
}