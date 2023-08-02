import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils/types";
import Expression from "../expression/data";
import FunctionCall from "./call";
import logs from "../../logs";

class Print extends FunctionCall {
  // CONSTRUCTOR
  constructor(
    public token: TokenInfo,
    public props: {
      params: Expression[];
      id: string;
    }
  ) {
    super(token, { ...props, id: "Print" }, true);
    this.props.id = "Print";
  }

  public compile(scope: Scope): boolean {
    this.props.params.forEach((exp) => {
      if (exp.compile(scope)) {
        if (exp.getValue(scope)?.compile(scope)) {
          logs.push(exp.getValue(scope)?.getValue(scope));
        }
      }
    });
    return true;
  }
}

export default Print;
