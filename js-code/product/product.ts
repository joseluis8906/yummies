import http from "@/http/http"
import money from "@/money/money"

namespace product {
  export type Product = {
    Code: string
    Name: string
    Description: string
    Price: money.Money
    Image: string
    Categories: string[]
    SpecialOffer: boolean
    PopularNow: boolean
    Discount: number
  }

  export class Gateway {
    public async specialOffer(): Promise<Product[]> {
      return http.client.request("product_specialOffer", null)
    }

    public async popularNow(): Promise<Product[]> {
      return http.client.request("product_popularNow", null)
    }

    public async categories(): Promise<string[]> {
      return http.client.request("product_categories", null)
    }
  }
}

export default product
