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
  SherlockChartReleaseV3,
  SherlockChartReleaseV3Create,
  SherlockChartReleaseV3Edit,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockChartReleaseV3FromJSON,
    SherlockChartReleaseV3ToJSON,
    SherlockChartReleaseV3CreateFromJSON,
    SherlockChartReleaseV3CreateToJSON,
    SherlockChartReleaseV3EditFromJSON,
    SherlockChartReleaseV3EditToJSON,
} from '../models/index';

export interface ApiChartReleasesV3GetRequest {
    appVersionBranch?: string;
    appVersionCommit?: string;
    appVersionExact?: string;
    appVersionFollowChartRelease?: string;
    appVersionReference?: string;
    appVersionResolver?: ApiChartReleasesV3GetAppVersionResolverEnum;
    chart?: string;
    chartVersionExact?: string;
    chartVersionFollowChartRelease?: string;
    chartVersionReference?: string;
    chartVersionResolver?: ApiChartReleasesV3GetChartVersionResolverEnum;
    cluster?: string;
    createdAt?: Date;
    destinationType?: string;
    environment?: string;
    helmfileRef?: string;
    helmfileRefEnabled?: boolean;
    id?: number;
    includedInBulkChangesets?: boolean;
    name?: string;
    namespace?: string;
    pagerdutyIntegration?: string;
    port?: number;
    protocol?: string;
    resolvedAt?: Date;
    subdomain?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiChartReleasesV3PostRequest {
    chartRelease: SherlockChartReleaseV3Create;
}

export interface ApiChartReleasesV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiChartReleasesV3SelectorGetRequest {
    selector: string;
}

export interface ApiChartReleasesV3SelectorPatchRequest {
    selector: string;
    chartRelease: SherlockChartReleaseV3Edit;
}

/**
 * 
 */
export class ChartReleasesApi extends runtime.BaseAPI {

