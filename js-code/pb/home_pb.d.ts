import * as jspb from 'google-protobuf'



export class HomeRequest extends jspb.Message {
  getCustomer(): string;
  setCustomer(value: string): HomeRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HomeRequest): HomeRequest.AsObject;
  static serializeBinaryToWriter(message: HomeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeRequest;
  static deserializeBinaryFromReader(message: HomeRequest, reader: jspb.BinaryReader): HomeRequest;
}

export namespace HomeRequest {
  export type AsObject = {
    customer: string,
  }
}

export class HomeResponse extends jspb.Message {
  getCategoriesList(): Array<string>;
  setCategoriesList(value: Array<string>): HomeResponse;
  clearCategoriesList(): HomeResponse;
  addCategories(value: string, index?: number): HomeResponse;

  getTodaysSpecialOffer(): HomeTodaysSpecialOffer | undefined;
  setTodaysSpecialOffer(value?: HomeTodaysSpecialOffer): HomeResponse;
  hasTodaysSpecialOffer(): boolean;
  clearTodaysSpecialOffer(): HomeResponse;

  getPopularNowList(): Array<HomePopularNow>;
  setPopularNowList(value: Array<HomePopularNow>): HomeResponse;
  clearPopularNowList(): HomeResponse;
  addPopularNow(value?: HomePopularNow, index?: number): HomePopularNow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HomeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HomeResponse): HomeResponse.AsObject;
  static serializeBinaryToWriter(message: HomeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HomeResponse;
  static deserializeBinaryFromReader(message: HomeResponse, reader: jspb.BinaryReader): HomeResponse;
}

export namespace HomeResponse {
  export type AsObject = {
    categoriesList: Array<string>,
    todaysSpecialOffer?: HomeTodaysSpecialOffer.AsObject,
    popularNowList: Array<HomePopularNow.AsObject>,
  }
}

export class HomeTodaysSpecialOffer extends jspb.Message {
  getName(): string;
  setName(value: string): HomeTodaysSpecialOffer;

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
    description: string,
    price?: HomeMoney.AsObject,
    discount: number,
    priceDiscounted?: HomeMoney.AsObject,
  }
}

export class HomePopularNow extends jspb.Message {
  getName(): string;
  setName(value: string): HomePopularNow;

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

