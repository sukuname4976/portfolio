import { render, screen } from '@testing-library/react'
import { describe, expect, it } from 'vitest'
import { RegisterForm } from './RegisterForm'

describe('RegisterForm', () => {
  it('名前・メール入力欄と登録ボタンが描画される', () => {
    render(<RegisterForm />)
    expect(screen.getByLabelText('Name:')).toBeInTheDocument()
    expect(screen.getByLabelText('Email:')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: '登録' })).toBeInTheDocument()
  })
})
