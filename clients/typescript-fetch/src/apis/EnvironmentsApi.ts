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
  PagerdutyAlertSummary,
  PagerdutySendAlertResponse,
  V2controllersCreatableEnvironment,
  V2controllersEditableEnvironment,
  V2controllersEnvironment,
} from '../models';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    PagerdutyAlertSummaryFromJSON,
    PagerdutyAlertSummaryToJSON,
    PagerdutySendAlertResponseFromJSON,
    PagerdutySendAlertResponseToJSON,
    V2controllersCreatableEnvironmentFromJSON,
    V2controllersCreatableEnvironmentToJSON,
    V2controllersEditableEnvironmentFromJSON,
    V2controllersEditableEnvironmentToJSON,
    V2controllersEnvironmentFromJSON,
    V2controllersEnvironmentToJSON,
} from '../models';

export interface ApiV2EnvironmentsGetRequest {
    autoPopulateChartReleases?: boolean;
    base?: string;
    baseDomain?: string;
    createdAt?: Date;
    defaultCluster?: string;
    defaultFirecloudDevelopRef?: string;
    defaultNamespace?: string;
    description?: string;
    helmfileRef?: string;
    id?: number;
    lifecycle?: string;
    name?: string;
    namePrefix?: string;
    namePrefixesDomain?: boolean;
    offline?: boolean;
    offlineScheduleBeginEnabled?: boolean;
    offlineScheduleBeginTime?: string;
    offlineScheduleEndEnabled?: boolean;
    offlineScheduleEndTime?: string;
    offlineScheduleEndWeekends?: boolean;
    owner?: string;
    pagerdutyIntegration?: string;
    preventDeletion?: boolean;
    requiresSuitability?: boolean;
    templateEnvironment?: string;
    uniqueResourcePrefix?: string;
    updatedAt?: Date;
    valuesName?: string;
    limit?: number;
}

export interface ApiV2EnvironmentsPostRequest {
    environment: V2controllersCreatableEnvironment;
}

export interface ApiV2EnvironmentsSelectorDeleteRequest {
    selector: string;
}

export interface ApiV2EnvironmentsSelectorGetRequest {
    selector: string;
}

export interface ApiV2EnvironmentsSelectorPatchRequest {
    selector: string;
    environment: V2controllersEditableEnvironment;
}

export interface ApiV2EnvironmentsSelectorPutRequest {
    selector: string;
    environment: V2controllersCreatableEnvironment;
}

export interface ApiV2ProceduresEnvironmentsTriggerIncidentSelectorPostRequest {
    selector: string;
    summary: PagerdutyAlertSummary;
}

export interface ApiV2SelectorsEnvironmentsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class EnvironmentsApi extends runtime.BaseAPI {

