import { useState, useEffect } from "react"
import layout from "@/components/layout"
import bar from "@/components/bar"
import button from "@/components/button"
import icon from "@/components/icon"
import text from "@/components/text"
import input from "@/components/input"
import spacing from "@/components/spacing"


export default function Header() {
  const [filter, setFilter] = useState<string>();
  useEffect(() => console.info(filter), [filter])
  const search = (e: React.ChangeEvent<HTMLInputElement>): void => setFilter(e.target.value)

  return (
    <layout.Header>
      <spacing.P4>
        <bar.Title>
          <button.Icon><icon.Bars /></button.Icon>
          <button.Icon><icon.Cart /></button.Icon>
        </bar.Title>
        {/*--greetings--*/}
        <spacing.P2 v={false}>
          <text.H5>Hello John,</text.H5>
          <text.H1>What would you like to <text.Strong>eat?</text.Strong></text.H1>
        </spacing.P2>
        <spacing.P2></spacing.P2>
        {/*form*/}
        <input.Form>
          <input.Search onChange={search} />
        </input.Form>
      </spacing.P4>
    </layout.Header>
  )
}
