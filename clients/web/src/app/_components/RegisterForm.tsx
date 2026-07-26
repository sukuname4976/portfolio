'use client'

import { type FormEvent, useState } from 'react'
import type { components } from '@/api/generated/schema'

type UserResponse = components['schemas']['UserResponse']
type ErrorResponse = components['schemas']['ErrorResponse']

type Status =
  | { kind: 'idle' }
  | { kind: 'loading' }
  | { kind: 'success'; data: UserResponse }
  | { kind: 'error'; message: string }

export function RegisterForm() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [status, setStatus] = useState<Status>({ kind: 'idle' })

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    setStatus({ kind: 'loading' })

    const res = await fetch('/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, email }),
    })

    if (!res.ok) {
      const errorBody = (await res.json().catch(() => null)) as ErrorResponse | null
      setStatus({
        kind: 'error',
        message: errorBody?.error ?? `HTTP ${res.status}`,
      })
      return
    }

    const data = (await res.json()) as UserResponse
    setStatus({ kind: 'success', data })
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: '8px' }}>
          <label htmlFor="name-input">Name:</label>
          <input
            id="name-input"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            style={{ marginLeft: '8px', padding: '4px 8px' }}
          />
        </div>
        <div style={{ marginBottom: '8px' }}>
          <label htmlFor="email-input">Email:</label>
          <input
            id="email-input"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{ marginLeft: '8px', padding: '4px 8px' }}
          />
        </div>
        <button type="submit" disabled={status.kind === 'loading'}>
          登録
        </button>
      </form>

      <div style={{ marginTop: '16px' }}>
        {status.kind === 'idle' && <p>名前とメールを入力して登録してください</p>}
        {status.kind === 'loading' && <p>Loading...</p>}
        {status.kind === 'error' && <p style={{ color: 'crimson' }}>Error: {status.message}</p>}
        {status.kind === 'success' && (
          <article>
            <p>登録しました</p>
            <p>ID: {status.data.user.id}</p>
            <p>Name: {status.data.user.name}</p>
            <p>Email: {status.data.user.email}</p>
          </article>
        )}
      </div>
    </div>
  )
}
