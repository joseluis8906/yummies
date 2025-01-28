import React from "react";

namespace input {
  interface ChildrenProps {
    children: React.ReactNode
  }

  export function Form(props: ChildrenProps) {
    return (
      <form className="flex justify-between">
        {props.children}
      </form>
    )
  }

  interface InputProps {
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  }

  export function Search(props: InputProps) {
    return (
      <div className="relative w-full">
        <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
          <svg className="w-4 h-4 dark:text-ctp-overlay2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
            <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
          </svg>
        </div>
        <input
          type="text"
          id="input-group-1"
          className="bg-gray-50 border text-sm rounded-lg block w-full ps-10 p-2.5 dark:bg-ctp-mantle dark:border-ctp-mantle dark:placeholder-ctp-overlay2 dark:text-ctp-text dark:focus:ring-ctp-peach dark:focus:border-ctp-peach outline-none"
          placeholder="Search ..."
          onChange={props.onChange}
        />
        <div className="absolute inset-y-0 end-3 flex items-center ps-3.5">
          <button type="button" className="dark:bg-ctp-surface1 dark:hover:bg-ctp-peach rounded-md p-1 group">
            <svg className="w-6 h-6 dark:text-ctp-text dark:group-hover:text-ctp-base" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
              <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 9v3a5.006 5.006 0 0 1-5 5h-4a5.006 5.006 0 0 1-5-5V9m7 9v3m-3 0h6M11 3h2a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3Z" />
            </svg>
          </button>
        </div>
      </div>
    )
  }
}

export default input
