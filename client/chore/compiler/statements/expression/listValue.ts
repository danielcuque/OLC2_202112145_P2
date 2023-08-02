import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import { Value } from "../expression";

class DynamicListValue extends Value {
  constructor(
    token: TokenInfo,
    public props: { value: DataValue; generic?: DataType; type: DataType }
  ) {
    super(token, {
      value: [],
      type: DataType.DYNAMICLIST,
      generic: props.type,
    });
    this.props.generic = props.type;
  }

  public compile(_scope: Scope): boolean {
    return true;
  }

  // OBTENER VALOR
  public getValue(): DataValue | undefined {
    return [];
  }
}

export default DynamicListValue;
