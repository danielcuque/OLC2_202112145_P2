import Scope from "../../runtime/scope";
import FunctionCall from "../methods/call";
import { TokenInfo } from "../../utils";
import Statement from "../models";

class Main extends Statement {
  constructor(public token: TokenInfo, public functionCall: FunctionCall) {
    super(token, "Main");
  }

  public compile(scope: Scope): boolean {
    this.functionCall.compile(scope);
    return true;
  }
}

export default Main;
