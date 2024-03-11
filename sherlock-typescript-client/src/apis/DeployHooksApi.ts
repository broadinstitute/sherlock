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
  SherlockGithubActionsDeployHookTestRunRequest,
  SherlockGithubActionsDeployHookTestRunResponse,
  SherlockGithubActionsDeployHookV3,
  SherlockGithubActionsDeployHookV3Create,
  SherlockGithubActionsDeployHookV3Edit,
  SherlockSlackDeployHookTestRunRequest,
  SherlockSlackDeployHookTestRunResponse,
  SherlockSlackDeployHookV3,
  SherlockSlackDeployHookV3Create,
  SherlockSlackDeployHookV3Edit,
} from '../models/index';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    SherlockGithubActionsDeployHookTestRunRequestFromJSON,
    SherlockGithubActionsDeployHookTestRunRequestToJSON,
    SherlockGithubActionsDeployHookTestRunResponseFromJSON,
    SherlockGithubActionsDeployHookTestRunResponseToJSON,
    SherlockGithubActionsDeployHookV3FromJSON,
    SherlockGithubActionsDeployHookV3ToJSON,
    SherlockGithubActionsDeployHookV3CreateFromJSON,
    SherlockGithubActionsDeployHookV3CreateToJSON,
    SherlockGithubActionsDeployHookV3EditFromJSON,
    SherlockGithubActionsDeployHookV3EditToJSON,
    SherlockSlackDeployHookTestRunRequestFromJSON,
    SherlockSlackDeployHookTestRunRequestToJSON,
    SherlockSlackDeployHookTestRunResponseFromJSON,
    SherlockSlackDeployHookTestRunResponseToJSON,
    SherlockSlackDeployHookV3FromJSON,
    SherlockSlackDeployHookV3ToJSON,
    SherlockSlackDeployHookV3CreateFromJSON,
    SherlockSlackDeployHookV3CreateToJSON,
    SherlockSlackDeployHookV3EditFromJSON,
    SherlockSlackDeployHookV3EditToJSON,
} from '../models/index';

export interface ApiDeployHooksGithubActionsProceduresV3TestSelectorPostRequest {
    selector: string;
    request: SherlockGithubActionsDeployHookTestRunRequest;
}

export interface ApiDeployHooksGithubActionsV3GetRequest {
    createdAt?: Date;
    githubActionsDefaultRef?: string;
    githubActionsOwner?: string;
    githubActionsRefBehavior?: ApiDeployHooksGithubActionsV3GetGithubActionsRefBehaviorEnum;
    githubActionsRepo?: string;
    githubActionsWorkflowPath?: string;
    id?: number;
    onChartRelease?: string;
    onEnvironment?: string;
    onFailure?: boolean;
    onSuccess?: boolean;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiDeployHooksGithubActionsV3PostRequest {
    githubActionsDeployHook: SherlockGithubActionsDeployHookV3Create;
}

export interface ApiDeployHooksGithubActionsV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiDeployHooksGithubActionsV3SelectorGetRequest {
    selector: string;
}

export interface ApiDeployHooksGithubActionsV3SelectorPatchRequest {
    selector: string;
    githubActionsDeployHook: SherlockGithubActionsDeployHookV3Edit;
}

export interface ApiDeployHooksSlackProceduresV3TestSelectorPostRequest {
    selector: string;
    request: SherlockSlackDeployHookTestRunRequest;
}

export interface ApiDeployHooksSlackV3GetRequest {
    createdAt?: Date;
    id?: number;
    mentionPeople?: boolean;
    onChartRelease?: string;
    onEnvironment?: string;
    onFailure?: boolean;
    onSuccess?: boolean;
    slackChannel?: string;
    updatedAt?: Date;
    limit?: number;
    offset?: number;
}

export interface ApiDeployHooksSlackV3PostRequest {
    slackDeployHook: SherlockSlackDeployHookV3Create;
}

export interface ApiDeployHooksSlackV3SelectorDeleteRequest {
    selector: string;
}

export interface ApiDeployHooksSlackV3SelectorGetRequest {
    selector: string;
}

export interface ApiDeployHooksSlackV3SelectorPatchRequest {
    selector: string;
    slackDeployHook: SherlockSlackDeployHookV3Edit;
}

/**
 * 
 */
export class DeployHooksApi extends runtime.BaseAPI {

