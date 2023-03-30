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
 * @interface V2controllersEditableDatabaseInstance
 */
export interface V2controllersEditableDatabaseInstance {
    /**
     * When creating, defaults to the chart name
     * @type {string}
     * @memberof V2controllersEditableDatabaseInstance
     */
    defaultDatabase?: string;
    /**
     * Required if platform is 'google'
     * @type {string}
     * @memberof V2controllersEditableDatabaseInstance
     */
    googleProject?: string;
    /**
     * Required if platform is 'google' or 'azure'
     * @type {string}
     * @memberof V2controllersEditableDatabaseInstance
     */
    instanceName?: string;
    /**
     * 'google', 'azure', or default 'kubernetes'
     * @type {string}
     * @memberof V2controllersEditableDatabaseInstance
     */
    platform?: string;
}

/**
 * Check if a given object implements the V2controllersEditableDatabaseInstance interface.
 */
export function instanceOfV2controllersEditableDatabaseInstance(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersEditableDatabaseInstanceFromJSON(json: any): V2controllersEditableDatabaseInstance {
    return V2controllersEditableDatabaseInstanceFromJSONTyped(json, false);
}

export function V2controllersEditableDatabaseInstanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersEditableDatabaseInstance {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'defaultDatabase': !exists(json, 'defaultDatabase') ? undefined : json['defaultDatabase'],
        'googleProject': !exists(json, 'googleProject') ? undefined : json['googleProject'],
        'instanceName': !exists(json, 'instanceName') ? undefined : json['instanceName'],
        'platform': !exists(json, 'platform') ? undefined : json['platform'],
    };
}

export function V2controllersEditableDatabaseInstanceToJSON(value?: V2controllersEditableDatabaseInstance | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'defaultDatabase': value.defaultDatabase,
        'googleProject': value.googleProject,
        'instanceName': value.instanceName,
        'platform': value.platform,
    };
}

