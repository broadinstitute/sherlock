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
  SherlockChartV3,
  SherlockChartV3Create,
  SherlockChartV3Edit,
  V2controllersChart,
  V2controllersCreatableChart,
  V2controllersEditableChart,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockChartV3FromJSON,
    SherlockChartV3ToJSON,
    SherlockChartV3CreateFromJSON,
    SherlockChartV3CreateToJSON,
    SherlockChartV3EditFromJSON,
    SherlockChartV3EditToJSON,
    V2controllersChartFromJSON,
    V2controllersChartToJSON,
    V2controllersCreatableChartFromJSON,
    V2controllersCreatableChartToJSON,
    V2controllersEditableChartFromJSON,
    V2controllersEditableChartToJSON,
} from '../models/index';

export interface ApiChartsV3GetRequest {
    appImageGitMainBranch?: string;
    appImageGitRepo?: string;
    chartExposesEndpoint?: boolean;
    chartRepo?: string;
    createdAt?: Date;
    defaultPort?: number;
    defaultProtocol?: string;
    defaultSubdomain?: string;
    description?: string;
    id?: number;
    name?: string;
    pactParticipant?: boolean;
    playbookURL?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiChartsV3PostRequest {
    chart: SherlockChartV3Create;
}

export interface ApiChartsV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiChartsV3SelectorGetRequest {
    selector: string;
}

export interface ApiChartsV3SelectorPatchRequest {
    selector: string;
    chart: SherlockChartV3Edit;
}

export interface ApiV2ChartsGetRequest {
    appImageGitMainBranch?: string;
    appImageGitRepo?: string;
    chartExposesEndpoint?: boolean;
    chartRepo?: string;
    createdAt?: Date;
    defaultPort?: number;
    defaultProtocol?: string;
    defaultSubdomain?: string;
    description?: string;
    id?: number;
    legacyConfigsEnabled?: boolean;
    name?: string;
    pactParticipant?: boolean;
    playbookURL?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2ChartsPostRequest {
    chart: V2controllersCreatableChart;
}

export interface ApiV2ChartsSelectorDeleteRequest {
    selector: string;
}

export interface ApiV2ChartsSelectorGetRequest {
    selector: string;
}

export interface ApiV2ChartsSelectorPatchRequest {
    selector: string;
    chart: V2controllersEditableChart;
}

export interface ApiV2ChartsSelectorPutRequest {
    selector: string;
    chart: V2controllersCreatableChart;
}

export interface ApiV2SelectorsChartsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class ChartsApi extends runtime.BaseAPI {

    /**
     * List Charts matching a filter.
     * List Charts matching a filter
     */
    async apiChartsV3GetRaw(requestParameters: ApiChartsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockChartV3>>> {
        const queryParameters: any = {};

        if (requestParameters.appImageGitMainBranch !== undefined) {
            queryParameters['appImageGitMainBranch'] = requestParameters.appImageGitMainBranch;
        }

        if (requestParameters.appImageGitRepo !== undefined) {
            queryParameters['appImageGitRepo'] = requestParameters.appImageGitRepo;
        }

        if (requestParameters.chartExposesEndpoint !== undefined) {
            queryParameters['chartExposesEndpoint'] = requestParameters.chartExposesEndpoint;
        }

        if (requestParameters.chartRepo !== undefined) {
            queryParameters['chartRepo'] = requestParameters.chartRepo;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.defaultPort !== undefined) {
            queryParameters['defaultPort'] = requestParameters.defaultPort;
        }

        if (requestParameters.defaultProtocol !== undefined) {
            queryParameters['defaultProtocol'] = requestParameters.defaultProtocol;
        }

        if (requestParameters.defaultSubdomain !== undefined) {
            queryParameters['defaultSubdomain'] = requestParameters.defaultSubdomain;
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.pactParticipant !== undefined) {
            queryParameters['pactParticipant'] = requestParameters.pactParticipant;
        }

        if (requestParameters.playbookURL !== undefined) {
            queryParameters['playbookURL'] = requestParameters.playbookURL;
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
            path: `/api/charts/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockChartV3FromJSON));
    }

    /**
     * List Charts matching a filter.
     * List Charts matching a filter
     */
    async apiChartsV3Get(requestParameters: ApiChartsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockChartV3>> {
        const response = await this.apiChartsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a Chart.
     * Create a Chart
     */
    async apiChartsV3PostRaw(requestParameters: ApiChartsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartV3>> {
        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiChartsV3Post.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/charts/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartV3CreateToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartV3FromJSON(jsonValue));
    }

    /**
     * Create a Chart.
     * Create a Chart
     */
    async apiChartsV3Post(requestParameters: ApiChartsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartV3> {
        const response = await this.apiChartsV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual Chart by its ID.
     * Delete an individual Chart
     */
    async apiChartsV3SelectorDeleteRaw(requestParameters: ApiChartsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiChartsV3SelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/charts/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual Chart by its ID.
     * Delete an individual Chart
     */
    async apiChartsV3SelectorDelete(requestParameters: ApiChartsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartV3> {
        const response = await this.apiChartsV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual Chart.
     * Get an individual Chart
     */
    async apiChartsV3SelectorGetRaw(requestParameters: ApiChartsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiChartsV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/charts/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartV3FromJSON(jsonValue));
    }

    /**
     * Get an individual Chart.
     * Get an individual Chart
     */
    async apiChartsV3SelectorGet(requestParameters: ApiChartsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartV3> {
        const response = await this.apiChartsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual Chart.
     * Edit an individual Chart
     */
    async apiChartsV3SelectorPatchRaw(requestParameters: ApiChartsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockChartV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiChartsV3SelectorPatch.');
        }

        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiChartsV3SelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/charts/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockChartV3EditToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockChartV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual Chart.
     * Edit an individual Chart
     */
    async apiChartsV3SelectorPatch(requestParameters: ApiChartsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockChartV3> {
        const response = await this.apiChartsV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing Chart entries, ordered by most recently updated.
     * List Chart entries
     */
    async apiV2ChartsGetRaw(requestParameters: ApiV2ChartsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChart>>> {
        const queryParameters: any = {};

        if (requestParameters.appImageGitMainBranch !== undefined) {
            queryParameters['appImageGitMainBranch'] = requestParameters.appImageGitMainBranch;
        }

        if (requestParameters.appImageGitRepo !== undefined) {
            queryParameters['appImageGitRepo'] = requestParameters.appImageGitRepo;
        }

        if (requestParameters.chartExposesEndpoint !== undefined) {
            queryParameters['chartExposesEndpoint'] = requestParameters.chartExposesEndpoint;
        }

        if (requestParameters.chartRepo !== undefined) {
            queryParameters['chartRepo'] = requestParameters.chartRepo;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.defaultPort !== undefined) {
            queryParameters['defaultPort'] = requestParameters.defaultPort;
        }

        if (requestParameters.defaultProtocol !== undefined) {
            queryParameters['defaultProtocol'] = requestParameters.defaultProtocol;
        }

        if (requestParameters.defaultSubdomain !== undefined) {
            queryParameters['defaultSubdomain'] = requestParameters.defaultSubdomain;
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.legacyConfigsEnabled !== undefined) {
            queryParameters['legacyConfigsEnabled'] = requestParameters.legacyConfigsEnabled;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.pactParticipant !== undefined) {
            queryParameters['pactParticipant'] = requestParameters.pactParticipant;
        }

        if (requestParameters.playbookURL !== undefined) {
            queryParameters['playbookURL'] = requestParameters.playbookURL;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/charts`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChartFromJSON));
    }

