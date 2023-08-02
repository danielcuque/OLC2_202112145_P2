export const validateFilename = (filename: string) => {
  // Validamos que el nombre no sea una cadena vacía y que tenga extensión .tw
  const regex = /^([a-zA-Z0-9\s_\\.\-\(\):])+(.tw)$/i;
  return regex.test(filename);
};
