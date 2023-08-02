import DataType, { TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Value from "../expression/value";
import Statement from "../models";
import errors from "../../error";

class AssignmentStatement extends Statement {
  public value?: Value;

  constructor(token: TokenInfo, public id?: string) {
    super(token, "Assignment");
  }

  public setValue(
    scope: Scope,
    type: DataType,
    value: Value | undefined,
    isNew: boolean = true
  ): boolean {
    if (this.id?.length) {
      if (value?.compile(scope)) {
        this.value = value;
        // Number exception
        const typeException =
          (type === DataType.DOUBLE && value.getType() === DataType.INTEGER) ||
          (type === DataType.INTEGER && value.getType() === DataType.DOUBLE) ||
          type === `${DataType.DYNAMICLIST}<${value.props.generic}>`;

        if (type === value.getType() || typeException) {
          if (isNew) {
            scope.addVariable(this.id, type, value);
          } else {
            scope.setVariable(this.id, value);
          }
          return true;
        } else {
          if (type) {
            errors.push({
              type: "Semántico",
              token: this.token,
              msg: `El tipo de dato de la variable ${
                this.id
              } es ${type} y no se puede asignar un valor de tipo ${value.getType()}.`,
            });
          } else {
            errors.push({
              type: "Semántico",
              token: this.token,
              msg: `La variable ${this.id} no ha sido declarada`,
            });
          }
          return false;
        }
      } else return false;
    } else return false;
  }

  public compile(_scope: Scope, _type?: DataType): boolean {
    return true;
  }
}
export default AssignmentStatement;
// Path: chore/compiler/statements/variable/variable.ts
