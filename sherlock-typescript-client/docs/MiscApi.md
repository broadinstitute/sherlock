# MiscApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**connectionCheckGet**](MiscApi.md#connectioncheckget) | **GET** /connection-check | Test the client\&#39;s connection to Sherlock |
| [**statusGet**](MiscApi.md#statusget) | **GET** /status | Get Sherlock\&#39;s current status |
| [**versionGet**](MiscApi.md#versionget) | **GET** /version | Get Sherlock\&#39;s own current version |



## connectionCheckGet

> MiscConnectionCheckResponse connectionCheckGet()

Test the client\&#39;s connection to Sherlock

Get a static response from Sherlock to verify connection through proxies like IAP.

### Example

```ts
import {
  Configuration,
  MiscApi,
} from '@sherlock-js-client/sherlock';
import type { ConnectionCheckGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new MiscApi();

  try {
    const data = await api.connectionCheckGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscConnectionCheckResponse**](MiscConnectionCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## statusGet

> MiscStatusResponse statusGet()

Get Sherlock\&#39;s current status

Get Sherlock\&#39;s current status. Right now, this endpoint always returned OK (if the server is online). This endpoint is acceptable to use for a readiness check.

### Example

```ts
import {
  Configuration,
  MiscApi,
} from '@sherlock-js-client/sherlock';
import type { StatusGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new MiscApi();

  try {
    const data = await api.statusGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscStatusResponse**](MiscStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## versionGet

> MiscVersionResponse versionGet()

Get Sherlock\&#39;s own current version

Get the build version of this Sherlock instance.

### Example

```ts
import {
  Configuration,
  MiscApi,
} from '@sherlock-js-client/sherlock';
import type { VersionGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new MiscApi();

  try {
    const data = await api.versionGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**MiscVersionResponse**](MiscVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

