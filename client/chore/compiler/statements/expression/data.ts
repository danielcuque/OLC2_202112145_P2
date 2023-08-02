import { TokenInfo, Operator } from "../../utils/types";
import Scope from "../../runtime/scope";
import Statement from "../models";
import operateValues from "./tools";
import Value from "./value";

class Expression extends Statement {
  private childToken: TokenInfo;

  constructor(
    token: TokenInfo,
    public props: Partial<{
      value: Value;
      left: Expression;
      right: Expression;
      operator: Operator;
      condition?: Expression;
    }>
  ) {
    super(token, "Expression");
    this.childToken = token;
  }

  public compile(_scope?: Scope): boolean {
    return true;
  }

  public getValue(scope: Scope): Value | undefined {
    const leftExp: Value | undefined = this.props.left?.getValue(scope);
    const rightExp: Value | undefined = this.props.right?.getValue(scope);
    const condition: Value | undefined = this.props.condition?.getValue(scope);

    if (leftExp) {
      if (this.props.operator) {
        const result: Value | undefined = operateValues(
          scope,
          this.childToken,
          leftExp,
          this.props.operator,
          rightExp,
          condition
        );
        if (result) return result;
      } else {
        if (leftExp.compile(scope)) return leftExp;
      }
    } else if (this.props.value) {
      if (this.props.value.compile(scope)) return this.props.value;
    }
  }
}

export default Expression;
