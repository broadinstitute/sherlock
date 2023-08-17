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
  SherlockCiRunV3,
  SherlockCiRunV3Upsert,
  V2controllersCiRun,
  V2controllersCreatableCiRun,
  V2controllersEditableCiRun,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockCiRunV3FromJSON,
    SherlockCiRunV3ToJSON,
    SherlockCiRunV3UpsertFromJSON,
    SherlockCiRunV3UpsertToJSON,
    V2controllersCiRunFromJSON,
    V2controllersCiRunToJSON,
    V2controllersCreatableCiRunFromJSON,
    V2controllersCreatableCiRunToJSON,
    V2controllersEditableCiRunFromJSON,
    V2controllersEditableCiRunToJSON,
} from '../models/index';

export interface ApiCiRunsV3GetRequest {
    argoWorkflowsName?: string;
    argoWorkflowsNamespace?: string;
    argoWorkflowsTemplate?: string;
    createdAt?: Date;
    deployHooksDispatchedAt?: Date;
    githubActionsAttemptNumber?: number;
    githubActionsOwner?: string;
    githubActionsRepo?: string;
    githubActionsRunID?: number;
    githubActionsWorkflowPath?: string;
    id?: number;
    platform?: string;
    startedAt?: string;
    status?: string;
    terminalAt?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiCiRunsV3PutRequest {
    ciRun: SherlockCiRunV3Upsert;
}

export interface ApiCiRunsV3SelectorGetRequest {
    selector: string;
}

export interface ApiV2CiRunsGetRequest {
    argoWorkflowsName?: string;
    argoWorkflowsNamespace?: string;
    argoWorkflowsTemplate?: string;
    createdAt?: Date;
    githubActionsAttemptNumber?: number;
    githubActionsOwner?: string;
    githubActionsRepo?: string;
    githubActionsRunID?: number;
    githubActionsWorkflowPath?: string;
    id?: number;
    platform?: string;
    startedAt?: string;
    status?: string;
    terminalAt?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2CiRunsPostRequest {
    ciRun: V2controllersCreatableCiRun;
}

export interface ApiV2CiRunsSelectorDeleteRequest {
    selector: string;
}

export interface ApiV2CiRunsSelectorGetRequest {
    selector: string;
}

export interface ApiV2CiRunsSelectorPatchRequest {
    selector: string;
    ciRun: V2controllersEditableCiRun;
}

export interface ApiV2CiRunsSelectorPutRequest {
    selector: string;
    ciRun: V2controllersCreatableCiRun;
}

export interface ApiV2SelectorsCiRunsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class CiRunsApi extends runtime.BaseAPI {

