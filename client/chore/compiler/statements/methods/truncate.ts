import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import FunctionCall from "./call";
import errors from "../../error";

class Truncate extends FunctionCall {
  // GLOBALES
  private refValue: DataValue | undefined;

  // CONSTRUCTOR
  constructor(token: TokenInfo, props: { id: string; params: Expression[] }) {
    super(token, { ...props, id: "Truncate" }, true);
    this.props.id = "Truncate";
  }

  // OBTENER TIPO
  public getType(): DataType {
    return DataType.INTEGER;
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
          parseInt(
            (this.props.params[0]
              .getValue(scope)
              ?.getValue(scope)
              ?.toString() || "0") as string,
            10
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
}

export default Truncate;
