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
  V2controllersGithubAccessPayload,
  V2controllersUser,
} from '../models';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    V2controllersGithubAccessPayloadFromJSON,
    V2controllersGithubAccessPayloadToJSON,
    V2controllersUserFromJSON,
    V2controllersUserToJSON,
} from '../models';

export interface ApiV2ProceduresUsersLinkGithubPostRequest {
    githubAccessPayloadRequest: V2controllersGithubAccessPayload;
}

/**
 * 
 */
export class UsersApi extends runtime.BaseAPI {

    /**
     * Update the authenticated User\'s associated personal GitHub account
     * Update the User\'s GitHub account link
     */
    async apiV2ProceduresUsersLinkGithubPostRaw(requestParameters: ApiV2ProceduresUsersLinkGithubPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersUser>> {
        if (requestParameters.githubAccessPayloadRequest === null || requestParameters.githubAccessPayloadRequest === undefined) {
            throw new runtime.RequiredError('githubAccessPayloadRequest','Required parameter requestParameters.githubAccessPayloadRequest was null or undefined when calling apiV2ProceduresUsersLinkGithubPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/procedures/users/link-github`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersGithubAccessPayloadToJSON(requestParameters.githubAccessPayloadRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersUserFromJSON(jsonValue));
    }

    /**
     * Update the authenticated User\'s associated personal GitHub account
     * Update the User\'s GitHub account link
     */
    async apiV2ProceduresUsersLinkGithubPost(requestParameters: ApiV2ProceduresUsersLinkGithubPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersUser> {
        const response = await this.apiV2ProceduresUsersLinkGithubPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
