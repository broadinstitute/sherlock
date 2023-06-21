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
 * The subset of Cluster fields that can be set upon creation
 * @export
 * @interface V2controllersCreatableCluster
 */
export interface V2controllersCreatableCluster {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    address?: string;
    /**
     * Required when creating if providers is 'azure'
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    azureSubscription?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    base?: string;
    /**
     * Required when creating if provider is 'google'
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    googleProject?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    helmfileRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    location?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableCluster
     */
    provider?: V2controllersCreatableClusterProviderEnum;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersCreatableCluster
     */
    requiresSuitability?: boolean;
}


/**
 * @export
 */
export const V2controllersCreatableClusterProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type V2controllersCreatableClusterProviderEnum = typeof V2controllersCreatableClusterProviderEnum[keyof typeof V2controllersCreatableClusterProviderEnum];


/**
 * Check if a given object implements the V2controllersCreatableCluster interface.
 */
export function instanceOfV2controllersCreatableCluster(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersCreatableClusterFromJSON(json: any): V2controllersCreatableCluster {
    return V2controllersCreatableClusterFromJSONTyped(json, false);
}

export function V2controllersCreatableClusterFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersCreatableCluster {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'address') ? undefined : json['address'],
        'azureSubscription': !exists(json, 'azureSubscription') ? undefined : json['azureSubscription'],
        'base': !exists(json, 'base') ? undefined : json['base'],
        'googleProject': !exists(json, 'googleProject') ? undefined : json['googleProject'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'location': !exists(json, 'location') ? undefined : json['location'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'provider': !exists(json, 'provider') ? undefined : json['provider'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
    };
}

export function V2controllersCreatableClusterToJSON(value?: V2controllersCreatableCluster | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'address': value.address,
        'azureSubscription': value.azureSubscription,
        'base': value.base,
        'googleProject': value.googleProject,
        'helmfileRef': value.helmfileRef,
        'location': value.location,
        'name': value.name,
        'provider': value.provider,
        'requiresSuitability': value.requiresSuitability,
    };
}
