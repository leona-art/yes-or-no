import { createPromiseClient } from "@connectrpc/connect";
import { YesOrNoService } from "../../gen/yesorno_connect";
import { createConnectTransport } from "@connectrpc/connect-node";



export async function ask(question: string) {
    "use server"
    const transport = createConnectTransport({
        baseUrl: "http://backend:8080",
        httpVersion: "1.1"
    });
    const client = createPromiseClient(YesOrNoService, transport);
    const res = await client.ask({
        question
    });
    return {
        answer: res.answer,
        finish:res.finish
    };
}