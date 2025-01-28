import React from "react"

namespace spacing {
  interface PProps {
    children?: React.ReactNode
    h?: boolean
    v?: boolean
  }

  export function P1(props: PProps) {
    let className = "p-1"
    if (props.h === false) className = "py-1"
    else if (props.v === false) className = "px-1"

    return (
      <div className={className} > {props.children}</div >
    )
  }

  P1.defaulProps = { h: true, v: true }

  export function P2(props: PProps) {
    let className = "p-2"
    if (props.h === false) className = "py-2"
    else if (props.v === false) className = "px-2"

    return (
      <div className={className} > {props.children}</div >
    )
  }

  P2.defaulProps = { h: true, v: true }

  export function P4(props: PProps) {
    let className = "p-4"
    if (props.h === false) className = "py-4"
    else if (props.v === false) className = "px-4"

    return (
      <div className={className} > {props.children}</div >
    )
  }

  P4.defaulProps = { h: true, v: true }
}

export default spacing
