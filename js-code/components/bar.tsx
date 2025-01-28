import React from "react"

namespace bar {
  interface BarProps {
    children: React.ReactNode
  }

  export function Title(props: BarProps) {
    return (
      <div className="w-full h-12 flex justify-between">
        {props.children}
      </div>
    )
  }
}

export default bar
