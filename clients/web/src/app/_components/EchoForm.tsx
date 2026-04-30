'use client'

import { type FormEvent, useState } from 'react'
import type { components } from '@/api/generated/schema'

type EchoResponse = components['schemas']['EchoResponse']
type ErrorResponse = components['schemas']['ErrorResponse']

type Status =
  | { kind: 'idle' }
  | { kind: 'loading' }
  | { kind: 'success'; data: EchoResponse }
  | { kind: 'error'; message: string }

export function EchoForm() {
  const [message, setMessage] = useState('Hello from web (Client Component)')
  const [status, setStatus] = useState<Status>({ kind: 'idle' })

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    setStatus({ kind: 'loading' })

    const res = await fetch('/api/echo', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message }),
    })

    if (!res.ok) {
      const errorBody = (await res.json().catch(() => null)) as ErrorResponse | null
      setStatus({
        kind: 'error',
        message: errorBody?.error ?? `HTTP ${res.status}`,
      })
      return
    }

    const data = (await res.json()) as EchoResponse
    setStatus({ kind: 'success', data })
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label htmlFor="message-input">Message:</label>
        <input
          id="message-input"
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          style={{ marginLeft: '8px', marginRight: '8px', padding: '4px 8px' }}
        />
        <button type="submit" disabled={status.kind === 'loading'}>
          送信
        </button>
      </form>

      <div style={{ marginTop: '16px' }}>
        {status.kind === 'idle' && <p>送信ボタンを押してください</p>}
        {status.kind === 'loading' && <p>Loading...</p>}
        {status.kind === 'error' && <p style={{ color: 'crimson' }}>Error: {status.message}</p>}
        {status.kind === 'success' && (
          <article>
            <p>Echo: {status.data.echo.message}</p>
            <p>
              Pokemon: #{status.data.pokemon.id} {status.data.pokemon.name} (
              {status.data.pokemon.types.map((t) => t.name).join(', ')})
            </p>
            <img
              src={status.data.pokemon.sprite}
              alt={status.data.pokemon.name}
              width={96}
              height={96}
            />
          </article>
        )}
      </div>
    </div>
  )
}
