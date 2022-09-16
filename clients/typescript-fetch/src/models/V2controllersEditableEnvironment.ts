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
 * @interface V2controllersEditableEnvironment
 */
export interface V2controllersEditableEnvironment {
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    baseDomain?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    defaultCluster?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    defaultNamespace?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    namePrefixesDomain?: boolean;
    /**
     * When creating, will be set to your email
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    owner?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    requiresSuitability?: boolean;
}

/**
 * Check if a given object implements the V2controllersEditableEnvironment interface.
 */
export function instanceOfV2controllersEditableEnvironment(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersEditableEnvironmentFromJSON(json: any): V2controllersEditableEnvironment {
    return V2controllersEditableEnvironmentFromJSONTyped(json, false);
}

export function V2controllersEditableEnvironmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersEditableEnvironment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'baseDomain': !exists(json, 'baseDomain') ? undefined : json['baseDomain'],
        'defaultCluster': !exists(json, 'defaultCluster') ? undefined : json['defaultCluster'],
        'defaultNamespace': !exists(json, 'defaultNamespace') ? undefined : json['defaultNamespace'],
        'namePrefixesDomain': !exists(json, 'namePrefixesDomain') ? undefined : json['namePrefixesDomain'],
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
    };
}

export function V2controllersEditableEnvironmentToJSON(value?: V2controllersEditableEnvironment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'baseDomain': value.baseDomain,
        'defaultCluster': value.defaultCluster,
        'defaultNamespace': value.defaultNamespace,
        'namePrefixesDomain': value.namePrefixesDomain,
        'owner': value.owner,
        'requiresSuitability': value.requiresSuitability,
    };
}

