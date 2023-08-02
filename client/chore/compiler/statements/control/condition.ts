import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils";
import Expression from "../expression/data";
import Statement from "../models";

interface ConditionExp {
  exp: Expression;
  body: Statement[];
}

class ConditionStatement extends Statement {
  constructor(
    public token: TokenInfo,
    private props: {
      valid: ConditionExp;
      inValid?: ConditionExp;
      fallback?: ConditionExp[];
    }
  ) {
    super(token, "Condition");
  }

  public compile(scope: Scope): boolean {
    let compile = true;

    compile = this.props.valid.exp.compile();
    if (compile) {
      const value = this.props.valid.exp.getValue(scope);
      compile = value?.compile(scope) ?? true;
      if (compile && value?.getValue(scope)) {
        const environment = new Scope("Condition", "if", scope);
        compile = this.props.valid.body
          .map((instruction) => instruction.compile(environment))
          .every((compile) => compile === true);
      } else {
        const inValidCondition = () => {
          if (this.props.inValid) {
            const environment = new Scope("Condition", "else", scope);
            compile = this.props.inValid.body
              .map((instruction) => instruction.compile(environment))
              .every((compile) => compile === true);
          }
        };
        if (this.props.fallback) {
          let foundValid = false;
          for (
            let conditionIndex = 0, length = this.props.fallback.length;
            conditionIndex < length;
            conditionIndex++
          ) {
            compile = this.props.fallback[conditionIndex].exp.compile();
            if (compile) {
              const value =
                this.props.fallback[conditionIndex].exp.getValue(scope);
              compile = value?.compile(scope) ?? true;
              if (compile && value?.getValue(scope)) {
                foundValid = true;
                const environment = new Scope("Condition", "else if", scope);
                compile = this.props.fallback[conditionIndex].body
                  .map((instruction) => instruction.compile(environment))
                  .every((compile) => compile === true);
                break;
              }
            }
          }

          if (!foundValid) inValidCondition();
        } else inValidCondition();
      }
    }

    return compile;
  }

  public getConditionExp(): ConditionExp {
    return this.props.valid;
  }
}

export default ConditionStatement;
