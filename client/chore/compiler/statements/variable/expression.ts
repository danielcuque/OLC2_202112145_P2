import DataType, { TokenInfo, DataValue } from "../../utils/types";
import Scope from "../../runtime/scope";
import { Expression, Value } from "../expression";
import AssignmentStatement from "./assignment";
import { defaultValues } from "../expression/tools";

class ExpresssionAssignment extends AssignmentStatement {
  constructor(
    token: TokenInfo,
    public props: {
      id: string;
      exp?: Expression;
    }
  ) {
    super(token, props.id);
  }

  public compile(scope: Scope, type: DataType): boolean {
    const nextValue: Value | undefined = this.getValue(scope, type);
    const value: Value | undefined = scope.getVariable(this.props.id);
    value?.compile(scope);
    return super.setValue(
      scope,
      value?.getType() ?? type,
      nextValue,
      type !== undefined
    );
  }

  public getValue(scope: Scope, type: DataType): Value | undefined {
    if (this.props.exp) {
      if (this.props.exp?.compile(scope)) {
        const value: Value | undefined = this.props.exp.getValue(scope);
        if (value) {
          return new Value(this.token, {
            value: value?.getValue(scope) as DataValue,
            type: value?.getType() as DataType,
            generic: value?.props.generic,
          });
        }
      }
    } else
      return new Value(this.token, {
        value: defaultValues(type),
        type: type,
      });
  }
}

export default ExpresssionAssignment;
