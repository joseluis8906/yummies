import React from "react"

namespace layout {
  interface ChildrenProps {
    children: React.ReactNode
  }

  export function Container(props: ChildrenProps) {
    return (
      <div className="w-full h-full">
        {props.children}
      </div>
    )
  }

  export function Header(props: ChildrenProps) {
    return (
      <div className="fixed top-0 left-0 w-full h-60 z-50">
        <div className="max-w-[430px] mx-auto bg-white border-t dark:bg-ctp-base dark:border-ctp-base shadow-md shadow-ctp-crust">
          {props.children}
        </div>
      </div>
    )
  }

  export function Content(props: ChildrenProps) {
    return (
      <div className="container mx-auto">
        <div className="max-w-[430px] mx-auto w-full h-full pt-60 pb-16 overflow-hidden">
          <div className="h-full">
            {props.children}
          </div>
        </div>
      </div>
    )
  }
}

export default layout
