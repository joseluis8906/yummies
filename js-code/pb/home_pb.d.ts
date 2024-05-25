import * as jspb from 'google-protobuf'



export class HomeIndexRequest extends jspb.Message {
  getCustomer(): string;
  setCustomer(value: string): HomeIndexRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeIndexRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HomeIndexRequest): HomeIndexRequest.AsObject;
  static serializeBinaryToWriter(message: HomeIndexRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeIndexRequest;
  static deserializeBinaryFromReader(message: HomeIndexRequest, reader: jspb.BinaryReader): HomeIndexRequest;
}

export namespace HomeIndexRequest {
  export type AsObject = {
    customer: string,
  }
}

export class HomeIndexResponse extends jspb.Message {
  getCategoriesList(): Array<string>;
  setCategoriesList(value: Array<string>): HomeIndexResponse;
  clearCategoriesList(): HomeIndexResponse;
  addCategories(value: string, index?: number): HomeIndexResponse;

  getTodaysSpecialOffer(): HomeTodaysSpecialOffer | undefined;
  setTodaysSpecialOffer(value?: HomeTodaysSpecialOffer): HomeIndexResponse;
  hasTodaysSpecialOffer(): boolean;
  clearTodaysSpecialOffer(): HomeIndexResponse;

  getPopularNowList(): Array<HomePopularNow>;
  setPopularNowList(value: Array<HomePopularNow>): HomeIndexResponse;
  clearPopularNowList(): HomeIndexResponse;
  addPopularNow(value?: HomePopularNow, index?: number): HomePopularNow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeIndexResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HomeIndexResponse): HomeIndexResponse.AsObject;
  static serializeBinaryToWriter(message: HomeIndexResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeIndexResponse;
  static deserializeBinaryFromReader(message: HomeIndexResponse, reader: jspb.BinaryReader): HomeIndexResponse;
}

export namespace HomeIndexResponse {
  export type AsObject = {
    categoriesList: Array<string>,
    todaysSpecialOffer?: HomeTodaysSpecialOffer.AsObject,
    popularNowList: Array<HomePopularNow.AsObject>,
  }
}

export class HomeTodaysSpecialOffer extends jspb.Message {
  getName(): string;
  setName(value: string): HomeTodaysSpecialOffer;

  getImg(): string;
  setImg(value: string): HomeTodaysSpecialOffer;

  getDescription(): string;
  setDescription(value: string): HomeTodaysSpecialOffer;

  getPrice(): HomeMoney | undefined;
  setPrice(value?: HomeMoney): HomeTodaysSpecialOffer;
  hasPrice(): boolean;
  clearPrice(): HomeTodaysSpecialOffer;

  getDiscount(): number;
  setDiscount(value: number): HomeTodaysSpecialOffer;

  getPriceDiscounted(): HomeMoney | undefined;
  setPriceDiscounted(value?: HomeMoney): HomeTodaysSpecialOffer;
  hasPriceDiscounted(): boolean;
  clearPriceDiscounted(): HomeTodaysSpecialOffer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeTodaysSpecialOffer.AsObject;
  static toObject(includeInstance: boolean, msg: HomeTodaysSpecialOffer): HomeTodaysSpecialOffer.AsObject;
  static serializeBinaryToWriter(message: HomeTodaysSpecialOffer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeTodaysSpecialOffer;
  static deserializeBinaryFromReader(message: HomeTodaysSpecialOffer, reader: jspb.BinaryReader): HomeTodaysSpecialOffer;
}

export namespace HomeTodaysSpecialOffer {
  export type AsObject = {
    name: string,
    img: string,
    description: string,
    price?: HomeMoney.AsObject,
    discount: number,
    priceDiscounted?: HomeMoney.AsObject,
  }
}

export class HomePopularNow extends jspb.Message {
  getName(): string;
  setName(value: string): HomePopularNow;

  getImg(): string;
  setImg(value: string): HomePopularNow;

  getPrice(): HomeMoney | undefined;
  setPrice(value?: HomeMoney): HomePopularNow;
  hasPrice(): boolean;
  clearPrice(): HomePopularNow;

  getIsFavorite(): boolean;
  setIsFavorite(value: boolean): HomePopularNow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomePopularNow.AsObject;
  static toObject(includeInstance: boolean, msg: HomePopularNow): HomePopularNow.AsObject;
  static serializeBinaryToWriter(message: HomePopularNow, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomePopularNow;
  static deserializeBinaryFromReader(message: HomePopularNow, reader: jspb.BinaryReader): HomePopularNow;
}

export namespace HomePopularNow {
  export type AsObject = {
    name: string,
    img: string,
    price?: HomeMoney.AsObject,
    isFavorite: boolean,
  }
}

export class HomeMoney extends jspb.Message {
  getAmount(): number;
  setAmount(value: number): HomeMoney;

  getDecimals(): number;
  setDecimals(value: number): HomeMoney;

  getCurrency(): string;
  setCurrency(value: string): HomeMoney;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeMoney.AsObject;
  static toObject(includeInstance: boolean, msg: HomeMoney): HomeMoney.AsObject;
  static serializeBinaryToWriter(message: HomeMoney, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeMoney;
  static deserializeBinaryFromReader(message: HomeMoney, reader: jspb.BinaryReader): HomeMoney;
}

export namespace HomeMoney {
  export type AsObject = {
    amount: number,
    decimals: number,
    currency: string,
  }
}

