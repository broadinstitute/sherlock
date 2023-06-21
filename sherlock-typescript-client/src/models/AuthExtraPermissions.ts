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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface AuthExtraPermissions
 */
export interface AuthExtraPermissions {
    /**
     * 
     * @type {boolean}
     * @memberof AuthExtraPermissions
     */
    suitable?: boolean;
}

/**
 * Check if a given object implements the AuthExtraPermissions interface.
 */
export function instanceOfAuthExtraPermissions(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AuthExtraPermissionsFromJSON(json: any): AuthExtraPermissions {
    return AuthExtraPermissionsFromJSONTyped(json, false);
}

export function AuthExtraPermissionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuthExtraPermissions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'suitable': !exists(json, 'suitable') ? undefined : json['suitable'],
    };
}

export function AuthExtraPermissionsToJSON(value?: AuthExtraPermissions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'suitable': value.suitable,
    };
}
