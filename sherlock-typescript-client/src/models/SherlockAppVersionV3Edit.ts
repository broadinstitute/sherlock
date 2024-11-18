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
 * @interface SherlockAppVersionV3Edit
 */
export interface SherlockAppVersionV3Edit {
    /**
     * Generally the Git commit message
     * @type {string}
     * @memberof SherlockAppVersionV3Edit
     */
    description?: string;
}

/**
 * Check if a given object implements the SherlockAppVersionV3Edit interface.
 */
export function instanceOfSherlockAppVersionV3Edit(value: object): value is SherlockAppVersionV3Edit {
    return true;
}

export function SherlockAppVersionV3EditFromJSON(json: any): SherlockAppVersionV3Edit {
    return SherlockAppVersionV3EditFromJSONTyped(json, false);
}

export function SherlockAppVersionV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockAppVersionV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'description': json['description'] == null ? undefined : json['description'],
    };
}

export function SherlockAppVersionV3EditToJSON(json: any): SherlockAppVersionV3Edit {
    return SherlockAppVersionV3EditToJSONTyped(json, false);
}

export function SherlockAppVersionV3EditToJSONTyped(value?: SherlockAppVersionV3Edit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'description': value['description'],
    };
}

