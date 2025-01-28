import React from "react"
import product from "@/product/product";
import customer from "@/customer/customer"
import layout from "@/components/layout"
import component from "@/components/component"
import scroll from "@/components/scroll"
import text from "@/components/text"
import spacing from "@/components/spacing"

interface Props {
  categories: string[] | undefined;
  specialOffer: product.Product | undefined;
  popularNow: product.Product[] | undefined;
}

export default function Content(props: Props) {
  return (
    <layout.Content>
      <spacing.P2></spacing.P2>
      <component.Categories value={props.categories} />
      <spacing.P2></spacing.P2>
      <SpecialOffer value={props.specialOffer} />
      <spacing.P2></spacing.P2>
      <PopularNow value={props.popularNow} />
    </layout.Content>
  )
}

function SpecialOffer(props: { value: product.Product | undefined }) {
  return (
    <>
      <spacing.P4 v={false}>
        <text.H3>Today's Special Offer</text.H3>
      </spacing.P4>
      <spacing.P1></spacing.P1>
      <scroll.H>
        <div className="w-full dark:bg-ctp-base rounded-lg flex flex-row shadow-md shadow-ctp-crust mb-2">
          <div className="basis-1/2">
            <img src={props.value?.Image} className="w-full h-full object-cover" />
          </div>
          <div className="basis-1/2 content-center justify-center text-center p-2">
            <h4 className="text-md dark:text-ctp-text font-semibold text-center">{props.value?.Name}</h4>
            <span className="dark:text-ctp-text text-sm">Now</span>
            <h3 className="dark:text-ctp-text text-2xl font-bold">{props.value?.Price?.Amount}</h3>
            <span className="dark:text-ctp-red text-sm">({props.value?.Discount}% off)</span>
            <button type="button" className="font-medium rounded-lg text-sm px-5 py-2.5 text-center border dark:border-ctp-peach dark:text-ctp-peach dark:hover:text-ctp-base dark:hover:bg-ctp-peach block mx-auto mt-3">Add to Cart</button>
          </div>
        </div>
      </scroll.H>
    </>
  )
}

function PopularNow(props: { value: product.Product[] | undefined }) {
  return (
    <>
      <spacing.P4 v={false}>
        <text.H3>Popular Now</text.H3>
      </spacing.P4>
      <spacing.P1></spacing.P1>
      <scroll.H>
        {props.value?.map((pn, key) =>
          <div key={key} className="relative min-w-48 min-h-48 p-3 dark:bg-ctp-base rounded-lg text-center content-center shadow-md shadow-ctp-crust">
            <button type="button" className="absolute top-1 right-1">
              {customer.favorite(pn) ? (
                <svg className="w-6 h-6 dark:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                  <path d="m12.75 20.66 6.184-7.098c2.677-2.884 2.559-6.506.754-8.705-.898-1.095-2.206-1.816-3.72-1.855-1.293-.034-2.652.43-3.963 1.442-1.315-1.012-2.678-1.476-3.973-1.442-1.515.04-2.825.76-3.724 1.855-1.806 2.201-1.915 5.823.772 8.706l6.183 7.097c.19.216.46.34.743.34a.985.985 0 0 0 .743-.34Z" />
                </svg>
              ) : (
                <svg className="w-6 h-6 dark:text-ctp-text" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                  <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12.01 6.001C6.5 1 1 8 5.782 13.001L12.011 20l6.23-7C23 8 17.5 1 12.01 6.002Z" />
                </svg>
              )}
            </button>
            <img src={pn.Image} className="object-cover size-32 rounded-full mx-auto" />
            <h5 className="dark:text-ctp-text font-medium mt-1">{pn.Name}</h5>
            <h3 className="dark:text-ctp-text text-lg font-bold">${pn.Price.Amount}</h3>
          </div>
        )}
      </scroll.H>
    </>
  )

}
