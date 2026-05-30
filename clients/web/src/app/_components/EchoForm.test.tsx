import { render, screen } from '@testing-library/react'
import { describe, expect, it } from 'vitest'
import { EchoForm } from './EchoForm'

describe('EchoForm', () => {
  it('入力欄と送信ボタンが描画される', () => {
    render(<EchoForm />)
    expect(screen.getByLabelText('Message:')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: '送信' })).toBeInTheDocument()
  })
})
