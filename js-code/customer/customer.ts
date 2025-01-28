import http from "@/http/http";
import product from "@/product/product";

var favorites: string[] = []

namespace customer {
  export function favorite(p: product.Product): boolean {
    return favorites.includes(p.Code)
  }


  export class Gateway {
    public async specialOffer(): Promise<string[]> {
      return http.client.request("customer_favorites", null)
    }
  }
}

export default customer
