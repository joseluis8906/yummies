const store = new Map()
store.set("server.url", process.env.NEXT_PUBLIC_SERVER_URL)

namespace config {
  export function getString(key: string): string {
    return store.get(key)
  }
}

export default config
