import React from "react"

namespace button {
  interface IconProps {
    children: React.ReactNode
  }

  export function Icon(props: IconProps) {
    return (
      <div className="content-center">
        <button type="button" className="p-2 dark:hover:bg-ctp-mantle rounded-full group">
          {props.children}
        </button>
      </div>
    )
  }
}

export default button
