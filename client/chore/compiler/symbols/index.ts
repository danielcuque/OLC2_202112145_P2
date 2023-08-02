export interface TokenSymbol {
  scope: string;
  type: string;
  name: string;
  dataType: string;
  value: string;
  params?: string[];
}

const symbols: TokenSymbol[] = [];
export default symbols;