    /**
     * Run a GitHub Action to simulate a GithubActionsDeployHook
     * Test a GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsProceduresV3TestSelectorPostRaw(requestParameters: ApiDeployHooksGithubActionsProceduresV3TestSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockGithubActionsDeployHookTestRunResponse>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksGithubActionsProceduresV3TestSelectorPost().'
            );
        }

        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiDeployHooksGithubActionsProceduresV3TestSelectorPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/github-actions/procedures/v3/test/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockGithubActionsDeployHookTestRunRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockGithubActionsDeployHookTestRunResponseFromJSON(jsonValue));
    }

    /**
     * Run a GitHub Action to simulate a GithubActionsDeployHook
     * Test a GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsProceduresV3TestSelectorPost(requestParameters: ApiDeployHooksGithubActionsProceduresV3TestSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockGithubActionsDeployHookTestRunResponse> {
        const response = await this.apiDeployHooksGithubActionsProceduresV3TestSelectorPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List GithubActionsDeployHooks matching a filter.
     * List GithubActionsDeployHooks matching a filter
     */
    async apiDeployHooksGithubActionsV3GetRaw(requestParameters: ApiDeployHooksGithubActionsV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockGithubActionsDeployHookV3>>> {
        const queryParameters: any = {};

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['githubActionsDefaultRef'] != null) {
            queryParameters['githubActionsDefaultRef'] = requestParameters['githubActionsDefaultRef'];
        }

        if (requestParameters['githubActionsOwner'] != null) {
            queryParameters['githubActionsOwner'] = requestParameters['githubActionsOwner'];
        }

        if (requestParameters['githubActionsRefBehavior'] != null) {
            queryParameters['githubActionsRefBehavior'] = requestParameters['githubActionsRefBehavior'];
        }

        if (requestParameters['githubActionsRepo'] != null) {
            queryParameters['githubActionsRepo'] = requestParameters['githubActionsRepo'];
        }

        if (requestParameters['githubActionsWorkflowPath'] != null) {
            queryParameters['githubActionsWorkflowPath'] = requestParameters['githubActionsWorkflowPath'];
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['onChartRelease'] != null) {
            queryParameters['onChartRelease'] = requestParameters['onChartRelease'];
        }

        if (requestParameters['onEnvironment'] != null) {
            queryParameters['onEnvironment'] = requestParameters['onEnvironment'];
        }

        if (requestParameters['onFailure'] != null) {
            queryParameters['onFailure'] = requestParameters['onFailure'];
        }

        if (requestParameters['onSuccess'] != null) {
            queryParameters['onSuccess'] = requestParameters['onSuccess'];
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
            path: `/api/deploy-hooks/github-actions/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockGithubActionsDeployHookV3FromJSON));
    }

    /**
     * List GithubActionsDeployHooks matching a filter.
     * List GithubActionsDeployHooks matching a filter
     */
    async apiDeployHooksGithubActionsV3Get(requestParameters: ApiDeployHooksGithubActionsV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockGithubActionsDeployHookV3>> {
        const response = await this.apiDeployHooksGithubActionsV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a GithubActionsDeployHook.
     * Create a GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3PostRaw(requestParameters: ApiDeployHooksGithubActionsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockGithubActionsDeployHookV3>> {
        if (requestParameters['githubActionsDeployHook'] == null) {
            throw new runtime.RequiredError(
                'githubActionsDeployHook',
                'Required parameter "githubActionsDeployHook" was null or undefined when calling apiDeployHooksGithubActionsV3Post().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/github-actions/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockGithubActionsDeployHookV3CreateToJSON(requestParameters['githubActionsDeployHook']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockGithubActionsDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Create a GithubActionsDeployHook.
     * Create a GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3Post(requestParameters: ApiDeployHooksGithubActionsV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockGithubActionsDeployHookV3> {
        const response = await this.apiDeployHooksGithubActionsV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual GithubActionsDeployHook by its ID.
     * Delete an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorDeleteRaw(requestParameters: ApiDeployHooksGithubActionsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockGithubActionsDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksGithubActionsV3SelectorDelete().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/deploy-hooks/github-actions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockGithubActionsDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual GithubActionsDeployHook by its ID.
     * Delete an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorDelete(requestParameters: ApiDeployHooksGithubActionsV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockGithubActionsDeployHookV3> {
        const response = await this.apiDeployHooksGithubActionsV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual GithubActionsDeployHook by its ID.
     * Get an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorGetRaw(requestParameters: ApiDeployHooksGithubActionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockGithubActionsDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksGithubActionsV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/deploy-hooks/github-actions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockGithubActionsDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Get an individual GithubActionsDeployHook by its ID.
     * Get an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorGet(requestParameters: ApiDeployHooksGithubActionsV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockGithubActionsDeployHookV3> {
        const response = await this.apiDeployHooksGithubActionsV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual GithubActionsDeployHook by its ID.
     * Edit an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorPatchRaw(requestParameters: ApiDeployHooksGithubActionsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockGithubActionsDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksGithubActionsV3SelectorPatch().'
            );
        }

        if (requestParameters['githubActionsDeployHook'] == null) {
            throw new runtime.RequiredError(
                'githubActionsDeployHook',
                'Required parameter "githubActionsDeployHook" was null or undefined when calling apiDeployHooksGithubActionsV3SelectorPatch().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/github-actions/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockGithubActionsDeployHookV3EditToJSON(requestParameters['githubActionsDeployHook']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockGithubActionsDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual GithubActionsDeployHook by its ID.
     * Edit an individual GithubActionsDeployHook
     */
    async apiDeployHooksGithubActionsV3SelectorPatch(requestParameters: ApiDeployHooksGithubActionsV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockGithubActionsDeployHookV3> {
        const response = await this.apiDeployHooksGithubActionsV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Send a Slack message to simulate a SlackDeployHook
     * Test a SlackDeployHook
     */
    async apiDeployHooksSlackProceduresV3TestSelectorPostRaw(requestParameters: ApiDeployHooksSlackProceduresV3TestSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockSlackDeployHookTestRunResponse>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksSlackProceduresV3TestSelectorPost().'
            );
        }

        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling apiDeployHooksSlackProceduresV3TestSelectorPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/slack/procedures/v3/test/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockSlackDeployHookTestRunRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockSlackDeployHookTestRunResponseFromJSON(jsonValue));
    }

    /**
     * Send a Slack message to simulate a SlackDeployHook
     * Test a SlackDeployHook
     */
    async apiDeployHooksSlackProceduresV3TestSelectorPost(requestParameters: ApiDeployHooksSlackProceduresV3TestSelectorPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockSlackDeployHookTestRunResponse> {
        const response = await this.apiDeployHooksSlackProceduresV3TestSelectorPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List SlackDeployHooks matching a filter.
     * List SlackDeployHooks matching a filter
     */
    async apiDeployHooksSlackV3GetRaw(requestParameters: ApiDeployHooksSlackV3GetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<SherlockSlackDeployHookV3>>> {
        const queryParameters: any = {};

        if (requestParameters['createdAt'] != null) {
            queryParameters['createdAt'] = (requestParameters['createdAt'] as any).toISOString();
        }

        if (requestParameters['id'] != null) {
            queryParameters['id'] = requestParameters['id'];
        }

        if (requestParameters['mentionPeople'] != null) {
            queryParameters['mentionPeople'] = requestParameters['mentionPeople'];
        }

        if (requestParameters['onChartRelease'] != null) {
            queryParameters['onChartRelease'] = requestParameters['onChartRelease'];
        }

        if (requestParameters['onEnvironment'] != null) {
            queryParameters['onEnvironment'] = requestParameters['onEnvironment'];
        }

        if (requestParameters['onFailure'] != null) {
            queryParameters['onFailure'] = requestParameters['onFailure'];
        }

        if (requestParameters['onSuccess'] != null) {
            queryParameters['onSuccess'] = requestParameters['onSuccess'];
        }

        if (requestParameters['slackChannel'] != null) {
            queryParameters['slackChannel'] = requestParameters['slackChannel'];
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
            path: `/api/deploy-hooks/slack/v3`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SherlockSlackDeployHookV3FromJSON));
    }

    /**
     * List SlackDeployHooks matching a filter.
     * List SlackDeployHooks matching a filter
     */
    async apiDeployHooksSlackV3Get(requestParameters: ApiDeployHooksSlackV3GetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<SherlockSlackDeployHookV3>> {
        const response = await this.apiDeployHooksSlackV3GetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a SlackDeployHook.
     * Create a SlackDeployHook
     */
    async apiDeployHooksSlackV3PostRaw(requestParameters: ApiDeployHooksSlackV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockSlackDeployHookV3>> {
        if (requestParameters['slackDeployHook'] == null) {
            throw new runtime.RequiredError(
                'slackDeployHook',
                'Required parameter "slackDeployHook" was null or undefined when calling apiDeployHooksSlackV3Post().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/slack/v3`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockSlackDeployHookV3CreateToJSON(requestParameters['slackDeployHook']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockSlackDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Create a SlackDeployHook.
     * Create a SlackDeployHook
     */
    async apiDeployHooksSlackV3Post(requestParameters: ApiDeployHooksSlackV3PostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockSlackDeployHookV3> {
        const response = await this.apiDeployHooksSlackV3PostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete an individual SlackDeployHook by its ID.
     * Delete an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorDeleteRaw(requestParameters: ApiDeployHooksSlackV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockSlackDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksSlackV3SelectorDelete().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/deploy-hooks/slack/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockSlackDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Delete an individual SlackDeployHook by its ID.
     * Delete an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorDelete(requestParameters: ApiDeployHooksSlackV3SelectorDeleteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockSlackDeployHookV3> {
        const response = await this.apiDeployHooksSlackV3SelectorDeleteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an individual SlackDeployHook by its ID.
     * Get an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorGetRaw(requestParameters: ApiDeployHooksSlackV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockSlackDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksSlackV3SelectorGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/deploy-hooks/slack/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockSlackDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Get an individual SlackDeployHook by its ID.
     * Get an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorGet(requestParameters: ApiDeployHooksSlackV3SelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockSlackDeployHookV3> {
        const response = await this.apiDeployHooksSlackV3SelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Edit an individual SlackDeployHook by its ID.
     * Edit an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorPatchRaw(requestParameters: ApiDeployHooksSlackV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SherlockSlackDeployHookV3>> {
        if (requestParameters['selector'] == null) {
            throw new runtime.RequiredError(
                'selector',
                'Required parameter "selector" was null or undefined when calling apiDeployHooksSlackV3SelectorPatch().'
            );
        }

        if (requestParameters['slackDeployHook'] == null) {
            throw new runtime.RequiredError(
                'slackDeployHook',
                'Required parameter "slackDeployHook" was null or undefined when calling apiDeployHooksSlackV3SelectorPatch().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/deploy-hooks/slack/v3/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters['selector']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SherlockSlackDeployHookV3EditToJSON(requestParameters['slackDeployHook']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SherlockSlackDeployHookV3FromJSON(jsonValue));
    }

    /**
     * Edit an individual SlackDeployHook by its ID.
     * Edit an individual SlackDeployHook
     */
    async apiDeployHooksSlackV3SelectorPatch(requestParameters: ApiDeployHooksSlackV3SelectorPatchRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SherlockSlackDeployHookV3> {
        const response = await this.apiDeployHooksSlackV3SelectorPatchRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const ApiDeployHooksGithubActionsV3GetGithubActionsRefBehaviorEnum = {
    AlwaysUseDefaultRef: 'always-use-default-ref',
    UseAppVersionAsRef: 'use-app-version-as-ref',
    UseAppVersionCommitAsRef: 'use-app-version-commit-as-ref'
} as const;
export type ApiDeployHooksGithubActionsV3GetGithubActionsRefBehaviorEnum = typeof ApiDeployHooksGithubActionsV3GetGithubActionsRefBehaviorEnum[keyof typeof ApiDeployHooksGithubActionsV3GetGithubActionsRefBehaviorEnum];
