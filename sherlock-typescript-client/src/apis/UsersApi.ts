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
  SherlockUserV3,
  SherlockUserV3DeactivateRequest,
  SherlockUserV3DeactivateResponse,
  SherlockUserV3Upsert,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockUserV3FromJSON,
    SherlockUserV3ToJSON,
    SherlockUserV3DeactivateRequestFromJSON,
    SherlockUserV3DeactivateRequestToJSON,
    SherlockUserV3DeactivateResponseFromJSON,
    SherlockUserV3DeactivateResponseToJSON,
    SherlockUserV3UpsertFromJSON,
    SherlockUserV3UpsertToJSON,
} from '../models/index';

export interface ApiUsersProceduresV3DeactivatePostRequest {
    users: SherlockUserV3DeactivateRequest;
}

export interface ApiUsersV3GetRequest {
    createdAt?: Date;
    deactivatedAt?: string;
    email?: string;
    githubID?: string;
    githubUsername?: string;
    googleID?: string;
    id?: number;
    name?: string;
    nameFrom?: ApiUsersV3GetNameFromEnum;
    slackID?: string;
    slackUsername?: string;
    suitabilityDescription?: string;
    suitable?: boolean;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
    includeDeactivated?: boolean;
}

export interface ApiUsersV3PutRequest {
    user?: SherlockUserV3Upsert;
}

export interface ApiUsersV3SelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class UsersApi extends runtime.BaseAPI {

    /**
     * Super-admin only method to deactivate users. Deactivated users will be removed from all roles and can\'t authenticate to Sherlock. This endpoint can optionally also attempt to suspend the same email handles across given Google Workspace domains, substituting email domains as necessary. It will do so by impersonating the caller in each given domain.
     * Deactivate Users
     */
    async apiUsersProceduresV3DeactivatePostRaw(requestParameters: ApiUsersProceduresV3DeactivatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockUserV3DeactivateResponse>> {
        if (requestParameters['users'] == null) {
            throw new runtime.RequiredError(
                'users',
                'Required parameter "users" was null or undefined when calling apiUsersProceduresV3DeactivatePost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';


        let urlPath = `/api/users/procedures/v3/deactivate`;

        const response = await this.request({
            path: urlPath,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockUserV3DeactivateRequestToJSON(requestParameters['users']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockUserV3DeactivateResponseFromJSON(jsonValue));
    }

    /**
     * Super-admin only method to deactivate users. Deactivated users will be removed from all roles and can\'t authenticate to Sherlock. This endpoint can optionally also attempt to suspend the same email handles across given Google Workspace domains, substituting email domains as necessary. It will do so by impersonating the caller in each given domain.
     * Deactivate Users
     */
    async apiUsersProceduresV3DeactivatePost(requestParameters: ApiUsersProceduresV3DeactivatePostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockUserV3DeactivateResponse> {
        const response = await this.apiUsersProceduresV3DeactivatePostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List Users matching a filter. The results will include suitability and other information. Note that the suitability info can\'t directly be filtered for at this time.
     * List Users matching a filter
     */
    async apiUsersV3GetRaw(requestParameters: ApiUsersV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockUserV3>>> {
        const queryParameters: any = {};

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['deactivatedAt'] != null) {
            queryParameters['deactivatedAt'] = requestParameters['deactivatedAt'];
        }

        if (requestParameters['email'] != null) {
            queryParameters['email'] = requestParameters['email'];
        }

        if (requestParameters['githubID'] != null) {
            queryParameters['githubID'] = requestParameters['githubID'];
        }

        if (requestParameters['githubUsername'] != null) {
            queryParameters['githubUsername'] = requestParameters['githubUsername'];
        }

        if (requestParameters['googleID'] != null) {
            queryParameters['googleID'] = requestParameters['googleID'];
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['name'] != null) {
            queryParameters['name'] = requestParameters['name'];
        }

        if (requestParameters['nameFrom'] != null) {
            queryParameters['nameFrom'] = requestParameters['nameFrom'];
        }

        if (requestParameters['slackID'] != null) {
            queryParameters['slackID'] = requestParameters['slackID'];
        }

        if (requestParameters['slackUsername'] != null) {
            queryParameters['slackUsername'] = requestParameters['slackUsername'];
        }

        if (requestParameters['suitabilityDescription'] != null) {
            queryParameters['suitabilityDescription'] = requestParameters['suitabilityDescription'];
        }

        if (requestParameters['suitable'] != null) {
            queryParameters['suitable'] = requestParameters['suitable'];
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

        if (requestParameters['includeDeactivated'] != null) {
            queryParameters['include-deactivated'] = requestParameters['includeDeactivated'];
        }

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/users/v3`;

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockUserV3FromJSON));
    }

    /**
     * List Users matching a filter. The results will include suitability and other information. Note that the suitability info can\'t directly be filtered for at this time.
     * List Users matching a filter
     */
    async apiUsersV3Get(requestParameters: ApiUsersV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockUserV3>> {
        const response = await this.apiUsersV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update the calling User\'s information. As with all authenticated Sherlock endpoints, newly-observed callers will have a User record added, meaning that this endpoint behaves like an upsert.
     * Update the calling User\'s information
     */
    async apiUsersV3PutRaw(requestParameters: ApiUsersV3PutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockUserV3>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';


        let urlPath = `/api/users/v3`;

        const response = await this.request({
            path: urlPath,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockUserV3UpsertToJSON(requestParameters['user']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockUserV3FromJSON(jsonValue));
    }

    /**
     * Update the calling User\'s information. As with all authenticated Sherlock endpoints, newly-observed callers will have a User record added, meaning that this endpoint behaves like an upsert.
     * Update the calling User\'s information
     */
    async apiUsersV3Put(requestParameters: ApiUsersV3PutRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockUserV3> {
        const response = await this.apiUsersV3PutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual User. As a special case, \"me\" or \"self\" can be passed as the selector to get the current user.
     * Get an individual User
     */
    async apiUsersV3SelectorGetRaw(requestParameters: ApiUsersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockUserV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiUsersV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/users/v3/{selector}`;
        urlPath = urlPath.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector'])));

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockUserV3FromJSON(jsonValue));
    }

    /**
     * Get an individual User. As a special case, \"me\" or \"self\" can be passed as the selector to get the current user.
     * Get an individual User
     */
    async apiUsersV3SelectorGet(requestParameters: ApiUsersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockUserV3> {
        const response = await this.apiUsersV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const ApiUsersV3GetNameFromEnum = {
    Sherlock: 'sherlock',
    Github: 'github',
    Slack: 'slack'
} as const;
export type ApiUsersV3GetNameFromEnum = typeof ApiUsersV3GetNameFromEnum[keyof typeof ApiUsersV3GetNameFromEnum];
