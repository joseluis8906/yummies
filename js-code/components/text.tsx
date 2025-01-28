import React from "react"

namespace text {
  interface ChildrenProps {
    children: React.ReactNode
  }

  export function H1(props: ChildrenProps) {
    return (
      <h1 className="dark:text-ctp-text text-4xl">{props.children}</h1>
    )
  }

  export function H3(props: ChildrenProps) {
    return (
      < h3 className="dark:text-ctp-text font-bold">{props.children}</h3 >
    )
  }

  export function H5(props: ChildrenProps) {
    return (
      <h5 className="dark:text-ctp-text">{props.children}</h5>
    )
  }

  export function Strong(props: ChildrenProps) {
    return (
      <strong className="font-bold text-ctp-peach">{props.children}</strong>
    )
  }
}

export default text
