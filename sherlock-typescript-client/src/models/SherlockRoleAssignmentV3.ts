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
 * @interface SherlockRoleAssignmentV3
 */
export interface SherlockRoleAssignmentV3 {
    /**
     * 
     * @type {Date}
     * @memberof SherlockRoleAssignmentV3
     */
    expiresAt?: Date;
    /**
     * 
     * @type {object}
     * @memberof SherlockRoleAssignmentV3
     */
    roleInfo?: object;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleAssignmentV3
     */
    suspended?: boolean;
    /**
     * 
     * @type {object}
     * @memberof SherlockRoleAssignmentV3
     */
    userInfo?: object;
}

/**
 * Check if a given object implements the SherlockRoleAssignmentV3 interface.
 */
export function instanceOfSherlockRoleAssignmentV3(value: object): value is SherlockRoleAssignmentV3 {
    return true;
}

export function SherlockRoleAssignmentV3FromJSON(json: any): SherlockRoleAssignmentV3 {
    return SherlockRoleAssignmentV3FromJSONTyped(json, false);
}

export function SherlockRoleAssignmentV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockRoleAssignmentV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'expiresAt': json['expiresAt'] == null ? undefined : (new Date(json['expiresAt'])),
        'roleInfo': json['roleInfo'] == null ? undefined : json['roleInfo'],
        'suspended': json['suspended'] == null ? undefined : json['suspended'],
        'userInfo': json['userInfo'] == null ? undefined : json['userInfo'],
    };
}

export function SherlockRoleAssignmentV3ToJSON(value?: SherlockRoleAssignmentV3 | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'expiresAt': value['expiresAt'] == null ? undefined : ((value['expiresAt']).toISOString()),
        'roleInfo': value['roleInfo'],
        'suspended': value['suspended'],
        'userInfo': value['userInfo'],
    };
}

