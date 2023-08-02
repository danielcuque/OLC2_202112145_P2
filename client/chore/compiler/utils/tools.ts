import type { JISONTokenInfo, TokenInfo } from "./types";

const getToken = (token: JISONTokenInfo): TokenInfo => ({
  line: token.first_line,
  col: token.first_column + 1,
});

export default getToken;
