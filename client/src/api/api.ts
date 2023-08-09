// Fetch to the backend with native fetch

import { ParserError, TokenSymbol } from "./types";

const url = import.meta.env.VITE_API_URL;

export const fetchAPI = async (code: string) => {
    const res = await fetch(`${url}/compile`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ code }),
    });

    // Return symbols, errors, logs and cst
    const data: CompileResponse = await res.json();

    return data;
}

interface CompileResponse {
    symbols: TokenSymbol[];
    errors: ParserError[];
    logs: string[];
    cst: string;
}
