/* eslint-disable indent */
// TIPOS
import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import FunctionCall from "../methods/call";
import Statement from "../models";

class Value extends Statement {
  private refType: DataType;

  constructor(
    token: TokenInfo,
    public props: {
      value: DataValue;
      generic?: DataType;
      type: DataType;
      fromCall?: FunctionCall;
    }
  ) {
    super(token, "Value");
    this.refType = this.props.type;
  }

  public compile(scope: Scope): boolean {
    if (this.props.type === DataType.ID) {
      if (scope) {
        const newValue: Value | undefined = scope.getVariable(
          this.props.value as string
        );
        if (newValue?.compile(scope)) {
          this.refType = newValue?.getType();
          this.props.generic = newValue.props.generic;
        }
      }
    }

    return true;
  }

  public getType(): DataType {
    return this.refType;
  }

  public getValue(env: Scope): DataValue | undefined {
    if (env) {
      if (typeof this.props.value !== "object") {
        switch (this.props.type) {
          case DataType.STRING:
            return this.props.value;
          case DataType.INTEGER:
            if (typeof this.props.value === "string")
              return parseInt(this.props.value as string, 10);
            else return this.props.value;
          case DataType.DOUBLE:
            if (typeof this.props.value === "string")
              return parseFloat(this.props.value as string);
            else return this.props.value;
          case DataType.BOOLEAN:
            if (typeof this.props.value === "string")
              return (this.props.value as string).toLowerCase() === "true";
            else return this.props.value;
          case DataType.CHARACTER:
            return (this.props.value as string).charAt(0);
          case DataType.ID:
            if (this.props.value) {
              const newValue: Value | undefined = env.getVariable(
                this.props.value as string
              );
              if (newValue?.compile(env)) {
                this.refType = newValue.getType();
                return newValue.getValue(env);
              }
            }
            break;
          default:
            return this.props.value;
        }
      } else return this.props.value;
    }
  }
}

export default Value;
