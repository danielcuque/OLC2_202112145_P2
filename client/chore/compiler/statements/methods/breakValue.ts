import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils/types";
import Statement from "../models";

class BreakValue extends Statement {
  constructor(public token: TokenInfo) {
    super(token, "Break");
  }

  public compile(scope: Scope): boolean {
    let currentEnvironment: Scope | undefined = scope;

    const searchEnvironment = () => {
      if (
        currentEnvironment?.getScopeName() !== "Function" &&
        currentEnvironment?.getScopeName() !== "Loop" &&
        currentEnvironment?.getScopeName() !== "Switch"
      ) {
        if (currentEnvironment?.getPrevScope()) {
          currentEnvironment = currentEnvironment?.getPrevScope();
          searchEnvironment();
        } else return;
      } else return;
    };
    searchEnvironment();

    if (currentEnvironment) {
      const breakFunction = currentEnvironment.getFunction("break");
      if (breakFunction) breakFunction.compile();
      return true;
    } else return false;
  }
}

export default BreakValue;
