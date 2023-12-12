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
export function instanceOfSherlockGitCommitV3(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockGitCommitV3FromJSON(json: any): SherlockGitCommitV3 {
    return SherlockGitCommitV3FromJSONTyped(json, false);
}

export function SherlockGitCommitV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockGitCommitV3 {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'committedAt': !exists(json, 'committedAt') ? undefined : json['committedAt'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'gitBranch': !exists(json, 'gitBranch') ? undefined : json['gitBranch'],
        'gitCommit': !exists(json, 'gitCommit') ? undefined : json['gitCommit'],
        'gitRepo': !exists(json, 'gitRepo') ? undefined : json['gitRepo'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'isMainBranch': !exists(json, 'isMainBranch') ? undefined : json['isMainBranch'],
        'secSincePrev': !exists(json, 'secSincePrev') ? undefined : json['secSincePrev'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockGitCommitV3ToJSON(value?: SherlockGitCommitV3 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'committedAt': value.committedAt,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'gitBranch': value.gitBranch,
        'gitCommit': value.gitCommit,
        'gitRepo': value.gitRepo,
        'id': value.id,
        'isMainBranch': value.isMainBranch,
        'secSincePrev': value.secSincePrev,
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}

