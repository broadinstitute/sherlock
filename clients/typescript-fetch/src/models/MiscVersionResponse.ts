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
 * @interface MiscVersionResponse
 */
export interface MiscVersionResponse {
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof MiscVersionResponse
     */
    buildInfo?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof MiscVersionResponse
     */
    goVersion?: string;
    /**
     * 
     * @type {string}
     * @memberof MiscVersionResponse
     */
    version?: string;
}

/**
 * Check if a given object implements the MiscVersionResponse interface.
 */
export function instanceOfMiscVersionResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MiscVersionResponseFromJSON(json: any): MiscVersionResponse {
    return MiscVersionResponseFromJSONTyped(json, false);
}

export function MiscVersionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MiscVersionResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'buildInfo': !exists(json, 'buildInfo') ? undefined : json['buildInfo'],
        'goVersion': !exists(json, 'goVersion') ? undefined : json['goVersion'],
        'version': !exists(json, 'version') ? undefined : json['version'],
    };
}

export function MiscVersionResponseToJSON(value?: MiscVersionResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'buildInfo': value.buildInfo,
        'goVersion': value.goVersion,
        'version': value.version,
    };
}

