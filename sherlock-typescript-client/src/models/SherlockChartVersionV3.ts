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
import type { SherlockUserV3 } from './SherlockUserV3';
import {
    SherlockUserV3FromJSON,
    SherlockUserV3FromJSONTyped,
    SherlockUserV3ToJSON,
    SherlockUserV3ToJSONTyped,
} from './SherlockUserV3';
import type { SherlockChartV3 } from './SherlockChartV3';
import {
    SherlockChartV3FromJSON,
    SherlockChartV3FromJSONTyped,
    SherlockChartV3ToJSON,
    SherlockChartV3ToJSONTyped,
} from './SherlockChartV3';
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
    SherlockCiIdentifierV3ToJSONTyped,
} from './SherlockCiIdentifierV3';

/**
 * 
 * @export
 * @interface SherlockChartVersionV3
 */
export interface SherlockChartVersionV3 {
    /**
     * 
     * @type {string}
     * @memberof SherlockChartVersionV3
     */
    authoredBy?: string;
    /**
     * 
     * @type {SherlockUserV3}
     * @memberof SherlockChartVersionV3
     */
    authoredByInfo?: SherlockUserV3;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartVersionV3
     */
    chart?: string;
    /**
     * 
     * @type {SherlockChartV3}
     * @memberof SherlockChartVersionV3
     */
    chartInfo?: SherlockChartV3;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartVersionV3
     */
    chartVersion?: string;
    /**
     * 
     * @type {SherlockCiIdentifierV3}
     * @memberof SherlockChartVersionV3
     */
    ciIdentifier?: SherlockCiIdentifierV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartVersionV3
     */
    createdAt?: Date;
    /**
     * Generally the Git commit message
     * @type {string}
     * @memberof SherlockChartVersionV3
     */
    description?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartVersionV3
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartVersionV3
     */
    parentChartVersion?: string;
    /**
     * 
     * @type {object}
     * @memberof SherlockChartVersionV3
     */
    parentChartVersionInfo?: object;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartVersionV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockChartVersionV3 interface.
 */
export function instanceOfSherlockChartVersionV3(value: object): value is SherlockChartVersionV3 {
    return true;
}

export function SherlockChartVersionV3FromJSON(json: any): SherlockChartVersionV3 {
    return SherlockChartVersionV3FromJSONTyped(json, false);
}

export function SherlockChartVersionV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartVersionV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'authoredBy': json['authoredBy'] == null ? undefined : json['authoredBy'],
        'authoredByInfo': json['authoredByInfo'] == null ? undefined : SherlockUserV3FromJSON(json['authoredByInfo']),
        'chart': json['chart'] == null ? undefined : json['chart'],
        'chartInfo': json['chartInfo'] == null ? undefined : SherlockChartV3FromJSON(json['chartInfo']),
        'chartVersion': json['chartVersion'] == null ? undefined : json['chartVersion'],
        'ciIdentifier': json['ciIdentifier'] == null ? undefined : SherlockCiIdentifierV3FromJSON(json['ciIdentifier']),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'parentChartVersion': json['parentChartVersion'] == null ? undefined : json['parentChartVersion'],
        'parentChartVersionInfo': json['parentChartVersionInfo'] == null ? undefined : json['parentChartVersionInfo'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockChartVersionV3ToJSON(json: any): SherlockChartVersionV3 {
    return SherlockChartVersionV3ToJSONTyped(json, false);
}

export function SherlockChartVersionV3ToJSONTyped(value?: SherlockChartVersionV3 | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'authoredBy': value['authoredBy'],
        'authoredByInfo': SherlockUserV3ToJSON(value['authoredByInfo']),
        'chart': value['chart'],
        'chartInfo': SherlockChartV3ToJSON(value['chartInfo']),
        'chartVersion': value['chartVersion'],
        'ciIdentifier': SherlockCiIdentifierV3ToJSON(value['ciIdentifier']),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'description': value['description'],
        'id': value['id'],
        'parentChartVersion': value['parentChartVersion'],
        'parentChartVersionInfo': value['parentChartVersionInfo'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

