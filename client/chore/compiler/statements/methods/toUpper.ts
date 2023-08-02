import DataType, { DataValue, TokenInfo } from "../../utils/types";
import { Expression, FunctionCall } from "..";
import Scope from "../../runtime/scope";
import errors from "../../error";

class ToUpper extends FunctionCall {
  // GLOBALES
  private refValue: DataValue | undefined;

  // CONSTRUCTOR
  constructor(token: TokenInfo, props: { id: string; params: Expression[] }) {
    super(token, { ...props, id: "ToUpper" }, true);
    this.props.id = "ToUpper";
  }

  // COMPILAR
  public compile(scope: Scope): boolean {
    if (this.props.params[0] && this.props.params[0].compile(scope)) {
      // VERIFICAR TIPO
      if (this.props.params[0].getValue(scope)?.getType() === DataType.STRING) {
        this.refValue = this.props.params[0]
          .getValue(scope)
          ?.getValue(scope)
          ?.toString()
          .toUpperCase();
        return true;
      } else {
        errors.push({
          token: this.token,
          type: "Semántico",
          msg: `La funcion espera un ${DataType.STRING} como parametro.`,
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
    return DataType.STRING;
  }
}

export default ToUpper;
