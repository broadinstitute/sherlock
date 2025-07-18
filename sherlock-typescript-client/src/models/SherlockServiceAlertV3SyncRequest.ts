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
 * @interface SherlockServiceAlertV3SyncRequest
 */
export interface SherlockServiceAlertV3SyncRequest {
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3SyncRequest
     */
    onEnvironment?: string;
}

/**
 * Check if a given object implements the SherlockServiceAlertV3SyncRequest interface.
 */
export function instanceOfSherlockServiceAlertV3SyncRequest(value: object): value is SherlockServiceAlertV3SyncRequest {
    return true;
}

export function SherlockServiceAlertV3SyncRequestFromJSON(json: any): SherlockServiceAlertV3SyncRequest {
    return SherlockServiceAlertV3SyncRequestFromJSONTyped(json, false);
}

export function SherlockServiceAlertV3SyncRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockServiceAlertV3SyncRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'onEnvironment': json['onEnvironment'] == null ? undefined : json['onEnvironment'],
    };
}

export function SherlockServiceAlertV3SyncRequestToJSON(json: any): SherlockServiceAlertV3SyncRequest {
    return SherlockServiceAlertV3SyncRequestToJSONTyped(json, false);
}

export function SherlockServiceAlertV3SyncRequestToJSONTyped(value?: SherlockServiceAlertV3SyncRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'onEnvironment': value['onEnvironment'],
    };
}

