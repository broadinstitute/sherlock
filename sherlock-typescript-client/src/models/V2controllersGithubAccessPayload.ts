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
 * @interface V2controllersGithubAccessPayload
 */
export interface V2controllersGithubAccessPayload {
    /**
     * 
     * @type {string}
     * @memberof V2controllersGithubAccessPayload
     */
    githubAccessToken?: string;
}

/**
 * Check if a given object implements the V2controllersGithubAccessPayload interface.
 */
export function instanceOfV2controllersGithubAccessPayload(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersGithubAccessPayloadFromJSON(json: any): V2controllersGithubAccessPayload {
    return V2controllersGithubAccessPayloadFromJSONTyped(json, false);
}

export function V2controllersGithubAccessPayloadFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersGithubAccessPayload {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'githubAccessToken': !exists(json, 'githubAccessToken') ? undefined : json['githubAccessToken'],
    };
}

export function V2controllersGithubAccessPayloadToJSON(value?: V2controllersGithubAccessPayload | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'githubAccessToken': value.githubAccessToken,
    };
}
