import { useCreateFile, useFileUpload, useSetTab } from "../../hooks";

export const SidebarFiles = () => {
  const { listFiles, fileInputRef, handleFileChange } = useFileUpload();
  const { setSelectTab } = useSetTab();
  const {
    isFormNewFileOpen,
    inputRef,
    onSubmit,
    setIsFormNewFileOpen,
    setIsInputFocused,
  } = useCreateFile();

  return (
    <>
      <div className="w-[200px] text-white border-r-2 border-gray-4 h-screen">
        <div className="flex justify-between p-2 border-b-2 border-gray-4 h-[5vh]">
          <h3 className="font-medium">Archivos</h3>
          <div className="flex h-full">
            <button
              onClick={() => {
                fileInputRef.current?.click();
              }}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="currentColor"
                className="files-navbar-icons"
              >
                <path
                  fillRule="evenodd"
                  d="M11.47 2.47a.75.75 0 011.06 0l3.75 3.75a.75.75 0 01-1.06 1.06l-2.47-2.47V21a.75.75 0 01-1.5 0V4.81L8.78 7.28a.75.75 0 01-1.06-1.06l3.75-3.75z"
                  clipRule="evenodd"
                />
              </svg>
            </button>
            <input
              type="file"
              className="hidden"
              ref={fileInputRef}
              onChange={handleFileChange}
            />

            <button
              onClick={() => {
                setIsFormNewFileOpen(true);
                setIsInputFocused(true);
              }}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className="files-navbar-icons"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M12 4.5v15m7.5-7.5h-15"
                />
              </svg>
            </button>
          </div>
        </div>
        <div className="flex flex-col gap-2 mt-4 text-gray-7">
          {listFiles.map(({ id, name }, index) => (
            <button
              className="hover:bg-gray-4 w-full text-start px-3 py-2 hover:text-white"
              key={index}
              onClick={() => {
                setSelectTab(id);
              }}
            >
              {name}
            </button>
          ))}
          {/* Vamos a mostrar un input para crear un nuevo archivo */}
          {isFormNewFileOpen && (
            <form onSubmit={onSubmit}>
              <input
                ref={inputRef}
                type="text"
                className="w-full px-3 py-2 bg-gray-4 text-white"
                placeholder="Nuevo archivo"
              />
            </form>
          )}
        </div>
      </div>
    </>
  );
};
