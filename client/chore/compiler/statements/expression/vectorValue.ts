import DataType, { DataValue, TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import errors from "../../error";
import Expression from "./data";
import Value from "./value";

class VectorValue extends Value {
  private index: number;
  private type: DataType;

  constructor(
    token: TokenInfo,
    public props: {
      index: Expression;
      value: string | DataValue[];
      type: DataType;
      generic?: DataType;
    }
  ) {
    super(token, props);
    this.index = -1;
    this.type = this.props.type;
  }

  public getIndex(): number {
    return this.index;
  }

  public getType(): DataType {
    return this.type;
  }

  // COMPILAR
  public compile(env: Scope): boolean {
    let compile = true;

    // COMPILAR INDICE
    if (this.props.index.compile()) {
      // VERIFICAR TIPO DE DATO DE EXPRESION INDEX
      const indexValue: Value | undefined = this.props.index.getValue(env);
      const index: number | undefined = indexValue?.getValue(env) as number;

      if (index !== undefined && indexValue?.getType() === DataType.INTEGER) {
        this.index = index;
      } else {
        compile = false;
        errors.push({
          type: "Semántico",
          token: this.token,
          msg: `La posicion del arreglo ${this.props.value} debe ser un ${DataType.INTEGER}.`,
        });
      }
    }

    // COMPILAR VARIABLE
    if (compile) {
      const newValue: Value | undefined = env.getVariable(
        this.props.value as string
      );
      if (newValue?.compile(env)) {
        this.type = newValue.props.generic || DataType.STRING;
        this.props.generic = undefined;
        this.props.type = newValue.props.generic || DataType.STRING;
      } else compile = false;
    }

    // RETORNAR VALIDACION
    return compile;
  }

  // OBTENER VALOR PRIMITIVO
  public getValue(env: Scope): DataValue | undefined {
    if (env) {
      if (this.props.value) {
        const newValue: Value | undefined = env.getVariable(
          this.props.value as string
        );
        if (newValue?.compile(env)) {
          const value = newValue.getValue(env) as DataValue[];

          // VERIFICAR POSICION
          if (this.index >= 0 && this.index < value.length) {
            return value[this.index];
          } else
            errors.push({
              type: "Semántico",
              token: this.token,
              msg: `La posición ${this.index} esta fuera de rango para la lista ${this.props.value}.`,
            });
        }
      }
    }
  }
}

export default VectorValue;
