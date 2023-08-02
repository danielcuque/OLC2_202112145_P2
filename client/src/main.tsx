import ReactDOM from "react-dom/client";
import { TypeWiseApp } from "./TypeWiseApp";
import "./index.css";
import { TabsProvider } from "./context";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <TabsProvider>
    <TypeWiseApp />
  </TabsProvider>
);
