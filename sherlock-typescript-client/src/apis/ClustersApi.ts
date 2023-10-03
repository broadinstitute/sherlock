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
  SherlockClusterV3,
  SherlockClusterV3Create,
  SherlockClusterV3Edit,
  V2controllersCluster,
  V2controllersCreatableCluster,
  V2controllersEditableCluster,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockClusterV3FromJSON,
    SherlockClusterV3ToJSON,
    SherlockClusterV3CreateFromJSON,
    SherlockClusterV3CreateToJSON,
    SherlockClusterV3EditFromJSON,
    SherlockClusterV3EditToJSON,
    V2controllersClusterFromJSON,
    V2controllersClusterToJSON,
    V2controllersCreatableClusterFromJSON,
    V2controllersCreatableClusterToJSON,
    V2controllersEditableClusterFromJSON,
    V2controllersEditableClusterToJSON,
} from '../models/index';

export interface ApiClustersV3GetRequest {
    address?: string;
    azureSubscription?: string;
    base?: string;
    createdAt?: Date;
    googleProject?: string;
    helmfileRef?: string;
    id?: number;
    location?: string;
    name?: string;
    provider?: ApiClustersV3GetProviderEnum;
    requiresSuitability?: boolean;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiClustersV3PostRequest {
    cluster: SherlockClusterV3Create;
}

export interface ApiClustersV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiClustersV3SelectorGetRequest {
    selector: string;
}

export interface ApiClustersV3SelectorPatchRequest {
    selector: string;
    cluster: SherlockClusterV3Edit;
}

export interface ApiV2ClustersGetRequest {
    address?: string;
    azureSubscription?: string;
    base?: string;
    createdAt?: Date;
    googleProject?: string;
    helmfileRef?: string;
    id?: number;
    location?: string;
    name?: string;
    provider?: ApiV2ClustersGetProviderEnum;
    requiresSuitability?: boolean;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2ClustersPostRequest {
    cluster: V2controllersCreatableCluster;
}

export interface ApiV2ClustersSelectorDeleteRequest {
    selector: string;
}

export interface ApiV2ClustersSelectorGetRequest {
    selector: string;
}

export interface ApiV2ClustersSelectorPatchRequest {
    selector: string;
    cluster: V2controllersEditableCluster;
}

export interface ApiV2ClustersSelectorPutRequest {
    selector: string;
    cluster: V2controllersCreatableCluster;
}

export interface ApiV2SelectorsClustersSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class ClustersApi extends runtime.BaseAPI {

    /**
     * List Clusters matching a filter.
     * List Clusters matching a filter
     */
    async apiClustersV3GetRaw(requestParameters: ApiClustersV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockClusterV3>>> {
        const queryParameters: any = {};

        if (requestParameters.address !== undefined) {
            queryParameters['address'] = requestParameters.address;
        }

        if (requestParameters.azureSubscription !== undefined) {
            queryParameters['azureSubscription'] = requestParameters.azureSubscription;
        }

        if (requestParameters.base !== undefined) {
            queryParameters['base'] = requestParameters.base;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.googleProject !== undefined) {
            queryParameters['googleProject'] = requestParameters.googleProject;
        }

        if (requestParameters.helmfileRef !== undefined) {
            queryParameters['helmfileRef'] = requestParameters.helmfileRef;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.location !== undefined) {
            queryParameters['location'] = requestParameters.location;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.provider !== undefined) {
            queryParameters['provider'] = requestParameters.provider;
        }

        if (requestParameters.requiresSuitability !== undefined) {
            queryParameters['requiresSuitability'] = requestParameters.requiresSuitability;
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
            path: `/api/clusters/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockClusterV3FromJSON));
    }

    /**
     * List Clusters matching a filter.
     * List Clusters matching a filter
     */
    async apiClustersV3Get(requestParameters: ApiClustersV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockClusterV3>> {
        const response = await this.apiClustersV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a Cluster.
     * Create a Cluster
     */
    async apiClustersV3PostRaw(requestParameters: ApiClustersV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockClusterV3>> {
        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling apiClustersV3Post.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/clusters/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockClusterV3CreateToJSON(requestParameters.cluster),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockClusterV3FromJSON(jsonValue));
    }

    /**
     * Create a Cluster.
     * Create a Cluster
     */
    async apiClustersV3Post(requestParameters: ApiClustersV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockClusterV3> {
        const response = await this.apiClustersV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual Cluster by its ID.
     * Delete an individual Cluster
     */
    async apiClustersV3SelectorDeleteRaw(requestParameters: ApiClustersV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockClusterV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiClustersV3SelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockClusterV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual Cluster by its ID.
     * Delete an individual Cluster
     */
    async apiClustersV3SelectorDelete(requestParameters: ApiClustersV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockClusterV3> {
        const response = await this.apiClustersV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual Cluster.
     * Get an individual Cluster
     */
    async apiClustersV3SelectorGetRaw(requestParameters: ApiClustersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockClusterV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiClustersV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockClusterV3FromJSON(jsonValue));
    }

    /**
     * Get an individual Cluster.
     * Get an individual Cluster
     */
    async apiClustersV3SelectorGet(requestParameters: ApiClustersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockClusterV3> {
        const response = await this.apiClustersV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual Cluster.
     * Edit an individual Cluster
     */
    async apiClustersV3SelectorPatchRaw(requestParameters: ApiClustersV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockClusterV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiClustersV3SelectorPatch.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling apiClustersV3SelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockClusterV3EditToJSON(requestParameters.cluster),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockClusterV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual Cluster.
     * Edit an individual Cluster
     */
    async apiClustersV3SelectorPatch(requestParameters: ApiClustersV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockClusterV3> {
        const response = await this.apiClustersV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing Cluster entries, ordered by most recently updated.
     * List Cluster entries
     */
    async apiV2ClustersGetRaw(requestParameters: ApiV2ClustersGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersCluster>>> {
        const queryParameters: any = {};

        if (requestParameters.address !== undefined) {
            queryParameters['address'] = requestParameters.address;
        }

        if (requestParameters.azureSubscription !== undefined) {
            queryParameters['azureSubscription'] = requestParameters.azureSubscription;
        }

        if (requestParameters.base !== undefined) {
            queryParameters['base'] = requestParameters.base;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.googleProject !== undefined) {
            queryParameters['googleProject'] = requestParameters.googleProject;
        }

        if (requestParameters.helmfileRef !== undefined) {
            queryParameters['helmfileRef'] = requestParameters.helmfileRef;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.location !== undefined) {
            queryParameters['location'] = requestParameters.location;
        }

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.provider !== undefined) {
            queryParameters['provider'] = requestParameters.provider;
        }

        if (requestParameters.requiresSuitability !== undefined) {
            queryParameters['requiresSuitability'] = requestParameters.requiresSuitability;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/clusters`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersClusterFromJSON));
    }

