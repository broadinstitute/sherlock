/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service
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
  V2controllersChartDeployRecord,
  V2controllersCreatableChartDeployRecord,
} from '../models';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    V2controllersChartDeployRecordFromJSON,
    V2controllersChartDeployRecordToJSON,
    V2controllersCreatableChartDeployRecordFromJSON,
    V2controllersCreatableChartDeployRecordToJSON,
} from '../models';

export interface ApiV2ChartDeployRecordsGetRequest {
    chartRelease?: string;
    createdAt?: string;
    exactAppVersion?: string;
    exactChartVersion?: string;
    helmfileRef?: string;
    id?: number;
    updatedAt?: string;
    limit?: number;
}

export interface ApiV2ChartDeployRecordsPostRequest {
    chartDeployRecord: V2controllersCreatableChartDeployRecord;
}

export interface ApiV2ChartDeployRecordsSelectorGetRequest {
    selector: string;
}

export interface ApiV2SelectorsChartDeployRecordsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class ChartDeployRecordsApi extends runtime.BaseAPI {

    /**
     * List existing ChartDeployRecord entries, ordered by most recently updated.
     * List ChartDeployRecord entries
     */
    async apiV2ChartDeployRecordsGetRaw(requestParameters: ApiV2ChartDeployRecordsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChartDeployRecord>>> {
        const queryParameters: any = {};

        if (requestParameters.chartRelease !== undefined) {
            queryParameters['chartRelease'] = requestParameters.chartRelease;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = requestParameters.createdAt;
        }

        if (requestParameters.exactAppVersion !== undefined) {
            queryParameters['exactAppVersion'] = requestParameters.exactAppVersion;
        }

        if (requestParameters.exactChartVersion !== undefined) {
            queryParameters['exactChartVersion'] = requestParameters.exactChartVersion;
        }

        if (requestParameters.helmfileRef !== undefined) {
            queryParameters['helmfileRef'] = requestParameters.helmfileRef;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = requestParameters.updatedAt;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/chart-deploy-records`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChartDeployRecordFromJSON));
    }

    /**
     * List existing ChartDeployRecord entries, ordered by most recently updated.
     * List ChartDeployRecord entries
     */
    async apiV2ChartDeployRecordsGet(requestParameters: ApiV2ChartDeployRecordsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChartDeployRecord>> {
        const response = await this.apiV2ChartDeployRecordsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new ChartDeployRecord entry. Note that fields are immutable after creation.
     * Create a new ChartDeployRecord entry
     */
    async apiV2ChartDeployRecordsPostRaw(requestParameters: ApiV2ChartDeployRecordsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartDeployRecord>> {
        if (requestParameters.chartDeployRecord === null || requestParameters.chartDeployRecord === undefined) {
            throw new runtime.RequiredError('chartDeployRecord','Required parameter requestParameters.chartDeployRecord was null or undefined when calling apiV2ChartDeployRecordsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/chart-deploy-records`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChartDeployRecordToJSON(requestParameters.chartDeployRecord),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartDeployRecordFromJSON(jsonValue));
    }

    /**
     * Create a new ChartDeployRecord entry. Note that fields are immutable after creation.
     * Create a new ChartDeployRecord entry
     */
    async apiV2ChartDeployRecordsPost(requestParameters: ApiV2ChartDeployRecordsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartDeployRecord> {
        const response = await this.apiV2ChartDeployRecordsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing ChartDeployRecord entry via one its \"selector\"--its numeric ID.
     * Get a ChartDeployRecord entry
     */
    async apiV2ChartDeployRecordsSelectorGetRaw(requestParameters: ApiV2ChartDeployRecordsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChartDeployRecord>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChartDeployRecordsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/chart-deploy-records/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChartDeployRecordFromJSON(jsonValue));
    }

    /**
     * Get an existing ChartDeployRecord entry via one its \"selector\"--its numeric ID.
     * Get a ChartDeployRecord entry
     */
    async apiV2ChartDeployRecordsSelectorGet(requestParameters: ApiV2ChartDeployRecordsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChartDeployRecord> {
        const response = await this.apiV2ChartDeployRecordsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given ChartDeployRecord selector and provide any other selectors that would match the same ChartDeployRecord.
     * List ChartDeployRecord selectors
     */
    async apiV2SelectorsChartDeployRecordsSelectorGetRaw(requestParameters: ApiV2SelectorsChartDeployRecordsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsChartDeployRecordsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/chart-deploy-records/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given ChartDeployRecord selector and provide any other selectors that would match the same ChartDeployRecord.
     * List ChartDeployRecord selectors
     */
    async apiV2SelectorsChartDeployRecordsSelectorGet(requestParameters: ApiV2SelectorsChartDeployRecordsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsChartDeployRecordsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
