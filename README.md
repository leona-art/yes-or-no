## setup
1. APIキーを取得する\
   [Google AI SDK](https://aistudio.google.com/app/apikey)にアクセスし、APIキーを取得する。
2. APIキーを環境変数に設定する\
   ```bash
   export GEMINI_API_KEY="取得したAPIキー"
   ```
   ```powershell
   $env:GEMINI_API_KEY="取得したAPIキー"
   ```
3. バックエンドを起動する\
   ```bash
   cd backend
   go mod tidy
   go run cmd/server/main.go
   ```

4. フロントエンドを起動する\
   ```bash
   cd web
   npm install
   npm run dev
   ```

5. ブラウザで`http://localhost:3000`にアクセスする

## デモの説明
- デモの概要\
   本デモは、Gemini APIを用いて、トピックが何かを当てるゲームを提供するWebアプリケーションです。\
   ユーザーは、トピックが何かを当てるために質問を行い、Gemini APIを用いてその質問に対するyes/no/neitherの回答を取得します。\
   デモではトピックの内容は**お寿司**に固定されています。\
   