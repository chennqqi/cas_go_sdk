# Signature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QSignAlgorithm** | **string** | 签名算法默认sha1 | [optional] 
**QAk** | **string** | 用于标识用户身份， SecretID 的字段，在腾讯云的 API 密钥页面中可以查看。 | [optional] 
**QSignTime** | **int64** | 签名的有效起止时间，其使用10位Unix时间戳来表示，有效效力精确到秒。该字段通过分号区分起止，起时在前、止时在后。 | [optional] 
**QKeyTime** | **int64** | 用户可以自定义 SignKey 有效时间，使用 10 位 Unix 时间戳来表示，有效效力精确到秒。该字段通过分号区分起止，起始时间在前、终止时间在后。一般 q-key-time 的时间范围大于等于 q-sign-time。 | [optional] 
**QHeaderList** | **string** | 提供密文中包含需要校验的 Headers 列表，必须是小写字符，且需要按字典序排序，以\&quot;;\&quot;分隔 | [optional] 
**QUrlParamList** | **string** | 提供密文中包含需要校验的 Parameters 列表，必须是小写字符，以\&quot;;\&quot;分隔 | [optional] 
**QSignature** | **string** | 经过 HMAC-SHA1 算法加密的请求校验信息。 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


