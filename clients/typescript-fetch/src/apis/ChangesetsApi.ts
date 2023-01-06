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
  V2controllersChangeset,
  V2controllersChangesetPlanRequest,
  V2controllersCreatableChangeset,
} from '../models';
import {
    ErrorsErrorResponseFromJSON,
    ErrorsErrorResponseToJSON,
    V2controllersChangesetFromJSON,
    V2controllersChangesetToJSON,
    V2controllersChangesetPlanRequestFromJSON,
    V2controllersChangesetPlanRequestToJSON,
    V2controllersCreatableChangesetFromJSON,
    V2controllersCreatableChangesetToJSON,
} from '../models';

export interface ApiV2ChangesetsGetRequest {
    appliedAt?: string;
    chartRelease?: string;
    createdAt?: Date;
    fromAppVersionBranch?: string;
    fromAppVersionCommit?: string;
    fromAppVersionExact?: string;
    fromAppVersionFollowChartRelease?: string;
    fromAppVersionReference?: string;
    fromAppVersionResolver?: string;
    fromChartVersionExact?: string;
    fromChartVersionFollowChartRelease?: string;
    fromChartVersionReference?: string;
    fromChartVersionResolver?: string;
    fromFirecloudDevelopRef?: string;
    fromHelmfileRef?: string;
    fromResolvedAt?: string;
    id?: number;
    supersededAt?: string;
    toAppVersionBranch?: string;
    toAppVersionCommit?: string;
    toAppVersionExact?: string;
    toAppVersionFollowChartRelease?: string;
    toAppVersionReference?: string;
    toAppVersionResolver?: string;
    toChartVersionExact?: string;
    toChartVersionFollowChartRelease?: string;
    toChartVersionReference?: string;
    toChartVersionResolver?: string;
    toFirecloudDevelopRef?: string;
    toHelmfileRef?: string;
    toResolvedAt?: string;
    updatedAt?: Date;
    limit?: number;
}

export interface ApiV2ChangesetsPostRequest {
    changeset: V2controllersCreatableChangeset;
}

export interface ApiV2ChangesetsSelectorGetRequest {
    selector: string;
}

export interface ApiV2ProceduresChangesetsApplyPostRequest {
    applyRequest: Array<string>;
}

export interface ApiV2ProceduresChangesetsPlanAndApplyPostRequest {
    changesetPlanRequest: V2controllersChangesetPlanRequest;
}

export interface ApiV2ProceduresChangesetsPlanPostRequest {
    changesetPlanRequest: V2controllersChangesetPlanRequest;
}

export interface ApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGetRequest {
    selector: string;
    offset?: number;
    limit?: number;
}

export interface ApiV2SelectorsChangesetsSelectorGetRequest {
    selector: string;
}

/**
 * 
 */
export class ChangesetsApi extends runtime.BaseAPI {

