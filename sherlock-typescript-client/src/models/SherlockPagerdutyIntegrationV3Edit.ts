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
 * @interface SherlockPagerdutyIntegrationV3Edit
 */
export interface SherlockPagerdutyIntegrationV3Edit {
    /**
     * 
     * @type {string}
     * @memberof SherlockPagerdutyIntegrationV3Edit
     */
    key?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockPagerdutyIntegrationV3Edit
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockPagerdutyIntegrationV3Edit
     */
    type?: string;
}

/**
 * Check if a given object implements the SherlockPagerdutyIntegrationV3Edit interface.
 */
export function instanceOfSherlockPagerdutyIntegrationV3Edit(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockPagerdutyIntegrationV3EditFromJSON(json: any): SherlockPagerdutyIntegrationV3Edit {
    return SherlockPagerdutyIntegrationV3EditFromJSONTyped(json, false);
}

export function SherlockPagerdutyIntegrationV3EditFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockPagerdutyIntegrationV3Edit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'key': !exists(json, 'key') ? undefined : json['key'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function SherlockPagerdutyIntegrationV3EditToJSON(value?: SherlockPagerdutyIntegrationV3Edit | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'key': value.key,
        'name': value.name,
        'type': value.type,
    };
}
