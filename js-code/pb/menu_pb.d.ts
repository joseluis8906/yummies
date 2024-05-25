import * as jspb from 'google-protobuf'



export class MenuIndexRequest extends jspb.Message {
  getCustomer(): string;
  setCustomer(value: string): MenuIndexRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MenuIndexRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MenuIndexRequest): MenuIndexRequest.AsObject;
  static serializeBinaryToWriter(message: MenuIndexRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MenuIndexRequest;
  static deserializeBinaryFromReader(message: MenuIndexRequest, reader: jspb.BinaryReader): MenuIndexRequest;
}

export namespace MenuIndexRequest {
  export type AsObject = {
    customer: string,
  }
}

export class MenuIndexResponse extends jspb.Message {
  getCategoriesList(): Array<string>;
  setCategoriesList(value: Array<string>): MenuIndexResponse;
  clearCategoriesList(): MenuIndexResponse;
  addCategories(value: string, index?: number): MenuIndexResponse;

  getProductsList(): Array<MenuProduct>;
  setProductsList(value: Array<MenuProduct>): MenuIndexResponse;
  clearProductsList(): MenuIndexResponse;
  addProducts(value?: MenuProduct, index?: number): MenuProduct;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MenuIndexResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MenuIndexResponse): MenuIndexResponse.AsObject;
  static serializeBinaryToWriter(message: MenuIndexResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MenuIndexResponse;
  static deserializeBinaryFromReader(message: MenuIndexResponse, reader: jspb.BinaryReader): MenuIndexResponse;
}

export namespace MenuIndexResponse {
  export type AsObject = {
    categoriesList: Array<string>,
    productsList: Array<MenuProduct.AsObject>,
  }
}

export class MenuProduct extends jspb.Message {
  getName(): string;
  setName(value: string): MenuProduct;

  getImg(): string;
  setImg(value: string): MenuProduct;

  getPrice(): MenuMoney | undefined;
  setPrice(value?: MenuMoney): MenuProduct;
  hasPrice(): boolean;
  clearPrice(): MenuProduct;

  getIsFavorite(): boolean;
  setIsFavorite(value: boolean): MenuProduct;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MenuProduct.AsObject;
  static toObject(includeInstance: boolean, msg: MenuProduct): MenuProduct.AsObject;
  static serializeBinaryToWriter(message: MenuProduct, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MenuProduct;
  static deserializeBinaryFromReader(message: MenuProduct, reader: jspb.BinaryReader): MenuProduct;
}

export namespace MenuProduct {
  export type AsObject = {
    name: string,
    img: string,
    price?: MenuMoney.AsObject,
    isFavorite: boolean,
  }
}

export class MenuMoney extends jspb.Message {
  getAmount(): number;
  setAmount(value: number): MenuMoney;

  getDecimals(): number;
  setDecimals(value: number): MenuMoney;

  getCurrency(): string;
  setCurrency(value: string): MenuMoney;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MenuMoney.AsObject;
  static toObject(includeInstance: boolean, msg: MenuMoney): MenuMoney.AsObject;
  static serializeBinaryToWriter(message: MenuMoney, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MenuMoney;
  static deserializeBinaryFromReader(message: MenuMoney, reader: jspb.BinaryReader): MenuMoney;
}

export namespace MenuMoney {
  export type AsObject = {
    amount: number,
    decimals: number,
    currency: string,
  }
}

