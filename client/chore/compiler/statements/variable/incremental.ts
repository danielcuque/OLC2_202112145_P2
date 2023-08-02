import DataType, { Operator, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Value from "../expression/value";
import AssignmentStatement from "./assignment";
import errors from "../../error";

class IncrementalAssignment extends AssignmentStatement {
  constructor(
    token: TokenInfo,
    public props: {
      id: string;
      operator: Operator;
    }
  ) {
    super(token, props.id);
  }

  public compile(scope: Scope): boolean {
    let compile = true;

    const refVariable = scope.getVariable(this.id ?? "");
    if (refVariable) {
      compile = refVariable?.compile(scope);
      if (compile) {
        if (
          refVariable.getType() === DataType.INTEGER ||
          refVariable.getType() === DataType.DOUBLE
        ) {
          // Verify data type
          compile = super.setValue(
            scope,
            refVariable.getType(),
            this.getValue(scope),
            false
          );
        } else {
          compile = false;
          errors.push({
            token: this.token,
            type: "Semántico",
            msg: `La variable ${this.id} debe ser del tipo ${DataType.INTEGER} o ${DataType.DOUBLE}`,
          });
        }
      }
    } else {
      compile = false;
      errors.push({
        token: this.token,
        type: "Semántico",
        msg: `La variable ${this.id} no está declarada.`,
      });
    }

    return compile;
  }

  public getValue(scope: Scope): Value | undefined {
    const refVariable = scope.getVariable(this.id ?? "");
    if (refVariable && refVariable?.compile(scope)) {
      return new Value(this.token, {
        value: (
          (refVariable?.getValue(scope) as number) +
          (this.props.operator === Operator.PLUSPLUS ? 1 : -1)
        ).toString(),
        type: refVariable.getType(),
      });
    }
  }
}

export default IncrementalAssignment;
