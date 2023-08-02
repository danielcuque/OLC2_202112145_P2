import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import FunctionCall from "./call";
import errors from "../../error";

class Round extends FunctionCall {
  // GLOBALES
  private refValue: DataValue | undefined;

  // CONSTRUCTOR
  constructor(token: TokenInfo, props: { id: string; params: Expression[] }) {
    super(token, { ...props, id: "Round" }, true);
    this.props.id = "Round";
  }

  // COMPILAR
  public compile(scope: Scope): boolean {
    if (this.props.params[0] && this.props.params[0].compile(scope)) {
      // VERIFICAR TIPO
      if (
        this.props.params[0].getValue(scope)?.getType() === DataType.DOUBLE ||
        this.props.params[0].getValue(scope)?.getType() === DataType.INTEGER
      ) {
        this.refValue =
          Math.round(
            (this.props.params[0].getValue(scope)?.getValue(scope) ||
              0) as number
          ) || 0;
        return true;
      } else {
        errors.push({
          token: this.token,
          type: "Semántico",
          msg: `La funcion espera un ${DataType.DOUBLE} | ${DataType.INTEGER} como parametro.`,
        });
        return false;
      }
    } else return false;
  }

  // OBTENER VALOR
  public getValue(): DataValue | undefined {
    return this.refValue;
  }

  // OBTENER TIPO
  public getType(): DataType {
    return DataType.INTEGER;
  }
}

export default Round;
