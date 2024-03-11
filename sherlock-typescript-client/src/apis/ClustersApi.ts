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

        if (requestParameters['address'] != null) {
            queryParameters['address'] = requestParameters['address'];
        }

        if (requestParameters['azureSubscription'] != null) {
            queryParameters['azureSubscription'] = requestParameters['azureSubscription'];
        }

        if (requestParameters['base'] != null) {
            queryParameters['base'] = requestParameters['base'];
        }

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['googleProject'] != null) {
            queryParameters['googleProject'] = requestParameters['googleProject'];
        }

        if (requestParameters['helmfileRef'] != null) {
            queryParameters['helmfileRef'] = requestParameters['helmfileRef'];
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['location'] != null) {
            queryParameters['location'] = requestParameters['location'];
        }

        if (requestParameters['name'] != null) {
            queryParameters['name'] = requestParameters['name'];
        }

        if (requestParameters['provider'] != null) {
            queryParameters['provider'] = requestParameters['provider'];
        }

        if (requestParameters['requiresSuitability'] != null) {
            queryParameters['requiresSuitability'] = requestParameters['requiresSuitability'];
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
        if (requestParameters['cluster'] == null) {
            throw new runtime.RequiredError(
                'cluster',
                'Required parameter "cluster" was null or undefined when calling apiClustersV3Post().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/clusters/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockClusterV3CreateToJSON(requestParameters['cluster']),
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
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiClustersV3SelectorDelete().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
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
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiClustersV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
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
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiClustersV3SelectorPatch().'
            );
        }

        if (requestParameters['cluster'] == null) {
            throw new runtime.RequiredError(
                'cluster',
                'Required parameter "cluster" was null or undefined when calling apiClustersV3SelectorPatch().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/clusters/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockClusterV3EditToJSON(requestParameters['cluster']),
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

}

/**
 * @export
 */
export const ApiClustersV3GetProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type ApiClustersV3GetProviderEnum = typeof ApiClustersV3GetProviderEnum[keyof typeof ApiClustersV3GetProviderEnum];