    /**
     * List ChartReleases matching a filter.
     * List ChartReleases matching a filter
     */
    async apiChartReleasesV3GetRaw(requestParameters: ApiChartReleasesV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockChartReleaseV3>>> {
        const queryParameters: any = {};

        if (requestParameters['appVersionBranch'] != null) {
            queryParameters['appVersionBranch'] = requestParameters['appVersionBranch'];
        }

        if (requestParameters['appVersionCommit'] != null) {
            queryParameters['appVersionCommit'] = requestParameters['appVersionCommit'];
        }

        if (requestParameters['appVersionExact'] != null) {
            queryParameters['appVersionExact'] = requestParameters['appVersionExact'];
        }

        if (requestParameters['appVersionFollowChartRelease'] != null) {
            queryParameters['appVersionFollowChartRelease'] = requestParameters['appVersionFollowChartRelease'];
        }

        if (requestParameters['appVersionReference'] != null) {
            queryParameters['appVersionReference'] = requestParameters['appVersionReference'];
        }

        if (requestParameters['appVersionResolver'] != null) {
            queryParameters['appVersionResolver'] = requestParameters['appVersionResolver'];
        }

        if (requestParameters['chart'] != null) {
            queryParameters['chart'] = requestParameters['chart'];
        }

        if (requestParameters['chartVersionExact'] != null) {
            queryParameters['chartVersionExact'] = requestParameters['chartVersionExact'];
        }

        if (requestParameters['chartVersionFollowChartRelease'] != null) {
            queryParameters['chartVersionFollowChartRelease'] = requestParameters['chartVersionFollowChartRelease'];
        }

        if (requestParameters['chartVersionReference'] != null) {
            queryParameters['chartVersionReference'] = requestParameters['chartVersionReference'];
        }

        if (requestParameters['chartVersionResolver'] != null) {
            queryParameters['chartVersionResolver'] = requestParameters['chartVersionResolver'];
        }

        if (requestParameters['cluster'] != null) {
            queryParameters['cluster'] = requestParameters['cluster'];
        }

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['destinationType'] != null) {
            queryParameters['destinationType'] = requestParameters['destinationType'];
        }

        if (requestParameters['environment'] != null) {
            queryParameters['environment'] = requestParameters['environment'];
        }

        if (requestParameters['helmfileRef'] != null) {
            queryParameters['helmfileRef'] = requestParameters['helmfileRef'];
        }

        if (requestParameters['helmfileRefEnabled'] != null) {
            queryParameters['helmfileRefEnabled'] = requestParameters['helmfileRefEnabled'];
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['includedInBulkChangesets'] != null) {
            queryParameters['includedInBulkChangesets'] = requestParameters['includedInBulkChangesets'];
        }

        if (requestParameters['name'] != null) {
            queryParameters['name'] = requestParameters['name'];
        }

        if (requestParameters['namespace'] != null) {
            queryParameters['namespace'] = requestParameters['namespace'];
        }

        if (requestParameters['pagerdutyIntegration'] != null) {
            queryParameters['pagerdutyIntegration'] = requestParameters['pagerdutyIntegration'];
        }

        if (requestParameters['port'] != null) {
            queryParameters['port'] = requestParameters['port'];
        }

        if (requestParameters['protocol'] != null) {
            queryParameters['protocol'] = requestParameters['protocol'];
        }

        if (requestParameters['resolvedAt'] != null) {
            queryParameters['resolvedAt'] = (requestParameters['resolvedAt'] as any).toISOString();
        }

        if (requestParameters['subdomain'] != null) {
            queryParameters['subdomain'] = requestParameters['subdomain'];
        }

        if (requestParameters['updatedAt'] != null) {
            queryParameters['updatedAt'] = (requestParameters['updatedAt'] as any).toISOString();
        }

        if (requestParameters['limit'] != null) {
            queryParameters['limit'] = requestParameters['limit'];
        }

        if (requestParameters['offset'] != null) {
            queryParameters['offset'] = requestParameters['offset'];
        }

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/chart-releases/v3`;

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockChartReleaseV3FromJSON));
    }

    /**
     * List ChartReleases matching a filter.
     * List ChartReleases matching a filter
     */
    async apiChartReleasesV3Get(requestParameters: ApiChartReleasesV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockChartReleaseV3>> {
        const response = await this.apiChartReleasesV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a ChartRelease.
     * Create a ChartRelease
     */
    async apiChartReleasesV3PostRaw(requestParameters: ApiChartReleasesV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartReleaseV3>> {
        if (requestParameters['chartRelease'] == null) {
            throw new runtime.RequiredError(
                'chartRelease',
                'Required parameter "chartRelease" was null or undefined when calling apiChartReleasesV3Post().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';


        let urlPath = `/api/chart-releases/v3`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartReleaseV3CreateToJSON(requestParameters['chartRelease']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartReleaseV3FromJSON(jsonValue));
    }

    /**
     * Create a ChartRelease.
     * Create a ChartRelease
     */
    async apiChartReleasesV3Post(requestParameters: ApiChartReleasesV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartReleaseV3> {
        const response = await this.apiChartReleasesV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual ChartRelease by its ID.
     * Delete an individual ChartRelease
     */
    async apiChartReleasesV3SelectorDeleteRaw(requestParameters: ApiChartReleasesV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartReleaseV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiChartReleasesV3SelectorDelete().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/chart-releases/v3/{selector}`;
        urlPath = urlPath.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector'])));

        const response = await this.request({
            path: urlPath,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartReleaseV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual ChartRelease by its ID.
     * Delete an individual ChartRelease
     */
    async apiChartReleasesV3SelectorDelete(requestParameters: ApiChartReleasesV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartReleaseV3> {
        const response = await this.apiChartReleasesV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual ChartRelease.
     * Get an individual ChartRelease
     */
    async apiChartReleasesV3SelectorGetRaw(requestParameters: ApiChartReleasesV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartReleaseV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiChartReleasesV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/chart-releases/v3/{selector}`;
        urlPath = urlPath.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector'])));

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartReleaseV3FromJSON(jsonValue));
    }

    /**
     * Get an individual ChartRelease.
     * Get an individual ChartRelease
     */
    async apiChartReleasesV3SelectorGet(requestParameters: ApiChartReleasesV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartReleaseV3> {
        const response = await this.apiChartReleasesV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual ChartRelease.
     * Edit an individual ChartRelease
     */
    async apiChartReleasesV3SelectorPatchRaw(requestParameters: ApiChartReleasesV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartReleaseV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiChartReleasesV3SelectorPatch().'
            );
        }

        if (requestParameters['chartRelease'] == null) {
            throw new runtime.RequiredError(
                'chartRelease',
                'Required parameter "chartRelease" was null or undefined when calling apiChartReleasesV3SelectorPatch().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';


        let urlPath = `/api/chart-releases/v3/{selector}`;
        urlPath = urlPath.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector'])));

        const response = await this.request({
            path: urlPath,
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartReleaseV3EditToJSON(requestParameters['chartRelease']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartReleaseV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual ChartRelease.
     * Edit an individual ChartRelease
     */
    async apiChartReleasesV3SelectorPatch(requestParameters: ApiChartReleasesV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartReleaseV3> {
        const response = await this.apiChartReleasesV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const ApiChartReleasesV3GetAppVersionResolverEnum = {
    Branch: 'branch',
    Commit: 'commit',
    Exact: 'exact',
    Follow: 'follow',
    None: 'none'
} as const;
export type ApiChartReleasesV3GetAppVersionResolverEnum = typeof ApiChartReleasesV3GetAppVersionResolverEnum[keyof typeof ApiChartReleasesV3GetAppVersionResolverEnum];
/**
 * @export
 */
export const ApiChartReleasesV3GetChartVersionResolverEnum = {
    Latest: 'latest',
    Exact: 'exact',
    Follow: 'follow'
} as const;
export type ApiChartReleasesV3GetChartVersionResolverEnum = typeof ApiChartReleasesV3GetChartVersionResolverEnum[keyof typeof ApiChartReleasesV3GetChartVersionResolverEnum];
