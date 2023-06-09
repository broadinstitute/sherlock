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
  V2controllersCreatableDatabaseInstance,
  V2controllersDatabaseInstance,
  V2controllersEditableDatabaseInstance,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    V2controllersCreatableDatabaseInstanceFromJSON,
    V2controllersCreatableDatabaseInstanceToJSON,
    V2controllersDatabaseInstanceFromJSON,
    V2controllersDatabaseInstanceToJSON,
    V2controllersEditableDatabaseInstanceFromJSON,
    V2controllersEditableDatabaseInstanceToJSON,
} from '../models/index';

export interface ApiV2DatabaseInstancesGetRequest {
    chartRelease?: string;
    createdAt?: Date;
    defaultDatabase?: string;
    googleProject?: string;
    id?: number;
    instanceName?: string;
    platform?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2DatabaseInstancesPostRequest {
    chart: V2controllersCreatableDatabaseInstance;
}

export interface ApiV2DatabaseInstancesSelectorDeleteRequest {
    selector: string;
}

export interface ApiV2DatabaseInstancesSelectorGetRequest {
    selector: string;
}

export interface ApiV2DatabaseInstancesSelectorPatchRequest {
    selector: string;
    chart: V2controllersEditableDatabaseInstance;
}

export interface ApiV2DatabaseInstancesSelectorPutRequest {
    selector: string;
    databaseInstance: V2controllersCreatableDatabaseInstance;
}

export interface ApiV2SelectorsDatabaseInstancesSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class DatabaseInstancesApi extends runtime.BaseAPI {

    /**
     * List existing DatabaseInstance entries, ordered by most recently updated.
     * List DatabaseInstance entries
     */
    async apiV2DatabaseInstancesGetRaw(requestParameters: ApiV2DatabaseInstancesGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersDatabaseInstance>>> {
        const queryParameters: any = {};

        if (requestParameters.chartRelease !== undefined) {
            queryParameters['chartRelease'] = requestParameters.chartRelease;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.defaultDatabase !== undefined) {
            queryParameters['defaultDatabase'] = requestParameters.defaultDatabase;
        }

        if (requestParameters.googleProject !== undefined) {
            queryParameters['googleProject'] = requestParameters.googleProject;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.instanceName !== undefined) {
            queryParameters['instanceName'] = requestParameters.instanceName;
        }

        if (requestParameters.platform !== undefined) {
            queryParameters['platform'] = requestParameters.platform;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/database-instances`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersDatabaseInstanceFromJSON));
    }

    /**
     * List existing DatabaseInstance entries, ordered by most recently updated.
     * List DatabaseInstance entries
     */
    async apiV2DatabaseInstancesGet(requestParameters: ApiV2DatabaseInstancesGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersDatabaseInstance>> {
        const response = await this.apiV2DatabaseInstancesGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new DatabaseInstance entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new DatabaseInstance entry
     */
    async apiV2DatabaseInstancesPostRaw(requestParameters: ApiV2DatabaseInstancesPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersDatabaseInstance>> {
        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiV2DatabaseInstancesPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/database-instances`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableDatabaseInstanceToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersDatabaseInstanceFromJSON(jsonValue));
    }

    /**
     * Create a new DatabaseInstance entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new DatabaseInstance entry
     */
    async apiV2DatabaseInstancesPost(requestParameters: ApiV2DatabaseInstancesPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersDatabaseInstance> {
        const response = await this.apiV2DatabaseInstancesPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector.
     * Delete a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorDeleteRaw(requestParameters: ApiV2DatabaseInstancesSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersDatabaseInstance>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2DatabaseInstancesSelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/database-instances/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersDatabaseInstanceFromJSON(jsonValue));
    }

    /**
     * Delete an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector.
     * Delete a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorDelete(requestParameters: ApiV2DatabaseInstancesSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersDatabaseInstance> {
        const response = await this.apiV2DatabaseInstancesSelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector.
     * Get a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorGetRaw(requestParameters: ApiV2DatabaseInstancesSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersDatabaseInstance>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2DatabaseInstancesSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/database-instances/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersDatabaseInstanceFromJSON(jsonValue));
    }

    /**
     * Get an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector.
     * Get a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorGet(requestParameters: ApiV2DatabaseInstancesSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersDatabaseInstance> {
        const response = await this.apiV2DatabaseInstancesSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorPatchRaw(requestParameters: ApiV2DatabaseInstancesSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersDatabaseInstance>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2DatabaseInstancesSelectorPatch.');
        }

        if (requestParameters.chart === null || requestParameters.chart === undefined) {
            throw new runtime.RequiredError('chart','Required parameter requestParameters.chart was null or undefined when calling apiV2DatabaseInstancesSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/database-instances/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableDatabaseInstanceToJSON(requestParameters.chart),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersDatabaseInstanceFromJSON(jsonValue));
    }

    /**
     * Edit an existing DatabaseInstance entry via one of its \"selectors\": numeric ID or \'chart-release/\' followed by a chart release selector. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorPatch(requestParameters: ApiV2DatabaseInstancesSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersDatabaseInstance> {
        const response = await this.apiV2DatabaseInstancesSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit a DatabaseInstance entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorPutRaw(requestParameters: ApiV2DatabaseInstancesSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersDatabaseInstance>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2DatabaseInstancesSelectorPut.');
        }

        if (requestParameters.databaseInstance === null || requestParameters.databaseInstance === undefined) {
            throw new runtime.RequiredError('databaseInstance','Required parameter requestParameters.databaseInstance was null or undefined when calling apiV2DatabaseInstancesSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/database-instances/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableDatabaseInstanceToJSON(requestParameters.databaseInstance),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersDatabaseInstanceFromJSON(jsonValue));
    }

    /**
     * Create or edit a DatabaseInstance entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a DatabaseInstance entry
     */
    async apiV2DatabaseInstancesSelectorPut(requestParameters: ApiV2DatabaseInstancesSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersDatabaseInstance> {
        const response = await this.apiV2DatabaseInstancesSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given DatabaseInstance selector and provide any other selectors that would match the same DatabaseInstance.
     * List DatabaseInstance selectors
     */
    async apiV2SelectorsDatabaseInstancesSelectorGetRaw(requestParameters: ApiV2SelectorsDatabaseInstancesSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsDatabaseInstancesSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/database-instances/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given DatabaseInstance selector and provide any other selectors that would match the same DatabaseInstance.
     * List DatabaseInstance selectors
     */
    async apiV2SelectorsDatabaseInstancesSelectorGet(requestParameters: ApiV2SelectorsDatabaseInstancesSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsDatabaseInstancesSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
