import FunctionBlock from "../statements/methods/functions";
import { Value } from "../statements/expression";
import { DataType, TokenInfo } from "../utils";
import errors from "../error";
import symbols from "../symbols";

type ScopeName =
  | "Function"
  | "Main"
  | "Global"
  | "Condition"
  | "Loop"
  | "Case"
  | "Switch";

class Scope {
  private variables: {
    [id: string]: {
      value: Value | undefined;
      type: DataType;
    };
  };

  private functions: {
    [id: string]: {
      value: FunctionBlock | undefined;
      type: DataType | "void";
    };
  };

  constructor(
    private name: ScopeName,
    private id: string,
    private prevEnv?: Scope
  ) {
    this.variables = {};
    this.functions = {};
  }

  public getPrevScope(): Scope | undefined {
    return this.prevEnv;
  }

  public getScopeName(): ScopeName {
    return this.name;
  }

  public getScopeId(): string {
    return this.id;
  }

  public getVariable(id: string): Value | undefined {
    if (id in this.variables) {
      return this.variables[id].value;
    }
    if (this.prevEnv) {
      return this.prevEnv.getVariable(id);
    } else {
      return undefined;
    }
  }

  public getFunction(id: string): FunctionBlock | undefined {
    if (id in this.functions) {
      if (this.functions[id]?.value) {
        return Object.create(
          Object.getPrototypeOf(this.functions[id]?.value),
          Object.getOwnPropertyDescriptors(this.functions[id]?.value)
        );
      }
    }
    if (this.prevEnv) {
      return this.prevEnv.getFunction(id);
    }
    return undefined;
  }

  public setVariable(id: string, newValue: Value): void {
    if (this.getVariable(id) !== undefined) {
      if (this.variables[id] !== undefined) {
        this.variables[id] = {
          value: newValue,
          type: this.variables[id].type,
        };
        const symbolIndex = symbols.findIndex(
          (symbol) =>
            symbol.name === id &&
            symbol.scope === this.name &&
            symbol.type === "Variable"
        );
        if (symbolIndex !== -1) {
          symbols[symbolIndex] = {
            ...symbols[symbolIndex],
            value:
              newValue.getValue(this) === undefined
                ? ""
                : (newValue.getValue(this) as string),
          };
        }
      } else {
        this.prevEnv?.setVariable(id, newValue);
      }
    } else {
      errors.push({
        type: "Semántico",
        msg: `La variable ${id} no ha sido declarada`,
        token: this.getVariable(id)?.token || ({} as TokenInfo),
      });
    }
  }

  public addVariable(id: string, type: DataType, value?: Value): void {
    if (this.variables[id] === undefined) {
      this.variables[id] = {
        value,
        type,
      };
      symbols.push({
        scope: this.name,
        type: "Variable",
        name: id,
        dataType: type,
        value: value
          ? value.getValue(this) === undefined
            ? ""
            : (value.getValue(this) as string)
          : "",
      });
    } else {
      errors.push({
        type: "Semántico",
        token: this.getVariable(id)?.token || ({} as TokenInfo),
        msg: `La variable ${id} ya ha sido declarada`,
      });
    }
  }

  public addFunction(
    id: string,
    type: DataType | "void",
    value: FunctionBlock
  ): void {
    if (this.functions[id] === undefined) {
      this.functions[id] = {
        value,
        type,
      };
      symbols.push({
        scope: this.name,
        type: "Función",
        name: id,
        dataType: type,
        value: "",
        params: value.props.params.map((param) => `${param.id}: ${param.type}`),
      });
    } else {
      errors.push({
        type: "Semántico",
        token: this.getVariable(id)?.token || ({} as TokenInfo),
        msg: `La función ${id} ya ha sido declarada`,
      });
    }
  }
}

export default Scope;
