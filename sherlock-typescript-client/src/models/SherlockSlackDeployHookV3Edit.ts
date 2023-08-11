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
 * @interface SherlockSlackDeployHookV3Edit
 */
export interface SherlockSlackDeployHookV3Edit {
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
export function instanceOfSherlockSlackDeployHookV3Edit(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockSlackDeployHookV3EditFromJSON(json: any): SherlockSlackDeployHookV3Edit {
    return SherlockSlackDeployHookV3EditFromJSONTyped(json, false);
}

export function SherlockSlackDeployHookV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockSlackDeployHookV3Edit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'onFailure': !exists(json, 'onFailure') ? undefined : json['onFailure'],
        'onSuccess': !exists(json, 'onSuccess') ? undefined : json['onSuccess'],
        'slackChannel': !exists(json, 'slackChannel') ? undefined : json['slackChannel'],
    };
}

export function SherlockSlackDeployHookV3EditToJSON(value?: SherlockSlackDeployHookV3Edit | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'onFailure': value.onFailure,
        'onSuccess': value.onSuccess,
        'slackChannel': value.slackChannel,
    };
}