    /**
     * List existing Chart entries, ordered by most recently updated.
     * List Chart entries
     */
    async apiV2ChartsGet(requestParameters: ApiV2ChartsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChart>> {
        const response = await this.apiV2ChartsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new Chart entry
     */
    async apiV2ChartsPostRaw(requestParameters: ApiV2ChartsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChart>> {
        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiV2ChartsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/charts`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChartToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartFromJSON(jsonValue));
    }

    /**
     * Create a new Chart entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new Chart entry
     */
    async apiV2ChartsPost(requestParameters: ApiV2ChartsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChart> {
        const response = await this.apiV2ChartsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an existing Chart entry via one of its \"selectors\": name or numeric ID.
     * Delete a Chart entry
     */
    async apiV2ChartsSelectorDeleteRaw(requestParameters: ApiV2ChartsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChart>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartsSelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/charts/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartFromJSON(jsonValue));
    }

    /**
     * Delete an existing Chart entry via one of its \"selectors\": name or numeric ID.
     * Delete a Chart entry
     */
    async apiV2ChartsSelectorDelete(requestParameters: ApiV2ChartsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChart> {
        const response = await this.apiV2ChartsSelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing Chart entry via one of its \"selectors\": name or numeric ID.
     * Get a Chart entry
     */
    async apiV2ChartsSelectorGetRaw(requestParameters: ApiV2ChartsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChart>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/charts/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartFromJSON(jsonValue));
    }

    /**
     * Get an existing Chart entry via one of its \"selectors\": name or numeric ID.
     * Get a Chart entry
     */
    async apiV2ChartsSelectorGet(requestParameters: ApiV2ChartsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChart> {
        const response = await this.apiV2ChartsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing Chart entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a Chart entry
     */
    async apiV2ChartsSelectorPatchRaw(requestParameters: ApiV2ChartsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChart>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartsSelectorPatch.');
        }

        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiV2ChartsSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/charts/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableChartToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartFromJSON(jsonValue));
    }

    /**
     * Edit an existing Chart entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a Chart entry
     */
    async apiV2ChartsSelectorPatch(requestParameters: ApiV2ChartsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChart> {
        const response = await this.apiV2ChartsSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit a Chart entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a Chart entry
     */
    async apiV2ChartsSelectorPutRaw(requestParameters: ApiV2ChartsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChart>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartsSelectorPut.');
        }

        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiV2ChartsSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/charts/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChartToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartFromJSON(jsonValue));
    }

    /**
     * Create or edit a Chart entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a Chart entry
     */
    async apiV2ChartsSelectorPut(requestParameters: ApiV2ChartsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChart> {
        const response = await this.apiV2ChartsSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given Chart selector and provide any other selectors that would match the same Chart.
     * List Chart selectors
     */
    async apiV2SelectorsChartsSelectorGetRaw(requestParameters: ApiV2SelectorsChartsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsChartsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/charts/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given Chart selector and provide any other selectors that would match the same Chart.
     * List Chart selectors
     */
    async apiV2SelectorsChartsSelectorGet(requestParameters: ApiV2SelectorsChartsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsChartsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
