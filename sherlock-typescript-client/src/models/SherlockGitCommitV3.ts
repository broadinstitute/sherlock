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
 * @interface SherlockGitCommitV3
 */
export interface SherlockGitCommitV3 {
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3
     */
    committedAt?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockGitCommitV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3
     */
    gitBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3
     */
    gitCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockGitCommitV3
     */
    gitRepo?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockGitCommitV3
     */
    id?: number;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockGitCommitV3
     */
    isMainBranch?: boolean;
    /**
     * 
     * @type {number}
     * @memberof SherlockGitCommitV3
     */
    secSincePrev?: number;
    /**
     * 
     * @type {Date}
     * @memberof SherlockGitCommitV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockGitCommitV3 interface.
 */
export function instanceOfSherlockGitCommitV3(value: object): value is SherlockGitCommitV3 {
    return true;
}

export function SherlockGitCommitV3FromJSON(json: any): SherlockGitCommitV3 {
    return SherlockGitCommitV3FromJSONTyped(json, false);
}

export function SherlockGitCommitV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockGitCommitV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'committedAt': json['committedAt'] == null ? undefined : json['committedAt'],
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'gitBranch': json['gitBranch'] == null ? undefined : json['gitBranch'],
        'gitCommit': json['gitCommit'] == null ? undefined : json['gitCommit'],
        'gitRepo': json['gitRepo'] == null ? undefined : json['gitRepo'],
        'id': json['id'] == null ? undefined : json['id'],
        'isMainBranch': json['isMainBranch'] == null ? undefined : json['isMainBranch'],
        'secSincePrev': json['secSincePrev'] == null ? undefined : json['secSincePrev'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockGitCommitV3ToJSON(value?: SherlockGitCommitV3 | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'committedAt': value['committedAt'],
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'gitBranch': value['gitBranch'],
        'gitCommit': value['gitCommit'],
        'gitRepo': value['gitRepo'],
        'id': value['id'],
        'isMainBranch': value['isMainBranch'],
        'secSincePrev': value['secSincePrev'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

