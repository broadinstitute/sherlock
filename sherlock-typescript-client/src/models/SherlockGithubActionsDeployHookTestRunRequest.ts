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
 * @interface SherlockGithubActionsDeployHookTestRunRequest
 */
export interface SherlockGithubActionsDeployHookTestRunRequest {
    /**
     * Required, whether to fully run the GHA
     * @type {boolean}
     * @memberof SherlockGithubActionsDeployHookTestRunRequest
     */
    execute?: boolean;
}

/**
 * Check if a given object implements the SherlockGithubActionsDeployHookTestRunRequest interface.
 */
export function instanceOfSherlockGithubActionsDeployHookTestRunRequest(value: object): value is SherlockGithubActionsDeployHookTestRunRequest {
    return true;
}

export function SherlockGithubActionsDeployHookTestRunRequestFromJSON(json: any): SherlockGithubActionsDeployHookTestRunRequest {
    return SherlockGithubActionsDeployHookTestRunRequestFromJSONTyped(json, false);
}

export function SherlockGithubActionsDeployHookTestRunRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockGithubActionsDeployHookTestRunRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'execute': json['execute'] == null ? undefined : json['execute'],
    };
}

export function SherlockGithubActionsDeployHookTestRunRequestToJSON(json: any): SherlockGithubActionsDeployHookTestRunRequest {
    return SherlockGithubActionsDeployHookTestRunRequestToJSONTyped(json, false);
}

export function SherlockGithubActionsDeployHookTestRunRequestToJSONTyped(value?: SherlockGithubActionsDeployHookTestRunRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'execute': value['execute'],
    };
}

