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
  SherlockEnvironmentV3,
  SherlockEnvironmentV3Create,
  SherlockEnvironmentV3Edit,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockEnvironmentV3FromJSON,
    SherlockEnvironmentV3ToJSON,
    SherlockEnvironmentV3CreateFromJSON,
    SherlockEnvironmentV3CreateToJSON,
    SherlockEnvironmentV3EditFromJSON,
    SherlockEnvironmentV3EditToJSON,
} from '../models/index';

export interface ApiEnvironmentsV3GetRequest {
    autoPopulateChartReleases?: boolean;
    base?: string;
    baseDomain?: string;
    createdAt?: Date;
    defaultCluster?: string;
    defaultNamespace?: string;
    deleteAfter?: Date;
    description?: string;
    helmfileRef?: string;
    id?: number;
    lifecycle?: string;
    name?: string;
    namePrefixesDomain?: boolean;
    offline?: boolean;
    offlineScheduleBeginEnabled?: boolean;
    offlineScheduleBeginTime?: Date;
    offlineScheduleEndEnabled?: boolean;
    offlineScheduleEndTime?: Date;
    offlineScheduleEndWeekends?: boolean;
    owner?: string;
    pactIdentifier?: string;
    pagerdutyIntegration?: string;
    preventDeletion?: boolean;
    requiresSuitability?: boolean;
    templateEnvironment?: string;
    uniqueResourcePrefix?: string;
    updatedAt?: Date;
    valuesName?: string;
    limit?: number;
    offset?: number;
}

export interface ApiEnvironmentsV3PostRequest {
    environment: SherlockEnvironmentV3Create;
}

export interface ApiEnvironmentsV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiEnvironmentsV3SelectorGetRequest {
    selector: string;
}

export interface ApiEnvironmentsV3SelectorPatchRequest {
    selector: string;
    environment: SherlockEnvironmentV3Edit;
}

/**
 * 
 */
export class EnvironmentsApi extends runtime.BaseAPI {

    /**
     * List Environments matching a filter.
     * List Environments matching a filter
     */
    async apiEnvironmentsV3GetRaw(requestParameters: ApiEnvironmentsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockEnvironmentV3>>> {
        const queryParameters: any = {};

        if (requestParameters.autoPopulateChartReleases !== undefined) {
            queryParameters['autoPopulateChartReleases'] = requestParameters.autoPopulateChartReleases;
        }

        if (requestParameters.base !== undefined) {
            queryParameters['base'] = requestParameters.base;
        }

        if (requestParameters.baseDomain !== undefined) {
            queryParameters['baseDomain'] = requestParameters.baseDomain;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.defaultCluster !== undefined) {
            queryParameters['defaultCluster'] = requestParameters.defaultCluster;
        }

        if (requestParameters.defaultNamespace !== undefined) {
            queryParameters['defaultNamespace'] = requestParameters.defaultNamespace;
        }

        if (requestParameters.deleteAfter !== undefined) {
            queryParameters['deleteAfter'] = (requestParameters.deleteAfter as any).toISOString();
        }

        if (requestParameters.description !== undefined) {
            queryParameters['description'] = requestParameters.description;
        }

        if (requestParameters.helmfileRef !== undefined) {
            queryParameters['helmfileRef'] = requestParameters.helmfileRef;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.lifecycle !== undefined) {
            queryParameters['lifecycle'] = requestParameters.lifecycle;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.namePrefixesDomain !== undefined) {
            queryParameters['namePrefixesDomain'] = requestParameters.namePrefixesDomain;
        }

        if (requestParameters.offline !== undefined) {
            queryParameters['offline'] = requestParameters.offline;
        }

        if (requestParameters.offlineScheduleBeginEnabled !== undefined) {
            queryParameters['offlineScheduleBeginEnabled'] = requestParameters.offlineScheduleBeginEnabled;
        }

        if (requestParameters.offlineScheduleBeginTime !== undefined) {
            queryParameters['offlineScheduleBeginTime'] = (requestParameters.offlineScheduleBeginTime as any).toISOString();
        }

        if (requestParameters.offlineScheduleEndEnabled !== undefined) {
            queryParameters['offlineScheduleEndEnabled'] = requestParameters.offlineScheduleEndEnabled;
        }

        if (requestParameters.offlineScheduleEndTime !== undefined) {
            queryParameters['offlineScheduleEndTime'] = (requestParameters.offlineScheduleEndTime as any).toISOString();
        }

        if (requestParameters.offlineScheduleEndWeekends !== undefined) {
            queryParameters['offlineScheduleEndWeekends'] = requestParameters.offlineScheduleEndWeekends;
        }

        if (requestParameters.owner !== undefined) {
            queryParameters['owner'] = requestParameters.owner;
        }

        if (requestParameters.pactIdentifier !== undefined) {
            queryParameters['PactIdentifier'] = requestParameters.pactIdentifier;
        }

        if (requestParameters.pagerdutyIntegration !== undefined) {
            queryParameters['pagerdutyIntegration'] = requestParameters.pagerdutyIntegration;
        }

        if (requestParameters.preventDeletion !== undefined) {
            queryParameters['preventDeletion'] = requestParameters.preventDeletion;
        }

        if (requestParameters.requiresSuitability !== undefined) {
            queryParameters['requiresSuitability'] = requestParameters.requiresSuitability;
        }

        if (requestParameters.templateEnvironment !== undefined) {
            queryParameters['templateEnvironment'] = requestParameters.templateEnvironment;
        }

        if (requestParameters.uniqueResourcePrefix !== undefined) {
            queryParameters['uniqueResourcePrefix'] = requestParameters.uniqueResourcePrefix;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.valuesName !== undefined) {
            queryParameters['valuesName'] = requestParameters.valuesName;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/environments/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockEnvironmentV3FromJSON));
    }

