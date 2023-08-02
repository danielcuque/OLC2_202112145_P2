import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import FunctionCall from "./call";
import errors from "../../error";

class GetValue extends FunctionCall {
  private refValue: DataValue | undefined;
  private listType: DataType;

  constructor(
    token: TokenInfo,
    public props: { id: string; params: [Expression]; generic?: DataType }
  ) {
    super(token, { ...props, id: "GetValue" }, true);
    this.listType = DataType.ID;
    this.props.id = "GetValue";
  }

  // COMPILAR
  public compile(env: Scope): boolean {
    // OBTENER LISTA
    let compile = true;
    const list = env.getVariable(this.props.id);
    if (list) {
      if (list.compile(env)) {
        // OBTENER VALOR
        if (this.props.params[0].compile(env)) {
          const indexValue = this.props.params[0].getValue(env);
          if (indexValue && indexValue.compile(env)) {
            // VERIFICAR TIPOS
            if (indexValue.getType() === DataType.INTEGER) {
              const indexNum = indexValue.getValue(env) as number;
              const currentList = list.getValue(env) as DataValue[];

              if (indexNum >= 0 && indexNum < currentList.length) {
                this.refValue = currentList[indexNum];
                this.props.generic = list.props.generic ?? DataType.STRING;
                this.listType = list.getType();
              } else {
                compile = false;
                errors.push({
                  token: this.token,
                  type: "Semántico",
                  msg: `La posicion ${indexNum} esta fuera de rango para la lista ${this.props.id}.`,
                });
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

  // OBTENER VALOR
  public getValue(): DataValue | undefined {
    return this.refValue;
  }

  // OBTENER TIPO
  public getType(): DataType {
    return this.listType;
  }
}

export default GetValue;
