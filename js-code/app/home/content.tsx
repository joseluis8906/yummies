import React from "react"
import product from "@/product/product";
import layout from "@/components/layout"
import component from "@/components/component"
import spacing from "@/components/spacing"

interface Props {
  categories: string[] | undefined;
  specialOffer: product.Product[] | undefined;
  popularNow: product.Product[] | undefined;
}

export default function Content(props: Props) {
  return (
    <layout.Content>
      <spacing.P2></spacing.P2>
      <component.Categories value={props.categories} />
      <spacing.P2></spacing.P2>
      <component.SpecialOffer value={props.specialOffer} />
      <spacing.P2></spacing.P2>
      <component.PopularNow value={props.popularNow} />
    </layout.Content>
  )
}
