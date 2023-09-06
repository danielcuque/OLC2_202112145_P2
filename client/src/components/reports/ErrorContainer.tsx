import { TabI } from "../../context";
import { CustomModal } from "../modal";

interface ErrorContainerProps {
  tab: TabI | undefined;
}

export const ErrorContainer = ({ tab }: ErrorContainerProps) => {
  const errors = tab?.parser?.errors || [];
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
            <path
              fillRule="evenodd"
              d="M9.401 3.003c1.155-2 4.043-2 5.197 0l7.355 12.748c1.154 2-.29 4.5-2.599 4.5H4.645c-2.309 0-3.752-2.5-2.598-4.5L9.4 3.003zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z"
              clipRule="evenodd"
            />
          </svg>
        }
        title="Tabla de errores"
      >
        <div className="w-full">
          <div className="bg-blue-600 w-full flex flex-row rounded-t-lg">
            <div className="w-full text-center p-4 text-white font-semibold text-lg">
              Tipo
            </div>
            <div className="w-full text-center p-4 text-white font-semibold text-lg">
              Mensaje
            </div>
            <div className="w-full text-center p-4 text-white font-semibold text-lg">
              Fila
            </div>
            <div className="w-full text-center p-4 text-white font-semibold text-lg">
              Columna
            </div>
          </div>
          <div className="h-full max-h-[30rem] w-full overflow-y-auto [&>*:nth-child(odd)]:bg-gray-100 [&>*:nth-child(even)]:bg-white">
            {errors.length > 0 ? (
              errors.map(({ Column, Line, Msg, Type }, index) => (
                <div
                  key={index}
                  className="w-full flex flex-row py-6 border border-gray-300"
                >
                  <div className="w-full text-center">{Type}</div>
                  <div className="w-full text-center"> {Msg}</div>
                  <div className="w-full text-center"> {Line}</div>
                  <div className="w-full text-center"> {Column}</div>
                </div>
              ))
            ) : (
              <div className="w-full flex flex-row py-6 border border-gray-300">
                <div className="w-full text-center">No hay errores</div>
              </div>
            )}
          </div>
        </div>
      </CustomModal>
    </>
  );
};
