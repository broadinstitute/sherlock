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
 * @interface SherlockUserV3Upsert
 */
export interface SherlockUserV3Upsert {
    /**
     * An access token for the GitHub account to associate with the calling user. The access token isn't stored.
     * The design here ensures that an association is only built when someone controls both accounts (Google via
     * IAP and GitHub via this access token).
     * @type {string}
     * @memberof SherlockUserV3Upsert
     */
    githubAccessToken?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3Upsert
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockUserV3Upsert
     */
    nameFrom?: SherlockUserV3UpsertNameFromEnum;
}


/**
 * @export
 */
export const SherlockUserV3UpsertNameFromEnum = {
    Sherlock: 'sherlock',
    Github: 'github',
    Slack: 'slack'
} as const;
export type SherlockUserV3UpsertNameFromEnum = typeof SherlockUserV3UpsertNameFromEnum[keyof typeof SherlockUserV3UpsertNameFromEnum];


/**
 * Check if a given object implements the SherlockUserV3Upsert interface.
 */
export function instanceOfSherlockUserV3Upsert(value: object): value is SherlockUserV3Upsert {
    return true;
}

export function SherlockUserV3UpsertFromJSON(json: any): SherlockUserV3Upsert {
    return SherlockUserV3UpsertFromJSONTyped(json, false);
}

export function SherlockUserV3UpsertFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockUserV3Upsert {
    if (json == null) {
        return json;
    }
    return {
        
        'githubAccessToken': json['githubAccessToken'] == null ? undefined : json['githubAccessToken'],
        'name': json['name'] == null ? undefined : json['name'],
        'nameFrom': json['nameFrom'] == null ? undefined : json['nameFrom'],
    };
}

export function SherlockUserV3UpsertToJSON(json: any): SherlockUserV3Upsert {
    return SherlockUserV3UpsertToJSONTyped(json, false);
}

export function SherlockUserV3UpsertToJSONTyped(value?: SherlockUserV3Upsert | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'githubAccessToken': value['githubAccessToken'],
        'name': value['name'],
        'nameFrom': value['nameFrom'],
    };
}

