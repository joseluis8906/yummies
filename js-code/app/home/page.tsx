import { HomeServiceClient } from "@/pb/HomeServiceClientPb";
import { HomeIndexRequest, HomeTodaysSpecialOffer, HomePopularNow, HomeMoney } from "@/pb/home_pb";

async function getData() {
    const client = new HomeServiceClient('http://localhost:8080');
    const headers = { "x-auth-email": "john.doe@example.com" }
    const req = new HomeIndexRequest();
    req.setCustomer('john.doe@example.com');
    const res = await client.index(req, headers);
    return { categories: res.getCategoriesList(), todaysSpecialOffer: res.getTodaysSpecialOffer(), popularNow: res.getPopularNowList() }
}

export default async function Home() {
    const data = await getData();
    return (
        <div className="w-full h-full">
            <AppBar></AppBar>
            <Content
                categories={data.categories}
                todaysSpecialOffer={data.todaysSpecialOffer}
                popularNow={data.popularNow} >
            </Content>
            <BottomBar></BottomBar>
        </div >
    )
}

function AppBar() {
    return (
        <div className="fixed top-0 left-0 w-full h-60 z-50">
            <div className="max-w-[430px] p-2 pb-4 mx-auto bg-white border-t dark:bg-ctp-base dark:border-ctp-base shadow-md shadow-ctp-crust">
                <div className="w-full h-12 flex justify-between">
                    <div className="content-center">
                        <button type="button" className="p-2 dark:hover:bg-ctp-mantle rounded-full group">
                            <svg className="w-6 h-6 dark:text-ctp-text group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                                <path stroke="currentColor" strokeLinecap="round" strokeWidth="2" d="M5 7h14M5 12h14M5 17h14" />
                            </svg>
                        </button>
                    </div>
                    <div className="content-center">
                        <button type="button" className="p-2 dark:hover:bg-ctp-mantle rounded-full group">
                            <svg className="w-6 h-6 dark:text-ctp-text group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                                <path fillRule="evenodd" d="M4 4a1 1 0 0 1 1-1h1.5a1 1 0 0 1 .979.796L7.939 6H19a1 1 0 0 1 .979 1.204l-1.25 6a1 1 0 0 1-.979.796H9.605l.208 1H17a3 3 0 1 1-2.83 2h-2.34a3 3 0 1 1-4.009-1.76L5.686 5H5a1 1 0 0 1-1-1Z" clipRule="evenodd" />
                            </svg>
                        </button>
                    </div>
                </div>
                {/*--greetings--*/}
                <div className="p-2">
                    <h5 className="dark:text-ctp-text">Hello John,</h5>
                    <h1 className="dark:text-ctp-text text-4xl">What would you like to <strong className="font-bold text-ctp-peach">eat?</strong></h1>
                </div>
                {/*--search bar--*/}
                <form className="px-2 flex justify-between">
                    <div className="relative w-full">
                        <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
                            <svg className="w-4 h-4 dark:text-ctp-overlay2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
                            </svg>
                        </div>
                        <input type="text" id="input-group-1" className="bg-gray-50 border text-sm rounded-lg block w-full ps-10 p-2.5 dark:bg-ctp-mantle dark:border-ctp-mantle dark:placeholder-ctp-overlay2 dark:text-ctp-text dark:focus:ring-ctp-peach dark:focus:border-ctp-peach outline-none" placeholder="Search ..." />
                        <div className="absolute inset-y-0 end-3 flex items-center ps-3.5">
                            <button type="button" className="dark:bg-ctp-surface1 dark:hover:bg-ctp-peach rounded-md p-1 group">
                                <svg className="w-6 h-6 dark:text-ctp-text dark:group-hover:text-ctp-base" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                                    <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 9v3a5.006 5.006 0 0 1-5 5h-4a5.006 5.006 0 0 1-5-5V9m7 9v3m-3 0h6M11 3h2a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3Z" />
                                </svg>
                            </button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    )
}

function toFloat(price: HomeMoney | undefined): number {
    if (price == undefined) {
        return 0.0
    }
    const divisor = 10 ** price.getDecimals()
    return price.getAmount() / divisor
}

interface Props {
    categories: Array<string> | undefined;
    todaysSpecialOffer: HomeTodaysSpecialOffer | undefined;
    popularNow: Array<HomePopularNow> | undefined;
}

