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
import type { SherlockChartV3 } from './SherlockChartV3';
import {
    SherlockChartV3FromJSON,
    SherlockChartV3FromJSONTyped,
    SherlockChartV3ToJSON,
} from './SherlockChartV3';
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
} from './SherlockCiIdentifierV3';
import type { SherlockUserV3 } from './SherlockUserV3';
import {
    SherlockUserV3FromJSON,
    SherlockUserV3FromJSONTyped,
    SherlockUserV3ToJSON,
} from './SherlockUserV3';

/**
 * 
 * @export
 * @interface SherlockAppVersionV3
 */
export interface SherlockAppVersionV3 {
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    appVersion?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    authoredBy?: string;
    /**
     * 
     * @type {SherlockUserV3}
     * @memberof SherlockAppVersionV3
     */
    authoredByInfo?: SherlockUserV3;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    chart?: string;
    /**
     * 
     * @type {SherlockChartV3}
     * @memberof SherlockAppVersionV3
     */
    chartInfo?: SherlockChartV3;
    /**
     * 
     * @type {SherlockCiIdentifierV3}
     * @memberof SherlockAppVersionV3
     */
    ciIdentifier?: SherlockCiIdentifierV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockAppVersionV3
     */
    createdAt?: Date;
    /**
     * Generally the Git commit message
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    gitBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    gitCommit?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockAppVersionV3
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockAppVersionV3
     */
    parentAppVersion?: string;
    /**
     * 
     * @type {object}
     * @memberof SherlockAppVersionV3
     */
    parentAppVersionInfo?: object;
    /**
     * 
     * @type {Date}
     * @memberof SherlockAppVersionV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockAppVersionV3 interface.
 */
export function instanceOfSherlockAppVersionV3(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockAppVersionV3FromJSON(json: any): SherlockAppVersionV3 {
    return SherlockAppVersionV3FromJSONTyped(json, false);
}

export function SherlockAppVersionV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockAppVersionV3 {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appVersion': !exists(json, 'appVersion') ? undefined : json['appVersion'],
        'authoredBy': !exists(json, 'authoredBy') ? undefined : json['authoredBy'],
        'authoredByInfo': !exists(json, 'authoredByInfo') ? undefined : SherlockUserV3FromJSON(json['authoredByInfo']),
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartInfo': !exists(json, 'chartInfo') ? undefined : SherlockChartV3FromJSON(json['chartInfo']),
        'ciIdentifier': !exists(json, 'ciIdentifier') ? undefined : SherlockCiIdentifierV3FromJSON(json['ciIdentifier']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'description': !exists(json, 'description') ? undefined : json['description'],
        'gitBranch': !exists(json, 'gitBranch') ? undefined : json['gitBranch'],
        'gitCommit': !exists(json, 'gitCommit') ? undefined : json['gitCommit'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'parentAppVersion': !exists(json, 'parentAppVersion') ? undefined : json['parentAppVersion'],
        'parentAppVersionInfo': !exists(json, 'parentAppVersionInfo') ? undefined : json['parentAppVersionInfo'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockAppVersionV3ToJSON(value?: SherlockAppVersionV3 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appVersion': value.appVersion,
        'authoredBy': value.authoredBy,
        'authoredByInfo': SherlockUserV3ToJSON(value.authoredByInfo),
        'chart': value.chart,
        'chartInfo': SherlockChartV3ToJSON(value.chartInfo),
        'ciIdentifier': SherlockCiIdentifierV3ToJSON(value.ciIdentifier),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'description': value.description,
        'gitBranch': value.gitBranch,
        'gitCommit': value.gitCommit,
        'id': value.id,
        'parentAppVersion': value.parentAppVersion,
        'parentAppVersionInfo': value.parentAppVersionInfo,
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}