    /**
     * List existing Environment entries, ordered by most recently updated.
     * List Environment entries
     */
    async apiV2EnvironmentsGetRaw(requestParameters: ApiV2EnvironmentsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersEnvironment>>> {
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

        if (requestParameters.defaultFirecloudDevelopRef !== undefined) {
            queryParameters['defaultFirecloudDevelopRef'] = requestParameters.defaultFirecloudDevelopRef;
        }

        if (requestParameters.defaultNamespace !== undefined) {
            queryParameters['defaultNamespace'] = requestParameters.defaultNamespace;
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

        if (requestParameters.namePrefix !== undefined) {
            queryParameters['namePrefix'] = requestParameters.namePrefix;
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
            queryParameters['offlineScheduleBeginTime'] = requestParameters.offlineScheduleBeginTime;
        }

        if (requestParameters.offlineScheduleEndEnabled !== undefined) {
            queryParameters['offlineScheduleEndEnabled'] = requestParameters.offlineScheduleEndEnabled;
        }

        if (requestParameters.offlineScheduleEndTime !== undefined) {
            queryParameters['offlineScheduleEndTime'] = requestParameters.offlineScheduleEndTime;
        }

        if (requestParameters.offlineScheduleEndWeekends !== undefined) {
            queryParameters['offlineScheduleEndWeekends'] = requestParameters.offlineScheduleEndWeekends;
        }

        if (requestParameters.owner !== undefined) {
            queryParameters['owner'] = requestParameters.owner;
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

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/environments`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersEnvironmentFromJSON));
    }

    /**
     * List existing Environment entries, ordered by most recently updated.
     * List Environment entries
     */
    async apiV2EnvironmentsGet(requestParameters: ApiV2EnvironmentsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersEnvironment>> {
        const response = await this.apiV2EnvironmentsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new Environment entry. Note that some fields are immutable after creation; /edit lists mutable fields. Creating a dynamic environment based on a template will also copy ChartReleases from the template.
     * Create a new Environment entry
     */
    async apiV2EnvironmentsPostRaw(requestParameters: ApiV2EnvironmentsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersEnvironment>> {
        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling apiV2EnvironmentsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/environments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableEnvironmentToJSON(requestParameters.environment),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersEnvironmentFromJSON(jsonValue));
    }

    /**
     * Create a new Environment entry. Note that some fields are immutable after creation; /edit lists mutable fields. Creating a dynamic environment based on a template will also copy ChartReleases from the template.
     * Create a new Environment entry
     */
    async apiV2EnvironmentsPost(requestParameters: ApiV2EnvironmentsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersEnvironment> {
        const response = await this.apiV2EnvironmentsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an existing Environment entry via one of its \"selectors\": name, numeric ID, or \"resource-prefix/\" + the unique resource prefix.
     * Delete a Environment entry
     */
    async apiV2EnvironmentsSelectorDeleteRaw(requestParameters: ApiV2EnvironmentsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersEnvironment>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2EnvironmentsSelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/environments/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersEnvironmentFromJSON(jsonValue));
    }

    /**
     * Delete an existing Environment entry via one of its \"selectors\": name, numeric ID, or \"resource-prefix/\" + the unique resource prefix.
     * Delete a Environment entry
     */
    async apiV2EnvironmentsSelectorDelete(requestParameters: ApiV2EnvironmentsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersEnvironment> {
        const response = await this.apiV2EnvironmentsSelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing Environment entry via one of its \"selectors\": name, numeric ID, or \"resource-prefix/\" + the unique resource prefix.
     * Get a Environment entry
     */
    async apiV2EnvironmentsSelectorGetRaw(requestParameters: ApiV2EnvironmentsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersEnvironment>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2EnvironmentsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/environments/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersEnvironmentFromJSON(jsonValue));
    }

    /**
     * Get an existing Environment entry via one of its \"selectors\": name, numeric ID, or \"resource-prefix/\" + the unique resource prefix.
     * Get a Environment entry
     */
    async apiV2EnvironmentsSelectorGet(requestParameters: ApiV2EnvironmentsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersEnvironment> {
        const response = await this.apiV2EnvironmentsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing Environment entry via one of its \"selectors\": name, numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create, or \"resource-prefix/\" + the unique resource prefix.
     * Edit a Environment entry
     */
    async apiV2EnvironmentsSelectorPatchRaw(requestParameters: ApiV2EnvironmentsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersEnvironment>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2EnvironmentsSelectorPatch.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling apiV2EnvironmentsSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/environments/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableEnvironmentToJSON(requestParameters.environment),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersEnvironmentFromJSON(jsonValue));
    }

    /**
     * Edit an existing Environment entry via one of its \"selectors\": name, numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create, or \"resource-prefix/\" + the unique resource prefix.
     * Edit a Environment entry
     */
    async apiV2EnvironmentsSelectorPatch(requestParameters: ApiV2EnvironmentsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersEnvironment> {
        const response = await this.apiV2EnvironmentsSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit an Environment entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit an Environment entry
     */
    async apiV2EnvironmentsSelectorPutRaw(requestParameters: ApiV2EnvironmentsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersEnvironment>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2EnvironmentsSelectorPut.');
        }

        if (requestParameters.environment === null || requestParameters.environment === undefined) {
            throw new runtime.RequiredError('environment','Required parameter requestParameters.environment was null or undefined when calling apiV2EnvironmentsSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/environments/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableEnvironmentToJSON(requestParameters.environment),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersEnvironmentFromJSON(jsonValue));
    }

    /**
     * Create or edit an Environment entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit an Environment entry
     */
    async apiV2EnvironmentsSelectorPut(requestParameters: ApiV2EnvironmentsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersEnvironment> {
        const response = await this.apiV2EnvironmentsSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Trigger an alert for the Pagerduty integration configured for a given Environment.
     * Trigger a Pagerduty incident for a given Environment
     */
    async apiV2ProceduresEnvironmentsTriggerIncidentSelectorPostRaw(requestParameters: ApiV2ProceduresEnvironmentsTriggerIncidentSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PagerdutySendAlertResponse>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ProceduresEnvironmentsTriggerIncidentSelectorPost.');
        }

        if (requestParameters.summary === null || requestParameters.summary === undefined) {
            throw new runtime.RequiredError('summary','Required parameter requestParameters.summary was null or undefined when calling apiV2ProceduresEnvironmentsTriggerIncidentSelectorPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/procedures/environments/trigger-incident/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PagerdutyAlertSummaryToJSON(requestParameters.summary),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PagerdutySendAlertResponseFromJSON(jsonValue));
    }

    /**
     * Trigger an alert for the Pagerduty integration configured for a given Environment.
     * Trigger a Pagerduty incident for a given Environment
     */
    async apiV2ProceduresEnvironmentsTriggerIncidentSelectorPost(requestParameters: ApiV2ProceduresEnvironmentsTriggerIncidentSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PagerdutySendAlertResponse> {
        const response = await this.apiV2ProceduresEnvironmentsTriggerIncidentSelectorPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given Environment selector and provide any other selectors that would match the same Environment.
     * List Environment selectors
     */
    async apiV2SelectorsEnvironmentsSelectorGetRaw(requestParameters: ApiV2SelectorsEnvironmentsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsEnvironmentsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/environments/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given Environment selector and provide any other selectors that would match the same Environment.
     * List Environment selectors
     */
    async apiV2SelectorsEnvironmentsSelectorGet(requestParameters: ApiV2SelectorsEnvironmentsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsEnvironmentsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
