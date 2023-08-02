import { DataType, TokenInfo } from "../../utils";
import Scope from "../../runtime/scope";
import Value from "../expression/value";
import Statement from "../models";
import errors from "../../error";

class FunctionBlock extends Statement {
  private scope: Scope | undefined;
  private functionValue: Value | undefined;
  private isOnBreak: boolean = false;

  constructor(
    public token: TokenInfo,
    public props: {
      id: string;
      type: DataType | "void";
      content: Statement[];
      generic?: DataType;
      params: {
        type: DataType;
        generic?: DataType;
        id: string;
      }[];
    }
  ) {
    super(token, "Function");
    this.scope = {} as Scope;
    this.functionValue = undefined;
  }

  public getFunctionValue(): Value | undefined {
    return this.functionValue;
  }

  public setScope(scope: Scope): void {
    this.scope = new Scope("Function", this.props.id, scope);
    this.isOnBreak = false;
    this.functionValue = undefined;
    this.scope.addFunction(
      "return",
      "void",
      new FunctionBlock(this.token, {
        id: "return",
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

  public compile(): boolean {
    const compiles: boolean[] = [];
    for (
      let instructionIndex = 0, length = this.props.content.length;
      instructionIndex < length;
      instructionIndex++
    ) {
      if (this.scope) {
        if (!this.isOnBreak) {
          compiles.push(
            this.props.content[instructionIndex].compile(this.scope)
          );
        } else {
          break;
        }
      }
    }

    if (this.scope && "getVariable" in this.scope) {
      this.functionValue = this.scope?.getVariable("return");
    }

    if (this.props.type !== "void") {
      if (
        this.props.type === this.functionValue?.getType() ||
        this.props.type ===
          `${DataType.DYNAMICLIST}<${this.functionValue?.props.generic}>`
      ) {
        if (this.props.type === DataType.ARRAY) {
          if (this.props.generic !== this.functionValue?.props.generic)
            return false;
        }

        return compiles.every((result: boolean) => result === true);
      } else {
        errors.push({
          type: "Semántico",
          token: this.token,
          msg: `La función retorna un ${this.functionValue?.getType()} pero se esperaba un dato de tipo ${
            this.props.type
          }`,
        });
        return false;
      }
    } else return true;
  }

  public getScope(): Scope | undefined {
    return this.scope;
  }
}

export default FunctionBlock;
