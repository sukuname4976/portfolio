import 'server-only'
import createClient, { type Client } from 'openapi-fetch'
import type { paths } from '@/api/generated/schema'

let cached: Client<paths> | undefined

// TODO(認証 issue): BFF が要求する API キーを Authorization ヘッダで付与
// TODO(認証 issue): cookie/Supabase session から user_token を抽出してリクエストに付与
export function getBffServerClient(): Client<paths> {
  if (cached) return cached

  const baseUrl = process.env.BFF_BASE_URL
  if (!baseUrl) {
    throw new Error('BFF_BASE_URL is not set')
  }
  cached = createClient<paths>({ baseUrl })
  return cached
}
