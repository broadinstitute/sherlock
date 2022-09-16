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
import type { AuthUser } from './AuthUser';
import {
    AuthUserFromJSON,
    AuthUserFromJSONTyped,
    AuthUserToJSON,
} from './AuthUser';

/**
 * 
 * @export
 * @interface MiscMyUserResponse
 */
export interface MiscMyUserResponse {
    /**
     * 
     * @type {string}
     * @memberof MiscMyUserResponse
     */
    email?: string;
    /**
     * 
     * @type {AuthUser}
     * @memberof MiscMyUserResponse
     */
    rawInfo?: AuthUser;
    /**
     * 
     * @type {string}
     * @memberof MiscMyUserResponse
     */
    suitability?: string;
}

/**
 * Check if a given object implements the MiscMyUserResponse interface.
 */
export function instanceOfMiscMyUserResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MiscMyUserResponseFromJSON(json: any): MiscMyUserResponse {
    return MiscMyUserResponseFromJSONTyped(json, false);
}

export function MiscMyUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MiscMyUserResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'email') ? undefined : json['email'],
        'rawInfo': !exists(json, 'rawInfo') ? undefined : AuthUserFromJSON(json['rawInfo']),
        'suitability': !exists(json, 'suitability') ? undefined : json['suitability'],
    };
}

export function MiscMyUserResponseToJSON(value?: MiscMyUserResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': value.email,
        'rawInfo': AuthUserToJSON(value.rawInfo),
        'suitability': value.suitability,
    };
}

