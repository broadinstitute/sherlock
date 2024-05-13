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
 * @interface SherlockRoleV3Edit
 */
export interface SherlockRoleV3Edit {
    /**
     * 
     * @type {number}
     * @memberof SherlockRoleV3Edit
     */
    canBeGlassBrokenByRole?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3Edit
     */
    defaultGlassBreakDuration?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3Edit
     */
    grantsDevAzureGroup?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3Edit
     */
    grantsDevFirecloudGroup?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3Edit
     */
    grantsSherlockSuperAdmin?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3Edit
     */
    name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3Edit
     */
    suspendNonSuitableUsers?: boolean;
}

/**
 * Check if a given object implements the SherlockRoleV3Edit interface.
 */
export function instanceOfSherlockRoleV3Edit(value: object): boolean {
    return true;
}

export function SherlockRoleV3EditFromJSON(json: any): SherlockRoleV3Edit {
    return SherlockRoleV3EditFromJSONTyped(json, false);
}

export function SherlockRoleV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockRoleV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'canBeGlassBrokenByRole': json['canBeGlassBrokenByRole'] == null ? undefined : json['canBeGlassBrokenByRole'],
        'defaultGlassBreakDuration': json['defaultGlassBreakDuration'] == null ? undefined : json['defaultGlassBreakDuration'],
        'grantsDevAzureGroup': json['grantsDevAzureGroup'] == null ? undefined : json['grantsDevAzureGroup'],
        'grantsDevFirecloudGroup': json['grantsDevFirecloudGroup'] == null ? undefined : json['grantsDevFirecloudGroup'],
        'grantsSherlockSuperAdmin': json['grantsSherlockSuperAdmin'] == null ? undefined : json['grantsSherlockSuperAdmin'],
        'name': json['name'] == null ? undefined : json['name'],
        'suspendNonSuitableUsers': json['suspendNonSuitableUsers'] == null ? undefined : json['suspendNonSuitableUsers'],
    };
}

export function SherlockRoleV3EditToJSON(value?: SherlockRoleV3Edit | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'canBeGlassBrokenByRole': value['canBeGlassBrokenByRole'],
        'defaultGlassBreakDuration': value['defaultGlassBreakDuration'],
        'grantsDevAzureGroup': value['grantsDevAzureGroup'],
        'grantsDevFirecloudGroup': value['grantsDevFirecloudGroup'],
        'grantsSherlockSuperAdmin': value['grantsSherlockSuperAdmin'],
        'name': value['name'],
        'suspendNonSuitableUsers': value['suspendNonSuitableUsers'],
    };
}
