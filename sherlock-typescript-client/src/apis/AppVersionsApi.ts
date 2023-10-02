/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service. Note: this API will try to load and return associations in responses, so clients won\'t need to make as many requests. This behavior isn\'t recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
 *
 * The version of the OpenAPI document: development
 * Contact: dsp-devops@broadinstitute.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ErrorsErrorResponse,
  SherlockAppVersionV3,
  V2controllersAppVersion,
  V2controllersCreatableAppVersion,
  V2controllersEditableAppVersion,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockAppVersionV3FromJSON,
    SherlockAppVersionV3ToJSON,
    V2controllersAppVersionFromJSON,
    V2controllersAppVersionToJSON,
    V2controllersCreatableAppVersionFromJSON,
    V2controllersCreatableAppVersionToJSON,
    V2controllersEditableAppVersionFromJSON,
    V2controllersEditableAppVersionToJSON,
} from '../models/index';

export interface ApiAppVersionsV3GetRequest {
    appVersion?: string;
    chart?: string;
    createdAt?: Date;
    description?: string;
    gitBranch?: string;
    gitCommit?: string;
    id?: number;
    parentAppVersion?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiAppVersionsV3SelectorGetRequest {
    selector: string;
}

export interface ApiV2AppVersionsGetRequest {
    appVersion?: string;
    chart?: string;
    createdAt?: Date;
    description?: string;
    gitBranch?: string;
    gitCommit?: string;
    id?: number;
    parentAppVersion?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2AppVersionsPostRequest {
    appVersion: V2controllersCreatableAppVersion;
}

export interface ApiV2AppVersionsSelectorGetRequest {
    selector: string;
}

export interface ApiV2AppVersionsSelectorPatchRequest {
    selector: string;
    appVersion: V2controllersEditableAppVersion;
}

export interface ApiV2AppVersionsSelectorPutRequest {
    selector: string;
    appVersion: V2controllersCreatableAppVersion;
}

export interface ApiV2ProceduresAppVersionsChildrenPathToParentGetRequest {
    child: string;
    parent: string;
}

export interface ApiV2SelectorsAppVersionsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class AppVersionsApi extends runtime.BaseAPI {

    /**
     * List AppVersions matching a filter.
     * List AppVersions matching a filter
     */
    async apiAppVersionsV3GetRaw(requestParameters: ApiAppVersionsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockAppVersionV3>>> {
        const queryParameters: any = {};

        if (requestParameters.appVersion !== undefined) {
            queryParameters['appVersion'] = requestParameters.appVersion;
        }

        if (requestParameters.chart !== undefined) {
            queryParameters['chart'] = requestParameters.chart;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.gitBranch !== undefined) {
            queryParameters['gitBranch'] = requestParameters.gitBranch;
        }

        if (requestParameters.gitCommit !== undefined) {
            queryParameters['gitCommit'] = requestParameters.gitCommit;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.parentAppVersion !== undefined) {
            queryParameters['parentAppVersion'] = requestParameters.parentAppVersion;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/app-versions/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockAppVersionV3FromJSON));
    }

