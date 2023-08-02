import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils/types";
import Statement from "../models";

class ContinueValue extends Statement {
  constructor(public token: TokenInfo) {
    super(token, "Continue");
  }

  public compile(env: Scope): boolean {
    let currentEnvironment: Scope | undefined = env;

    const searchEnvironment = () => {
      if (currentEnvironment?.getScopeName() !== "Loop") {
        if (currentEnvironment?.getPrevScope()) {
          currentEnvironment = currentEnvironment?.getPrevScope();
          searchEnvironment();
        } else return;
      } else return;
    };
    searchEnvironment();

    if (currentEnvironment) {
      // EJECUTAR BREAK
      const continueFunction = currentEnvironment.getFunction("continue");
      if (continueFunction) continueFunction.compile();
      return true;
    } else return false;
  }
}

export default ContinueValue;