    /**
     * List existing Changeset entries, ordered by most recently updated.
     * List Changeset entries
     */
    async apiV2ChangesetsGetRaw(requestParameters: ApiV2ChangesetsGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChangeset>>> {
        const queryParameters: any = {};

        if (requestParameters.appliedAt !== undefined) {
            queryParameters['appliedAt'] = requestParameters.appliedAt;
        }

        if (requestParameters.chartRelease !== undefined) {
            queryParameters['chartRelease'] = requestParameters.chartRelease;
        }

        if (requestParameters.createdAt !== undefined) {
            queryParameters['createdAt'] = (requestParameters.createdAt as any).toISOString();
        }

        if (requestParameters.fromAppVersionBranch !== undefined) {
            queryParameters['fromAppVersionBranch'] = requestParameters.fromAppVersionBranch;
        }

        if (requestParameters.fromAppVersionCommit !== undefined) {
            queryParameters['fromAppVersionCommit'] = requestParameters.fromAppVersionCommit;
        }

        if (requestParameters.fromAppVersionExact !== undefined) {
            queryParameters['fromAppVersionExact'] = requestParameters.fromAppVersionExact;
        }

        if (requestParameters.fromAppVersionFollowChartRelease !== undefined) {
            queryParameters['fromAppVersionFollowChartRelease'] = requestParameters.fromAppVersionFollowChartRelease;
        }

        if (requestParameters.fromAppVersionReference !== undefined) {
            queryParameters['fromAppVersionReference'] = requestParameters.fromAppVersionReference;
        }

        if (requestParameters.fromAppVersionResolver !== undefined) {
            queryParameters['fromAppVersionResolver'] = requestParameters.fromAppVersionResolver;
        }

        if (requestParameters.fromChartVersionExact !== undefined) {
            queryParameters['fromChartVersionExact'] = requestParameters.fromChartVersionExact;
        }

        if (requestParameters.fromChartVersionFollowChartRelease !== undefined) {
            queryParameters['fromChartVersionFollowChartRelease'] = requestParameters.fromChartVersionFollowChartRelease;
        }

        if (requestParameters.fromChartVersionReference !== undefined) {
            queryParameters['fromChartVersionReference'] = requestParameters.fromChartVersionReference;
        }

        if (requestParameters.fromChartVersionResolver !== undefined) {
            queryParameters['fromChartVersionResolver'] = requestParameters.fromChartVersionResolver;
        }

        if (requestParameters.fromFirecloudDevelopRef !== undefined) {
            queryParameters['fromFirecloudDevelopRef'] = requestParameters.fromFirecloudDevelopRef;
        }

        if (requestParameters.fromHelmfileRef !== undefined) {
            queryParameters['fromHelmfileRef'] = requestParameters.fromHelmfileRef;
        }

        if (requestParameters.fromResolvedAt !== undefined) {
            queryParameters['fromResolvedAt'] = requestParameters.fromResolvedAt;
        }

        if (requestParameters.id !== undefined) {
            queryParameters['id'] = requestParameters.id;
        }

        if (requestParameters.supersededAt !== undefined) {
            queryParameters['supersededAt'] = requestParameters.supersededAt;
        }

        if (requestParameters.toAppVersionBranch !== undefined) {
            queryParameters['toAppVersionBranch'] = requestParameters.toAppVersionBranch;
        }

        if (requestParameters.toAppVersionCommit !== undefined) {
            queryParameters['toAppVersionCommit'] = requestParameters.toAppVersionCommit;
        }

        if (requestParameters.toAppVersionExact !== undefined) {
            queryParameters['toAppVersionExact'] = requestParameters.toAppVersionExact;
        }

        if (requestParameters.toAppVersionFollowChartRelease !== undefined) {
            queryParameters['toAppVersionFollowChartRelease'] = requestParameters.toAppVersionFollowChartRelease;
        }

        if (requestParameters.toAppVersionReference !== undefined) {
            queryParameters['toAppVersionReference'] = requestParameters.toAppVersionReference;
        }

        if (requestParameters.toAppVersionResolver !== undefined) {
            queryParameters['toAppVersionResolver'] = requestParameters.toAppVersionResolver;
        }

        if (requestParameters.toChartVersionExact !== undefined) {
            queryParameters['toChartVersionExact'] = requestParameters.toChartVersionExact;
        }

        if (requestParameters.toChartVersionFollowChartRelease !== undefined) {
            queryParameters['toChartVersionFollowChartRelease'] = requestParameters.toChartVersionFollowChartRelease;
        }

        if (requestParameters.toChartVersionReference !== undefined) {
            queryParameters['toChartVersionReference'] = requestParameters.toChartVersionReference;
        }

        if (requestParameters.toChartVersionResolver !== undefined) {
            queryParameters['toChartVersionResolver'] = requestParameters.toChartVersionResolver;
        }

        if (requestParameters.toFirecloudDevelopRef !== undefined) {
            queryParameters['toFirecloudDevelopRef'] = requestParameters.toFirecloudDevelopRef;
        }

        if (requestParameters.toHelmfileRef !== undefined) {
            queryParameters['toHelmfileRef'] = requestParameters.toHelmfileRef;
        }

        if (requestParameters.toResolvedAt !== undefined) {
            queryParameters['toResolvedAt'] = requestParameters.toResolvedAt;
        }

        if (requestParameters.updatedAt !== undefined) {
            queryParameters['updatedAt'] = (requestParameters.updatedAt as any).toISOString();
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/changesets`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChangesetFromJSON));
    }

    /**
     * List existing Changeset entries, ordered by most recently updated.
     * List Changeset entries
     */
    async apiV2ChangesetsGet(requestParameters: ApiV2ChangesetsGetRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChangeset>> {
        const response = await this.apiV2ChangesetsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create a new Changeset entry. Note that fields are immutable after creation. You\'ll likely want to use the plan endpoint instead, which conditionally creates a Changeset based on there actually being a version diff.
     * Create a new Changeset entry
     */
    async apiV2ChangesetsPostRaw(requestParameters: ApiV2ChangesetsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChangeset>> {
        if (requestParameters.changeset === null || requestParameters.changeset === undefined) {
            throw new runtime.RequiredError('changeset','Required parameter requestParameters.changeset was null or undefined when calling apiV2ChangesetsPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/changesets`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersCreatableChangesetToJSON(requestParameters.changeset),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChangesetFromJSON(jsonValue));
    }

    /**
     * Create a new Changeset entry. Note that fields are immutable after creation. You\'ll likely want to use the plan endpoint instead, which conditionally creates a Changeset based on there actually being a version diff.
     * Create a new Changeset entry
     */
    async apiV2ChangesetsPost(requestParameters: ApiV2ChangesetsPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChangeset> {
        const response = await this.apiV2ChangesetsPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get an existing Changeset entry via its \"selector\"--its numeric ID.
     * Get a Changeset entry
     */
    async apiV2ChangesetsSelectorGetRaw(requestParameters: ApiV2ChangesetsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<V2controllersChangeset>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ChangesetsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/changesets/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => V2controllersChangesetFromJSON(jsonValue));
    }

    /**
     * Get an existing Changeset entry via its \"selector\"--its numeric ID.
     * Get a Changeset entry
     */
    async apiV2ChangesetsSelectorGet(requestParameters: ApiV2ChangesetsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<V2controllersChangeset> {
        const response = await this.apiV2ChangesetsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded. Multiple Changesets can be specified simply by passing multiple IDs in the list.
     * Apply previously planned version changes to Chart Releases
     */
    async apiV2ProceduresChangesetsApplyPostRaw(requestParameters: ApiV2ProceduresChangesetsApplyPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChangeset>>> {
        if (requestParameters.applyRequest === null || requestParameters.applyRequest === undefined) {
            throw new runtime.RequiredError('applyRequest','Required parameter requestParameters.applyRequest was null or undefined when calling apiV2ProceduresChangesetsApplyPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/procedures/changesets/apply`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.applyRequest,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChangesetFromJSON));
    }

    /**
     * Looks up and applies previously-planned version diffs given by the ID. Other stored plans against the same Chart Releases are marked as superseded. Multiple Changesets can be specified simply by passing multiple IDs in the list.
     * Apply previously planned version changes to Chart Releases
     */
    async apiV2ProceduresChangesetsApplyPost(requestParameters: ApiV2ProceduresChangesetsApplyPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChangeset>> {
        const response = await this.apiV2ProceduresChangesetsApplyPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Like the plan and apply endpoints immediately in sequence.
     * Plan and apply version changes in one step
     */
    async apiV2ProceduresChangesetsPlanAndApplyPostRaw(requestParameters: ApiV2ProceduresChangesetsPlanAndApplyPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChangeset>>> {
        if (requestParameters.changesetPlanRequest === null || requestParameters.changesetPlanRequest === undefined) {
            throw new runtime.RequiredError('changesetPlanRequest','Required parameter requestParameters.changesetPlanRequest was null or undefined when calling apiV2ProceduresChangesetsPlanAndApplyPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/procedures/changesets/plan-and-apply`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersChangesetPlanRequestToJSON(requestParameters.changesetPlanRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChangesetFromJSON));
    }

    /**
     * Like the plan and apply endpoints immediately in sequence.
     * Plan and apply version changes in one step
     */
    async apiV2ProceduresChangesetsPlanAndApplyPost(requestParameters: ApiV2ProceduresChangesetsPlanAndApplyPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChangeset>> {
        const response = await this.apiV2ProceduresChangesetsPlanAndApplyPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Refreshes and calculates version diffs for Chart Releases. If there\'s a diff, the plan is stored and returned so it can be applied later. Multiple Chart Releases can be specified--as can groups of Chart Releases from multiple Environments.
     * Plan--but do not apply--version changes to Chart Releases
     */
    async apiV2ProceduresChangesetsPlanPostRaw(requestParameters: ApiV2ProceduresChangesetsPlanPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChangeset>>> {
        if (requestParameters.changesetPlanRequest === null || requestParameters.changesetPlanRequest === undefined) {
            throw new runtime.RequiredError('changesetPlanRequest','Required parameter requestParameters.changesetPlanRequest was null or undefined when calling apiV2ProceduresChangesetsPlanPost.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/api/v2/procedures/changesets/plan`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V2controllersChangesetPlanRequestToJSON(requestParameters.changesetPlanRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChangesetFromJSON));
    }

    /**
     * Refreshes and calculates version diffs for Chart Releases. If there\'s a diff, the plan is stored and returned so it can be applied later. Multiple Chart Releases can be specified--as can groups of Chart Releases from multiple Environments.
     * Plan--but do not apply--version changes to Chart Releases
     */
    async apiV2ProceduresChangesetsPlanPost(requestParameters: ApiV2ProceduresChangesetsPlanPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChangeset>> {
        const response = await this.apiV2ProceduresChangesetsPlanPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List existing applied Changesets for a particular Chart Release, ordered by most recently applied.
     * List applied Changesets for a Chart Release
     */
    async apiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGetRaw(requestParameters: ApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<V2controllersChangeset>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGet.');
        }

        const queryParameters: any = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/procedures/changesets/query-applied-for-chart-release/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(V2controllersChangesetFromJSON));
    }

    /**
     * List existing applied Changesets for a particular Chart Release, ordered by most recently applied.
     * List applied Changesets for a Chart Release
     */
    async apiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGet(requestParameters: ApiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<V2controllersChangeset>> {
        const response = await this.apiV2ProceduresChangesetsQueryAppliedForChartReleaseSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Validate a given Changeset selector and provide any other selectors that would match the same Changeset.
     * List Changeset selectors
     */
    async apiV2SelectorsChangesetsSelectorGetRaw(requestParameters: ApiV2SelectorsChangesetsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<string>>> {
        if (requestParameters.selector === null || requestParameters.selector === undefined) {
            throw new runtime.RequiredError('selector','Required parameter requestParameters.selector was null or undefined when calling apiV2SelectorsChangesetsSelectorGet.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/selectors/changesets/{selector}`.replace(`{${"selector"}}`, encodeURIComponent(String(requestParameters.selector))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * Validate a given Changeset selector and provide any other selectors that would match the same Changeset.
     * List Changeset selectors
     */
    async apiV2SelectorsChangesetsSelectorGet(requestParameters: ApiV2SelectorsChangesetsSelectorGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<string>> {
        const response = await this.apiV2SelectorsChangesetsSelectorGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
