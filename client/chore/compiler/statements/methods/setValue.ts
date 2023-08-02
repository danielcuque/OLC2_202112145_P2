import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import Value from "../expression/value";
import FunctionCall from "./call";
import errors from "../../error";

class SetValue extends FunctionCall {
  constructor(
    token: TokenInfo,
    public props: { id: string; params: [Expression, Expression] }
  ) {
    super(token, { ...props, id: "SetValue" }, true);
    this.props.id = "SetValue";
  }

  public compile(scope: Scope): boolean {
    // OBTENER LISTA
    let compile = true;
    const list = scope.getVariable(this.props.id);
    if (list) {
      if (list.compile(scope)) {
        // OBTENER VALOR
        if (this.props.params[0].compile(scope)) {
          const indexValue = this.props.params[0].getValue(scope);

          if (indexValue && indexValue.compile(scope)) {
            if (indexValue.getType() === DataType.INTEGER) {
              const newValue = this.props.params[1].getValue(scope);
              if (newValue && newValue.compile(scope)) {
                // VERIFICAR TIPOS
                if (
                  newValue.getType() === list.props.generic ||
                  `${DataType.DYNAMICLIST}<${newValue.props.generic}>` ===
                    list.props.generic
                ) {
                  // AGREGAR
                  const temporal = list.getValue(scope) as DataValue[];
                  const indexNum = indexValue.getValue(scope) as number;
                  temporal[indexNum] = newValue.getValue(scope) as DataValue;

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
                    }`,
                  });
                }
              }
            } else {
              compile = false;
              errors.push({
                token: this.token,
                type: "Semántico",
                msg: `La posicion de la lista dinamica ${this.props.id} debe ser un ${DataType.INTEGER}.`,
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
        msg: `La lista dinamica ${this.props.id} no existe.`,
      });
    }

    return compile;
  }
}

export default SetValue;
