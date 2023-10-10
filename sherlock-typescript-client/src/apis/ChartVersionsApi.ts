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
  SherlockChartVersionV3,
  SherlockChartVersionV3ChangelogResponse,
  SherlockChartVersionV3Create,
  SherlockChartVersionV3Edit,
  V2controllersChartVersion,
  V2controllersCreatableChartVersion,
  V2controllersEditableChartVersion,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockChartVersionV3FromJSON,
    SherlockChartVersionV3ToJSON,
    SherlockChartVersionV3ChangelogResponseFromJSON,
    SherlockChartVersionV3ChangelogResponseToJSON,
    SherlockChartVersionV3CreateFromJSON,
    SherlockChartVersionV3CreateToJSON,
    SherlockChartVersionV3EditFromJSON,
    SherlockChartVersionV3EditToJSON,
    V2controllersChartVersionFromJSON,
    V2controllersChartVersionToJSON,
    V2controllersCreatableChartVersionFromJSON,
    V2controllersCreatableChartVersionToJSON,
    V2controllersEditableChartVersionFromJSON,
    V2controllersEditableChartVersionToJSON,
} from '../models/index';

export interface ApiChartVersionsProceduresV3ChangelogGetRequest {
    child: string;
    parent: string;
}

export interface ApiChartVersionsV3GetRequest {
    authoredBy?: string;
    chart?: string;
    chartVersion?: string;
    createdAt?: Date;
    description?: string;
    id?: number;
    parentChartVersion?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiChartVersionsV3PostRequest {
    chartVersion: SherlockChartVersionV3Create;
}

export interface ApiChartVersionsV3SelectorGetRequest {
    selector: string;
}

export interface ApiChartVersionsV3SelectorPatchRequest {
    selector: string;
    chartVersion: SherlockChartVersionV3Edit;
}

export interface ApiV2ChartVersionsGetRequest {
    chart?: string;
    chartVersion?: string;
    createdAt?: Date;
    description?: string;
    id?: number;
    parentChartVersion?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2ChartVersionsPostRequest {
    chartVersion: V2controllersCreatableChartVersion;
}

export interface ApiV2ChartVersionsSelectorGetRequest {
    selector: string;
}

export interface ApiV2ChartVersionsSelectorPatchRequest {
    selector: string;
    chartVersion: V2controllersEditableChartVersion;
}

export interface ApiV2ChartVersionsSelectorPutRequest {
    selector: string;
    chartVersion: V2controllersCreatableChartVersion;
}

export interface ApiV2ProceduresChartVersionsChildrenPathToParentGetRequest {
    child: string;
    parent: string;
}

export interface ApiV2SelectorsChartVersionsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class ChartVersionsApi extends runtime.BaseAPI {

