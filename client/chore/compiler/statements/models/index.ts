import Scope from "../../runtime/scope";
import { TokenInfo } from "../../utils/types";

export type StatementName =
  | "Declaration"
  | "Assignment"
  | "Expression"
  | "Value"
  | "Function"
  | "FunctionCall"
  | "Main"
  | "Return"
  | "Condition"
  | "Loop"
  | "Break"
  | "Continue"
  | "Switch";

export default abstract class Statement {
  constructor(public token: TokenInfo, public name: StatementName) {}
  public abstract compile(scope: Scope, type?: string): boolean;
}