    /**
     * List CiRuns matching a filter. The CiRuns would have to re-queried directly to load any related resources. Results are ordered by start time, starting at most recent.
     * List CiRuns matching a filter
     */
    async apiCiRunsV3GetRaw(requestParameters: ApiCiRunsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockCiRunV3>>> {
        const queryParameters: any = {};

        if (requestParameters.argoWorkflowsName !== undefined) {
            queryParameters['argoWorkflowsName'] = requestParameters.argoWorkflowsName;
        }

        if (requestParameters.argoWorkflowsNamespace !== undefined) {
            queryParameters['argoWorkflowsNamespace'] = requestParameters.argoWorkflowsNamespace;
        }

        if (requestParameters.argoWorkflowsTemplate !== undefined) {
            queryParameters['argoWorkflowsTemplate'] = requestParameters.argoWorkflowsTemplate;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.deployHooksDispatchedAt !== undefined) {
            queryParameters['deployHooksDispatchedAt'] = (requestParameters.deployHooksDispatchedAt as any).toISOString();
        }

        if (requestParameters.githubActionsAttemptNumber !== undefined) {
            queryParameters['githubActionsAttemptNumber'] = requestParameters.githubActionsAttemptNumber;
        }

        if (requestParameters.githubActionsOwner !== undefined) {
            queryParameters['githubActionsOwner'] = requestParameters.githubActionsOwner;
        }

        if (requestParameters.githubActionsRepo !== undefined) {
            queryParameters['githubActionsRepo'] = requestParameters.githubActionsRepo;
        }

        if (requestParameters.githubActionsRunID !== undefined) {
            queryParameters['githubActionsRunID'] = requestParameters.githubActionsRunID;
        }

        if (requestParameters.githubActionsWorkflowPath !== undefined) {
            queryParameters['githubActionsWorkflowPath'] = requestParameters.githubActionsWorkflowPath;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.platform !== undefined) {
            queryParameters['platform'] = requestParameters.platform;
        }

        if (requestParameters.startedAt !== undefined) {
            queryParameters['startedAt'] = requestParameters.startedAt;
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.terminalAt !== undefined) {
            queryParameters['terminalAt'] = requestParameters.terminalAt;
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
            path: `/api/ci-runs/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockCiRunV3FromJSON));
    }

    /**
     * List CiRuns matching a filter. The CiRuns would have to re-queried directly to load any related resources. Results are ordered by start time, starting at most recent.
     * List CiRuns matching a filter
     */
    async apiCiRunsV3Get(requestParameters: ApiCiRunsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockCiRunV3>> {
        const response = await this.apiCiRunsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent. The fields for clusters, charts, chart releases, environments, etc. all accept selectors, and they will be smart about \"spreading\" to indirect relations. More info is available on the CiRunV3Upsert data type, but the gist is that specifying a changeset implies its chart release (and optionally app/chart versions), specifying or implying a chart release implies its environment/cluster, and specifying an environment/cluster implies all chart releases they contain.
     * Create or update a CiRun
     */
    async apiCiRunsV3PutRaw(requestParameters: ApiCiRunsV3PutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockCiRunV3>> {
        if (requestParameters.ciRun === null || requestParameters.ciRun === undefined) {
            throw new runtime.RequiredError('ciRun','Required parameter requestParameters.ciRun was null or undefined when calling apiCiRunsV3Put.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/ci-runs/v3`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockCiRunV3UpsertToJSON(requestParameters.ciRun),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockCiRunV3FromJSON(jsonValue));
    }

    /**
     * Create or update a CiRun with timing, status, and related resource information. This endpoint is idempotent. The fields for clusters, charts, chart releases, environments, etc. all accept selectors, and they will be smart about \"spreading\" to indirect relations. More info is available on the CiRunV3Upsert data type, but the gist is that specifying a changeset implies its chart release (and optionally app/chart versions), specifying or implying a chart release implies its environment/cluster, and specifying an environment/cluster implies all chart releases they contain.
     * Create or update a CiRun
     */
    async apiCiRunsV3Put(requestParameters: ApiCiRunsV3PutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockCiRunV3> {
        const response = await this.apiCiRunsV3PutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get a CiRun, including CiIdentifiers representing related resources or resources it affected.
     * Get a CiRun, including CiIdentifiers for related resources
     */
    async apiCiRunsV3SelectorGetRaw(requestParameters: ApiCiRunsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockCiRunV3>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiCiRunsV3SelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/ci-runs/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockCiRunV3FromJSON(jsonValue));
    }

    /**
     * Get a CiRun, including CiIdentifiers representing related resources or resources it affected.
     * Get a CiRun, including CiIdentifiers for related resources
     */
    async apiCiRunsV3SelectorGet(requestParameters: ApiCiRunsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockCiRunV3> {
        const response = await this.apiCiRunsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing CiRun entries, ordered by most recently updated.
     * List CiRun entries
     */
    async apiV2CiRunsGetRaw(requestParameters: ApiV2CiRunsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersCiRun>>> {
        const queryParameters: any = {};

        if (requestParameters.argoWorkflowsName !== undefined) {
            queryParameters['argoWorkflowsName'] = requestParameters.argoWorkflowsName;
        }

        if (requestParameters.argoWorkflowsNamespace !== undefined) {
            queryParameters['argoWorkflowsNamespace'] = requestParameters.argoWorkflowsNamespace;
        }

        if (requestParameters.argoWorkflowsTemplate !== undefined) {
            queryParameters['argoWorkflowsTemplate'] = requestParameters.argoWorkflowsTemplate;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.githubActionsAttemptNumber !== undefined) {
            queryParameters['githubActionsAttemptNumber'] = requestParameters.githubActionsAttemptNumber;
        }

        if (requestParameters.githubActionsOwner !== undefined) {
            queryParameters['githubActionsOwner'] = requestParameters.githubActionsOwner;
        }

        if (requestParameters.githubActionsRepo !== undefined) {
            queryParameters['githubActionsRepo'] = requestParameters.githubActionsRepo;
        }

        if (requestParameters.githubActionsRunID !== undefined) {
            queryParameters['githubActionsRunID'] = requestParameters.githubActionsRunID;
        }

        if (requestParameters.githubActionsWorkflowPath !== undefined) {
            queryParameters['githubActionsWorkflowPath'] = requestParameters.githubActionsWorkflowPath;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.platform !== undefined) {
            queryParameters['platform'] = requestParameters.platform;
        }

        if (requestParameters.startedAt !== undefined) {
            queryParameters['startedAt'] = requestParameters.startedAt;
        }

        if (requestParameters.status !== undefined) {
            queryParameters['status'] = requestParameters.status;
        }

        if (requestParameters.terminalAt !== undefined) {
            queryParameters['terminalAt'] = requestParameters.terminalAt;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/ci-runs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersCiRunFromJSON));
    }

    /**
     * List existing CiRun entries, ordered by most recently updated.
     * List CiRun entries
     */
    async apiV2CiRunsGet(requestParameters: ApiV2CiRunsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersCiRun>> {
        const response = await this.apiV2CiRunsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new CiRun entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new CiRun entry
     */
    async apiV2CiRunsPostRaw(requestParameters: ApiV2CiRunsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCiRun>> {
        if (requestParameters.ciRun === null || requestParameters.ciRun === undefined) {
            throw new runtime.RequiredError('ciRun','Required parameter requestParameters.ciRun was null or undefined when calling apiV2CiRunsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/ci-runs`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableCiRunToJSON(requestParameters.ciRun),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersCiRunFromJSON(jsonValue));
    }

    /**
     * Create a new CiRun entry. Note that some fields are immutable after creation; /edit lists mutable fields.
     * Create a new CiRun entry
     */
    async apiV2CiRunsPost(requestParameters: ApiV2CiRunsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCiRun> {
        const response = await this.apiV2CiRunsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name.
     * Delete a CiRun entry
     */
    async apiV2CiRunsSelectorDeleteRaw(requestParameters: ApiV2CiRunsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCiRun>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2CiRunsSelectorDelete.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/ci-runs/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersCiRunFromJSON(jsonValue));
    }

    /**
     * Delete an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name.
     * Delete a CiRun entry
     */
    async apiV2CiRunsSelectorDelete(requestParameters: ApiV2CiRunsSelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCiRun> {
        const response = await this.apiV2CiRunsSelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name.
     * Get a CiRun entry
     */
    async apiV2CiRunsSelectorGetRaw(requestParameters: ApiV2CiRunsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCiRun>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2CiRunsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/ci-runs/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersCiRunFromJSON(jsonValue));
    }

    /**
     * Get an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name.
     * Get a CiRun entry
     */
    async apiV2CiRunsSelectorGet(requestParameters: ApiV2CiRunsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCiRun> {
        const response = await this.apiV2CiRunsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a CiRun entry
     */
    async apiV2CiRunsSelectorPatchRaw(requestParameters: ApiV2CiRunsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCiRun>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2CiRunsSelectorPatch.');
        }

        if (requestParameters.ciRun === null || requestParameters.ciRun === undefined) {
            throw new runtime.RequiredError('ciRun','Required parameter requestParameters.ciRun was null or undefined when calling apiV2CiRunsSelectorPatch.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/ci-runs/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersEditableCiRunToJSON(requestParameters.ciRun),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersCiRunFromJSON(jsonValue));
    }

    /**
     * Edit an existing CiRun entry via one of its \"selectors\": ID, \'github-actions/\' + owner + repo + run ID + attempt number, or \'argo-workflows/\' + namespace + name. Note that only mutable fields are available here, immutable fields can only be set using /create.
     * Edit a CiRun entry
     */
    async apiV2CiRunsSelectorPatch(requestParameters: ApiV2CiRunsSelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCiRun> {
        const response = await this.apiV2CiRunsSelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or edit a CiRun entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a CiRun entry
     */
    async apiV2CiRunsSelectorPutRaw(requestParameters: ApiV2CiRunsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersCiRun>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2CiRunsSelectorPut.');
        }

        if (requestParameters.ciRun === null || requestParameters.ciRun === undefined) {
            throw new runtime.RequiredError('ciRun','Required parameter requestParameters.ciRun was null or undefined when calling apiV2CiRunsSelectorPut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/ci-runs/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableCiRunToJSON(requestParameters.ciRun),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersCiRunFromJSON(jsonValue));
    }

    /**
     * Create or edit a CiRun entry. Attempts to edit and will attempt to create upon an error. If an edit was made or the creation process de-duplicates, this method will return normally with a 200.
     * Create or edit a CiRun entry
     */
    async apiV2CiRunsSelectorPut(requestParameters: ApiV2CiRunsSelectorPutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersCiRun> {
        const response = await this.apiV2CiRunsSelectorPutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given CiRun selector and provide any other selectors that would match the same CiRun.
     * List CiRun selectors
     */
    async apiV2SelectorsCiRunsSelectorGetRaw(requestParameters: ApiV2SelectorsCiRunsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsCiRunsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/ci-runs/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given CiRun selector and provide any other selectors that would match the same CiRun.
     * List CiRun selectors
     */
    async apiV2SelectorsCiRunsSelectorGet(requestParameters: ApiV2SelectorsCiRunsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsCiRunsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
