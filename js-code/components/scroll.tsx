import React from "react"

namespace scroll {
  export function H(props?: { children?: React.ReactNode }) {
    return (
      <div className="w-full overflow-x-clip px-4">
        <div className="flex flex-nowrap gap-2 mx-auto overflow-x-auto no-scrollbar" role="group">
          {props?.children}
        </div>
      </div>
    )
  }

}

export default scroll
