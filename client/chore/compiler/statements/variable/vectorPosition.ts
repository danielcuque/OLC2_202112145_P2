import { Expression, Value, VectorValue } from "../expression";
import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import AssignmentStatement from "./assignment";
import errors from "../../error";

class VectorPosition extends AssignmentStatement {
  constructor(
    token: TokenInfo,
    public props: {
      exp: Expression;
      value: VectorValue;
    }
  ) {
    super(token, props.value.props.value as string);
    this.id = props.value.props.value as string;
  }

  // COMPILAR
  public compile(env: Scope): boolean {
    if (this.id) {
      const refValue: Value | undefined = env.getVariable(this.id);
      if (refValue?.compile(env)) {
        return this.setValue(env, refValue.getType(), this.getValue(env));
      } else return false;
    } else return false;
  }

  // ASIGNAR VALOR
  public setValue(
    env: Scope,
    _type: DataType,
    value: Value | undefined
  ): boolean {
    let compile = true;

    if (this.id) {
      const refValue: Value | undefined = env.getVariable(this.id);
      if (refValue) {
        const temporal: DataValue[] | undefined = refValue.getValue(
          env
        ) as DataValue[];
        if (temporal) {
          if (this.props.value.compile(env)) {
            const index: number = this.props.value.getIndex();
            if (index >= 0 && index < temporal.length) {
              if (refValue.props.generic === value?.getType()) {
                const newValue = value?.getValue(env);

                if (newValue !== undefined && value?.compile(env) === true) {
                  temporal[index] = newValue;
                  env.setVariable(
                    this.id,
                    new Value(this.token, {
                      value: [...temporal],
                      type: refValue.getType(),
                      generic: refValue.props.generic,
                    })
                  );
                }
              } else
                errors.push({
                  type: "Semántico",
                  token: this.token,
                  msg: `No se puede asignar el tipo ${value?.getType()} a ${
                    refValue.props.generic
                  }`,
                });
            } else {
              compile = false;
              errors.push({
                type: "Semántico",
                token: this.token,
                msg: `La posicion ${index} esta fuera de rango para la lista ${this.props.value.props.value}.`,
              });
            }
          }
        } else {
          compile = false;
          errors.push({
            token: this.token,
            type: "Semántico",
            msg: `El arreglo ${this.id} esta vacio.`,
          });
        }
      } else {
        compile = false;
        errors.push({
          token: this.token,
          type: "Semántico",
          msg: `La variable ${this.id} no existe.`,
        });
      }
    }

    return compile;
  }

  public getValue(env: Scope): Value | undefined {
    if (this.props.exp.compile(env)) {
      const value: Value | undefined = this.props.exp.getValue(env);
      return value;
    }
  }
}

export default VectorPosition;
