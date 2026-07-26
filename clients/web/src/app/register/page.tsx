import Link from 'next/link'
import { RegisterForm } from '../_components/RegisterForm'

export default function RegisterPage() {
  return (
    <main style={{ padding: '24px', fontFamily: 'sans-serif' }}>
      <h1>ユーザー登録</h1>
      <p>
        <Link href="/">← ホームへ戻る</Link>
      </p>

      <section style={{ marginTop: '24px' }}>
        <h2>登録フォーム (ブラウザ → Next Route Handler → BFF → user サービス)</h2>
        <RegisterForm />
      </section>
    </main>
  )
}
