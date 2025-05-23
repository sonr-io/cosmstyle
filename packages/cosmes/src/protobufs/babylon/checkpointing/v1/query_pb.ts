// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file babylon/checkpointing/v1/query.proto (package babylon.checkpointing.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { CheckpointStatus } from "./checkpoint_pb.js";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination_pb.js";
import { ValidatorWithBlsKey } from "./bls_key_pb.js";

/**
 * QueryRawCheckpointListRequest is the request type for the
 * Query/RawCheckpoints RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointListRequest
 */
export class QueryRawCheckpointListRequest extends Message<QueryRawCheckpointListRequest> {
  /**
   * status defines the status of the raw checkpoints of the query
   *
   * @generated from field: babylon.checkpointing.v1.CheckpointStatus status = 1;
   */
  status = CheckpointStatus.CKPT_STATUS_ACCUMULATING;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryRawCheckpointListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(CheckpointStatus) },
    { no: 2, name: "pagination", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointListRequest {
    return new QueryRawCheckpointListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointListRequest {
    return new QueryRawCheckpointListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointListRequest {
    return new QueryRawCheckpointListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointListRequest | PlainMessage<QueryRawCheckpointListRequest> | undefined, b: QueryRawCheckpointListRequest | PlainMessage<QueryRawCheckpointListRequest> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointListRequest, a, b);
  }
}

/**
 * QueryRawCheckpointListResponse is the response type for the
 * Query/RawCheckpoints RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointListResponse
 */
export class QueryRawCheckpointListResponse extends Message<QueryRawCheckpointListResponse> {
  /**
   * the order is going from the newest to oldest based on the epoch number
   *
   * @generated from field: repeated babylon.checkpointing.v1.RawCheckpointWithMetaResponse raw_checkpoints = 1;
   */
  rawCheckpoints: RawCheckpointWithMetaResponse[] = [];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryRawCheckpointListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "raw_checkpoints", kind: "message", T: RawCheckpointWithMetaResponse, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointListResponse {
    return new QueryRawCheckpointListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointListResponse {
    return new QueryRawCheckpointListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointListResponse {
    return new QueryRawCheckpointListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointListResponse | PlainMessage<QueryRawCheckpointListResponse> | undefined, b: QueryRawCheckpointListResponse | PlainMessage<QueryRawCheckpointListResponse> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointListResponse, a, b);
  }
}

/**
 * QueryRawCheckpointRequest is the request type for the Query/RawCheckpoint
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointRequest
 */
export class QueryRawCheckpointRequest extends Message<QueryRawCheckpointRequest> {
  /**
   * epoch_num defines the epoch for the queried checkpoint
   *
   * @generated from field: uint64 epoch_num = 1;
   */
  epochNum = protoInt64.zero;

  constructor(data?: PartialMessage<QueryRawCheckpointRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "epoch_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointRequest {
    return new QueryRawCheckpointRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointRequest {
    return new QueryRawCheckpointRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointRequest {
    return new QueryRawCheckpointRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointRequest | PlainMessage<QueryRawCheckpointRequest> | undefined, b: QueryRawCheckpointRequest | PlainMessage<QueryRawCheckpointRequest> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointRequest, a, b);
  }
}

/**
 * QueryRawCheckpointResponse is the response type for the Query/RawCheckpoint
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointResponse
 */
export class QueryRawCheckpointResponse extends Message<QueryRawCheckpointResponse> {
  /**
   * @generated from field: babylon.checkpointing.v1.RawCheckpointWithMetaResponse raw_checkpoint = 1;
   */
  rawCheckpoint?: RawCheckpointWithMetaResponse;

  constructor(data?: PartialMessage<QueryRawCheckpointResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "raw_checkpoint", kind: "message", T: RawCheckpointWithMetaResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointResponse {
    return new QueryRawCheckpointResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointResponse {
    return new QueryRawCheckpointResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointResponse {
    return new QueryRawCheckpointResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointResponse | PlainMessage<QueryRawCheckpointResponse> | undefined, b: QueryRawCheckpointResponse | PlainMessage<QueryRawCheckpointResponse> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointResponse, a, b);
  }
}

/**
 * QueryRawCheckpointsRequest is the request type for the Query/RawCheckpoints
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointsRequest
 */
export class QueryRawCheckpointsRequest extends Message<QueryRawCheckpointsRequest> {
  /**
   * pagination defines whether to have the pagination in the request
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryRawCheckpointsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointsRequest {
    return new QueryRawCheckpointsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointsRequest {
    return new QueryRawCheckpointsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointsRequest {
    return new QueryRawCheckpointsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointsRequest | PlainMessage<QueryRawCheckpointsRequest> | undefined, b: QueryRawCheckpointsRequest | PlainMessage<QueryRawCheckpointsRequest> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointsRequest, a, b);
  }
}

/**
 * QueryRawCheckpointsResponse is the response type for the Query/RawCheckpoints
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRawCheckpointsResponse
 */
export class QueryRawCheckpointsResponse extends Message<QueryRawCheckpointsResponse> {
  /**
   * the order is going from the newest to oldest based on the epoch number
   *
   * @generated from field: repeated babylon.checkpointing.v1.RawCheckpointWithMetaResponse raw_checkpoints = 1;
   */
  rawCheckpoints: RawCheckpointWithMetaResponse[] = [];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryRawCheckpointsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRawCheckpointsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "raw_checkpoints", kind: "message", T: RawCheckpointWithMetaResponse, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRawCheckpointsResponse {
    return new QueryRawCheckpointsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRawCheckpointsResponse {
    return new QueryRawCheckpointsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRawCheckpointsResponse {
    return new QueryRawCheckpointsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRawCheckpointsResponse | PlainMessage<QueryRawCheckpointsResponse> | undefined, b: QueryRawCheckpointsResponse | PlainMessage<QueryRawCheckpointsResponse> | undefined): boolean {
    return proto3.util.equals(QueryRawCheckpointsResponse, a, b);
  }
}

/**
 * QueryBlsPublicKeyListRequest is the request type for the Query/BlsPublicKeys
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryBlsPublicKeyListRequest
 */
export class QueryBlsPublicKeyListRequest extends Message<QueryBlsPublicKeyListRequest> {
  /**
   * epoch_num defines the epoch for the queried bls public keys
   *
   * @generated from field: uint64 epoch_num = 1;
   */
  epochNum = protoInt64.zero;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;

  constructor(data?: PartialMessage<QueryBlsPublicKeyListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryBlsPublicKeyListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "epoch_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "pagination", kind: "message", T: PageRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBlsPublicKeyListRequest {
    return new QueryBlsPublicKeyListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBlsPublicKeyListRequest {
    return new QueryBlsPublicKeyListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBlsPublicKeyListRequest {
    return new QueryBlsPublicKeyListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBlsPublicKeyListRequest | PlainMessage<QueryBlsPublicKeyListRequest> | undefined, b: QueryBlsPublicKeyListRequest | PlainMessage<QueryBlsPublicKeyListRequest> | undefined): boolean {
    return proto3.util.equals(QueryBlsPublicKeyListRequest, a, b);
  }
}

/**
 * QueryBlsPublicKeyListResponse is the response type for the
 * Query/BlsPublicKeys RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryBlsPublicKeyListResponse
 */
export class QueryBlsPublicKeyListResponse extends Message<QueryBlsPublicKeyListResponse> {
  /**
   * @generated from field: repeated babylon.checkpointing.v1.ValidatorWithBlsKey validator_with_bls_keys = 1;
   */
  validatorWithBlsKeys: ValidatorWithBlsKey[] = [];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;

  constructor(data?: PartialMessage<QueryBlsPublicKeyListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryBlsPublicKeyListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "validator_with_bls_keys", kind: "message", T: ValidatorWithBlsKey, repeated: true },
    { no: 2, name: "pagination", kind: "message", T: PageResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBlsPublicKeyListResponse {
    return new QueryBlsPublicKeyListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBlsPublicKeyListResponse {
    return new QueryBlsPublicKeyListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBlsPublicKeyListResponse {
    return new QueryBlsPublicKeyListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBlsPublicKeyListResponse | PlainMessage<QueryBlsPublicKeyListResponse> | undefined, b: QueryBlsPublicKeyListResponse | PlainMessage<QueryBlsPublicKeyListResponse> | undefined): boolean {
    return proto3.util.equals(QueryBlsPublicKeyListResponse, a, b);
  }
}

/**
 * QueryEpochStatusRequest is the request type for the Query/EpochStatus
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryEpochStatusRequest
 */
export class QueryEpochStatusRequest extends Message<QueryEpochStatusRequest> {
  /**
   * @generated from field: uint64 epoch_num = 1;
   */
  epochNum = protoInt64.zero;

  constructor(data?: PartialMessage<QueryEpochStatusRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryEpochStatusRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "epoch_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryEpochStatusRequest {
    return new QueryEpochStatusRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryEpochStatusRequest {
    return new QueryEpochStatusRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryEpochStatusRequest {
    return new QueryEpochStatusRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryEpochStatusRequest | PlainMessage<QueryEpochStatusRequest> | undefined, b: QueryEpochStatusRequest | PlainMessage<QueryEpochStatusRequest> | undefined): boolean {
    return proto3.util.equals(QueryEpochStatusRequest, a, b);
  }
}

/**
 * QueryEpochStatusResponse is the response type for the Query/EpochStatus
 * RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryEpochStatusResponse
 */
export class QueryEpochStatusResponse extends Message<QueryEpochStatusResponse> {
  /**
   * @generated from field: babylon.checkpointing.v1.CheckpointStatus status = 1;
   */
  status = CheckpointStatus.CKPT_STATUS_ACCUMULATING;

  constructor(data?: PartialMessage<QueryEpochStatusResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryEpochStatusResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(CheckpointStatus) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryEpochStatusResponse {
    return new QueryEpochStatusResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryEpochStatusResponse {
    return new QueryEpochStatusResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryEpochStatusResponse {
    return new QueryEpochStatusResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryEpochStatusResponse | PlainMessage<QueryEpochStatusResponse> | undefined, b: QueryEpochStatusResponse | PlainMessage<QueryEpochStatusResponse> | undefined): boolean {
    return proto3.util.equals(QueryEpochStatusResponse, a, b);
  }
}

/**
 * QueryRecentEpochStatusCountRequest is the request type for the
 * Query/EpochStatusCount RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRecentEpochStatusCountRequest
 */
export class QueryRecentEpochStatusCountRequest extends Message<QueryRecentEpochStatusCountRequest> {
  /**
   * epoch_count is the number of the most recent epochs to include in the
   * aggregation
   *
   * @generated from field: uint64 epoch_count = 1;
   */
  epochCount = protoInt64.zero;

  constructor(data?: PartialMessage<QueryRecentEpochStatusCountRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRecentEpochStatusCountRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "epoch_count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRecentEpochStatusCountRequest {
    return new QueryRecentEpochStatusCountRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRecentEpochStatusCountRequest {
    return new QueryRecentEpochStatusCountRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRecentEpochStatusCountRequest {
    return new QueryRecentEpochStatusCountRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRecentEpochStatusCountRequest | PlainMessage<QueryRecentEpochStatusCountRequest> | undefined, b: QueryRecentEpochStatusCountRequest | PlainMessage<QueryRecentEpochStatusCountRequest> | undefined): boolean {
    return proto3.util.equals(QueryRecentEpochStatusCountRequest, a, b);
  }
}

/**
 * QueryRecentEpochStatusCountResponse is the response type for the
 * Query/EpochStatusCount RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryRecentEpochStatusCountResponse
 */
export class QueryRecentEpochStatusCountResponse extends Message<QueryRecentEpochStatusCountResponse> {
  /**
   * @generated from field: uint64 tip_epoch = 1;
   */
  tipEpoch = protoInt64.zero;

  /**
   * @generated from field: uint64 epoch_count = 2;
   */
  epochCount = protoInt64.zero;

  /**
   * @generated from field: map<string, uint64> status_count = 3;
   */
  statusCount: { [key: string]: bigint } = {};

  constructor(data?: PartialMessage<QueryRecentEpochStatusCountResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryRecentEpochStatusCountResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tip_epoch", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "epoch_count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "status_count", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 4 /* ScalarType.UINT64 */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRecentEpochStatusCountResponse {
    return new QueryRecentEpochStatusCountResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRecentEpochStatusCountResponse {
    return new QueryRecentEpochStatusCountResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRecentEpochStatusCountResponse {
    return new QueryRecentEpochStatusCountResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRecentEpochStatusCountResponse | PlainMessage<QueryRecentEpochStatusCountResponse> | undefined, b: QueryRecentEpochStatusCountResponse | PlainMessage<QueryRecentEpochStatusCountResponse> | undefined): boolean {
    return proto3.util.equals(QueryRecentEpochStatusCountResponse, a, b);
  }
}

/**
 * QueryLastCheckpointWithStatusRequest is the request type for the
 * Query/LastCheckpointWithStatus RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryLastCheckpointWithStatusRequest
 */
export class QueryLastCheckpointWithStatusRequest extends Message<QueryLastCheckpointWithStatusRequest> {
  /**
   * @generated from field: babylon.checkpointing.v1.CheckpointStatus status = 1;
   */
  status = CheckpointStatus.CKPT_STATUS_ACCUMULATING;

  constructor(data?: PartialMessage<QueryLastCheckpointWithStatusRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryLastCheckpointWithStatusRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(CheckpointStatus) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryLastCheckpointWithStatusRequest {
    return new QueryLastCheckpointWithStatusRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryLastCheckpointWithStatusRequest {
    return new QueryLastCheckpointWithStatusRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryLastCheckpointWithStatusRequest {
    return new QueryLastCheckpointWithStatusRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryLastCheckpointWithStatusRequest | PlainMessage<QueryLastCheckpointWithStatusRequest> | undefined, b: QueryLastCheckpointWithStatusRequest | PlainMessage<QueryLastCheckpointWithStatusRequest> | undefined): boolean {
    return proto3.util.equals(QueryLastCheckpointWithStatusRequest, a, b);
  }
}

/**
 * QueryLastCheckpointWithStatusResponse is the response type for the
 * Query/LastCheckpointWithStatus RPC method.
 *
 * @generated from message babylon.checkpointing.v1.QueryLastCheckpointWithStatusResponse
 */
export class QueryLastCheckpointWithStatusResponse extends Message<QueryLastCheckpointWithStatusResponse> {
  /**
   * @generated from field: babylon.checkpointing.v1.RawCheckpointResponse raw_checkpoint = 1;
   */
  rawCheckpoint?: RawCheckpointResponse;

  constructor(data?: PartialMessage<QueryLastCheckpointWithStatusResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.QueryLastCheckpointWithStatusResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "raw_checkpoint", kind: "message", T: RawCheckpointResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryLastCheckpointWithStatusResponse {
    return new QueryLastCheckpointWithStatusResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryLastCheckpointWithStatusResponse {
    return new QueryLastCheckpointWithStatusResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryLastCheckpointWithStatusResponse {
    return new QueryLastCheckpointWithStatusResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryLastCheckpointWithStatusResponse | PlainMessage<QueryLastCheckpointWithStatusResponse> | undefined, b: QueryLastCheckpointWithStatusResponse | PlainMessage<QueryLastCheckpointWithStatusResponse> | undefined): boolean {
    return proto3.util.equals(QueryLastCheckpointWithStatusResponse, a, b);
  }
}

/**
 * RawCheckpointResponse wraps the BLS multi sig with metadata
 *
 * @generated from message babylon.checkpointing.v1.RawCheckpointResponse
 */
export class RawCheckpointResponse extends Message<RawCheckpointResponse> {
  /**
   * epoch_num defines the epoch number the raw checkpoint is for
   *
   * @generated from field: uint64 epoch_num = 1;
   */
  epochNum = protoInt64.zero;

  /**
   * block_hash_hex defines the 'BlockID.Hash', which is the hash of
   * the block that individual BLS sigs are signed on as hex string
   *
   * @generated from field: string block_hash_hex = 2;
   */
  blockHashHex = "";

  /**
   * bitmap defines the bitmap that indicates the signers of the BLS multi sig
   *
   * @generated from field: bytes bitmap = 3;
   */
  bitmap = new Uint8Array(0);

  /**
   * bls_multi_sig defines the multi sig that is aggregated from individual BLS
   * sigs
   *
   * @generated from field: bytes bls_multi_sig = 4;
   */
  blsMultiSig = new Uint8Array(0);

  constructor(data?: PartialMessage<RawCheckpointResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.RawCheckpointResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "epoch_num", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "block_hash_hex", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "bitmap", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "bls_multi_sig", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RawCheckpointResponse {
    return new RawCheckpointResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RawCheckpointResponse {
    return new RawCheckpointResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RawCheckpointResponse {
    return new RawCheckpointResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RawCheckpointResponse | PlainMessage<RawCheckpointResponse> | undefined, b: RawCheckpointResponse | PlainMessage<RawCheckpointResponse> | undefined): boolean {
    return proto3.util.equals(RawCheckpointResponse, a, b);
  }
}

/**
 * CheckpointStateUpdateResponse defines a state transition on the checkpoint.
 *
 * @generated from message babylon.checkpointing.v1.CheckpointStateUpdateResponse
 */
export class CheckpointStateUpdateResponse extends Message<CheckpointStateUpdateResponse> {
  /**
   * state defines the event of a state transition towards this state
   *
   * @generated from field: babylon.checkpointing.v1.CheckpointStatus state = 1;
   */
  state = CheckpointStatus.CKPT_STATUS_ACCUMULATING;

  /**
   * status_desc respresents the description of status enum.
   *
   * @generated from field: string status_desc = 2;
   */
  statusDesc = "";

  /**
   * block_height is the height of the Babylon block that triggers the state
   * update
   *
   * @generated from field: uint64 block_height = 3;
   */
  blockHeight = protoInt64.zero;

  /**
   * block_time is the timestamp in the Babylon block that triggers the state
   * update
   *
   * @generated from field: google.protobuf.Timestamp block_time = 4;
   */
  blockTime?: Timestamp;

  constructor(data?: PartialMessage<CheckpointStateUpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.CheckpointStateUpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "state", kind: "enum", T: proto3.getEnumType(CheckpointStatus) },
    { no: 2, name: "status_desc", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "block_height", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "block_time", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckpointStateUpdateResponse {
    return new CheckpointStateUpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckpointStateUpdateResponse {
    return new CheckpointStateUpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckpointStateUpdateResponse {
    return new CheckpointStateUpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CheckpointStateUpdateResponse | PlainMessage<CheckpointStateUpdateResponse> | undefined, b: CheckpointStateUpdateResponse | PlainMessage<CheckpointStateUpdateResponse> | undefined): boolean {
    return proto3.util.equals(CheckpointStateUpdateResponse, a, b);
  }
}

/**
 * RawCheckpointWithMetaResponse wraps the raw checkpoint with metadata.
 *
 * @generated from message babylon.checkpointing.v1.RawCheckpointWithMetaResponse
 */
export class RawCheckpointWithMetaResponse extends Message<RawCheckpointWithMetaResponse> {
  /**
   * @generated from field: babylon.checkpointing.v1.RawCheckpointResponse ckpt = 1;
   */
  ckpt?: RawCheckpointResponse;

  /**
   * status defines the status of the checkpoint
   *
   * @generated from field: babylon.checkpointing.v1.CheckpointStatus status = 2;
   */
  status = CheckpointStatus.CKPT_STATUS_ACCUMULATING;

  /**
   * status_desc respresents the description of status enum.
   *
   * @generated from field: string status_desc = 3;
   */
  statusDesc = "";

  /**
   * bls_aggr_pk defines the aggregated BLS public key
   *
   * @generated from field: bytes bls_aggr_pk = 4;
   */
  blsAggrPk = new Uint8Array(0);

  /**
   * power_sum defines the accumulated voting power for the checkpoint
   *
   * @generated from field: uint64 power_sum = 5;
   */
  powerSum = protoInt64.zero;

  /**
   * lifecycle defines the lifecycle of this checkpoint, i.e., each state
   * transition and the time (in both timestamp and block height) of this
   * transition.
   *
   * @generated from field: repeated babylon.checkpointing.v1.CheckpointStateUpdateResponse lifecycle = 6;
   */
  lifecycle: CheckpointStateUpdateResponse[] = [];

  constructor(data?: PartialMessage<RawCheckpointWithMetaResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "babylon.checkpointing.v1.RawCheckpointWithMetaResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ckpt", kind: "message", T: RawCheckpointResponse },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(CheckpointStatus) },
    { no: 3, name: "status_desc", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "bls_aggr_pk", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 5, name: "power_sum", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 6, name: "lifecycle", kind: "message", T: CheckpointStateUpdateResponse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RawCheckpointWithMetaResponse {
    return new RawCheckpointWithMetaResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RawCheckpointWithMetaResponse {
    return new RawCheckpointWithMetaResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RawCheckpointWithMetaResponse {
    return new RawCheckpointWithMetaResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RawCheckpointWithMetaResponse | PlainMessage<RawCheckpointWithMetaResponse> | undefined, b: RawCheckpointWithMetaResponse | PlainMessage<RawCheckpointWithMetaResponse> | undefined): boolean {
    return proto3.util.equals(RawCheckpointWithMetaResponse, a, b);
  }
}

