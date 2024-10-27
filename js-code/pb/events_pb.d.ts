import * as jspb from 'google-protobuf'



export class V1 extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): V1.AsObject;
  static toObject(includeInstance: boolean, msg: V1): V1.AsObject;
  static serializeBinaryToWriter(message: V1, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): V1;
  static deserializeBinaryFromReader(message: V1, reader: jspb.BinaryReader): V1;
}

export namespace V1 {
  export type AsObject = {
  }

  export class Tested extends jspb.Message {
    getId(): string;
    setId(value: string): Tested;

    getOccurredOn(): number;
    setOccurredOn(value: number): Tested;

    getMsg(): string;
    setMsg(value: string): Tested;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Tested.AsObject;
    static toObject(includeInstance: boolean, msg: Tested): Tested.AsObject;
    static serializeBinaryToWriter(message: Tested, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Tested;
    static deserializeBinaryFromReader(message: Tested, reader: jspb.BinaryReader): Tested;
  }

  export namespace Tested {
    export type AsObject = {
      id: string,
      occurredOn: number,
      msg: string,
    }
  }

}

