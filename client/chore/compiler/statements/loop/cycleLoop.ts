import Scope from "../../runtime/scope";
import FunctionBlock from "../methods/functions";
import { TokenInfo } from "../../utils/types";
import Expression from "../expression/data";
import Statement from "../models";

class CycleControl extends Statement {
  private isOnContinue = false;
  private handleBreak = false;
  private isOnBreak = false;

  // CONSTRUCTOR
  constructor(
    token: TokenInfo,
    public props: {
      condition: Expression;
      body: Statement[];
      isDoLoop?: boolean;
    }
  ) {
    super(token, "Loop");
  }

  // AGREGAR FUNCION DE SALIDA
  private addControlFunction(
    scope: Scope,
    name: "return" | "break" | "continue"
  ) {
    scope.addFunction(
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
              if (name === "continue") this.isOnContinue = true;
              else this.isOnBreak = true;
              return true;
            },
          },
        ],
        params: [],
      })
    );
  }

  public compile(env: Scope): boolean {
    const containerScope: Scope = new Scope("Loop", "while-container", env);
    this.addControlFunction(containerScope, "continue");
    this.addControlFunction(containerScope, "return");
    this.addControlFunction(containerScope, "break");

    const runInstructions = (id: string) => {
      const localEnv: Scope = new Scope("Loop", id, containerScope);
      for (
        let instructionCount = 0, length = this.props.body.length;
        instructionCount < length;
        instructionCount++
      ) {
        const instruction = this.props.body[instructionCount];
        if (!this.isOnBreak && !this.isOnContinue)
          this.handleBreak = !instruction.compile(localEnv);
        else {
          if (this.isOnContinue) this.isOnContinue = false;
          break;
        }
      }
    };

    if (!this.props.isDoLoop)
      while (
        this.props.condition.compile(containerScope) &&
        this.props.condition
          .getValue(containerScope)
          ?.compile(containerScope) &&
        this.props.condition
          .getValue(containerScope)
          ?.getValue(containerScope) &&
        !this.handleBreak &&
        !this.isOnBreak
      )
        runInstructions("while-content");
    else
      do runInstructions("do while-content");
      while (
        this.props.condition.compile(env) &&
        this.props.condition.getValue(env)?.compile(env) &&
        this.props.condition.getValue(env)?.getValue(env) &&
        !this.handleBreak &&
        !this.isOnBreak
      );

    return true;
  }
}

export default CycleControl;