    /**
     * Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent.
     * Get a changelog between two ChartVersions
     */
    async apiChartVersionsProceduresV3ChangelogGetRaw(requestParameters: ApiChartVersionsProceduresV3ChangelogGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartVersionV3ChangelogResponse>> {
        if (requestParameters.child === null || requestParameters.child === undefined) {
            throw new runtime.RequiredError('child','Required parameter requestParameters.child was null or undefined when calling apiChartVersionsProceduresV3ChangelogGet.');
        }

        if (requestParameters.parent === null || requestParameters.parent === undefined) {
            throw new runtime.RequiredError('parent','Required parameter requestParameters.parent was null or undefined when calling apiChartVersionsProceduresV3ChangelogGet.');
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
            path: `/api/chart-versions/procedures/v3/changelog`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartVersionV3ChangelogResponseFromJSON(jsonValue));
    }

    /**
     * Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent.
     * Get a changelog between two ChartVersions
     */
    async apiChartVersionsProceduresV3ChangelogGet(requestParameters: ApiChartVersionsProceduresV3ChangelogGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartVersionV3ChangelogResponse> {
        const response = await this.apiChartVersionsProceduresV3ChangelogGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List ChartVersions matching a filter.
     * List ChartVersions matching a filter
     */
    async apiChartVersionsV3GetRaw(requestParameters: ApiChartVersionsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockChartVersionV3>>> {
        const queryParameters: any = {};

        if (requestParameters.authoredBy !== undefined) {
            queryParameters['authoredBy'] = requestParameters.authoredBy;
        }

        if (requestParameters.chart !== undefined) {
            queryParameters['chart'] = requestParameters.chart;
        }

        if (requestParameters.chartVersion !== undefined) {
            queryParameters['chartVersion'] = requestParameters.chartVersion;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.parentChartVersion !== undefined) {
            queryParameters['parentChartVersion'] = requestParameters.parentChartVersion;
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
            path: `/api/chart-versions/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockChartVersionV3FromJSON));
    }

    /**
     * List ChartVersions matching a filter.
     * List ChartVersions matching a filter
     */
    async apiChartVersionsV3Get(requestParameters: ApiChartVersionsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockChartVersionV3>> {
        const response = await this.apiChartVersionsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Upsert a ChartVersion.
     * Upsert a ChartVersion
     */
    async apiChartVersionsV3PostRaw(requestParameters: ApiChartVersionsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartVersionV3>> {
        if (requestParameters.chartVersion === null || requestParameters.chartVersion === undefined) {
            throw new runtime.RequiredError('chartVersion','Required parameter requestParameters.chartVersion was null or undefined when calling apiChartVersionsV3Post.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/chartVersions/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartVersionV3CreateToJSON(requestParameters.chartVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartVersionV3FromJSON(jsonValue));
    }

    /**
     * Upsert a ChartVersion.
     * Upsert a ChartVersion
     */
    async apiChartVersionsV3Post(requestParameters: ApiChartVersionsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartVersionV3> {
        const response = await this.apiChartVersionsV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual ChartVersion.
     * Get an individual ChartVersion
     */
    async apiChartVersionsV3SelectorGetRaw(requestParameters: ApiChartVersionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartVersionV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiChartVersionsV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/chart-versions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartVersionV3FromJSON(jsonValue));
    }

    /**
     * Get an individual ChartVersion.
     * Get an individual ChartVersion
     */
    async apiChartVersionsV3SelectorGet(requestParameters: ApiChartVersionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartVersionV3> {
        const response = await this.apiChartVersionsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual ChartVersion.
     * Edit an individual ChartVersion
     */
    async apiChartVersionsV3SelectorPatchRaw(requestParameters: ApiChartVersionsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartVersionV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiChartVersionsV3SelectorPatch.');
        }

        if (requestParameters.chartVersion === null || requestParameters.chartVersion === undefined) {
            throw new runtime.RequiredError('chartVersion','Required parameter requestParameters.chartVersion was null or undefined when calling apiChartVersionsV3SelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/chart-versions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartVersionV3EditToJSON(requestParameters.chartVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartVersionV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual ChartVersion.
     * Edit an individual ChartVersion
     */
    async apiChartVersionsV3SelectorPatch(requestParameters: ApiChartVersionsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartVersionV3> {
        const response = await this.apiChartVersionsV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing ChartVersion entries, ordered by most recently updated.
     * List ChartVersion entries
     */
    async apiV2ChartVersionsGetRaw(requestParameters: ApiV2ChartVersionsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChartVersion>>> {
        const queryParameters: any = {};

        if (requestParameters.chart !== undefined) {
            queryParameters['chart'] = requestParameters.chart;
        }

        if (requestParameters.chartVersion !== undefined) {
            queryParameters['chartVersion'] = requestParameters.chartVersion;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.parentChartVersion !== undefined) {
            queryParameters['parentChartVersion'] = requestParameters.parentChartVersion;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/chart-versions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChartVersionFromJSON));
    }

    /**
     * List existing ChartVersion entries, ordered by most recently updated.
     * List ChartVersion entries
     */
    async apiV2ChartVersionsGet(requestParameters: ApiV2ChartVersionsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChartVersion>> {
        const response = await this.apiV2ChartVersionsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new ChartVersion entry. Note that fields are immutable after creation. If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
     * Create a new ChartVersion entry
     */
    async apiV2ChartVersionsPostRaw(requestParameters: ApiV2ChartVersionsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartVersion>> {
        if (requestParameters.chartVersion === null || requestParameters.chartVersion === undefined) {
            throw new runtime.RequiredError('chartVersion','Required parameter requestParameters.chartVersion was null or undefined when calling apiV2ChartVersionsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/chart-versions`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChartVersionToJSON(requestParameters.chartVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartVersionFromJSON(jsonValue));
    }

    /**
     * Create a new ChartVersion entry. Note that fields are immutable after creation. If the new entry is a duplicate of one already in the database, the database will not be altered and the call will return normally but with a 200 code.
     * Create a new ChartVersion entry
     */
    async apiV2ChartVersionsPost(requestParameters: ApiV2ChartVersionsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartVersion> {
        const response = await this.apiV2ChartVersionsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing ChartVersion entry via one its \"selectors\": chart/version or numeric ID.
     * Get a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorGetRaw(requestParameters: ApiV2ChartVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartVersionsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/chart-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartVersionFromJSON(jsonValue));
    }

    /**
     * Get an existing ChartVersion entry via one its \"selectors\": chart/version or numeric ID.
     * Get a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorGet(requestParameters: ApiV2ChartVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartVersion> {
        const response = await this.apiV2ChartVersionsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing ChartVersion entry via one its \"selectors\": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorPatchRaw(requestParameters: ApiV2ChartVersionsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartVersionsSelectorPatch.');
        }

        if (requestParameters.chartVersion === null || requestParameters.chartVersion === undefined) {
            throw new runtime.RequiredError('chartVersion','Required parameter requestParameters.chartVersion was null or undefined when calling apiV2ChartVersionsSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/chart-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableChartVersionToJSON(requestParameters.chartVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartVersionFromJSON(jsonValue));
    }

    /**
     * Edit an existing ChartVersion entry via one its \"selectors\": chart/version or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorPatch(requestParameters: ApiV2ChartVersionsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartVersion> {
        const response = await this.apiV2ChartVersionsSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit a ChartVersion entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorPutRaw(requestParameters: ApiV2ChartVersionsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartVersion>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartVersionsSelectorPut.');
        }

        if (requestParameters.chartVersion === null || requestParameters.chartVersion === undefined) {
            throw new runtime.RequiredError('chartVersion','Required parameter requestParameters.chartVersion was null or undefined when calling apiV2ChartVersionsSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/chart-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChartVersionToJSON(requestParameters.chartVersion),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartVersionFromJSON(jsonValue));
    }

    /**
     * Create or edit a ChartVersion entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a ChartVersion entry
     */
    async apiV2ChartVersionsSelectorPut(requestParameters: ApiV2ChartVersionsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartVersion> {
        const response = await this.apiV2ChartVersionsSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent. If the child can\'t be connected to the parent, just the child will be returned with a 204 code.
     * Get a changelog between two ChartVersions
     */
    async apiV2ProceduresChartVersionsChildrenPathToParentGetRaw(requestParameters: ApiV2ProceduresChartVersionsChildrenPathToParentGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChartVersion>>> {
        if (requestParameters.child === null || requestParameters.child === undefined) {
            throw new runtime.RequiredError('child','Required parameter requestParameters.child was null or undefined when calling apiV2ProceduresChartVersionsChildrenPathToParentGet.');
        }

        if (requestParameters.parent === null || requestParameters.parent === undefined) {
            throw new runtime.RequiredError('parent','Required parameter requestParameters.parent was null or undefined when calling apiV2ProceduresChartVersionsChildrenPathToParentGet.');
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
            path: `/api/v2/procedures/chart-versions/children-path-to-parent`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChartVersionFromJSON));
    }

    /**
     * Get the path through parent references from a child ChartVersion (inclusive) to a parent ChartVersion (exclusive), if possible. Because parent references point from newer children to older parents, the newer ChartVersion should be the child. The result will always exclude the parent. If the child can\'t be connected to the parent, just the child will be returned with a 204 code.
     * Get a changelog between two ChartVersions
     */
    async apiV2ProceduresChartVersionsChildrenPathToParentGet(requestParameters: ApiV2ProceduresChartVersionsChildrenPathToParentGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChartVersion>> {
        const response = await this.apiV2ProceduresChartVersionsChildrenPathToParentGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.
     * List ChartVersion selectors
     */
    async apiV2SelectorsChartVersionsSelectorGetRaw(requestParameters: ApiV2SelectorsChartVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsChartVersionsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/chart-versions/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given ChartVersion selector and provide any other selectors that would match the same ChartVersion.
     * List ChartVersion selectors
     */
    async apiV2SelectorsChartVersionsSelectorGet(requestParameters: ApiV2SelectorsChartVersionsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsChartVersionsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
