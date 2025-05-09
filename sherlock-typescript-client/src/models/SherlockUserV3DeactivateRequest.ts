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
 * @interface SherlockUserV3DeactivateRequest
 */
export interface SherlockUserV3DeactivateRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof SherlockUserV3DeactivateRequest
     */
    suspendEmailHandlesAcrossGoogleWorkspaceDomains?: Array<string>;
    /**
     * Domain of UserEmails that can be swapped out for the domains in SuspendEmailHandlesAcrossGoogleWorkspaceDomains
     * @type {string}
     * @memberof SherlockUserV3DeactivateRequest
     */
    userEmailHomeDomain?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SherlockUserV3DeactivateRequest
     */
    userEmails?: Array<string>;
}

/**
 * Check if a given object implements the SherlockUserV3DeactivateRequest interface.
 */
export function instanceOfSherlockUserV3DeactivateRequest(value: object): value is SherlockUserV3DeactivateRequest {
    return true;
}

export function SherlockUserV3DeactivateRequestFromJSON(json: any): SherlockUserV3DeactivateRequest {
    return SherlockUserV3DeactivateRequestFromJSONTyped(json, false);
}

export function SherlockUserV3DeactivateRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockUserV3DeactivateRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'suspendEmailHandlesAcrossGoogleWorkspaceDomains': json['suspendEmailHandlesAcrossGoogleWorkspaceDomains'] == null ? undefined : json['suspendEmailHandlesAcrossGoogleWorkspaceDomains'],
        'userEmailHomeDomain': json['userEmailHomeDomain'] == null ? undefined : json['userEmailHomeDomain'],
        'userEmails': json['userEmails'] == null ? undefined : json['userEmails'],
    };
}

export function SherlockUserV3DeactivateRequestToJSON(json: any): SherlockUserV3DeactivateRequest {
    return SherlockUserV3DeactivateRequestToJSONTyped(json, false);
}

export function SherlockUserV3DeactivateRequestToJSONTyped(value?: SherlockUserV3DeactivateRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'suspendEmailHandlesAcrossGoogleWorkspaceDomains': value['suspendEmailHandlesAcrossGoogleWorkspaceDomains'],
        'userEmailHomeDomain': value['userEmailHomeDomain'],
        'userEmails': value['userEmails'],
    };
}

