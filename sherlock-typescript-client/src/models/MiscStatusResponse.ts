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
 * @interface MiscStatusResponse
 */
export interface MiscStatusResponse {
    /**
     * 
     * @type {boolean}
     * @memberof MiscStatusResponse
     */
    ok?: boolean;
}

/**
 * Check if a given object implements the MiscStatusResponse interface.
 */
export function instanceOfMiscStatusResponse(value: object): value is MiscStatusResponse {
    return true;
}

export function MiscStatusResponseFromJSON(json: any): MiscStatusResponse {
    return MiscStatusResponseFromJSONTyped(json, false);
}

export function MiscStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MiscStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'ok': json['ok'] == null ? undefined : json['ok'],
    };
}

export function MiscStatusResponseToJSON(json: any): MiscStatusResponse {
    return MiscStatusResponseToJSONTyped(json, false);
}

export function MiscStatusResponseToJSONTyped(value?: MiscStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ok': value['ok'],
    };
}

