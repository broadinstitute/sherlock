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
  V2controllersCreatableEnvironment,
  V2controllersEditableEnvironment,
  V2controllersEnvironment,
} from '../models';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    V2controllersCreatableEnvironmentFromJSON,
    V2controllersCreatableEnvironmentToJSON,
    V2controllersEditableEnvironmentFromJSON,
    V2controllersEditableEnvironmentToJSON,
    V2controllersEnvironmentFromJSON,
    V2controllersEnvironmentToJSON,
} from '../models';

export interface ApiV2EnvironmentsGetRequest {
    base?: string;
    baseDomain?: string;
    chartReleasesFromTemplate?: boolean;
    createdAt?: Date;
    defaultCluster?: string;
    defaultFirecloudDevelopRef?: string;
    defaultNamespace?: string;
    helmfileRef?: string;
    id?: number;
    lifecycle?: string;
    name?: string;
    namePrefix?: string;
    namePrefixesDomain?: boolean;
    owner?: string;
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

        if (requestParameters.base !== undefined) {
            queryParameters['base'] = requestParameters.base;
        }

        if (requestParameters.baseDomain !== undefined) {
            queryParameters['baseDomain'] = requestParameters.baseDomain;
        }

        if (requestParameters.chartReleasesFromTemplate !== undefined) {
            queryParameters['chartReleasesFromTemplate'] = requestParameters.chartReleasesFromTemplate;
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

        if (requestParameters.owner !== undefined) {
            queryParameters['owner'] = requestParameters.owner;
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
