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
 * @interface SherlockEnvironmentV3Create
 */
export interface SherlockEnvironmentV3Create {
    /**
     * If true when creating, dynamic environments copy from template and template environments get the honeycomb chart
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    autoPopulateChartReleases?: boolean;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    base?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    baseDomain?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    defaultCluster?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    defaultNamespace?: string;
    /**
     * If set, the BEE will be automatically deleted after this time (thelma checks this field)
     * @type {Date}
     * @memberof SherlockEnvironmentV3Create
     */
    deleteAfter?: Date;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    helmfileRef?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    lifecycle?: string;
    /**
     * When creating, will be calculated if dynamic, required otherwise
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    name?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    namePrefixesDomain?: boolean;
    /**
     * Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    offline?: boolean;
    /**
     * When enabled, the BEE will be slated to go offline around the begin time each day
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    offlineScheduleBeginEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof SherlockEnvironmentV3Create
     */
    offlineScheduleBeginTime?: Date;
    /**
     * When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    offlineScheduleEndEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof SherlockEnvironmentV3Create
     */
    offlineScheduleEndTime?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    offlineScheduleEndWeekends?: boolean;
    /**
     * When creating, will default to you
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    pactIdentifier?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    pagerdutyIntegration?: string;
    /**
     * Used to protect specific BEEs from deletion (thelma checks this field)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    preventDeletion?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Create
     */
    requiresSuitability?: boolean;
    /**
     * Required for dynamic environments
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    templateEnvironment?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    uniqueResourcePrefix?: string;
    /**
     * When creating, defaults to template name or environment name
     * @type {string}
     * @memberof SherlockEnvironmentV3Create
     */
    valuesName?: string;
}

/**
 * Check if a given object implements the SherlockEnvironmentV3Create interface.
 */
export function instanceOfSherlockEnvironmentV3Create(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockEnvironmentV3CreateFromJSON(json: any): SherlockEnvironmentV3Create {
    return SherlockEnvironmentV3CreateFromJSONTyped(json, false);
}

export function SherlockEnvironmentV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockEnvironmentV3Create {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'autoPopulateChartReleases': !exists(json, 'autoPopulateChartReleases') ? undefined : json['autoPopulateChartReleases'],
        'base': !exists(json, 'base') ? undefined : json['base'],
        'baseDomain': !exists(json, 'baseDomain') ? undefined : json['baseDomain'],
        'defaultCluster': !exists(json, 'defaultCluster') ? undefined : json['defaultCluster'],
        'defaultNamespace': !exists(json, 'defaultNamespace') ? undefined : json['defaultNamespace'],
        'deleteAfter': !exists(json, 'deleteAfter') ? undefined : (new Date(json['deleteAfter'])),
        'description': !exists(json, 'description') ? undefined : json['description'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'lifecycle': !exists(json, 'lifecycle') ? undefined : json['lifecycle'],
        'name': !exists(json, 'name') ? undefined : json['name'],
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

export function SherlockEnvironmentV3CreateToJSON(value?: SherlockEnvironmentV3Create | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'autoPopulateChartReleases': value.autoPopulateChartReleases,
        'base': value.base,
        'baseDomain': value.baseDomain,
        'defaultCluster': value.defaultCluster,
        'defaultNamespace': value.defaultNamespace,
        'deleteAfter': value.deleteAfter === undefined ? undefined : (value.deleteAfter.toISOString()),
        'description': value.description,
        'helmfileRef': value.helmfileRef,
        'lifecycle': value.lifecycle,
        'name': value.name,
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

