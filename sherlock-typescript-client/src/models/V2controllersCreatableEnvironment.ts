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
 * @interface V2controllersCreatableEnvironment
 */
export interface V2controllersCreatableEnvironment {
    /**
     * 
     * @type {EnvironmentAutoDelete}
     * @memberof V2controllersCreatableEnvironment
     */
    autoDelete?: EnvironmentAutoDelete;
    /**
     * If true when creating, dynamic environments copy from template and template environments get the honeycomb chart
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    autoPopulateChartReleases?: boolean;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    base?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    baseDomain?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    defaultCluster?: string;
    /**
     * should be the environment branch for live envs. Is usually dev for template/dynamic but not necessarily
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    defaultFirecloudDevelopRef?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    defaultNamespace?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    helmfileRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    lifecycle?: string;
    /**
     * When creating, will be calculated if dynamic, required otherwise
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    name?: string;
    /**
     * Used for dynamic environment name generation only, to override using the owner email handle and template name
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    namePrefix?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    namePrefixesDomain?: boolean;
    /**
     * Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    offline?: boolean;
    /**
     * When enabled, the BEE will be slated to go offline around the begin time each day
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    offlineScheduleBeginEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof V2controllersCreatableEnvironment
     */
    offlineScheduleBeginTime?: Date;
    /**
     * When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    offlineScheduleEndEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof V2controllersCreatableEnvironment
     */
    offlineScheduleEndTime?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    offlineScheduleEndWeekends?: boolean;
    /**
     * When creating, will default to you
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    pactIdentifier?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    pagerdutyIntegration?: string;
    /**
     * Used to protect specific BEEs from deletion (thelma checks this field)
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    preventDeletion?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersCreatableEnvironment
     */
    requiresSuitability?: boolean;
    /**
     * Required for dynamic environments
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    templateEnvironment?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    uniqueResourcePrefix?: string;
    /**
     * When creating, defaults to template name or environment name
     * @type {string}
     * @memberof V2controllersCreatableEnvironment
     */
    valuesName?: string;
}

/**
 * Check if a given object implements the V2controllersCreatableEnvironment interface.
 */
export function instanceOfV2controllersCreatableEnvironment(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersCreatableEnvironmentFromJSON(json: any): V2controllersCreatableEnvironment {
    return V2controllersCreatableEnvironmentFromJSONTyped(json, false);
}

export function V2controllersCreatableEnvironmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersCreatableEnvironment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'autoDelete': !exists(json, 'autoDelete') ? undefined : EnvironmentAutoDeleteFromJSON(json['autoDelete']),
        'autoPopulateChartReleases': !exists(json, 'autoPopulateChartReleases') ? undefined : json['autoPopulateChartReleases'],
        'base': !exists(json, 'base') ? undefined : json['base'],
        'baseDomain': !exists(json, 'baseDomain') ? undefined : json['baseDomain'],
        'defaultCluster': !exists(json, 'defaultCluster') ? undefined : json['defaultCluster'],
        'defaultFirecloudDevelopRef': !exists(json, 'defaultFirecloudDevelopRef') ? undefined : json['defaultFirecloudDevelopRef'],
        'defaultNamespace': !exists(json, 'defaultNamespace') ? undefined : json['defaultNamespace'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'lifecycle': !exists(json, 'lifecycle') ? undefined : json['lifecycle'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'namePrefix': !exists(json, 'namePrefix') ? undefined : json['namePrefix'],
        'namePrefixesDomain': !exists(json, 'namePrefixesDomain') ? undefined : json['namePrefixesDomain'],
        'offline': !exists(json, 'offline') ? undefined : json['offline'],
        'offlineScheduleBeginEnabled': !exists(json, 'offlineScheduleBeginEnabled') ? undefined : json['offlineScheduleBeginEnabled'],
        'offlineScheduleBeginTime': !exists(json, 'offlineScheduleBeginTime') ? undefined : (new Date(json['offlineScheduleBeginTime'])),
        'offlineScheduleEndEnabled': !exists(json, 'offlineScheduleEndEnabled') ? undefined : json['offlineScheduleEndEnabled'],
        'offlineScheduleEndTime': !exists(json, 'offlineScheduleEndTime') ? undefined : (new Date(json['offlineScheduleEndTime'])),
        'offlineScheduleEndWeekends': !exists(json, 'offlineScheduleEndWeekends') ? undefined : json['offlineScheduleEndWeekends'],
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'pactIdentifier': !exists(json, 'pactIdentifier') ? undefined : json['pactIdentifier'],
        'pagerdutyIntegration': !exists(json, 'pagerdutyIntegration') ? undefined : json['pagerdutyIntegration'],
        'preventDeletion': !exists(json, 'preventDeletion') ? undefined : json['preventDeletion'],
        'requiresSuitability': !exists(json, 'requiresSuitability') ? undefined : json['requiresSuitability'],
        'templateEnvironment': !exists(json, 'templateEnvironment') ? undefined : json['templateEnvironment'],
        'uniqueResourcePrefix': !exists(json, 'uniqueResourcePrefix') ? undefined : json['uniqueResourcePrefix'],
        'valuesName': !exists(json, 'valuesName') ? undefined : json['valuesName'],
    };
}

export function V2controllersCreatableEnvironmentToJSON(value?: V2controllersCreatableEnvironment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'autoDelete': EnvironmentAutoDeleteToJSON(value.autoDelete),
        'autoPopulateChartReleases': value.autoPopulateChartReleases,
        'base': value.base,
        'baseDomain': value.baseDomain,
        'defaultCluster': value.defaultCluster,
        'defaultFirecloudDevelopRef': value.defaultFirecloudDevelopRef,
        'defaultNamespace': value.defaultNamespace,
        'description': value.description,
        'helmfileRef': value.helmfileRef,
        'lifecycle': value.lifecycle,
        'name': value.name,
        'namePrefix': value.namePrefix,
        'namePrefixesDomain': value.namePrefixesDomain,
        'offline': value.offline,
        'offlineScheduleBeginEnabled': value.offlineScheduleBeginEnabled,
        'offlineScheduleBeginTime': value.offlineScheduleBeginTime === undefined ? undefined : (value.offlineScheduleBeginTime.toISOString()),
        'offlineScheduleEndEnabled': value.offlineScheduleEndEnabled,
        'offlineScheduleEndTime': value.offlineScheduleEndTime === undefined ? undefined : (value.offlineScheduleEndTime.toISOString()),
        'offlineScheduleEndWeekends': value.offlineScheduleEndWeekends,
        'owner': value.owner,
        'pactIdentifier': value.pactIdentifier,
        'pagerdutyIntegration': value.pagerdutyIntegration,
        'preventDeletion': value.preventDeletion,
        'requiresSuitability': value.requiresSuitability,
        'templateEnvironment': value.templateEnvironment,
        'uniqueResourcePrefix': value.uniqueResourcePrefix,
        'valuesName': value.valuesName,
    };
}

