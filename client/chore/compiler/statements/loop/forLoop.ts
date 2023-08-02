import Scope from "../../runtime/scope";
import Declaration from "../variable/declaration";
import FunctionBlock from "../methods/functions";
import AssignmentStatement from "../variable/assignment";
import Expression from "../expression/data";
import { TokenInfo } from "../../utils";
import Statement from "../models";
import CycleControl from "./cycleLoop";

class ForLoop extends CycleControl {
  private isOnLoopContinue = false;
  private loopHandleBreak = false;
  private isOnLoopBreak = false;

  // CONSTRUCTOR
  constructor(
    token: TokenInfo,
    public props: {
      assignments: (AssignmentStatement | Declaration)[];
      withDeclarations?: boolean;
      modifiers: AssignmentStatement;
      condition: Expression;
      body: Statement[];
    }
  ) {
    super(token, { condition: props.condition, body: props.body });
  }

  private addLoopControlFunction(
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
              if (name == "continue") this.isOnLoopContinue = true;
              else this.isOnLoopBreak = true;
              return true;
            },
          },
        ],
        params: [],
      })
    );
  }

  public compile(scope: Scope): boolean {
    const containerEnv: Scope = new Scope("Loop", "for-container", scope);

    this.addLoopControlFunction(containerEnv, "continue");
    this.addLoopControlFunction(containerEnv, "return");
    this.addLoopControlFunction(containerEnv, "break");

    this.props.assignments.forEach((assignment) =>
      assignment.compile(this.props.withDeclarations ? containerEnv : scope)
    );

    while (
      this.props.condition.compile(containerEnv) &&
      this.props.condition.getValue(containerEnv)?.compile(containerEnv) &&
      this.props.condition.getValue(containerEnv)?.getValue(containerEnv) &&
      !this.loopHandleBreak &&
      !this.isOnLoopBreak
    ) {
      const localScope: Scope = new Scope("Loop", "for-content", containerEnv);

      for (
        let instructionCount = 0, length = this.props.body.length;
        instructionCount < length;
        instructionCount++
      ) {
        const instruction = this.props.body[instructionCount];
        if (!this.isOnLoopBreak && !this.isOnLoopContinue)
          this.loopHandleBreak = !instruction.compile(localScope);
        else {
          if (this.isOnLoopContinue) this.isOnLoopContinue = false;
          break;
        }
      }

      this.loopHandleBreak = !this.props.modifiers.compile(containerEnv);
    }
    return true;
  }
}

export default ForLoop;
