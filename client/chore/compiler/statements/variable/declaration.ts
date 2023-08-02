import DataType, { TokenInfo } from "../../utils/types";
import Scope from "../../runtime/scope";
import AssignmentStatement from "./assignment";
import Statement from "../models";

class DeclarationStatement extends Statement {
  constructor(
    token: TokenInfo,
    public props: {
      type: DataType;
      assignments: AssignmentStatement[];
    }
  ) {
    super(token, "Declaration");
  }

  public compile(scope: Scope): boolean {
    const compiles = this.props.assignments.map(
      (assignment: AssignmentStatement) =>
        assignment.compile(scope, this.props.type)
    );
    return compiles.every((compile: boolean) => compile === true);
  }
}

export default DeclarationStatement;
