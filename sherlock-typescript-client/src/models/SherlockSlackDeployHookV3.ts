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
 * @interface SherlockSlackDeployHookV3
 */
export interface SherlockSlackDeployHookV3 {
    /**
     * 
     * @type {Date}
     * @memberof SherlockSlackDeployHookV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof SherlockSlackDeployHookV3
     */
    id?: number;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3
     */
    mentionPeople?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockSlackDeployHookV3
     */
    onChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockSlackDeployHookV3
     */
    onEnvironment?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3
     */
    onFailure?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3
     */
    onSuccess?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockSlackDeployHookV3
     */
    slackChannel?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockSlackDeployHookV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockSlackDeployHookV3 interface.
 */
export function instanceOfSherlockSlackDeployHookV3(value: object): value is SherlockSlackDeployHookV3 {
    return true;
}

export function SherlockSlackDeployHookV3FromJSON(json: any): SherlockSlackDeployHookV3 {
    return SherlockSlackDeployHookV3FromJSONTyped(json, false);
}

export function SherlockSlackDeployHookV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockSlackDeployHookV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'id': json['id'] == null ? undefined : json['id'],
        'mentionPeople': json['mentionPeople'] == null ? undefined : json['mentionPeople'],
        'onChartRelease': json['onChartRelease'] == null ? undefined : json['onChartRelease'],
        'onEnvironment': json['onEnvironment'] == null ? undefined : json['onEnvironment'],
        'onFailure': json['onFailure'] == null ? undefined : json['onFailure'],
        'onSuccess': json['onSuccess'] == null ? undefined : json['onSuccess'],
        'slackChannel': json['slackChannel'] == null ? undefined : json['slackChannel'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockSlackDeployHookV3ToJSON(json: any): SherlockSlackDeployHookV3 {
    return SherlockSlackDeployHookV3ToJSONTyped(json, false);
}

export function SherlockSlackDeployHookV3ToJSONTyped(value?: SherlockSlackDeployHookV3 | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'id': value['id'],
        'mentionPeople': value['mentionPeople'],
        'onChartRelease': value['onChartRelease'],
        'onEnvironment': value['onEnvironment'],
        'onFailure': value['onFailure'],
        'onSuccess': value['onSuccess'],
        'slackChannel': value['slackChannel'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

