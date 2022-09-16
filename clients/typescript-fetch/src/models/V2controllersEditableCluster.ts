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
 * The subset of Cluster fields that can be edited after creation
 * @export
 * @interface V2controllersEditableCluster
 */
export interface V2controllersEditableCluster {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersEditableCluster
     */
    address?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersEditableCluster
     */
    base?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEditableCluster
     */
    requiresSuitability?: boolean;
}

/**
 * Check if a given object implements the V2controllersEditableCluster interface.
 */
export function instanceOfV2controllersEditableCluster(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersEditableClusterFromJSON(json: any): V2controllersEditableCluster {
    return V2controllersEditableClusterFromJSONTyped(json, false);
}

export function V2controllersEditableClusterFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersEditableCluster {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'address') ? undefined : json['address'],
        'base': !exists(json, 'base') ? undefined : json['base'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
    };
}

export function V2controllersEditableClusterToJSON(value?: V2controllersEditableCluster | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'address': value.address,
        'base': value.base,
        'requiresSuitability': value.requiresSuitability,
    };
}

