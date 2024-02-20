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
  SherlockIncidentV3,
  SherlockIncidentV3Create,
  SherlockIncidentV3Edit,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockIncidentV3FromJSON,
    SherlockIncidentV3ToJSON,
    SherlockIncidentV3CreateFromJSON,
    SherlockIncidentV3CreateToJSON,
    SherlockIncidentV3EditFromJSON,
    SherlockIncidentV3EditToJSON,
} from '../models/index';

export interface ApiIncidentsV3GetRequest {
    createdAt?: Date;
    description?: string;
    id?: number;
    remediatedAt?: string;
    reviewCompletedAt?: string;
    startedAt?: string;
    ticket?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiIncidentsV3PostRequest {
    incident: SherlockIncidentV3Create;
}

export interface ApiIncidentsV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiIncidentsV3SelectorGetRequest {
    selector: string;
}

export interface ApiIncidentsV3SelectorPatchRequest {
    selector: string;
    incident: SherlockIncidentV3Edit;
}

/**
 * 
 */
export class IncidentsApi extends runtime.BaseAPI {

    /**
     * List Incidents matching a filter.
     * List Incidents matching a filter
     */
    async apiIncidentsV3GetRaw(requestParameters: ApiIncidentsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockIncidentV3>>> {
        const queryParameters: any = {};

        if (runtime.exists(requestParameters, 'createdAt')) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (runtime.exists(requestParameters, 'description')) {
            queryParameters['description'] = requestParameters['description'];
        }

        if (runtime.exists(requestParameters, 'id')) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (runtime.exists(requestParameters, 'remediatedAt')) {
            queryParameters['remediatedAt'] = requestParameters['remediatedAt'];
        }

        if (runtime.exists(requestParameters, 'reviewCompletedAt')) {
            queryParameters['reviewCompletedAt'] = requestParameters['reviewCompletedAt'];
        }

        if (runtime.exists(requestParameters, 'startedAt')) {
            queryParameters['startedAt'] = requestParameters['startedAt'];
        }

        if (runtime.exists(requestParameters, 'ticket')) {
            queryParameters['ticket'] = requestParameters['ticket'];
        }

        if (runtime.exists(requestParameters, 'updatedAt')) {
            queryParameters['updatedAt'] = (requestParameters['updatedAt'] as any).toISOString();
        }

        if (runtime.exists(requestParameters, 'limit')) {
            queryParameters['limit'] = requestParameters['limit'];
        }

        if (runtime.exists(requestParameters, 'offset')) {
            queryParameters['offset'] = requestParameters['offset'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/incidents/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockIncidentV3FromJSON));
    }

    /**
     * List Incidents matching a filter.
     * List Incidents matching a filter
     */
    async apiIncidentsV3Get(requestParameters: ApiIncidentsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockIncidentV3>> {
        const response = await this.apiIncidentsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a Incident.
     * Create a Incident
     */
    async apiIncidentsV3PostRaw(requestParameters: ApiIncidentsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockIncidentV3>> {
        if (!runtime.exists(requestParameters, 'incident')) {
            throw new runtime.RequiredError(
                'incident',
                'Required parameter "incident" was null or undefined when calling apiIncidentsV3Post().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/incidents/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockIncidentV3CreateToJSON(requestParameters['incident']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockIncidentV3FromJSON(jsonValue));
    }

    /**
     * Create a Incident.
     * Create a Incident
     */
    async apiIncidentsV3Post(requestParameters: ApiIncidentsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockIncidentV3> {
        const response = await this.apiIncidentsV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual Incident by its ID.
     * Delete an individual Incident
     */
    async apiIncidentsV3SelectorDeleteRaw(requestParameters: ApiIncidentsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockIncidentV3>> {
        if (!runtime.exists(requestParameters, 'selector')) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiIncidentsV3SelectorDelete().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/incidents/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockIncidentV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual Incident by its ID.
     * Delete an individual Incident
     */
    async apiIncidentsV3SelectorDelete(requestParameters: ApiIncidentsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockIncidentV3> {
        const response = await this.apiIncidentsV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual Incident.
     * Get an individual Incident
     */
    async apiIncidentsV3SelectorGetRaw(requestParameters: ApiIncidentsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockIncidentV3>> {
        if (!runtime.exists(requestParameters, 'selector')) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiIncidentsV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/incidents/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockIncidentV3FromJSON(jsonValue));
    }

    /**
     * Get an individual Incident.
     * Get an individual Incident
     */
    async apiIncidentsV3SelectorGet(requestParameters: ApiIncidentsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockIncidentV3> {
        const response = await this.apiIncidentsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual Incident.
     * Edit an individual Incident
     */
    async apiIncidentsV3SelectorPatchRaw(requestParameters: ApiIncidentsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockIncidentV3>> {
        if (!runtime.exists(requestParameters, 'selector')) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiIncidentsV3SelectorPatch().'
            );
        }

        if (!runtime.exists(requestParameters, 'incident')) {
            throw new runtime.RequiredError(
                'incident',
                'Required parameter "incident" was null or undefined when calling apiIncidentsV3SelectorPatch().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/incidents/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockIncidentV3EditToJSON(requestParameters['incident']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockIncidentV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual Incident.
     * Edit an individual Incident
     */
    async apiIncidentsV3SelectorPatch(requestParameters: ApiIncidentsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockIncidentV3> {
        const response = await this.apiIncidentsV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
