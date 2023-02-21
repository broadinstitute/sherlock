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
import type { EnvironmentAutoDelete } from './EnvironmentAutoDelete';
import {
    EnvironmentAutoDeleteFromJSON,
    EnvironmentAutoDeleteFromJSONTyped,
    EnvironmentAutoDeleteToJSON,
} from './EnvironmentAutoDelete';

/**
 * 
 * @export
 * @interface V2controllersEditableEnvironment
 */
export interface V2controllersEditableEnvironment {
    /**
     * 
     * @type {EnvironmentAutoDelete}
     * @memberof V2controllersEditableEnvironment
     */
    autoDelete?: EnvironmentAutoDelete;
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
     * should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    defaultFirecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    helmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    namePrefixesDomain?: boolean;
    /**
     * Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    offline?: boolean;
    /**
     * When enabled, the BEE will be slated to go offline around the begin time each day
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    offlineScheduleBeginEnabled?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersEditableEnvironment
     */
    offlineScheduleBeginTime?: Date;
    /**
     * When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    offlineScheduleEndEnabled?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersEditableEnvironment
     */
    offlineScheduleEndTime?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    offlineScheduleEndWeekends?: boolean;
    /**
     * When creating, will be set to your email
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersEditableEnvironment
     */
    pagerdutyIntegration?: string;
    /**
     * Used to protect specific BEEs from deletion (thelma checks this field)
     * @type {boolean}
     * @memberof V2controllersEditableEnvironment
     */
    preventDeletion?: boolean;
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
        
        'autoDelete': !exists(json, 'autoDelete') ? undefined : EnvironmentAutoDeleteFromJSON(json['autoDelete']),
        'baseDomain': !exists(json, 'baseDomain') ? undefined : json['baseDomain'],
        'defaultCluster': !exists(json, 'defaultCluster') ? undefined : json['defaultCluster'],
        'defaultFirecloudDevelopRef': !exists(json, 'defaultFirecloudDevelopRef') ? undefined : json['defaultFirecloudDevelopRef'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'namePrefixesDomain': !exists(json, 'namePrefixesDomain') ? undefined : json['namePrefixesDomain'],
        'offline': !exists(json, 'offline') ? undefined : json['offline'],
        'offlineScheduleBeginEnabled': !exists(json, 'offlineScheduleBeginEnabled') ? undefined : json['offlineScheduleBeginEnabled'],
        'offlineScheduleBeginTime': !exists(json, 'offlineScheduleBeginTime') ? undefined : (new Date(json['offlineScheduleBeginTime'])),
        'offlineScheduleEndEnabled': !exists(json, 'offlineScheduleEndEnabled') ? undefined : json['offlineScheduleEndEnabled'],
        'offlineScheduleEndTime': !exists(json, 'offlineScheduleEndTime') ? undefined : (new Date(json['offlineScheduleEndTime'])),
        'offlineScheduleEndWeekends': !exists(json, 'offlineScheduleEndWeekends') ? undefined : json['offlineScheduleEndWeekends'],
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'pagerdutyIntegration': !exists(json, 'pagerdutyIntegration') ? undefined : json['pagerdutyIntegration'],
        'preventDeletion': !exists(json, 'preventDeletion') ? undefined : json['preventDeletion'],
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
        
        'autoDelete': EnvironmentAutoDeleteToJSON(value.autoDelete),
        'baseDomain': value.baseDomain,
        'defaultCluster': value.defaultCluster,
        'defaultFirecloudDevelopRef': value.defaultFirecloudDevelopRef,
        'description': value.description,
        'helmfileRef': value.helmfileRef,
        'namePrefixesDomain': value.namePrefixesDomain,
        'offline': value.offline,
        'offlineScheduleBeginEnabled': value.offlineScheduleBeginEnabled,
        'offlineScheduleBeginTime': value.offlineScheduleBeginTime === undefined ? undefined : (value.offlineScheduleBeginTime.toISOString()),
        'offlineScheduleEndEnabled': value.offlineScheduleEndEnabled,
        'offlineScheduleEndTime': value.offlineScheduleEndTime === undefined ? undefined : (value.offlineScheduleEndTime.toISOString()),
        'offlineScheduleEndWeekends': value.offlineScheduleEndWeekends,
        'owner': value.owner,
        'pagerdutyIntegration': value.pagerdutyIntegration,
        'preventDeletion': value.preventDeletion,
        'requiresSuitability': value.requiresSuitability,
    };
}

