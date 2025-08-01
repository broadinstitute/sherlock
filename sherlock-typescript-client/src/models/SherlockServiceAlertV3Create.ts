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
 * @interface SherlockServiceAlertV3Create
 */
export interface SherlockServiceAlertV3Create {
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3Create
     */
    link?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3Create
     */
    message?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3Create
     */
    onEnvironment?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3Create
     */
    severity?: SherlockServiceAlertV3CreateSeverityEnum;
    /**
     * 
     * @type {string}
     * @memberof SherlockServiceAlertV3Create
     */
    title?: string;
}


/**
 * @export
 */
export const SherlockServiceAlertV3CreateSeverityEnum = {
    Blocker: 'blocker',
    Critical: ' critical',
    Minor: ' minor'
} as const;
export type SherlockServiceAlertV3CreateSeverityEnum = typeof SherlockServiceAlertV3CreateSeverityEnum[keyof typeof SherlockServiceAlertV3CreateSeverityEnum];


/**
 * Check if a given object implements the SherlockServiceAlertV3Create interface.
 */
export function instanceOfSherlockServiceAlertV3Create(value: object): value is SherlockServiceAlertV3Create {
    return true;
}

export function SherlockServiceAlertV3CreateFromJSON(json: any): SherlockServiceAlertV3Create {
    return SherlockServiceAlertV3CreateFromJSONTyped(json, false);
}

export function SherlockServiceAlertV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockServiceAlertV3Create {
    if (json == null) {
        return json;
    }
    return {
        
        'link': json['link'] == null ? undefined : json['link'],
        'message': json['message'] == null ? undefined : json['message'],
        'onEnvironment': json['onEnvironment'] == null ? undefined : json['onEnvironment'],
        'severity': json['severity'] == null ? undefined : json['severity'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function SherlockServiceAlertV3CreateToJSON(json: any): SherlockServiceAlertV3Create {
    return SherlockServiceAlertV3CreateToJSONTyped(json, false);
}

export function SherlockServiceAlertV3CreateToJSONTyped(value?: SherlockServiceAlertV3Create | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'link': value['link'],
        'message': value['message'],
        'onEnvironment': value['onEnvironment'],
        'severity': value['severity'],
        'title': value['title'],
    };
}

