// Fetch to the backend with native fetch

import { ParserError, TokenSymbol } from "./types";

const url = import.meta.env.VITE_API_URL;

export const fetchAPI = async (code: string) => {
    const body = new FormData();
    body.append("code", code);
    const res = await fetch(`${url}/compile`, {
        method: "POST",
        body: body,
    });

    // Return symbols, errors, logs and cst
    const data: CompileResponse = await res.json();

    return data;
}

interface CompileResponse {
    symbols: TokenSymbol[];
    errors: ParserError[];
    compiled: string;
    optimized: string;
}