    /**
     * List AppVersions matching a filter.
     * List AppVersions matching a filter
     */
    async apiAppVersionsV3Get(requestParameters: ApiAppVersionsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockAppVersionV3>> {
        const response = await this.apiAppVersionsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual AppVersion.
     * Get an individual AppVersion
     */
    async apiAppVersionsV3SelectorGetRaw(requestParameters: ApiAppVersionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockAppVersionV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiAppVersionsV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/app-versions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockAppVersionV3FromJSON(jsonValue));
    }

    /**
     * Get an individual AppVersion.
     * Get an individual AppVersion
     */
    async apiAppVersionsV3SelectorGet(requestParameters: ApiAppVersionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockAppVersionV3> {
        const response = await this.apiAppVersionsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing AppVersion entries, ordered by most recently updated.
     * List AppVersion entries
     */
    async apiV2AppVersionsGetRaw(requestParameters: ApiV2AppVersionsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersAppVersion>>> {
        const queryParameters: any = {};

        if (requestParameters.appVersion !== undefined) {
            queryParameters['appVersion'] = requestParameters.appVersion;
        }

        if (requestParameters.chart !== undefined) {
            queryParameters['chart'] = requestParameters.chart;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.gitBranch !== undefined) {
            queryParameters['gitBranch'] = requestParameters.gitBranch;
        }

        if (requestParameters.gitCommit !== undefined) {
            queryParameters['gitCommit'] = requestParameters.gitCommit;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.parentAppVersion !== undefined) {
            queryParameters['parentAppVersion'] = requestParameters.parentAppVersion;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/app-versions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersAppVersionFromJSON));
    }

    /**
     * List existing AppVersion entries, ordered by most recently updated.
     * List AppVersion entries
     */
    async apiV2AppVersionsGet(requestParameters: ApiV2AppVersionsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersAppVersion>> {
        const response = await this.apiV2AppVersionsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new AppVersion entry. Note that fields are immutable after creation. If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
     * Create a new AppVersion entry
     */
    async apiV2AppVersionsPostRaw(requestParameters: ApiV2AppVersionsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersAppVersion>> {
        if (requestParameters.appVersion === null || requestParameters.appVersion === undefined) {
            throw new runtime.RequiredError('appVersion','Required parameter requestParameters.appVersion was null or undefined when calling apiV2AppVersionsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/app-versions`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableAppVersionToJSON(requestParameters.appVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersAppVersionFromJSON(jsonValue));
    }

    /**
     * Create a new AppVersion entry. Note that fields are immutable after creation. If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
     * Create a new AppVersion entry
     */
    async apiV2AppVersionsPost(requestParameters: ApiV2AppVersionsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersAppVersion> {
        const response = await this.apiV2AppVersionsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing AppVersion entry via one its \"selectors\": chart/version or numeric ID.
     * Get a AppVersion entry
     */
    async apiV2AppVersionsSelectorGetRaw(requestParameters: ApiV2AppVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersAppVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2AppVersionsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/app-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersAppVersionFromJSON(jsonValue));
    }

    /**
     * Get an existing AppVersion entry via one its \"selectors\": chart/version or numeric ID.
     * Get a AppVersion entry
     */
    async apiV2AppVersionsSelectorGet(requestParameters: ApiV2AppVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersAppVersion> {
        const response = await this.apiV2AppVersionsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing AppVersion entry via one its \"selectors\": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a AppVersion entry
     */
    async apiV2AppVersionsSelectorPatchRaw(requestParameters: ApiV2AppVersionsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersAppVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2AppVersionsSelectorPatch.');
        }

        if (requestParameters.appVersion === null || requestParameters.appVersion === undefined) {
            throw new runtime.RequiredError('appVersion','Required parameter requestParameters.appVersion was null or undefined when calling apiV2AppVersionsSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/app-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableAppVersionToJSON(requestParameters.appVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersAppVersionFromJSON(jsonValue));
    }

    /**
     * Edit an existing AppVersion entry via one its \"selectors\": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a AppVersion entry
     */
    async apiV2AppVersionsSelectorPatch(requestParameters: ApiV2AppVersionsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersAppVersion> {
        const response = await this.apiV2AppVersionsSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit an AppVersion entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit an AppVersion entry
     */
    async apiV2AppVersionsSelectorPutRaw(requestParameters: ApiV2AppVersionsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersAppVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2AppVersionsSelectorPut.');
        }

        if (requestParameters.appVersion === null || requestParameters.appVersion === undefined) {
            throw new runtime.RequiredError('appVersion','Required parameter requestParameters.appVersion was null or undefined when calling apiV2AppVersionsSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/app-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableAppVersionToJSON(requestParameters.appVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersAppVersionFromJSON(jsonValue));
    }

    /**
     * Create or edit an AppVersion entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit an AppVersion entry
     */
    async apiV2AppVersionsSelectorPut(requestParameters: ApiV2AppVersionsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersAppVersion> {
        const response = await this.apiV2AppVersionsSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent. If the child can\'t be connected to the parent, just the child will be returned with a 204 code.
     * Get a changelog between two AppVersions
     */
    async apiV2ProceduresAppVersionsChildrenPathToParentGetRaw(requestParameters: ApiV2ProceduresAppVersionsChildrenPathToParentGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersAppVersion>>> {
        if (requestParameters.child === null || requestParameters.child === undefined) {
            throw new runtime.RequiredError('child','Required parameter requestParameters.child was null or undefined when calling apiV2ProceduresAppVersionsChildrenPathToParentGet.');
        }

        if (requestParameters.parent === null || requestParameters.parent === undefined) {
            throw new runtime.RequiredError('parent','Required parameter requestParameters.parent was null or undefined when calling apiV2ProceduresAppVersionsChildrenPathToParentGet.');
        }

        const queryParameters: any = {};

        if (requestParameters.child !== undefined) {
            queryParameters['child'] = requestParameters.child;
        }

        if (requestParameters.parent !== undefined) {
            queryParameters['parent'] = requestParameters.parent;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/procedures/app-versions/children-path-to-parent`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersAppVersionFromJSON));
    }

    /**
     * Get the path through parent references from a child AppVersion (inclusive) to a parent AppVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer AppVersion should be the child. The result will always exclude the parent. If the child can\'t be connected to the parent, just the child will be returned with a 204 code.
     * Get a changelog between two AppVersions
     */
    async apiV2ProceduresAppVersionsChildrenPathToParentGet(requestParameters: ApiV2ProceduresAppVersionsChildrenPathToParentGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersAppVersion>> {
        const response = await this.apiV2ProceduresAppVersionsChildrenPathToParentGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
     * List AppVersion selectors
     */
    async apiV2SelectorsAppVersionsSelectorGetRaw(requestParameters: ApiV2SelectorsAppVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsAppVersionsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/app-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given AppVersion selector and provide any other selectors that would match the same AppVersion.
     * List AppVersion selectors
     */
    async apiV2SelectorsAppVersionsSelectorGet(requestParameters: ApiV2SelectorsAppVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsAppVersionsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
