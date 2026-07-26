import { type NextRequest, NextResponse } from 'next/server'
import { getBffServerClient } from '@/api/bff-server-client'
import type { components } from '@/api/generated/schema'

type CreateUserRequest = components['schemas']['CreateUserRequest']

// TODO(認証 issue): リクエスト cookie から user_token を抽出して BFF へ転送
// TODO(認証 issue): BFF が要求する API キーを Authorization ヘッダで付与
export async function POST(request: NextRequest) {
  let body: CreateUserRequest
  try {
    body = (await request.json()) as CreateUserRequest
  } catch {
    return NextResponse.json({ error: 'Invalid JSON' }, { status: 400 })
  }

  const { data, error, response } = await getBffServerClient().POST('/api/v1/users', {
    body,
  })

  if (error) {
    return NextResponse.json(error, { status: response.status })
  }
  return NextResponse.json(data, { status: 201 })
}
