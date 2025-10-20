# UsersApi

All URIs are relative to *https://sherlock.dsp-devops-prod.broadinstitute.org*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiUsersProceduresV3DeactivatePost**](UsersApi.md#apiusersproceduresv3deactivatepost) | **POST** /api/users/procedures/v3/deactivate | Deactivate Users |
| [**apiUsersV3Get**](UsersApi.md#apiusersv3get) | **GET** /api/users/v3 | List Users matching a filter |
| [**apiUsersV3Put**](UsersApi.md#apiusersv3put) | **PUT** /api/users/v3 | Update the calling User\&#39;s information |
| [**apiUsersV3SelectorGet**](UsersApi.md#apiusersv3selectorget) | **GET** /api/users/v3/{selector} | Get an individual User |



## apiUsersProceduresV3DeactivatePost

> SherlockUserV3DeactivateResponse apiUsersProceduresV3DeactivatePost(users)

Deactivate Users

Super-admin only method to deactivate users. Deactivated users will be removed from all roles and can\&#39;t authenticate to Sherlock. This endpoint can optionally also attempt to suspend the same email handles across given Google Workspace domains, substituting email domains as necessary. It will do so by impersonating the caller in each given domain.

### Example

```ts
import {
  Configuration,
  UsersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiUsersProceduresV3DeactivatePostRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new UsersApi();

  const body = {
    // SherlockUserV3DeactivateRequest | Information on the users to deactivate
    users: ...,
  } satisfies ApiUsersProceduresV3DeactivatePostRequest;

  try {
    const data = await api.apiUsersProceduresV3DeactivatePost(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **users** | [SherlockUserV3DeactivateRequest](SherlockUserV3DeactivateRequest.md) | Information on the users to deactivate | |

### Return type

[**SherlockUserV3DeactivateResponse**](SherlockUserV3DeactivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiUsersV3Get

> Array&lt;SherlockUserV3&gt; apiUsersV3Get(createdAt, deactivatedAt, email, githubID, githubUsername, googleID, id, name, nameFrom, slackID, slackUsername, suitabilityDescription, suitable, updatedAt, limit, offset, includeDeactivated)

List Users matching a filter

List Users matching a filter. The results will include suitability and other information. Note that the suitability info can\&#39;t directly be filtered for at this time.

### Example

```ts
import {
  Configuration,
  UsersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiUsersV3GetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new UsersApi();

  const body = {
    // Date (optional)
    createdAt: 2013-10-20T19:20:30+01:00,
    // string | If set, indicates that the user is currently deactivated (optional)
    deactivatedAt: deactivatedAt_example,
    // string (optional)
    email: email_example,
    // string (optional)
    githubID: githubID_example,
    // string (optional)
    githubUsername: githubUsername_example,
    // string (optional)
    googleID: googleID_example,
    // number (optional)
    id: 56,
    // string (optional)
    name: name_example,
    // 'sherlock' | 'github' | 'slack' (optional)
    nameFrom: nameFrom_example,
    // string (optional)
    slackID: slackID_example,
    // string (optional)
    slackUsername: slackUsername_example,
    // string | Available only in responses; describes the user\'s production-suitability (optional)
    suitabilityDescription: suitabilityDescription_example,
    // boolean | Available only in responses; indicates whether the user is production-suitable (optional)
    suitable: true,
    // Date (optional)
    updatedAt: 2013-10-20T19:20:30+01:00,
    // number | Control how many Users are returned (default 0, no limit) (optional)
    limit: 56,
    // number | Control the offset for the returned Users (default 0) (optional)
    offset: 56,
    // boolean | Include deactivated users in the results (default false) (optional)
    includeDeactivated: true,
  } satisfies ApiUsersV3GetRequest;

  try {
    const data = await api.apiUsersV3Get(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createdAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **deactivatedAt** | `string` | If set, indicates that the user is currently deactivated | [Optional] [Defaults to `undefined`] |
| **email** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubID** | `string` |  | [Optional] [Defaults to `undefined`] |
| **githubUsername** | `string` |  | [Optional] [Defaults to `undefined`] |
| **googleID** | `string` |  | [Optional] [Defaults to `undefined`] |
| **id** | `number` |  | [Optional] [Defaults to `undefined`] |
| **name** | `string` |  | [Optional] [Defaults to `undefined`] |
| **nameFrom** | `sherlock`, `github`, `slack` |  | [Optional] [Defaults to `undefined`] [Enum: sherlock, github, slack] |
| **slackID** | `string` |  | [Optional] [Defaults to `undefined`] |
| **slackUsername** | `string` |  | [Optional] [Defaults to `undefined`] |
| **suitabilityDescription** | `string` | Available only in responses; describes the user\&#39;s production-suitability | [Optional] [Defaults to `undefined`] |
| **suitable** | `boolean` | Available only in responses; indicates whether the user is production-suitable | [Optional] [Defaults to `undefined`] |
| **updatedAt** | `Date` |  | [Optional] [Defaults to `undefined`] |
| **limit** | `number` | Control how many Users are returned (default 0, no limit) | [Optional] [Defaults to `undefined`] |
| **offset** | `number` | Control the offset for the returned Users (default 0) | [Optional] [Defaults to `undefined`] |
| **includeDeactivated** | `boolean` | Include deactivated users in the results (default false) | [Optional] [Defaults to `undefined`] |

### Return type

[**Array&lt;SherlockUserV3&gt;**](SherlockUserV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiUsersV3Put

> SherlockUserV3 apiUsersV3Put(user)

Update the calling User\&#39;s information

Update the calling User\&#39;s information. As with all authenticated Sherlock endpoints, newly-observed callers will have a User record added, meaning that this endpoint behaves like an upsert.

### Example

```ts
import {
  Configuration,
  UsersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiUsersV3PutRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new UsersApi();

  const body = {
    // SherlockUserV3Upsert | The User data to update (optional)
    user: ...,
  } satisfies ApiUsersV3PutRequest;

  try {
    const data = await api.apiUsersV3Put(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user** | [SherlockUserV3Upsert](SherlockUserV3Upsert.md) | The User data to update | [Optional] |

### Return type

[**SherlockUserV3**](SherlockUserV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **201** | Created |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## apiUsersV3SelectorGet

> SherlockUserV3 apiUsersV3SelectorGet(selector)

Get an individual User

Get an individual User. As a special case, \&quot;me\&quot; or \&quot;self\&quot; can be passed as the selector to get the current user.

### Example

```ts
import {
  Configuration,
  UsersApi,
} from '@sherlock-js-client/sherlock';
import type { ApiUsersV3SelectorGetRequest } from '@sherlock-js-client/sherlock';

async function example() {
  console.log("🚀 Testing @sherlock-js-client/sherlock SDK...");
  const api = new UsersApi();

  const body = {
    // string | The selector of the User, which can be either a numeric ID, the email, \'google-id/{google subject ID}\', \'github/{github username}\', or \'github-id/{github numeric ID}\'. As a special case, \'me\' or \'self\' can be passed to get the calling user.
    selector: selector_example,
  } satisfies ApiUsersV3SelectorGetRequest;

  try {
    const data = await api.apiUsersV3SelectorGet(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **selector** | `string` | The selector of the User, which can be either a numeric ID, the email, \&#39;google-id/{google subject ID}\&#39;, \&#39;github/{github username}\&#39;, or \&#39;github-id/{github numeric ID}\&#39;. As a special case, \&#39;me\&#39; or \&#39;self\&#39; can be passed to get the calling user. | [Defaults to `undefined`] |

### Return type

[**SherlockUserV3**](SherlockUserV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Bad Request |  -  |
| **403** | Forbidden |  -  |
| **404** | Not Found |  -  |
| **407** | Proxy Authentication Required |  -  |
| **409** | Conflict |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

