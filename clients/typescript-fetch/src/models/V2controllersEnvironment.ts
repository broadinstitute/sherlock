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
import type { V2controllersCluster } from './V2controllersCluster';
import {
    V2controllersClusterFromJSON,
    V2controllersClusterFromJSONTyped,
    V2controllersClusterToJSON,
} from './V2controllersCluster';

/**
 * 
 * @export
 * @interface V2controllersEnvironment
 */
export interface V2controllersEnvironment {
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    base?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    baseDomain?: string;
    /**
     * Upon creation of a dynamic environment, if this is true the template's chart releases will be copied to the new environment
     * @type {boolean}
     * @memberof V2controllersEnvironment
     */
    chartReleasesFromTemplate?: boolean;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    defaultCluster?: string;
    /**
     * 
     * @type {V2controllersCluster}
     * @memberof V2controllersEnvironment
     */
    defaultClusterInfo?: V2controllersCluster;
    /**
     * should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    defaultFirecloudDevelopRef?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    defaultNamespace?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    helmfileRef?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersEnvironment
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    lifecycle?: string;
    /**
     * When creating, will be calculated if dynamic, required otherwise
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEnvironment
     */
    namePrefixesDomain?: boolean;
    /**
     * When creating, will be set to your email
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    owner?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEnvironment
     */
    requiresSuitability?: boolean;
    /**
     * Required for dynamic environments
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    templateEnvironment?: string;
    /**
     * Single-layer recursive; provides info of the template environment if this environment has one
     * @type {object}
     * @memberof V2controllersEnvironment
     */
    templateEnvironmentInfo?: object;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    uniqueResourcePrefix?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    updatedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEnvironment
     */
    valuesName?: string;
}

/**
 * Check if a given object implements the V2controllersEnvironment interface.
 */
export function instanceOfV2controllersEnvironment(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersEnvironmentFromJSON(json: any): V2controllersEnvironment {
    return V2controllersEnvironmentFromJSONTyped(json, false);
}

export function V2controllersEnvironmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersEnvironment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'base': !exists(json, 'base') ? undefined : json['base'],
        'baseDomain': !exists(json, 'baseDomain') ? undefined : json['baseDomain'],
        'chartReleasesFromTemplate': !exists(json, 'chartReleasesFromTemplate') ? undefined : json['chartReleasesFromTemplate'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'defaultCluster': !exists(json, 'defaultCluster') ? undefined : json['defaultCluster'],
        'defaultClusterInfo': !exists(json, 'defaultClusterInfo') ? undefined : V2controllersClusterFromJSON(json['defaultClusterInfo']),
        'defaultFirecloudDevelopRef': !exists(json, 'defaultFirecloudDevelopRef') ? undefined : json['defaultFirecloudDevelopRef'],
        'defaultNamespace': !exists(json, 'defaultNamespace') ? undefined : json['defaultNamespace'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'lifecycle': !exists(json, 'lifecycle') ? undefined : json['lifecycle'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'namePrefixesDomain': !exists(json, 'namePrefixesDomain') ? undefined : json['namePrefixesDomain'],
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
        'templateEnvironment': !exists(json, 'templateEnvironment') ? undefined : json['templateEnvironment'],
        'templateEnvironmentInfo': !exists(json, 'templateEnvironmentInfo') ? undefined : json['templateEnvironmentInfo'],
        'uniqueResourcePrefix': !exists(json, 'uniqueResourcePrefix') ? undefined : json['uniqueResourcePrefix'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
        'valuesName': !exists(json, 'valuesName') ? undefined : json['valuesName'],
    };
}

export function V2controllersEnvironmentToJSON(value?: V2controllersEnvironment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'base': value.base,
        'baseDomain': value.baseDomain,
        'chartReleasesFromTemplate': value.chartReleasesFromTemplate,
        'createdAt': value.createdAt,
        'defaultCluster': value.defaultCluster,
        'defaultClusterInfo': V2controllersClusterToJSON(value.defaultClusterInfo),
        'defaultFirecloudDevelopRef': value.defaultFirecloudDevelopRef,
        'defaultNamespace': value.defaultNamespace,
        'helmfileRef': value.helmfileRef,
        'id': value.id,
        'lifecycle': value.lifecycle,
        'name': value.name,
        'namePrefixesDomain': value.namePrefixesDomain,
        'owner': value.owner,
        'requiresSuitability': value.requiresSuitability,
        'templateEnvironment': value.templateEnvironment,
        'templateEnvironmentInfo': value.templateEnvironmentInfo,
        'uniqueResourcePrefix': value.uniqueResourcePrefix,
        'updatedAt': value.updatedAt,
        'valuesName': value.valuesName,
    };
}

