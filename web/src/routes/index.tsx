import { Title } from "@solidjs/meta";
import { action, useAction, useSubmission } from "@solidjs/router";
import { For, batch, createEffect, createSignal } from "solid-js";
import { ask as askGrpc } from "~/lib/grpc";

const askAction = action(askGrpc)

export default function Home() {
  const [question, setQuestion] = createSignal("")
  const asking = useSubmission(askAction)
  const [history, setHistory] = createSignal<{
    question: string
    answer: string
  }[]>([])

  createEffect(() => {
    if (asking.result) {
      batch(() => {
        setHistory(history => [ {
          question: question(),
          answer: asking.result ?? ""
        },...history])
        asking.clear()
      })
    }
  })
  return (
    <main>
      <Title>yes or no</Title>
      <h1>yes or no</h1>
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