function Content({ categories, todaysSpecialOffer, popularNow }: Props) {
    return (
        <div className="container mx-auto">
            <div className="max-w-[430px] mx-auto w-full h-full pt-60 pb-16 overflow-hidden">
                <div className="h-full">
                    {/*categories*/}
                    <div className="w-full overflow-x-clip">
                        <div className="flex flex-nowrap px-4 gap-2 mx-auto my-2 rounded-lg overflow-x-auto no-scrollbar" role="group">
                            {categories?.map((category, key) =>
                                <button
                                    key={key}
                                    type="button"
                                    className="px-5 py-1.5 text-xs font-medium dark:bg-ctp-surface0 dark:hover:bg-ctp-peach dark:text-ctp-text dark:hover:text-ctp-base rounded-lg text-nowrap"
                                >
                                    {category}
                                </button>
                            )}
                        </div>
                    </div>
                    {/*--todays special offer--*/}
                    <div className="w-full px-4 mt-3">
                        <h3 className="dark:text-ctp-text font-bold">Today's Special Offer</h3>
                        <div className="w-full mt-2 dark:bg-ctp-base rounded-lg flex flex-row overflow-hidden shadow-md shadow-ctp-crust">
                            <div className="basis-1/2">
                                <img src={todaysSpecialOffer?.getImg()} className="w-full h-full object-cover" />
                            </div>
                            <div className="basis-1/2 content-center justify-center text-center p-2">
                                <h4 className="text-md dark:text-ctp-text font-semibold text-center">{todaysSpecialOffer?.getName()}</h4>
                                <span className="dark:text-ctp-text text-sm">Now</span>
                                <h3 className="dark:text-ctp-text text-2xl font-bold">${toFloat(todaysSpecialOffer?.getPrice()).toFixed(todaysSpecialOffer?.getPrice()?.getDecimals())}</h3>
                                <span className="dark:text-ctp-red text-sm">({todaysSpecialOffer?.getDiscount()}% off)</span>
                                <button type="button" className="font-medium rounded-lg text-sm px-5 py-2.5 text-center border dark:border-ctp-peach dark:text-ctp-peach dark:hover:text-ctp-base dark:hover:bg-ctp-peach block mx-auto mt-3">Add to Cart</button>
                            </div>
                        </div>
                    </div>
                    {/*--popular now--*/}
                    <div className="w-full p-4 mt-5 overflow-x-clip">
                        <h3 className="dark:text-ctp-text font-bold">Popular Now</h3>
                        <div className="flex flex-nowrap gap-3 mx-auto rounded-lg overflow-x-auto no-scrollbar py-2 pr-3">
                            {popularNow?.map((pn, key) =>
                                <div key={key} className="relative min-w-48 min-h-48 p-3 dark:bg-ctp-base rounded-lg text-center content-center shadow-md shadow-ctp-crust">
                                    <button type="button" className="absolute top-1 right-1">
                                        {pn.getIsFavorite() ? (
                                            <svg className="w-6 h-6 dark:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                                                <path d="m12.75 20.66 6.184-7.098c2.677-2.884 2.559-6.506.754-8.705-.898-1.095-2.206-1.816-3.72-1.855-1.293-.034-2.652.43-3.963 1.442-1.315-1.012-2.678-1.476-3.973-1.442-1.515.04-2.825.76-3.724 1.855-1.806 2.201-1.915 5.823.772 8.706l6.183 7.097c.19.216.46.34.743.34a.985.985 0 0 0 .743-.34Z" />
                                            </svg>
                                        ) : (
                                            <svg className="w-6 h-6 dark:text-ctp-text" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                                                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12.01 6.001C6.5 1 1 8 5.782 13.001L12.011 20l6.23-7C23 8 17.5 1 12.01 6.002Z" />
                                            </svg>
                                        )}
                                    </button>
                                    <img src={pn.getImg()} className="object-cover size-32 rounded-full mx-auto" />
                                    <h5 className="dark:text-ctp-text font-medium mt-1">{pn.getName()}</h5>
                                    <h3 className="dark:text-ctp-text text-lg font-bold">${toFloat(pn.getPrice()).toFixed(pn.getPrice()?.getDecimals())}</h3>
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

function BottomBar() {
    return (
        <div className="fixed bottom-0 left-0 w-full h-16 z-[1000]">
            <div className="max-w-[430px] h-full mx-auto border-t dark:bg-ctp-base dark:border-ctp-base dark:shadow-ctp-crust">
                <div className="grid h-full max-w-lg grid-cols-5 mx-auto">
                    <button type="button" className="inline-flex flex-col items-center justify-center font-medium px-2 dark:hover:bg-ctp-mantle group">
                        <svg className="w-6 h-6 dark:text-ctp-text dark:group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path fillRule="evenodd" d="M4 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1v2a1 1 0 0 0 1.707.707L9.414 13H15a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4Z" clipRule="evenodd" />
                            <path fillRule="evenodd" d="M8.023 17.215c.033-.03.066-.062.098-.094L10.243 15H15a3 3 0 0 0 3-3V8h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-1v2a1 1 0 0 1-1.707.707L14.586 18H9a1 1 0 0 1-.977-.785Z" clipRule="evenodd" />
                        </svg>
                        <span className="text-sm dark:text-ctp-text dark:group-hover:text-ctp-peach">Chat</span>
                    </button>
                    <button type="button" className="inline-flex flex-col items-center justify-center font-medium px-2 dark:hover:bg-ctp-mantle group">
                        <svg className="w-6 h-6 text-gray-500 dark:text-ctp-text dark:group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path fillRule="evenodd" d="M12 4a4 4 0 1 0 0 8 4 4 0 0 0 0-8Zm-2 9a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1a4 4 0 0 0-4-4h-4Z" clipRule="evenodd" />
                        </svg>
                        <span className="text-sm dark:text-ctp-text dark:group-hover:text-ctp-peach">Profile</span>
                    </button>
                    <button type="button" className="inline-flex flex-col items-center justify-center font-medium px-2 dark:hover:bg-ctp-mantle group">
                        <svg className="w-6 h-6 dark:text-ctp-text dark:group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path fillRule="evenodd" d="M11.293 3.293a1 1 0 0 1 1.414 0l6 6 2 2a1 1 0 0 1-1.414 1.414L19 12.414V19a2 2 0 0 1-2 2h-3a1 1 0 0 1-1-1v-3h-2v3a1 1 0 0 1-1 1H7a2 2 0 0 1-2-2v-6.586l-.293.293a1 1 0 0 1-1.414-1.414l2-2 6-6Z" clipRule="evenodd" />
                        </svg>
                        <span className="text-sm dark:text-ctp-text dark:group-hover:text-ctp-peach">Home</span>
                    </button>
                    <button type="button" className="inline-flex flex-col items-center justify-center font-medium px-2 dark:hover:bg-ctp-mantle group">
                        <svg className="w-6 h-6 dark:text-ctp-text dark:group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 45.001 45.001">
                            <path d="M43.604,41.146c0.703,0.64,0.945,2.207,0.037,3.107c-0.902,0.909-2.472,0.668-3.109-0.034 c-5.893-6.328-11.783-12.659-17.676-18.989l-0.67-0.715c-6.164,6.618-12.326,13.239-18.488,19.86 c-0.635,0.701-2.202,0.943-3.104,0.034c-0.91-0.901-0.67-2.472,0.032-3.109c6.642-6.178,13.279-12.358,19.92-18.539 c0,0-3.214-3.467-4.6-4.955c-3.243,2.325-8.049,1.441-10.938-2.126C1.901,11.917-0.779,5.187,1.815,2.431 c2.759-2.596,9.485,0.09,13.248,3.194c3.566,2.892,4.455,7.691,2.125,10.936c1.67,1.556,3.34,3.108,5.012,4.663 c-0.004,0.003-0.006,0.005-0.01,0.008c0,0,0.104,0.082,0.01-0.008c1.562-1.455,3.125-2.909,4.688-4.364 c-1.32-2.254-0.443-5.415,1.701-7.219c2.459-2.061,11.57-9.883,11.815-9.635c0.244,0.243-6.146,6.281-10.415,10.563l0.712,0.711 c4.557-3.994,10.99-9.984,11.238-9.74c0.244,0.245-5.946,6.48-10.076,10.905l0.707,0.705C36.988,9.017,43.225,2.826,43.467,3.069 c0.245,0.246-5.746,6.68-9.74,11.237l0.707,0.706c4.28-4.269,10.315-10.659,10.562-10.417c0.247,0.246-9.618,11.805-9.632,11.822 c-0.002-0.002,2.094-2.462-0.008-0.006c-2.006,2.347-5.002,3.057-7.217,1.705c-1.438,1.543-2.873,3.086-4.309,4.629 C23.704,22.626,37.014,35.013,43.604,41.146z"></path>
                        </svg>
                        <span className="text-sm dark:text-ctp-text dark:group-hover:text-ctp-peach">Menu</span>
                    </button>
                    <button type="button" className="inline-flex flex-col items-center justify-center font-medium px-2 dark:hover:bg-ctp-mantle group">
                        <svg className="w-6 h-6 mb-1 dark:text-ctp-text dark:group-hover:text-ctp-peach" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 22 20">
                            <path d="M20.924 7.625a1.523 1.523 0 0 0-1.238-1.044l-5.051-.734-2.259-4.577a1.534 1.534 0 0 0-2.752 0L7.365 5.847l-5.051.734A1.535 1.535 0 0 0 1.463 9.2l3.656 3.563-.863 5.031a1.532 1.532 0 0 0 2.226 1.616L11 17.033l4.518 2.375a1.534 1.534 0 0 0 2.226-1.617l-.863-5.03L20.537 9.2a1.523 1.523 0 0 0 .387-1.575Z" />
                        </svg>
                        <span className="text-sm dark:text-ctp-text dark:group-hover:text-ctp-peach">Favorites</span>
                    </button>
                </div>
            </div>
        </div>
    )
}
