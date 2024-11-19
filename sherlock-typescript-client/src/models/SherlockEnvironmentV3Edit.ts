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
 * @interface SherlockEnvironmentV3Edit
 */
export interface SherlockEnvironmentV3Edit {
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    baseDomain?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    defaultCluster?: string;
    /**
     * If set, the BEE will be automatically deleted after this time. Can be set to "" or Go's zero time value to clear the field.
     * @type {Date}
     * @memberof SherlockEnvironmentV3Edit
     */
    deleteAfter?: Date;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    description?: string;
    /**
     * If true, janitor resource cleanup will be enabled for this environment. BEEs default to template's value, templates default to true, and static/live environments default to false.
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    enableJanitor?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    helmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    namePrefixesDomain?: boolean;
    /**
     * Applicable for BEEs only, whether Thelma should render the BEE as "offline" zero replicas (this field is a target state, not a status)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    offline?: boolean;
    /**
     * When enabled, the BEE will be slated to go offline around the begin time each day
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    offlineScheduleBeginEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof SherlockEnvironmentV3Edit
     */
    offlineScheduleBeginTime?: Date;
    /**
     * When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    offlineScheduleEndEnabled?: boolean;
    /**
     * Stored with timezone to determine day of the week
     * @type {Date}
     * @memberof SherlockEnvironmentV3Edit
     */
    offlineScheduleEndTime?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    offlineScheduleEndWeekends?: boolean;
    /**
     * When creating, will default to you
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    pactIdentifier?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    pagerdutyIntegration?: string;
    /**
     * Used to protect specific BEEs from deletion (thelma checks this field)
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    preventDeletion?: boolean;
    /**
     * If present, requires membership in the given role for mutations. Set to an empty string to clear.
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    requiredRole?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockEnvironmentV3Edit
     */
    requiresSuitability?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockEnvironmentV3Edit
     */
    serviceBannerBucket?: string;
}

/**
 * Check if a given object implements the SherlockEnvironmentV3Edit interface.
 */
export function instanceOfSherlockEnvironmentV3Edit(value: object): value is SherlockEnvironmentV3Edit {
    return true;
}

export function SherlockEnvironmentV3EditFromJSON(json: any): SherlockEnvironmentV3Edit {
    return SherlockEnvironmentV3EditFromJSONTyped(json, false);
}

export function SherlockEnvironmentV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockEnvironmentV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'baseDomain': json['baseDomain'] == null ? undefined : json['baseDomain'],
        'defaultCluster': json['defaultCluster'] == null ? undefined : json['defaultCluster'],
        'deleteAfter': json['deleteAfter'] == null ? undefined : (new Date(json['deleteAfter'])),
        'description': json['description'] == null ? undefined : json['description'],
        'enableJanitor': json['enableJanitor'] == null ? undefined : json['enableJanitor'],
        'helmfileRef': json['helmfileRef'] == null ? undefined : json['helmfileRef'],
        'namePrefixesDomain': json['namePrefixesDomain'] == null ? undefined : json['namePrefixesDomain'],
        'offline': json['offline'] == null ? undefined : json['offline'],
        'offlineScheduleBeginEnabled': json['offlineScheduleBeginEnabled'] == null ? undefined : json['offlineScheduleBeginEnabled'],
        'offlineScheduleBeginTime': json['offlineScheduleBeginTime'] == null ? undefined : (new Date(json['offlineScheduleBeginTime'])),
        'offlineScheduleEndEnabled': json['offlineScheduleEndEnabled'] == null ? undefined : json['offlineScheduleEndEnabled'],
        'offlineScheduleEndTime': json['offlineScheduleEndTime'] == null ? undefined : (new Date(json['offlineScheduleEndTime'])),
        'offlineScheduleEndWeekends': json['offlineScheduleEndWeekends'] == null ? undefined : json['offlineScheduleEndWeekends'],
        'owner': json['owner'] == null ? undefined : json['owner'],
        'pactIdentifier': json['pactIdentifier'] == null ? undefined : json['pactIdentifier'],
        'pagerdutyIntegration': json['pagerdutyIntegration'] == null ? undefined : json['pagerdutyIntegration'],
        'preventDeletion': json['preventDeletion'] == null ? undefined : json['preventDeletion'],
        'requiredRole': json['requiredRole'] == null ? undefined : json['requiredRole'],
        'requiresSuitability': json['requiresSuitability'] == null ? undefined : json['requiresSuitability'],
        'serviceBannerBucket': json['serviceBannerBucket'] == null ? undefined : json['serviceBannerBucket'],
    };
}

export function SherlockEnvironmentV3EditToJSON(json: any): SherlockEnvironmentV3Edit {
    return SherlockEnvironmentV3EditToJSONTyped(json, false);
}

export function SherlockEnvironmentV3EditToJSONTyped(value?: SherlockEnvironmentV3Edit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'baseDomain': value['baseDomain'],
        'defaultCluster': value['defaultCluster'],
        'deleteAfter': value['deleteAfter'] == null ? undefined : ((value['deleteAfter']).toISOString()),
        'description': value['description'],
        'enableJanitor': value['enableJanitor'],
        'helmfileRef': value['helmfileRef'],
        'namePrefixesDomain': value['namePrefixesDomain'],
        'offline': value['offline'],
        'offlineScheduleBeginEnabled': value['offlineScheduleBeginEnabled'],
        'offlineScheduleBeginTime': value['offlineScheduleBeginTime'] == null ? undefined : ((value['offlineScheduleBeginTime']).toISOString()),
        'offlineScheduleEndEnabled': value['offlineScheduleEndEnabled'],
        'offlineScheduleEndTime': value['offlineScheduleEndTime'] == null ? undefined : ((value['offlineScheduleEndTime']).toISOString()),
        'offlineScheduleEndWeekends': value['offlineScheduleEndWeekends'],
        'owner': value['owner'],
        'pactIdentifier': value['pactIdentifier'],
        'pagerdutyIntegration': value['pagerdutyIntegration'],
        'preventDeletion': value['preventDeletion'],
        'requiredRole': value['requiredRole'],
        'requiresSuitability': value['requiresSuitability'],
        'serviceBannerBucket': value['serviceBannerBucket'],
    };
}

