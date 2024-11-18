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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface SherlockGithubActionsDeployHookTestRunResponse
 */
export interface SherlockGithubActionsDeployHookTestRunResponse {
    /**
     * 
     * @type {boolean}
     * @memberof SherlockGithubActionsDeployHookTestRunResponse
     */
    ok?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockGithubActionsDeployHookTestRunResponse
     */
    url?: string;
}

/**
 * Check if a given object implements the SherlockGithubActionsDeployHookTestRunResponse interface.
 */
export function instanceOfSherlockGithubActionsDeployHookTestRunResponse(value: object): value is SherlockGithubActionsDeployHookTestRunResponse {
    return true;
}

export function SherlockGithubActionsDeployHookTestRunResponseFromJSON(json: any): SherlockGithubActionsDeployHookTestRunResponse {
    return SherlockGithubActionsDeployHookTestRunResponseFromJSONTyped(json, false);
}

export function SherlockGithubActionsDeployHookTestRunResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockGithubActionsDeployHookTestRunResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'ok': json['ok'] == null ? undefined : json['ok'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function SherlockGithubActionsDeployHookTestRunResponseToJSON(json: any): SherlockGithubActionsDeployHookTestRunResponse {
    return SherlockGithubActionsDeployHookTestRunResponseToJSONTyped(json, false);
}

export function SherlockGithubActionsDeployHookTestRunResponseToJSONTyped(value?: SherlockGithubActionsDeployHookTestRunResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ok': value['ok'],
        'url': value['url'],
    };
}

