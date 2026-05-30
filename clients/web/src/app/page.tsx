import { getBffServerClient } from '@/api/bff-server-client'
import { EchoForm } from './_components/EchoForm'

export const dynamic = 'force-dynamic'

export default async function Home() {
  // ネットワーク例外 (BFF 未起動・DNS 失敗・タイムアウト等) は openapi-fetch の error と同形にフォールバックさせ、ページ全体の 500 化を避ける
  const { data, error } = await getBffServerClient()
    .POST('/api/v1/echo', {
      body: { message: 'Hello from web (Server Component)' },
    })
    .catch(() => ({
      data: undefined,
      error: { error: 'BFF に接続できません' },
    }))

  return (
    <main style={{ padding: '24px', fontFamily: 'sans-serif' }}>
      <h1>Portfolio Web Application</h1>

      <section style={{ marginTop: '24px' }}>
        <h2>初期表示 (Server Component から BFF を直接呼び出し)</h2>
        {error ? (
          <p style={{ color: 'crimson' }}>Error: {error.error}</p>
        ) : (
          <article>
            <p>Echo: {data.echo.message}</p>
            <p>
              Pokemon: #{data.pokemon.id} {data.pokemon.name} (
              {data.pokemon.types.map((t) => t.name).join(', ')})
            </p>
            <img src={data.pokemon.sprite} alt={data.pokemon.name} width={96} height={96} />
          </article>
        )}
      </section>

      <section style={{ marginTop: '32px' }}>
        <h2>サンプル: フォームから BFF を呼び出し (ブラウザ → Next Route Handler → BFF)</h2>
        <EchoForm />
      </section>
    </main>
  )
}
