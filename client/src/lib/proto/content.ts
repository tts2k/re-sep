// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.174.0
//   protoc               v4.25.3
// source: content.proto

/* eslint-disable */
import {
  type CallOptions,
  ChannelCredentials,
  Client,
  type ClientOptions,
  type ClientUnaryCall,
  type handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  type ServiceError,
  type UntypedServiceImplementation,
} from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "content";

export interface EntryName {
  entryName: string;
}

export interface TOCItem {
  id: string;
  label: string;
  subItems: TOCItem[];
}

export interface Article {
  entryName: string;
  title: string;
  issued: Date | undefined;
  modified: Date | undefined;
  htmlText: string;
  authors: string[];
  toc: TOCItem[];
}

function createBaseEntryName(): EntryName {
  return { entryName: "" };
}

export const EntryName = {
  encode(message: EntryName, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entryName !== "") {
      writer.uint32(10).string(message.entryName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EntryName {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEntryName();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.entryName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EntryName {
    return { entryName: isSet(object.entryName) ? globalThis.String(object.entryName) : "" };
  },

  toJSON(message: EntryName): unknown {
    const obj: any = {};
    if (message.entryName !== "") {
      obj.entryName = message.entryName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EntryName>, I>>(base?: I): EntryName {
    return EntryName.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<EntryName>, I>>(object: I): EntryName {
    const message = createBaseEntryName();
    message.entryName = object.entryName ?? "";
    return message;
  },
};

function createBaseTOCItem(): TOCItem {
  return { id: "", label: "", subItems: [] };
}

export const TOCItem = {
  encode(message: TOCItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    for (const v of message.subItems) {
      TOCItem.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TOCItem {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTOCItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.label = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.subItems.push(TOCItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TOCItem {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      label: isSet(object.label) ? globalThis.String(object.label) : "",
      subItems: globalThis.Array.isArray(object?.subItems) ? object.subItems.map((e: any) => TOCItem.fromJSON(e)) : [],
    };
  },

  toJSON(message: TOCItem): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.label !== "") {
      obj.label = message.label;
    }
    if (message.subItems?.length) {
      obj.subItems = message.subItems.map((e) => TOCItem.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TOCItem>, I>>(base?: I): TOCItem {
    return TOCItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TOCItem>, I>>(object: I): TOCItem {
    const message = createBaseTOCItem();
    message.id = object.id ?? "";
    message.label = object.label ?? "";
    message.subItems = object.subItems?.map((e) => TOCItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseArticle(): Article {
  return { entryName: "", title: "", issued: undefined, modified: undefined, htmlText: "", authors: [], toc: [] };
}

export const Article = {
  encode(message: Article, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.entryName !== "") {
      writer.uint32(10).string(message.entryName);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.issued !== undefined) {
      Timestamp.encode(toTimestamp(message.issued), writer.uint32(26).fork()).ldelim();
    }
    if (message.modified !== undefined) {
      Timestamp.encode(toTimestamp(message.modified), writer.uint32(34).fork()).ldelim();
    }
    if (message.htmlText !== "") {
      writer.uint32(42).string(message.htmlText);
    }
    for (const v of message.authors) {
      writer.uint32(50).string(v!);
    }
    for (const v of message.toc) {
      TOCItem.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Article {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseArticle();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.entryName = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.issued = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.modified = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.htmlText = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.authors.push(reader.string());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.toc.push(TOCItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Article {
    return {
      entryName: isSet(object.entryName) ? globalThis.String(object.entryName) : "",
      title: isSet(object.title) ? globalThis.String(object.title) : "",
      issued: isSet(object.issued) ? fromJsonTimestamp(object.issued) : undefined,
      modified: isSet(object.modified) ? fromJsonTimestamp(object.modified) : undefined,
      htmlText: isSet(object.htmlText) ? globalThis.String(object.htmlText) : "",
      authors: globalThis.Array.isArray(object?.authors) ? object.authors.map((e: any) => globalThis.String(e)) : [],
      toc: globalThis.Array.isArray(object?.toc) ? object.toc.map((e: any) => TOCItem.fromJSON(e)) : [],
    };
  },

  toJSON(message: Article): unknown {
    const obj: any = {};
    if (message.entryName !== "") {
      obj.entryName = message.entryName;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.issued !== undefined) {
      obj.issued = message.issued.toISOString();
    }
    if (message.modified !== undefined) {
      obj.modified = message.modified.toISOString();
    }
    if (message.htmlText !== "") {
      obj.htmlText = message.htmlText;
    }
    if (message.authors?.length) {
      obj.authors = message.authors;
    }
    if (message.toc?.length) {
      obj.toc = message.toc.map((e) => TOCItem.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Article>, I>>(base?: I): Article {
    return Article.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Article>, I>>(object: I): Article {
    const message = createBaseArticle();
    message.entryName = object.entryName ?? "";
    message.title = object.title ?? "";
    message.issued = object.issued ?? undefined;
    message.modified = object.modified ?? undefined;
    message.htmlText = object.htmlText ?? "";
    message.authors = object.authors?.map((e) => e) || [];
    message.toc = object.toc?.map((e) => TOCItem.fromPartial(e)) || [];
    return message;
  },
};

export type ContentService = typeof ContentService;
export const ContentService = {
  getArticle: {
    path: "/content.Content/GetArticle",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: EntryName) => Buffer.from(EntryName.encode(value).finish()),
    requestDeserialize: (value: Buffer) => EntryName.decode(value),
    responseSerialize: (value: Article) => Buffer.from(Article.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Article.decode(value),
  },
} as const;

export interface ContentServer extends UntypedServiceImplementation {
  getArticle: handleUnaryCall<EntryName, Article>;
}

export interface ContentClient extends Client {
  getArticle(request: EntryName, callback: (error: ServiceError | null, response: Article) => void): ClientUnaryCall;
  getArticle(
    request: EntryName,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Article) => void,
  ): ClientUnaryCall;
  getArticle(
    request: EntryName,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Article) => void,
  ): ClientUnaryCall;
}

export const ContentClient = makeGenericClientConstructor(ContentService, "content.Content") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): ContentClient;
  service: typeof ContentService;
  serviceName: string;
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof globalThis.Date) {
    return o;
  } else if (typeof o === "string") {
    return new globalThis.Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
