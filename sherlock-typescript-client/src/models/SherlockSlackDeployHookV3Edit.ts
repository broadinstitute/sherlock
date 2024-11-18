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
 * @interface SherlockSlackDeployHookV3Edit
 */
export interface SherlockSlackDeployHookV3Edit {
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3Edit
     */
    mentionPeople?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3Edit
     */
    onFailure?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockSlackDeployHookV3Edit
     */
    onSuccess?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockSlackDeployHookV3Edit
     */
    slackChannel?: string;
}

/**
 * Check if a given object implements the SherlockSlackDeployHookV3Edit interface.
 */
export function instanceOfSherlockSlackDeployHookV3Edit(value: object): value is SherlockSlackDeployHookV3Edit {
    return true;
}

export function SherlockSlackDeployHookV3EditFromJSON(json: any): SherlockSlackDeployHookV3Edit {
    return SherlockSlackDeployHookV3EditFromJSONTyped(json, false);
}

export function SherlockSlackDeployHookV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockSlackDeployHookV3Edit {
    if (json == null) {
        return json;
    }
    return {
        
        'mentionPeople': json['mentionPeople'] == null ? undefined : json['mentionPeople'],
        'onFailure': json['onFailure'] == null ? undefined : json['onFailure'],
        'onSuccess': json['onSuccess'] == null ? undefined : json['onSuccess'],
        'slackChannel': json['slackChannel'] == null ? undefined : json['slackChannel'],
    };
}

export function SherlockSlackDeployHookV3EditToJSON(json: any): SherlockSlackDeployHookV3Edit {
    return SherlockSlackDeployHookV3EditToJSONTyped(json, false);
}

export function SherlockSlackDeployHookV3EditToJSONTyped(value?: SherlockSlackDeployHookV3Edit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'mentionPeople': value['mentionPeople'],
        'onFailure': value['onFailure'],
        'onSuccess': value['onSuccess'],
        'slackChannel': value['slackChannel'],
    };
}

