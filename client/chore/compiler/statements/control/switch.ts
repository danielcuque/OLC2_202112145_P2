import Scope from "../../runtime/scope";
import FunctionBlock from "../methods/functions";
import Expression from "../expression/data";
import { TokenInfo } from "../../utils";
import Statement from "../models";

// PROPS
interface CaseBody {
  case: Expression;
  body: Statement[];
}

class SwitchStatement extends Statement {
  // GLOBALES
  private isOnBreak = false;

  // CONSTRUCTOR
  constructor(
    token: TokenInfo,
    public props: {
      value: Expression;
      cases?: CaseBody[];
      default?: Omit<CaseBody, "case">;
    }
  ) {
    super(token, "Switch");
  }

  private addControlFunction(
    env: Scope,
    name: "return" | "break" | "continue"
  ) {
    env.addFunction(
      name,
      "void",
      new FunctionBlock(this.token, {
        id: name,
        type: "void",
        content: [
          {
            token: this.token,
            name: "FunctionCall",
            compile: () => {
              this.isOnBreak = true;
              return true;
            },
          },
        ],
        params: [],
      })
    );
  }

  public compile(scope: Scope) {
    const localScope = new Scope("Switch", "switch", scope);
    this.addControlFunction(localScope, "break");
    this.addControlFunction(localScope, "return");

    let compile = true;
    if (this.props.value.compile(localScope)) {
      let foundCase = false;
      for (
        let caseIndex = 0, length = this.props.cases?.length || 0;
        caseIndex < length;
        caseIndex++
      ) {
        if (this.props.cases?.length) {
          if (this.props.cases[caseIndex].case?.compile(localScope)) {
            if (
              this.props.cases[caseIndex].case
                .getValue(localScope)
                ?.getValue(localScope) ===
              this.props.value.getValue(localScope)?.getValue(localScope)
            ) {
              foundCase = true;
              const caseEnv = new Scope(
                "Case",
                `case_${caseIndex}`,
                localScope
              );
              for (
                let instructionIndex = 0,
                  instructionsLength = this.props.cases[caseIndex].body.length;
                instructionIndex < instructionsLength;
                instructionIndex++
              ) {
                if (!this.isOnBreak)
                  compile =
                    this.props.cases[caseIndex].body[instructionIndex].compile(
                      caseEnv
                    );
                else break;
              }
              if (this.isOnBreak) break;
            }
          }
        }
      }
      if (!foundCase) {
        if (this.props.default && this.props.default.body) {
          const defEnv = new Scope("Case", "default", localScope);
          for (
            let instructionIndex = 0,
              instructionsLength = this.props.default.body.length;
            instructionIndex < instructionsLength;
            instructionIndex++
          ) {
            if (!this.isOnBreak)
              compile =
                this.props.default.body[instructionIndex].compile(defEnv);
            else break;
          }
        }
      }
    } else compile = false;

    return compile;
  }
}

export default SwitchStatement;
