// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file ibc/core/client/v2/counterparty.proto (package ibc.core.client.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * CounterpartyInfo defines the key that the counterparty will use to message our client
 *
 * @generated from message ibc.core.client.v2.CounterpartyInfo
 */
export class CounterpartyInfo extends Message<CounterpartyInfo> {
  /**
   * merkle prefix key is the prefix that ics provable keys are stored under
   *
   * @generated from field: repeated bytes merkle_prefix = 1;
   */
  merklePrefix: Uint8Array[] = [];

  /**
   * client identifier is the identifier used to send packet messages to our client
   *
   * @generated from field: string client_id = 2;
   */
  clientId = "";

  constructor(data?: PartialMessage<CounterpartyInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ibc.core.client.v2.CounterpartyInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "merkle_prefix", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
    { no: 2, name: "client_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CounterpartyInfo {
    return new CounterpartyInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CounterpartyInfo {
    return new CounterpartyInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CounterpartyInfo {
    return new CounterpartyInfo().fromJsonString(jsonString, options);
  }

  static equals(a: CounterpartyInfo | PlainMessage<CounterpartyInfo> | undefined, b: CounterpartyInfo | PlainMessage<CounterpartyInfo> | undefined): boolean {
    return proto3.util.equals(CounterpartyInfo, a, b);
  }
}