    /**
     * List existing Cluster entries, ordered by most recently updated.
     * List Cluster entries
     */
    async apiV2ClustersGet(requestParameters: ApiV2ClustersGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersCluster>> {
        const response = await this.apiV2ClustersGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new Cluster entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new Cluster entry
     */
    async apiV2ClustersPostRaw(requestParameters: ApiV2ClustersPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCluster>> {
        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling apiV2ClustersPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/clusters`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableClusterToJSON(requestParameters.cluster),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersClusterFromJSON(jsonValue));
    }

    /**
     * Create a new Cluster entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new Cluster entry
     */
    async apiV2ClustersPost(requestParameters: ApiV2ClustersPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCluster> {
        const response = await this.apiV2ClustersPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an existing Cluster entry via one of its \"selectors\": name or numeric ID.
     * Delete a Cluster entry
     */
    async apiV2ClustersSelectorDeleteRaw(requestParameters: ApiV2ClustersSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCluster>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ClustersSelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/clusters/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersClusterFromJSON(jsonValue));
    }

    /**
     * Delete an existing Cluster entry via one of its \"selectors\": name or numeric ID.
     * Delete a Cluster entry
     */
    async apiV2ClustersSelectorDelete(requestParameters: ApiV2ClustersSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCluster> {
        const response = await this.apiV2ClustersSelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing Cluster entry via one of its \"selectors\": name or numeric ID.
     * Get a Cluster entry
     */
    async apiV2ClustersSelectorGetRaw(requestParameters: ApiV2ClustersSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCluster>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ClustersSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/clusters/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersClusterFromJSON(jsonValue));
    }

    /**
     * Get an existing Cluster entry via one of its \"selectors\": name or numeric ID.
     * Get a Cluster entry
     */
    async apiV2ClustersSelectorGet(requestParameters: ApiV2ClustersSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCluster> {
        const response = await this.apiV2ClustersSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing Cluster entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a Cluster entry
     */
    async apiV2ClustersSelectorPatchRaw(requestParameters: ApiV2ClustersSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCluster>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ClustersSelectorPatch.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling apiV2ClustersSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/clusters/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableClusterToJSON(requestParameters.cluster),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersClusterFromJSON(jsonValue));
    }

    /**
     * Edit an existing Cluster entry via one of its \"selectors\": name or numeric ID. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a Cluster entry
     */
    async apiV2ClustersSelectorPatch(requestParameters: ApiV2ClustersSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCluster> {
        const response = await this.apiV2ClustersSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit a Cluster entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a Cluster entry
     */
    async apiV2ClustersSelectorPutRaw(requestParameters: ApiV2ClustersSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCluster>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ClustersSelectorPut.');
        }

        if (requestParameters.cluster === null || requestParameters.cluster === undefined) {
            throw new runtime.RequiredError('cluster','Required parameter requestParameters.cluster was null or undefined when calling apiV2ClustersSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/clusters/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableClusterToJSON(requestParameters.cluster),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersClusterFromJSON(jsonValue));
    }

    /**
     * Create or edit a Cluster entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a Cluster entry
     */
    async apiV2ClustersSelectorPut(requestParameters: ApiV2ClustersSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCluster> {
        const response = await this.apiV2ClustersSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given Cluster selector and provide any other selectors that would match the same Cluster.
     * List Cluster selectors
     */
    async apiV2SelectorsClustersSelectorGetRaw(requestParameters: ApiV2SelectorsClustersSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsClustersSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/clusters/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given Cluster selector and provide any other selectors that would match the same Cluster.
     * List Cluster selectors
     */
    async apiV2SelectorsClustersSelectorGet(requestParameters: ApiV2SelectorsClustersSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsClustersSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const ApiClustersV3GetProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type ApiClustersV3GetProviderEnum = typeof ApiClustersV3GetProviderEnum[keyof typeof ApiClustersV3GetProviderEnum];
/**
 * @export
 */
export const ApiV2ClustersGetProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type ApiV2ClustersGetProviderEnum = typeof ApiV2ClustersGetProviderEnum[keyof typeof ApiV2ClustersGetProviderEnum];
