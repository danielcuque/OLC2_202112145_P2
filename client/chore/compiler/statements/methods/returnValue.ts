import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils/types";
import Expression from "../expression/data";
import Value from "../expression/value";
import Statement from "../models";

class ReturnValue extends Statement {
  constructor(public token: TokenInfo, public props: { content: Expression }) {
    super(token, "Return");
  }

  public compile(scope: Scope): boolean {
    if (this.props.content.compile(scope)) {
      // BUSCAR ENTORNO DE FUNCION
      let currentScope: Scope | undefined = scope;

      const searchScope = () => {
        if (
          currentScope?.getScopeName() !== "Function" &&
          currentScope?.getScopeName() !== "Loop" &&
          currentScope?.getScopeName() !== "Switch"
        ) {
          if (currentScope?.getPrevScope()) {
            currentScope = currentScope?.getPrevScope();
            searchScope();
          } else return;
        } else return;
      };
      searchScope();

      // ASIGNAR RETORNO A FUNCION
      if (currentScope) {
        const value = this.props.content.getValue(scope);
        if (value?.compile(scope)) {
          // EVALUAR Y GUARDAR
          const newValue = new Value(this.token, {
            value: value.getValue(scope) as string,
            type: value.getType(),
            generic: value.props.generic,
          });
          currentScope.addVariable("return", value.getType(), newValue);
          const returnFunction = currentScope.getFunction("return");

          // EJECUTAR RETURN
          if (returnFunction) returnFunction.compile();
          return true;
        } else return false;
      } else return false;
    } else return false;
  }

  // OBTENER VALOR
  public getValue(env: Scope): Value | undefined {
    return this.props.content.getValue(env);
  }

  public getDot(): string {
    return "";
  }
}

export default ReturnValue;
