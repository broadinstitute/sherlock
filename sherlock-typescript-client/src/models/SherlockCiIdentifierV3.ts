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
import type { SherlockCiRunV3 } from './SherlockCiRunV3';
import {
    SherlockCiRunV3FromJSON,
    SherlockCiRunV3FromJSONTyped,
    SherlockCiRunV3ToJSON,
    SherlockCiRunV3ToJSONTyped,
} from './SherlockCiRunV3';

/**
 * 
 * @export
 * @interface SherlockCiIdentifierV3
 */
export interface SherlockCiIdentifierV3 {
    /**
     * 
     * @type {Array<SherlockCiRunV3>}
     * @memberof SherlockCiIdentifierV3
     */
    ciRuns?: Array<SherlockCiRunV3>;
    /**
     * 
     * @type {Date}
     * @memberof SherlockCiIdentifierV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof SherlockCiIdentifierV3
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof SherlockCiIdentifierV3
     */
    resourceID?: number;
    /**
     * Available only when querying a CiIdentifier via a CiRun, indicates the status of the run for that resource
     * @type {string}
     * @memberof SherlockCiIdentifierV3
     */
    resourceStatus?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiIdentifierV3
     */
    resourceType?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockCiIdentifierV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockCiIdentifierV3 interface.
 */
export function instanceOfSherlockCiIdentifierV3(value: object): value is SherlockCiIdentifierV3 {
    return true;
}

export function SherlockCiIdentifierV3FromJSON(json: any): SherlockCiIdentifierV3 {
    return SherlockCiIdentifierV3FromJSONTyped(json, false);
}

export function SherlockCiIdentifierV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockCiIdentifierV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'ciRuns': json['ciRuns'] == null ? undefined : ((json['ciRuns'] as Array<any>).map(SherlockCiRunV3FromJSON)),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'id': json['id'] == null ? undefined : json['id'],
        'resourceID': json['resourceID'] == null ? undefined : json['resourceID'],
        'resourceStatus': json['resourceStatus'] == null ? undefined : json['resourceStatus'],
        'resourceType': json['resourceType'] == null ? undefined : json['resourceType'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockCiIdentifierV3ToJSON(json: any): SherlockCiIdentifierV3 {
    return SherlockCiIdentifierV3ToJSONTyped(json, false);
}

export function SherlockCiIdentifierV3ToJSONTyped(value?: SherlockCiIdentifierV3 | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ciRuns': value['ciRuns'] == null ? undefined : ((value['ciRuns'] as Array<any>).map(SherlockCiRunV3ToJSON)),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'id': value['id'],
        'resourceID': value['resourceID'],
        'resourceStatus': value['resourceStatus'],
        'resourceType': value['resourceType'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