    /**
     * List Environments matching a filter.
     * List Environments matching a filter
     */
    async apiEnvironmentsV3Get(requestParameters: ApiEnvironmentsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockEnvironmentV3>> {
        const response = await this.apiEnvironmentsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a Environment.
     * Create a Environment
     */
    async apiEnvironmentsV3PostRaw(requestParameters: ApiEnvironmentsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockEnvironmentV3>> {
        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling apiEnvironmentsV3Post.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/environments/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockEnvironmentV3CreateToJSON(requestParameters.environment),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockEnvironmentV3FromJSON(jsonValue));
    }

    /**
     * Create a Environment.
     * Create a Environment
     */
    async apiEnvironmentsV3Post(requestParameters: ApiEnvironmentsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockEnvironmentV3> {
        const response = await this.apiEnvironmentsV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual Environment by its ID.
     * Delete an individual Environment
     */
    async apiEnvironmentsV3SelectorDeleteRaw(requestParameters: ApiEnvironmentsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockEnvironmentV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiEnvironmentsV3SelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/environments/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockEnvironmentV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual Environment by its ID.
     * Delete an individual Environment
     */
    async apiEnvironmentsV3SelectorDelete(requestParameters: ApiEnvironmentsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockEnvironmentV3> {
        const response = await this.apiEnvironmentsV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual Environment.
     * Get an individual Environment
     */
    async apiEnvironmentsV3SelectorGetRaw(requestParameters: ApiEnvironmentsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockEnvironmentV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiEnvironmentsV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/environments/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockEnvironmentV3FromJSON(jsonValue));
    }

    /**
     * Get an individual Environment.
     * Get an individual Environment
     */
    async apiEnvironmentsV3SelectorGet(requestParameters: ApiEnvironmentsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockEnvironmentV3> {
        const response = await this.apiEnvironmentsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual Environment.
     * Edit an individual Environment
     */
    async apiEnvironmentsV3SelectorPatchRaw(requestParameters: ApiEnvironmentsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockEnvironmentV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiEnvironmentsV3SelectorPatch.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling apiEnvironmentsV3SelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/environments/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockEnvironmentV3EditToJSON(requestParameters.environment),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockEnvironmentV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual Environment.
     * Edit an individual Environment
     */
    async apiEnvironmentsV3SelectorPatch(requestParameters: ApiEnvironmentsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockEnvironmentV3> {
        const response = await this.apiEnvironmentsV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
