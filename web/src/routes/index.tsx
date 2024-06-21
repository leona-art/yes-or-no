import { Title } from "@solidjs/meta";
import { action, useAction, useSubmission } from "@solidjs/router";
import { For, batch, createEffect, createMemo, createSignal } from "solid-js";
import { ask as askGrpc } from "~/lib/grpc";

const askAction = action(askGrpc)

export default function Home() {
  const [question, setQuestion] = createSignal("")
  const asking = useSubmission(askAction)
  type AskRecord={
    question: string
    answer: string
  }
  const [history, setHistory] = createSignal<AskRecord[]>([])
  
  // const record=createMemo(()=>({
  //   question: question(),
  //   answer: asking.result ?? ""
  // }))
  createEffect<AskRecord[]>((prev) => {
    if (asking.result) {
      const record={
        question: question(),
        answer: asking.result
      } satisfies AskRecord
      setHistory([...prev,record].reverse())
      asking.clear()
      setQuestion("")
      return [...prev,record]
    }else{
      return prev
    }
  },[] as AskRecord[])
  return (
    <main>
      <Title>yes or no</Title>
      <h2>これは何でしょうか？</h2>
      <form action={askAction.with(question())} method="post">
        <input type="text" value={question()} onInput={e => setQuestion(e.currentTarget.value)} />
        <button type="submit">Ask</button>
      </form>

      {asking.pending&&<p>Asking...</p>}
      {asking.result && <p>{asking.result}</p>}
      
      <For each={history()}>
        {({question, answer}) => (
          <div style={{ 
            "display":"flex"
           }}>
            <p>{question}</p>
            <p>{answer}</p>
          </div>
        )}
      </For>
    </main>
  );
}
