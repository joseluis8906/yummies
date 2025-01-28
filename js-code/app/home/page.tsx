'use client'
import { useEffect, useState } from "react"
import Header from "@/app/home/header"
import Content from "@/app/home/content"
import component from "@/components/component"
import product from "@/product/product"
import layout from "@/components/layout"

export default function Home() {
  const [categories, setCategories] = useState<string[]>()
  const [specialOffer, setSpecialOffer] = useState<product.Product[]>()
  const [popularNow, setPopularNow] = useState<product.Product[]>()

  useEffect(() => {
    const productGW = new product.Gateway()
    productGW.specialOffer().then(res => setSpecialOffer(res)).catch(err => console.error(err))
    productGW.popularNow().then(res => setPopularNow(res)).catch(err => console.error(err))
    productGW.categories().then(res => setCategories(res)).catch(err => console.error(err))
  }, [])

  return (
    <layout.Container>
      <Header />
      <Content
        categories={categories}
        specialOffer={specialOffer?.[0]}
        popularNow={popularNow}
      />
      <component.BottomBar />
    </layout.Container>
  )
}
