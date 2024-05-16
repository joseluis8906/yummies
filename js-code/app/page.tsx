import Image from "next/image";

export default function Index() {
    return (
        <div className="dark:bg-gray-900">
            <div className="container mx-auto h-screen">
                <div className="w-[430px] mx-auto bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 h-full splash-screen place-content-center relative">
                    <div className="px-5 text-center absolute bottom-[25%]">
                        <h1 className="mb-2 text-6xl font-bold tracking-tight text-gray-900 dark:text-white">Yummies</h1>
                        <div className="px-20">
                            <h6 className="mb-2 text-xl font-medium tracking-tight text-gray-900 dark:text-white" >Tasty meals delivered to your doorstep</h6>
                        </div>
                        <button type="button" className="text-white bg-orange-500 hover:bg-orange-600 focus:outline-none focus:ring-4 focus:ring-yellow-300 font-bold rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:focus:ring-yellow-900">
                            Get Started
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
