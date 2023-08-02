import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import Value from "../expression/value";
import FunctionCall from "./call";
import errors from "../../error";

class Append extends FunctionCall {
  constructor(
    token: TokenInfo,
    public props: { id: string; params: [Expression] }
  ) {
    super(token, { ...props, id: "Append" }, true);
    this.props.id = "Append";
  }

  public compile(scope: Scope): boolean {
    // OBTENER LISTA
    let compile = true;
    const list = scope.getVariable(this.props.id);
    if (list) {
      if (list.compile(scope)) {
        // OBTENER VALOR
        if (this.props.params[0].compile(scope)) {
          const newValue = this.props.params[0].getValue(scope);
          if (newValue && newValue.compile(scope)) {
            // VERIFICAR TIPOS
            if (
              newValue.getType() === list.props.generic ||
              `${DataType.DYNAMICLIST}<${newValue.props.generic}>` ===
                list.props.generic
            ) {
              // AGREGAR
              const temporal = list.getValue(scope) as DataValue[];
              temporal.push(newValue.getValue(scope) as DataValue);

              // GUARDAR
              scope.setVariable(
                this.props.id,
                new Value(this.token, {
                  value: temporal,
                  type: DataType.DYNAMICLIST,
                  generic: list.props.generic,
                })
              );
            } else {
              compile = false;
              errors.push({
                token: this.token,
                type: "Semántico",
                msg: `El no se puede asignar el tipo ${newValue.getType()} a ${
                  list.props.generic
                }.`,
              });
            }
          }
        }
      }
    } else {
      compile = false;
      errors.push({
        token: this.token,
        type: "Semántico",
        msg: `La lista ${this.props.id} no existe.`,
      });
    }

    return compile;
  }
}

export default Append;
