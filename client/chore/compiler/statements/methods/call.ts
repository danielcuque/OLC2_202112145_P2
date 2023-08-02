import { DataType, DataValue, TokenInfo } from "../../utils";
import Scope from "../../runtime/scope";
import Expression from "../expression/data";
import Value from "../expression/value";
import FunctionBlock from "./functions";
import Statement from "../models";
import errors from "../../error";

class FunctionCall extends Statement {
  private functionValue: DataValue | undefined;
  private refType: DataType;

  constructor(
    token: TokenInfo,
    public props: {
      params: Expression[];
      id: string;
      generic?: DataType;
    },
    private builtIn: boolean = false
  ) {
    super(token, "FunctionCall");
    this.refType = DataType.ID;
  }

  public getValue(): DataValue | undefined {
    return this.functionValue;
  }

  public getType(): DataType | undefined {
    return this.refType;
  }

  public isBuiltIn(): boolean {
    return this.builtIn;
  }

  public compile(env: Scope): boolean {
    const functionBlock: FunctionBlock | undefined = env.getFunction(
      this.props.id
    );
    if (functionBlock) {
      functionBlock.setScope(env);
      const functionEnv: Scope | undefined = functionBlock.getScope();

      if (functionEnv) {
        const values: { value: Value; type: DataType }[] =
          this.props.params.map((exp: Expression) => ({
            value: exp.getValue(env),
            type: exp.getValue(env)?.getType(),
          })) as { value: Value; type: DataType }[];
        if (
          this.props.params
            .map((exp) => exp.compile(functionEnv))
            .every((exp) => exp === true)
        ) {
          if (functionBlock.props.params.length === this.props.params.length) {
            let compile = true;
            values.forEach((value, index: number) => {
              compile = value.value.compile(env);

              if (compile) {
                if (
                  value.type === functionBlock.props.params[index].type ||
                  value.value.props.generic ===
                    functionBlock.props.params[index].type ||
                  `${DataType.DYNAMICLIST}<${value.value.props.generic}>` ===
                    functionBlock.props.params[index].type
                ) {
                  if (value.value.compile(env)) {
                    const copy = new Value(this.token, {
                      value: value.value.getValue(env) ?? "",
                      type: value.value.getType(),
                      generic: value.value.props.generic,
                    });
                    functionEnv.addVariable(
                      functionBlock.props.params[index].id,
                      value.type,
                      copy
                    );
                  }
                } else {
                  errors.push({
                    type: "Semántico",
                    token: this.token,
                    msg: `Se esperaba un ${
                      functionBlock.props.params[index].type
                    } en el parámetro ${index + 1} en la función.`,
                  });
                  compile = false;
                }
              }
            });

            compile = functionBlock.compile();
            if (compile) {
              const functionValue = functionBlock.getFunctionValue();
              if (functionValue) {
                this.functionValue = functionValue?.getValue(env);
                this.refType = functionValue?.getType();
                this.props.generic = functionValue.props.generic;
              }
            }
            return compile;
          } else {
            errors.push({
              type: "Semántico",
              token: this.token,
              msg: `Se esperaban ${functionBlock.props.params.length} parametros pero se obtuvieron ${this.props.params.length} en la funcion ${this.props.id}`,
            });
            return false;
          }
        } else return false;
      } else return false;
    } else {
      errors.push({
        type: "Semántico",
        token: this.token,
        msg: `La funcion ${this.props.id} no existe.`,
      });
      return false;
    }
  }

  public getDot(): string {
    return "";
  }
}

export default FunctionCall;
