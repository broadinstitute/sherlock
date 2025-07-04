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
  SherlockCiIdentifierV3,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3ToJSON,
} from '../models/index';

export interface ApiCiIdentifiersV3GetRequest {
    createdAt?: Date;
    id?: number;
    resourceID?: number;
    resourceStatus?: string;
    resourceType?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiCiIdentifiersV3SelectorGetRequest {
    selector: string;
    limitCiRuns?: number;
    offsetCiRuns?: number;
    allowStubCiRuns?: boolean;
}

/**
 * 
 */
export class CiIdentifiersApi extends runtime.BaseAPI {

    /**
     * List CiIdentifiers matching a filter. The CiRuns would have to re-queried directly to load the CiRuns. This is mainly helpful for debugging and directly querying the existence of a CiIdentifier. Results are ordered by creation date, starting at most recent.
     * List CiIdentifiers matching a filter
     */
    async apiCiIdentifiersV3GetRaw(requestParameters: ApiCiIdentifiersV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockCiIdentifierV3>>> {
        const queryParameters: any = {};

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['resourceID'] != null) {
            queryParameters['resourceID'] = requestParameters['resourceID'];
        }

        if (requestParameters['resourceStatus'] != null) {
            queryParameters['resourceStatus'] = requestParameters['resourceStatus'];
        }

        if (requestParameters['resourceType'] != null) {
            queryParameters['resourceType'] = requestParameters['resourceType'];
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


        let urlPath = `/api/ci-identifiers/v3`;

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockCiIdentifierV3FromJSON));
    }

    /**
     * List CiIdentifiers matching a filter. The CiRuns would have to re-queried directly to load the CiRuns. This is mainly helpful for debugging and directly querying the existence of a CiIdentifier. Results are ordered by creation date, starting at most recent.
     * List CiIdentifiers matching a filter
     */
    async apiCiIdentifiersV3Get(requestParameters: ApiCiIdentifiersV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockCiIdentifierV3>> {
        const response = await this.apiCiIdentifiersV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get CiRuns for a resource by its CiIdentifier, which can be referenced by \'{type}/{selector...}\'.
     * Get CiRuns for a resource by its CiIdentifier
     */
    async apiCiIdentifiersV3SelectorGetRaw(requestParameters: ApiCiIdentifiersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockCiIdentifierV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiCiIdentifiersV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['limitCiRuns'] != null) {
            queryParameters['limitCiRuns'] = requestParameters['limitCiRuns'];
        }

        if (requestParameters['offsetCiRuns'] != null) {
            queryParameters['offsetCiRuns'] = requestParameters['offsetCiRuns'];
        }

        if (requestParameters['allowStubCiRuns'] != null) {
            queryParameters['allowStubCiRuns'] = requestParameters['allowStubCiRuns'];
        }

        const headerParameters: runtime.HTTPHeaders = {};


        let urlPath = `/api/ci-identifiers/v3/{selector}`;
        urlPath = urlPath.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector'])));

        const response = await this.request({
            path: urlPath,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockCiIdentifierV3FromJSON(jsonValue));
    }

    /**
     * Get CiRuns for a resource by its CiIdentifier, which can be referenced by \'{type}/{selector...}\'.
     * Get CiRuns for a resource by its CiIdentifier
     */
    async apiCiIdentifiersV3SelectorGet(requestParameters: ApiCiIdentifiersV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockCiIdentifierV3> {
        const response = await this.apiCiIdentifiersV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
