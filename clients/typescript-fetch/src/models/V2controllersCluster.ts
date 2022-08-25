/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service
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
 * The full set of Cluster fields that can be read or used for filtering queries
 * @export
 * @interface V2controllersCluster
 */
export interface V2controllersCluster {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCluster
     */
    address?: string;
    /**
     * Required when creating if providers is 'azure'
     * @type {string}
     * @memberof V2controllersCluster
     */
    azureSubscription?: string;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCluster
     */
    base?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCluster
     */
    createdAt?: string;
    /**
     * Required when creating if provider is 'google'
     * @type {string}
     * @memberof V2controllersCluster
     */
    googleProject?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersCluster
     */
    id?: number;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCluster
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCluster
     */
    provider?: V2controllersClusterProviderEnum;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersCluster
     */
    requiresSuitability?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCluster
     */
    updatedAt?: string;
}


/**
 * @export
 */
export const V2controllersClusterProviderEnum = {
    Google: 'google',
    Azure: 'azure'
} as const;
export type V2controllersClusterProviderEnum = typeof V2controllersClusterProviderEnum[keyof typeof V2controllersClusterProviderEnum];


/**
 * Check if a given object implements the V2controllersCluster interface.
 */
export function instanceOfV2controllersCluster(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersClusterFromJSON(json: any): V2controllersCluster {
    return V2controllersClusterFromJSONTyped(json, false);
}

export function V2controllersClusterFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersCluster {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'address': !exists(json, 'address') ? undefined : json['address'],
        'azureSubscription': !exists(json, 'azureSubscription') ? undefined : json['azureSubscription'],
        'base': !exists(json, 'base') ? undefined : json['base'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'googleProject': !exists(json, 'googleProject') ? undefined : json['googleProject'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'provider': !exists(json, 'provider') ? undefined : json['provider'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}

export function V2controllersClusterToJSON(value?: V2controllersCluster | null): any {
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
        'createdAt': value.createdAt,
        'googleProject': value.googleProject,
        'id': value.id,
        'name': value.name,
        'provider': value.provider,
        'requiresSuitability': value.requiresSuitability,
        'updatedAt': value.updatedAt,
    };
}

